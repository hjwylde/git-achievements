package gexec

import "log"
import "os/exec"
import "strings"

func AddNote(message string, sha string, ref string) error {
	log.Printf("addNote(%s, %s, %s)\n", message, sha, ref)

	cmd := exec.Command("git", "notes", "--ref="+ref, "add", "-f", "-m", message, sha)

	err := cmd.Run()

	return err
}

func Fetch(remote string, ref string) error {
	log.Printf("fetch(%s, %s)\n", remote, ref)

	cmd := exec.Command("git", "fetch", remote, ref)

	err := cmd.Run()

	return err
}

func getBranch() (string, error) {
	log.Printf("getBranch()\n")

	cmd := exec.Command("git", "rev-parse", "--abbrev-head", "--symbolic-full-name", "head")

	b, err := cmd.Output()

	return string(b), err
}

func GetUpstreamRemote() (string, error) {
	log.Printf("getUpstreamRemote()\n")

	cmd := exec.Command("git", "rev-parse", "--abbrev-head", "--symbolic-full-name", "@{u}")

	b, err := cmd.Output()
	if err != nil {
		return "", err
	}

	branch, err := getBranch()
	if err != nil {
		return "", err
	}

	remote := strings.TrimSuffix(string(b), "/"+branch)

	return remote, nil
}

func GetPushRemote() (string, error) {
	log.Printf("getPushRemote()\n")

	cmd := exec.Command("git", "rev-parse", "--abbrev-head", "--symbolic-full-name", "@{push}")

	b, err := cmd.Output()
	if err != nil {
		return "", err
	}

	branch, err := getBranch()
	if err != nil {
		return "", err
	}

	remote := strings.TrimSuffix(string(b), "/"+branch)

	return remote, nil
}

func ListNotes(ref string) ([]string, error) {
	log.Printf("listNotes(%s)\n", ref)

	cmd := exec.Command("git", "log", "--format=%N", "--notes="+ref)

	b, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	notes := strings.Fields(
		strings.Replace(string(b), "\n\n", "\n", -1),
	)

	return notes, nil
}

func ListRevisions(ref string) ([]string, error) {
	log.Printf("listRevisions(%s)\n", ref)

	cmd := exec.Command("git", "rev-list", ref)

	b, err := cmd.Output()

	revisions := strings.Fields(string(b))

	return revisions, err
}

func MergeNotes(remote string, ref string) error {
	log.Printf("mergeNotes(%s, %s)\n", remote, ref)

	cmd := exec.Command("git", "notes", "merge", "-s cat_sort_uniq", remote+"/"+ref)

	err := cmd.Run()

	return err
}

func PruneNotes(ref string) error {
	log.Printf("pruneNotes(%s)\n", ref)

	cmd := exec.Command("git", "notes", "--ref="+ref, "prune")

	err := cmd.Run()

	return err
}

func Push(remote string, ref string) error {
	log.Printf("push(%s, %s)\n", remote, ref)

	cmd := exec.Command("git", "push", remote, ref)

	err := cmd.Run()

	return err
}
