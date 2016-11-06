package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/toqueteos/clgen"
)

func main() {
	var (
		ref     string
		tpl     string
		mdOut   string
		htmlOut string
		title   string

		err error
	)

	flag.StringVar(&ref, "ref", "HEAD", `ref to start going back, use HEAD~100 for last 100 commits`)
	flag.StringVar(&tpl, "tpl", "changelog.tpl", `template input file`)
	flag.StringVar(&mdOut, "md-out", "changelog.md", `Markdown output file`)
	flag.StringVar(&htmlOut, "html-out", "changelog.html", `HTML output file`)
	flag.StringVar(&title, "title", "Changelog", `HTML document title`)
	flag.Parse()

	start := time.Now()

	commits := clgen.GitLog(ref)
	err = clgen.WriteTemplate(tpl, mdOut, commits)
	if err != nil {
		log.Fatalln("clgen:", err)
	}

	err = clgen.TemplateToHTML(title, mdOut, htmlOut)
	if err != nil {
		log.Fatalln("clgen:", err)
	}

	elapsed := time.Since(start)

	fmt.Println("Done! Elapsed", elapsed)
}
