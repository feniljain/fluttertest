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

func getAllTests(c []byte) ([]string, []string) {
	var completeTests []string
	var tests []string
	content := string(c)
	initialContent := ""
	re := regexp.MustCompile("test[()]['][A-Za-z0-9 ]*[']")
	locs := re.FindAllIndex([]byte(content), -1)
	for _, loc := range locs {
		test, err := isValidBlock(content[loc[0]:])
		if err != nil {
			continue
		}
		title := getTitle(test)
		if title == "" {
			continue
		}
		if initialContent == "" {
			initialContent = content[:loc[0]]
			re1 := regexp.MustCompile("group[()]['][A-Za-z0-9 ]*['][,][ ]*[()]{2}[ ]*[{]")
			initialContent = string(re1.ReplaceAll([]byte(initialContent), []byte("")))
		}
		completeTests = append(completeTests, initialContent+test+");}")
		tests = append(tests, test)
	}
	return tests, completeTests
}

//Parse function parses the given dart code
func Parse(fileName string) error {
	c, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	groups, groupTests := getGroupTests(c)
	fmt.Println()
	fmt.Println("Select group(s) to be executed:")
	fmt.Println()
	i := 0
	for _, group := range groups {
		fmt.Printf("%d. %s\n", i+1, getTitle(group))
		i++
	}
	fmt.Println()
	fmt.Println("----OR----")
	fmt.Println()
	fmt.Println("Select test(s) to be executed:")
	fmt.Println()
	tests, completeTests := getAllTests(c)
	for _, tests := range tests {
		fmt.Printf("%d. %s\n", i+1, getTitle(tests))
		i++
	}
	fmt.Println()
	fmt.Print("Your choice: ")
	var choice int
	fmt.Scanf("%d\n", &choice)
	fmt.Println()
	if choice > len(groupTests) {
		err = runTest(completeTests[choice-len(groupTests)-1])
	} else {
		err = runTest(groupTests[choice-1])
	}
	if err != nil {
		return err
	}
	return nil
}

func getGroupTests(c []byte) ([]string, []string) {
	var groupTests []string
	var groups []string
	content := string(c)
	initialContent := ""
	re := regexp.MustCompile("group()")
	locs := re.FindAllIndex([]byte(content), -1)
	for _, loc := range locs {
		group, err := isValidBlock(content[loc[0]:])
		if err != nil {
			continue
		}
		if initialContent == "" {
			initialContent = content[:loc[0]]
		}
		groupTests = append(groupTests, initialContent+group+");}")
		groups = append(groups, group)
	}
	return groups, groupTests
}

func getTitle(content string) string {
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

func runTest(content string) error {
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

func isValidBlock(content string) (string, error) {
	s := stack.Init()
	var content1 string
	for _, r := range content {
		ch := string(r)
		if ch == "(" || ch == "{" {
			s.Push(ch)
		}
		var r string
		if ch == ")" {
			c, err := s.TopElement()
			if err == nil && c != "(" {
				return "", errors.New("Mismatch of ()")
			}
			if err != nil {
				continue
			}
			r, err = s.Pop()
			if err != nil {
				return "", err
			}
		}
		if ch == "}" {
			c, err := s.TopElement()
			if err == nil && c != "{" {
				return "", errors.New("Mismatch of {}")
			}
			if err != nil {
				continue
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
