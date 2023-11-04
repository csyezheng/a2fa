package fs

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

// findFile join directory with filename, and check if exists, return empty string if not exists
func findFile(dir string, filename string) string {
	path := filepath.Join(dir, filename)
	if _, err := os.Stat(path); err != nil {
		return ""
	}
	return path
}

// findExeDir find the filename in executable directory
// <a2fa_exe_dir>/<filename>
func findExeDir(filename string) (dir string, filePath string) {
	if exePath, err := os.Executable(); err == nil {
		dir = filepath.Dir(exePath)
		filePath = findFile(dir, filename)
	}
	return
}

// findAppDataDir get path to Windows AppData config subdirectory for a2fa and look for the filename
// $AppData/a2fa/<filename>
func findAppDataDir(filename string) (dir string, filePath string) {
	if appDataDir := os.Getenv("APPDATA"); appDataDir != "" {
		dir = filepath.Join(appDataDir, "a2fa")
		filePath = findFile(dir, filename)
	} else {
		slog.Debug("Environment variable APPDATA is not defined and cannot be used as configuration location")
	}
	return
}

// findXDGConfig get path to XDG config subdirectory for a2fa and look for the filename
// $XDG_CONFIG_HOME\a2fa\<filename>
func findXDGConfig(filename string) (dir string, filePath string) {
	if xdgConfigDir := os.Getenv("XDG_CONFIG_HOME"); xdgConfigDir != "" {
		dir = filepath.Join(xdgConfigDir, "a2fa")
		filePath = findFile(filePath, filename)
	}
	return
}

// findHomeDir find current user's home directory
// ~/.config/a2fa/<filename>
func findHomeConfigDir(filename string) (dir string, filePath string) {
	home, err := homedir.Dir()
	if err != nil {
		slog.Debug("Home directory lookup failed and cannot be used as configuration location: %v", err)
		return
	} else if home == "" {
		// On Unix homedir return success but empty string for user with empty home configured in passwd file
		slog.Debug("Home directory not defined and cannot be used as configuration location")
		return
	}
	dir = filepath.Join(home, ".config", "a2fa")
	filePath = findFile(dir, filename)
	return
}

// MakeFilenamePath return the path to the filename
// looking for existing file in prioritized list of known locations
// or set a directory to use for the file when the file is no exist.
func MakeFilenamePath(filename string) (filePath string) {

	var (
		dir           string
		defaultDir    string
		homeConfigDir string
	)
	// <a2fa_exe_dir>/<filename>
	if _, filePath = findExeDir(filename); filePath != "" {
		return
	}

	// windows: $AppData/a2fa/<filename>
	// this is also the default location for new file when no existing is found
	if runtime.GOOS == "windows" {
		if defaultDir, filePath = findAppDataDir(filename); filePath != "" {
			return
		}
	}

	// $XDG_CONFIG_HOME/a2fa/<filename>
	// also looking for this on Windows, for backwards compatibility reasons.
	if dir, filePath = findXDGConfig(filename); filePath != "" {
		return
	}

	if runtime.GOOS != "windows" {
		// On Unix this is also the default location for new config when no existing is found
		defaultDir = dir
	}

	// ~/.config/a2fa/<filename>
	// This is also the fallback location for new config
	// (when $AppData on Windows and $XDG_CONFIG_HOME on Unix is not defined)
	if homeConfigDir, filePath = findHomeConfigDir(filename); filePath != "" {
		return
	}

	slog.Debug(fmt.Sprintf("no existing %s found, create a new one", filename))

	if defaultDir != "" {
		filePath = filepath.Join(defaultDir, filename)
		if err := os.MkdirAll(defaultDir, os.ModePerm); err == nil {
			return
		}
	} else if homeConfigDir != "" {
		filePath = filepath.Join(homeConfigDir, filename)
		if err := os.MkdirAll(homeConfigDir, os.ModePerm); err == nil {
			return
		}
	}
	return filename
}
