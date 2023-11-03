# Contributing
Please, before sending patches, read these brief comments. They are here to
help the project have both its users happy using the program and the
developers/maintainers feel good when trying to change code that other
people contributed.

### Write good commit messages

When you write your pull request and your commit messages, please, be
detailed, explaining why you are doing what you are doing. Don't be afraid
of being too verbose here.  Also, please follow the highly recommended
guidelines on how to write good [good commit messages][commit-msgs].

When in doubt, follow the model of the Linux kernel commit logs. Their
commit messages are some of the best that I have seen. Also, the ffmpeg has
some good messages that I believe that should be followed. If you are in a
hurry, read the section named
["Contributing" from subsurface's README][contributing].

[commit-msgs]: https://robots.thoughtbot.com/5-useful-tips-for-a-better-commit-message
[contributing]: https://github.com/torvalds/subsurface/blob/master/README#L71-L114


### Test that your changes don't break existing functionality

Run the test suite with

    go test

If some test fails, please don't send your changes yet. Fix what broke
before sending your pull request.

If you need to change the test suite, explain in the commit message why it
needs to be changed.


### Dependency management

The Prometheus project uses [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) to manage dependencies on external packages.

To add or update a new dependency, use the `go get` command:

```bash
# Pick the latest tagged release.
go get example.com/some/module/pkg@latest

# Pick a specific version.
go get example.com/some/module/pkg@vX.Y.Z
```

Tidy up the `go.mod` and `go.sum` files:

```bash
# The GO111MODULE variable can be omitted when the code isn't located in GOPATH.
GO111MODULE=on go mod tidy
```

You have to commit the changes to `go.mod` and `go.sum` before submitting the pull request.


### Check for potential bugs

Please, help keep the code tidy by checking for any potential bugs with the
help of golangci-lint. Fixing the code to comply with the linter's recommendation
is in general the preferred course of action.

If you happen to find any issue reported by these programs, I welcome you to
fix them.  Many of the issues are usually very easy to fix and they are a
great way to start contributing to this (and other projects in general).
Furthermore, we all benefit from a better code base.