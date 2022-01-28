// Nicholas Ferreira
// 28/01/22

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

//check if array s contains string str
func contains(s []string, str string) bool {
	for _, i := range s {
		if i == str {
			return true
		}
	}
	return false
}

//check if len(argv) == 2
func parseArgs() string {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		os.Exit(1)
	}
	return args[1]
}

//make a get request to url and return body content
func get_request(url string) *bufio.Scanner {
	get, _ := http.Get(url)
	resp, _ := ioutil.ReadAll(get.Body)
	code := bufio.NewScanner(strings.NewReader(string(resp)))
	return code
}

//extract body lines containing domain
func parse_body(body *bufio.Scanner, domain string) string {
	var urls string
	for body.Scan() {
		if strings.Contains(body.Text(), "<TD>") && strings.Contains(body.Text(), domain) {
			//replace <td>,</td>,<br>, </br> and double spaces for a single space
			u := regexp.MustCompile("(</*TD>|</*BR>| +|\t+)").ReplaceAllString(body.Text(), " ")
			urls += u
		}
	}
	return urls
}

//remove duplicate itens in array and print them
func deduplicate(array []string) {
	var final_subs []string
	for _, u := range array {
		if !contains(final_subs, u) {
			final_subs = append(final_subs, u)
			fmt.Println(u)
		}
	}
}

func main() {
	arg := parseArgs()
	url := "https://crt.sh/?q=" + arg
	body := get_request(url)
	urls := strings.Split(parse_body(body, arg), " ")
	sort.Strings(urls)
	fmt.Printf("[+] Subdomains found:")
	deduplicate(urls)
}
