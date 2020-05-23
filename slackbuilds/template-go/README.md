## Slackbuild for Go things

1. Edit the top of the file. (copyright, project name, import path)

2. Download the .zip or .tar.gz from git host (or pack it yourself)

3. Have Go installed (I prefer manually install to /usr/local/go, and
manually symlink `/usr/local/go/bin/* /usr/local/bin/`

4. `./yourprogram.SlackBuild`


### Clever script

looks for `./cmd/{things}`, falls back to building `./`

(define CMDS to build only one or more commands from ./cmd/)

can be ran from a directory containing only a `slack-desc` file

(define SBO_GOGET=1, IMPORT_PATH, and PRGNAM)
  

### TODO

  * use go list to figure out if `./` is a main package
  * handle repos with no go.mod file?
  * handle librarys (no main packages in repo)


### Example

Build markdownd from latest master branch commit:

```
sudo env SBO_GOGET=1 IMPORT_PATH=github.com/aerth/markdownd PRGNAM=markdownd bash ../template-go/go-template.SlackBuild
```

Set version override:

```
sudo env VERSION=foo SBO_GOGET=1 IMPORT_PATH=github.com/aerth/markdownd PRGNAM=markdownd ../template-go/go-template.SlackBuild
```

Build commands foo and bar from the master branch ./cmd/foo and ./cmd/bar:

Will build foobase-${commithash slackbuild.

In this example, go-template.SlackBuild is executable and lives in ~/bin/

```
sudo env SBO_GOGET=1 CMDS='foo bar' IMPORT_PATH=github.com/user/foobase ~/bin/go-template.SlackBuild
```

