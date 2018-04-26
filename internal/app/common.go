package app

import "log"
import "os/exec"
import "strings"

func addNote(message string, sha string, ref string) error {
	log.Printf("addNote(%s, %s, %s)\n", message, sha, ref)

	cmd := exec.Command("git", "notes", "--ref="+ref, "add", "-f", "-m", message, sha)

	err := cmd.Run()

	return err
}

func fetch(remote string, ref string) error {
	log.Printf("fetch(%s, %s)\n", remote, ref)

	cmd := exec.Command("git", "fetch", remote, ref)

	err := cmd.Run()

	return err
}

func getBranch() (string, error) {
	log.Printf("getBranch()\n")

	cmd := exec.Command("git", "rev-parse", "--abbrev-head", "--symbolic-full-name", "head")

	output, err := cmd.Output()

	return string(output), err
}

func getUpstreamRemote() (string, error) {
	log.Printf("getUpstreamRemote()\n")

	cmd := exec.Command("git", "rev-parse", "--abbrev-head", "--symbolic-full-name", "@{u}")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	branch, err := getBranch()
	if err != nil {
		return "", err
	}

	remote := strings.TrimSuffix(string(output), "/"+branch)

	return remote, nil
}

func getPushRemote() (string, error) {
	log.Printf("getPushRemote()\n")

	cmd := exec.Command("git", "rev-parse", "--abbrev-head", "--symbolic-full-name", "@{push}")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	branch, err := getBranch()
	if err != nil {
		return "", err
	}

	remote := strings.TrimSuffix(string(output), "/"+branch)

	return remote, nil
}

func listRevisions(ref string) ([]string, error) {
	log.Printf("listRevisions(%s)\n", ref)

	cmd := exec.Command("git", "rev-list", ref)

	output, err := cmd.Output()

	revisions := strings.Fields(string(output))

	return revisions, err
}

func mergeNotes(remote string, ref string) error {
	log.Printf("mergeNotes(%s, %s)\n", remote, ref)

	cmd := exec.Command("git", "notes", "merge", "-s cat_sort_uniq", remote+"/"+ref)

	err := cmd.Run()

	return err
}

func pruneNotes(ref string) error {
	log.Printf("pruneNotes(%s)\n", ref)

	cmd := exec.Command("git", "notes", "--ref="+ref, "prune")

	err := cmd.Run()

	return err
}

func push(remote string, ref string) error {
	log.Printf("push(%s, %s)\n", remote, ref)

	cmd := exec.Command("git", "push", remote, ref)

	err := cmd.Run()

	return err
}
