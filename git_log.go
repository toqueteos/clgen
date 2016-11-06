package clgen

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type Commit struct {
	Hash  string
	Date  time.Time
	Title string
	Body  string
	Tags  []string
}

func (c Commit) String() string {
	return fmt.Sprintf("Commit{Hash: %q, Date: %q, Title: %q, Body: %q, Tags: %q}",
		c.Hash, c.Date, c.Title, c.Body, c.Tags)
}

const (
	RawCommitSeparator = "¬¬¬"
	RawCommitSuffix    = "···"
	TagSeparator       = ", "
	TagPrefix          = "tag: "
)

func GitLog(ref string) []Commit {
	logfmt := []string{"%H", "%cI", "%s", "%b", "%D"}
	pretty := strings.Join(logfmt, RawCommitSeparator) + RawCommitSuffix

	cmd := exec.Command(
		"git",
		"--no-pager",
		"log",
		"--all",
		// "--graph",
		"--pretty="+pretty,
		ref,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("git log:", err)
	}

	return parseLines(out)
}

func parseLines(output []byte) []Commit {
	lines := bytes.Split(output, []byte(RawCommitSuffix))

	var commits []Commit
	for _, line := range lines {
		rawCommit := strings.Split(string(line), RawCommitSeparator)

		for idx, part := range rawCommit {
			rawCommit[idx] = trim(part)
		}

		if len(rawCommit) != 5 {
			// fmt.Println("err?", rawCommit)
			continue
		}

		commits = append(commits, lineToCommit(rawCommit))
	}

	return commits
}

func trim(input string) string {
	return strings.TrimSpace(input)
}

func lineToCommit(line []string) Commit {
	date, _ := time.Parse(time.RFC3339, line[1])
	return Commit{
		Hash:  line[0],
		Date:  date,
		Title: line[2],
		Body:  line[3],
		Tags:  parseTags(line[4]),
	}
}

func parseTags(rawRefs string) []string {
	var tags []string

	if len(rawRefs) == 0 {
		return nil
	}

	refs := strings.Split(rawRefs, TagSeparator)
	for _, ref := range refs {
		if !strings.HasPrefix(ref, TagPrefix) {
			continue
		}

		tags = append(tags, strings.TrimPrefix(ref, TagPrefix))
	}

	return tags
}
