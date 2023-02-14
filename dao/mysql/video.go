package mysql

import (
	"gorm.io/gorm"
	"log"
	"tiktok/model"
	"time"
)

func InsertVideo(v *model.Video) (err error) {
	if err = DB.Create(v).Error; err != nil {
		log.Println("InsertVideo failed to insert", err)
	}
	return
}

func QueryVideoList(id string, start interface{}) ([]model.Video, error) {
	var list []model.Video
	var tx *gorm.DB
	if id == "" {
		tx = DB.Select("id, author_id, play_url, cover_url, favorite_count, comment_count, title, create_time").Where("create_time < ?", start.(time.Time)).Order("create_time desc").Limit(30).Find(&list)
	} else {
		tx = DB.Select("id, author_id, play_url, cover_url, favorite_count, comment_count, title").Where("author_id = ?", id).Order("create_time desc").Find(&list)

	}
	return list, tx.Error
}

func QueryFavorite(userId, videioId interface{}) bool {
	n := 0
	DB.Table("user_favorite_video").Select("count(*)").Where("`user_id` = ? and `videio_id` = ?", userId, videioId).Find(&n)
	return n == 1
}
