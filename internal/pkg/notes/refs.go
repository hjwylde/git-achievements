package notes

var BaseRef = "achievements"
var ProgressRef = BaseRef + "/progress"

func AchievementProgressRef(achievement Achievement) string {
	return ProgressRef + "/" + achievement.Group + "-" + achievement.Code
}
