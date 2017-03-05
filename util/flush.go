package util

import (
	"io/ioutil"
	"os"

	"scrape-sd/showdown"
)

// FlushAsCSV flushes []showdown.BattleLink as CSV file.
func FlushAsCSV(links []showdown.BattleLink, toRelativePath string) {
	var lines []byte
	for i := range links {
		line := links[i].URL + "," + links[i].Text + "\n"
		lines = append(lines, []byte(line)...)
	}
	ioutil.WriteFile(toRelativePath, lines, os.ModeAppend)
}
