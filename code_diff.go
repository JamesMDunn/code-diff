package main

import (
	"log"
	"strings"

	git "github.com/libgit2/git2go/v31"
)

func main() {

	colorRed := "\033[31m"
	colorWhite := "\033[37m"
	colorGreen := "\033[32m"
	//test repo
	repoPath := "/Users/jamesdunn/Documents/projects/VybinQL"
	gitRepo, err := git.OpenRepository(repoPath)
	if err != nil {
		log.Fatal(err)
	}
	// commit SHA-1 checksum do i care to commit this?
	commitID := `eb833ae612537d7b2c15b8184b5938ee448d2ab8`
	commitOid, err := git.NewOid(commitID)
	if err != nil {
		log.Fatal(err)
	}
	commit, err := gitRepo.LookupCommit(commitOid)
	if err != nil {
		log.Fatal(err)
	}
	commitTree, err := commit.Tree()
	if err != nil {
		log.Fatal(err)
	}
	options, err := git.DefaultDiffOptions()
	if err != nil {
		log.Fatal(err)
	}
	options.IdAbbrev = 40
	var parentTree *git.Tree
	if commit.ParentCount() > 0 {
		parentTree, err = commit.Parent(0).Tree()
		if err != nil {
			log.Fatal(err)
		}
	}
	gitDiff, err := gitRepo.DiffTreeToTree(parentTree, commitTree, &options)
	if err != nil {
		log.Fatal(err)
	}

	// Show all file patch diffs in a commit.
	numDeltas, err := gitDiff.NumDeltas()
	if err != nil {
		log.Fatal(err)
	}
	codeDiff := []string{}
	for d := 0; d < numDeltas; d++ {
		patch, err := gitDiff.Patch(d)
		if err != nil {
			log.Fatal(err)
		}
		patchString, err := patch.String()
		if err != nil {
			log.Fatal(err)
		}
		//make this a function
		patchSplit := strings.Split(patchString, "\n")
		for i := 0; i < len(patchSplit); i++ {

			if strings.HasPrefix(patchSplit[i], "+") {
				codeDiff = append(codeDiff, colorGreen+patchSplit[i])
			} else if strings.HasPrefix(patchSplit[i], "-") {
				codeDiff = append(codeDiff, colorRed+patchSplit[i])
			} else {
				codeDiff = append(codeDiff, colorWhite+patchSplit[i])
			}
			println("this is split stuff", patchSplit[i])
		}
		patch.Free()
	}
	println("testing codeDiff", strings.Join(codeDiff, "\n"))
}
