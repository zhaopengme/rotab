package akismet

import (
	"os"
	"testing"
)

var chromeUA = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.146 Safari/537.36"
var notSpam = "But the good of the scorpion is not the good of the frog, yes?"

type testCaseCheck struct {
	name     string
	c        Comment
	isSpam   bool
	hasError bool
}

func TestCheck(t *testing.T) {
	tests := []func(*testing.T, string){
		testInvalids,
		testSpam,
	}

	if os.Getenv("AKISMET_KEY") == "" {
		t.Errorf("AKISMET_KEY not set")
		return
	}

	key := os.Getenv("AKISMET_KEY")
	for _, test := range tests {
		test(t, key)
	}
}

func runTests(t *testing.T, testCases []testCaseCheck, key string) {
	for _, tc := range testCases {
		isSpam, err := Check(&tc.c, key)

		if isSpam != tc.isSpam {
			t.Errorf("%s: expected isSpam=%v, got isSpam=%v", tc.name, tc.isSpam, isSpam)
		}

		if tc.hasError != (err != nil) {
			t.Errorf("%s: hasError=%v, got err=%v", tc.name, tc.hasError, err)
		}
	}
}

func testInvalids(t *testing.T, key string) {
	testCases := []testCaseCheck{
		testCaseCheck{
			"Comment{} missing everything",
			Comment{},
			true,
			true,
		},

		testCaseCheck{
			"Comment{} missing blog",
			Comment{UserIP: "8.8.8.8", UserAgent: chromeUA, CommentContent: "Hello!"},
			true,
			true,
		},
	}

	runTests(t, testCases, key)
}

func testSpam(t *testing.T, key string) {
	testCases := []testCaseCheck{
		testCaseCheck{
			"Typical 419 scam",
			Comment{Blog: "https://example.com", UserIP: "8.8.8.8", UserAgent: chromeUA, CommentContent: "Send $6,321 to my western union account and receive $1,000,000 today http://419.com http://419.com"},
			true,
			false,
		},

		testCaseCheck{
			"Outed by user agent",
			Comment{Blog: "https://example.com", UserIP: "8.8.8.8", UserAgent: "Python-urllib/2.1", CommentContent: notSpam},
			true,
			false,
		},

		testCaseCheck{
			"Known to be a spammer by email",
			Comment{Blog: "https://example.com", UserIP: "8.8.8.8", UserAgent: chromeUA, CommentContent: notSpam, CommentAuthorEmail: "akismet-guaranteed-spam@example.com"},
			true,
			false,
		},
	}

	runTests(t, testCases, key)
}

func testHam(t *testing.T, key string) {
	testCases := []testCaseCheck{
		testCaseCheck{
			"Comment{} missing blog",
			Comment{Blog: "https://example.com", UserIP: "8.8.8.8", UserAgent: chromeUA, CommentContent: notSpam},
			false,
			false,
		},
	}

	runTests(t, testCases, key)
}
