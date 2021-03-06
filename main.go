package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("failed to open git repo: %v", err)
	}

	head, err := repo.Head()
	if err != nil {
		log.Fatalf("failed to get repo HEAD: %v", err)
	}

	iter, err := repo.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		log.Fatalf("failed to get repo log: %v", err)
	}

	lastCommit, err := iter.Next()
	if err != nil {
		log.Fatalf("failed to get last commit: %v", err)
	}

	fmt.Println("# generated by benchcontext")
	fmt.Printf("by: %s\n", lastCommit.Author.Email)
	fmt.Printf("branch: %s\n", head.Name().Short())
	fmt.Printf("commit: %s\n", lastCommit.Hash)
	fmt.Printf("commit-time: %s\n", lastCommit.Author.When.Format(time.RFC3339))
}
