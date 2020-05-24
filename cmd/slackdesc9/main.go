// Copyright (c) 2020 aerth <aerth@memeware.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.

// slackdesc9 generates slack-desc files for SlackBuilds
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		pkgname     = flag.String("name", "", "")
		pkgdescrip  = flag.String("short", "", "")
		pkglongdesc = flag.String("long", "", "")
		pkghomepage = flag.String("web", "", "")
		pkgversion  = flag.String("version", "", "")
	)

	flag.Parse()
	log.SetFlags(0)

	var config = Config{
		Name:    *pkgname,
		Short:   *pkgdescrip,
		Long:    *pkglongdesc,
		Web:     *pkghomepage,
		Version: *pkgversion,
	}

	_, err := os.Stat("slack-desc")
	if !os.IsNotExist(err) {
		if err != nil {
			log.Fatalln("fatal:", err)
		}
		log.Fatalln("fatal:", "'slack-desc' file exists")
	}

	log.Println("creating new 'slack-desc' file")

	err = interactive(&config)
	if err != nil {
		log.Fatalln("fatal:", err)
	}

	f, err := os.Create("slack-desc")
	if err != nil {
		log.Fatalln("fatal:", err)
	}

	if err := writeFile(f, config); err != nil {
		f.Close()
		log.Fatalln("fatal:", err)
	}

	if err := f.Close(); err != nil {
		log.Fatalln("fatal:", err)
	}

	log.Println("slack-desc file created")

	os.Exit(0)
}

type Config struct {
	Name    string
	Short   string
	Long    string
	Web     string
	Version string
}

func printerr(f string, i ...interface{}) {
	fmt.Fprintf(os.Stderr, f, i...)
}

func ErrEmpty(s string) error {
	return fmt.Errorf("%q is empty", s)
}

func interactive(c *Config) error {
	printerr("package name: [%s]: ", c.Name)
	c.Name = readline(c.Name)
	if c.Name == "" {
		return ErrEmpty("name")
	}

	printerr("package version: [%s]: ", c.Version)
	c.Version = readline(c.Version)
	if c.Version == "" {
		return ErrEmpty("version")
	}

	printerr("package description (short): [%s]: ", c.Short)
	c.Short = readline(c.Short)
	if c.Short == "" {
		return ErrEmpty("short-description")
	}

	printerr("package description (long): [%s]: ", c.Long)
	c.Long = readline(c.Long)
	if c.Long == "" {
		return ErrEmpty("long-description")
	}

	printerr("package website: [%s]: ", c.Web)
	c.Web = readline(c.Web)
	if c.Web == "" {
		return ErrEmpty("website")
	}

	return nil
}

func readline(def string) string {
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalln("read stdin error:", err)
	}
	s = strings.TrimSuffix(s, "\n")
	if s != "" {
		return s
	}
	return def
}

func readmultiline(escapestring string) string {
	return ""
}

// TODO: go templates?
const fmtstr = `
1prgName|-----handy-ruler--9ruler
1prgName: 1prgName 2prgVersion 3pkgShortDesc
1prgName:
4pkgLongDesc
1prgName: 5pkgWeb
`

func writeFile(f *os.File, c Config) error {
	var t = fmtstr
	t = strings.Replace(t, "1prgName", c.Name, -1)
	t = strings.Replace(t, "2prgVersion", c.Version, -1)
	t = strings.Replace(t, "3pkgShortDesc", c.Short, -1)
	t = strings.Replace(t, "4pkgLongDesc", c.GetLongDescription(), -1)
	t = strings.Replace(t, "5pkgWeb", c.Web, -1)
	t = strings.Replace(t, "9ruler", strings.Repeat("-", 79-(19+len(c.Name))), 1)
	_, err := f.Write([]byte(t))
	return err
}

func (c Config) GetLongDescription() string {
	lines := [10]string{}
	long := strings.Replace(c.Long, "\n", " ", -1)
	long = strings.Replace(long, "\t", " ", -1)
	chars := long[:]
	l := len(chars)
	max := 79 - len(c.Name) - 2 // max each line

	ind := 0
	for i := range lines {
		start := i * max
		end := start + max
		if start > l {
			break
		}
		if end >= l {
			end = l
		}
		lines[i] = chars[start:end]
		ind += max
	}

	buf := new(strings.Builder)
	for i := 0; i < len(lines); i++ {
		fmt.Fprintf(buf, "%s: %s\n", c.Name, lines[i])
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
