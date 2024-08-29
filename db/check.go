package db

import "core-cli/model"

func PlayerExistsByIntraName(intraName string) bool {
	var count int64
	db.Model(&model.Player{}).Where("LOWER(intra_name) = LOWER(?)", intraName).Count(&count)
	return count > 0
}

func PlayerExistsByGithubName(githubName string) bool {
	var count int64
	db.Model(&model.Player{}).Where("LOWER(github_name) = LOWER(?)", githubName).Count(&count)
	return count > 0
}

func TeamExistsByName(name string) bool {
	var count int64
	db.Model(&model.Team{}).Where("LOWER(name) = LOWER(?)", name).Count(&count)
	return count > 0
}
