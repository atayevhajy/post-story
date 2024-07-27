package main

import (
	_ "github.com/lib/pq"
)

func getFollowers(userID string) ([]string, error) {
	rows, err := db.Query("SELECT follower_id FROM followers WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followerIDs []string
	for rows.Next() {
		var followerID string
		if err := rows.Scan(&followerID); err != nil {
			return nil, err
		}
		followerIDs = append(followerIDs, followerID)
	}
	return followerIDs, nil
}
