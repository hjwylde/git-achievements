package notes

var BaseRef = "achievements"
var ProgressRef = BaseRef + "/progress"
var UnlockedRef = BaseRef + "/unlocked"

func AchievementProgressRef(achievement Achievement) string {
	return ProgressRef + "/" + achievement.Group + "-" + achievement.Code
}

func AchievementUnlockedRef(achievement Achievement) string {
	return UnlockedRef + "/" + achievement.Group + "-" + achievement.Code
}
