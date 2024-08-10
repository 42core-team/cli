package db

import "core-cli/model"

func PlayerExistsByIntraName(intraName string) bool {
	var count int64
	db.Model(&model.Player{}).Where("intra_name = ?", intraName).Count(&count)
	return count > 0
}

func PlayerExistsByGithubName(githubName string) bool {
	var count int64
	db.Model(&model.Player{}).Where("github_name = ?", githubName).Count(&count)
	return count > 0
}
