package sentences

import (
	"io/ioutil"
	"strings"
	"testing"

	td "github.com/neurosnap/sentences/data"
)

func loadTokenizer(data string) *DefaultSentenceTokenizer {
	b, err := td.Asset(data)
	if err != nil {
		panic(err)
	}

	training, err := LoadTraining(b)
	if err != nil {
		panic(err)
	}

	return NewSentenceTokenizer(training)
}

func readFile(fname string) string {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func getFileLocation(prefix, original, expected string) []string {
	origText := strings.Join([]string{prefix, original}, "")
	expectedText := strings.Join([]string{prefix, expected}, "")
	return []string{origText, expectedText}
}

func TestEnglish(t *testing.T) {
	t.Log("Starting test suite ...")

	tokenizer := loadTokenizer("data/english.json")

	prefix := "test_files/english/"

	testFiles := [][]string{
		getFileLocation(prefix, "carolyn.txt", "carolyn_s.txt"),
		getFileLocation(prefix, "ecig.txt", "ecig_s.txt"),
		getFileLocation(prefix, "foul_ball.txt", "foul_ball_s.txt"),
		getFileLocation(prefix, "fbi.txt", "fbi_s.txt"),
		getFileLocation(prefix, "dre.txt", "dre_s.txt"),
		getFileLocation(prefix, "dr.txt", "dr_s.txt"),
		getFileLocation(prefix, "quotes.txt", "quotes_s.txt"),
		getFileLocation(prefix, "kiss.txt", "kiss_s.txt"),
		getFileLocation(prefix, "kentucky.txt", "kentucky_s.txt"),
		getFileLocation(prefix, "iphone6s.txt", "iphone6s_s.txt"),
		getFileLocation(prefix, "lebanon.txt", "lebanon_s.txt"),
		getFileLocation(prefix, "duma.txt", "duma_s.txt"),
		getFileLocation(prefix, "demolitions.txt", "demolitions_s.txt"),
		getFileLocation(prefix, "qa.txt", "qa_s.txt"),
		getFileLocation(prefix, "anarchy.txt", "anarchy_s.txt"),
		getFileLocation(prefix, "ethicist.txt", "ethicist_s.txt"),
		getFileLocation(prefix, "self_reliance.txt", "self_reliance_s.txt"),
		getFileLocation(prefix, "punct.txt", "punct_s.txt"),
		getFileLocation(prefix, "clinton.txt", "clinton_s.txt"),
		getFileLocation(prefix, "markets.txt", "markets_s.txt"),
		getFileLocation(prefix, "nyfed.txt", "nyfed_s.txt"),
	}

	for _, f := range testFiles {
		actualText := readFile(f[0])
		expectedText := readFile(f[1])
		expected := strings.Split(expectedText, "{{sentence_break}}")

		t.Log(f[0])
		sentences := tokenizer.Tokenize(actualText)
		for index, s := range sentences {
			sentence := strings.TrimSpace(s.Text)
			if sentence != strings.TrimSpace(expected[index]) {
				t.Logf("Actual  : %q", sentence)
				t.Log("--------")
				t.Logf("Expected: %q", strings.TrimSpace(expected[index]))
				t.Fatalf("%s line %d: Actual sentence does not match expected sentence", f[0], index)
			}
		}
	}

}
