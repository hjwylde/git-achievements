package gexec

import "encoding/json"
import "github.com/hjwylde/git-achievements/internal/pkg/git"
import "github.com/hjwylde/git-achievements/internal/pkg/notes"
import "log"
import "os/exec"
import "strings"
import "time"

func AddNote(message string, sha string, ref string) error {
	log.Printf("AddNote(%s, %s, %s)\n", message, sha, ref)

	cmd := exec.Command("git", "notes", "--ref="+ref, "add", "-f", "-m", message, sha)

	err := cmd.Run()

	return err
}

func Fetch(remote string, ref string) error {
	log.Printf("Fetch(%s, %s)\n", remote, ref)

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
	log.Printf("GetUpstreamRemote()\n")

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

func GetProgress(ref string) (map[git.Commit]notes.Progress, error) {
	log.Printf("GetProgress(%s)\n", ref)

	cmd := exec.Command("git", "log", "--format=%H	%aD	%N", "--notes="+ref)

	b, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	m := make(map[git.Commit]notes.Progress)
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}

		fields := strings.Split(line, "	")

		authorDate, err := time.Parse(time.RFC1123Z, fields[1])
		if err != nil {
			return nil, err
		}

		commit := git.Commit{
			Sha:        fields[0],
			AuthorDate: authorDate,
		}

		var progress notes.Progress
		if err = json.Unmarshal([]byte(fields[2]), &progress); err != nil {
			return nil, err
		}

		m[commit] = progress
	}

	return m, nil
}

func GetPushRemote() (string, error) {
	log.Printf("GetPushRemote()\n")

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

func GetUnlocked(ref string) (map[git.Commit]notes.Unlocked, error) {
	log.Printf("GetUnlocked(%s)\n", ref)

	cmd := exec.Command("git", "log", "--format=%H	%aD	%N", "--notes="+ref)

	b, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	m := make(map[git.Commit]notes.Unlocked)
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}

		fields := strings.Split(line, "	")

		authorDate, err := time.Parse(time.RFC1123Z, fields[1])
		if err != nil {
			return nil, err
		}

		commit := git.Commit{
			Sha:        fields[0],
			AuthorDate: authorDate,
		}

		var unlocked notes.Unlocked
		if err = json.Unmarshal([]byte(fields[2]), &unlocked); err != nil {
			return nil, err
		}

		m[commit] = unlocked
	}

	return m, nil
}

func ListRevisions(ref string) ([]string, error) {
	log.Printf("ListRevisions(%s)\n", ref)

	cmd := exec.Command("git", "rev-list", ref)

	b, err := cmd.Output()

	revisions := strings.Fields(string(b))

	return revisions, err
}

func MergeNotes(remote string, ref string) error {
	log.Printf("MergeNotes(%s, %s)\n", remote, ref)

	cmd := exec.Command("git", "notes", "merge", "-s cat_sort_uniq", remote+"/"+ref)

	err := cmd.Run()

	return err
}

func PruneNotes(ref string) error {
	log.Printf("PruneNotes(%s)\n", ref)

	cmd := exec.Command("git", "notes", "--ref="+ref, "prune")

	err := cmd.Run()

	return err
}

func Push(remote string, ref string) error {
	log.Printf("Push(%s, %s)\n", remote, ref)

	cmd := exec.Command("git", "push", remote, ref)

	err := cmd.Run()

	return err
}
