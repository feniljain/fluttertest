package parser

import (
	"errors"
	"flutterTest/stack"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

//Parse function parses the given dart code
func Parse(fileName string) error {
	c, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	var tests []string
	var groups []string
	content := string(c)
	initialContent := ""
	re := regexp.MustCompile("group()")
	locs := re.FindAllIndex([]byte(content), -1)
	for _, loc := range locs {
		group, err := isValidGroup(content[loc[0]:])
		if err != nil {
			continue
		}
		if initialContent == "" {
			initialContent = content[:loc[0]]
		}
		tests = append(tests, initialContent+group+");}")
		groups = append(groups, group)
	}
	fmt.Println("Select group(s) to be executed:")
	for i, group := range groups {
		fmt.Printf("%d. %s\n", i+1, getGroupTitle(group))
	}
	var groupChosen int
	fmt.Scanf("%d\n", &groupChosen)
	err = runGroupTest(tests[groupChosen-1])
	if err != nil {
		return err
	}
	return nil
}

func getGroupTitle(content string) string {
	title := ""
	record := false
	quotesCnt := 0
	for _, r := range content {
		ch := string(r)
		if ch == "(" {
			record = true
		}
		if record && (ch == "\"" || ch == "'") {
			quotesCnt++
		}
		if record && quotesCnt == 1 {
			if ch != "\"" && ch != "'" {
				title += ch
			}
		}
	}
	return title
}

func runGroupTest(content string) error {
	f, err := os.Create("test/temp_test.dart")
	if err != nil {
		return err
	}
	r := strings.NewReader(content)
	io.Copy(f, r)
	cmd := exec.Command("flutter", "test", "test/temp_test.dart")
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	return nil
}

func isValidGroup(content string) (string, error) {
	s := stack.Init()
	var content1 string
	for _, r := range content {
		ch := string(r)
		if ch == "(" || ch == "{" {
			s.Push(ch)
		}
		var r string
		var err error
		if ch == ")" {
			if s.TopElement() != "(" {
				return "", errors.New("Mismatch of ()")
			}
			r, err = s.Pop()
			if err != nil {
				return "", err
			}
		}
		if ch == "}" {
			if s.TopElement() != "{" {
				return "", errors.New("Mismatch of {}")
			}
			r, err = s.Pop()
			if err != nil {
				return "", err
			}
		}
		content1 += string(ch)
		if s.Top == 0 && r == "{" {
			return content1, nil
		}
	}
	return content1, nil
}
