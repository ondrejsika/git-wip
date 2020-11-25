// git-wip
//
// Author: Ondrej Sika <ondrej@ondrejsika.com>
// Source: https://github.com/ondrejsika/git-wip
//

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

type stateStruct map[string]int

func main() {
	// Get stateFilePath ~/.gitwip.state.json
	usr, _ := user.Current()
	stateFilePath := filepath.Join(usr.HomeDir, ".git-wip.state.json")

	// Setup state
	state := stateStruct{}

	// Load state if state file exists
	_, err := os.Stat(stateFilePath)
	if !os.IsNotExist(err) {
		jsonFile, _ := os.Open(stateFilePath)
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &state)
	}

	// Get latest WIP count from state and increase by 1
	repoPath, _ := filepath.Abs(".")
	if _, ok := state[repoPath]; !ok {
		state[repoPath] = 0
	}
	state[repoPath] = state[repoPath] + 1

	// Commit all changes
	r, _ := git.PlainOpen(".")
	w, _ := r.Worktree()

	ps, _ := gitignore.ReadPatterns(w.Filesystem, []string{})
	w.Excludes = ps
	w.AddWithOptions(&git.AddOptions{
		All: true,
	})

	commit, _ := w.Commit("WIP "+strconv.Itoa(state[repoPath])+"\n\nCommited by https://github.com/ondrejsika/git-wip", &git.CommitOptions{})
	obj, _ := r.CommitObject(commit)
	fmt.Println(obj)

	// Save new state
	file, _ := json.Marshal(state)
	ioutil.WriteFile(stateFilePath, file, 0644)
}
