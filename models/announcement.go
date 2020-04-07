package models

import (
	"fmt"
	"time"
)

type Announcement struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Publisher   string    `json:"publisher"`
	PublishTime time.Time `json:"publish_time"`
	Content     string    `json:"content"`
	Image       string    `json:"image"`
	FilePath    string    `json:"file_path"`
	HaveRead    bool      `json:"have_read"`
}

// 添加公告
func CreateAnnouncement(an *Announcement) error {
	err := db.Create(&an).Error
	if err != nil {
		return err
	}
	return nil
}

// 通过id查询公告
func FetchById(id int) (*Announcement, error) {
	tx := db.Begin()
	var an Announcement
	err := tx.Debug().Where("id=?", id).First(&an).Error
	if err != nil {
		fmt.Println("FetById error find")
		fmt.Printf("%#v", err)
	}
	tx.Commit()
	return &an, nil
}

// 通过选项获取公告 false-未读 true-已读
func FetchAnnByOption(option bool) ([]Announcement, error) {
	tx := db.Begin()
	var an []Announcement
	err := tx.Debug().Where("have_read=?", option).Find(&an).Error
	if err != nil {
		fmt.Println("FetByOption error find")
		fmt.Printf("%#v", err)
	}
	tx.Commit()
	return an, nil
}

// 获取所有公告
func FetchAllAnnouncement() (ans []*Announcement, err error) {
	ans = []*Announcement{}
	err = db.Find(&ans).Error
	if err != nil {
		return nil, err
	}
	return ans, nil

}

// 更新公告
func UpdateAnn(an *Announcement) (*Announcement, error) {
	err := db.Save(&an).Error
	return an, err
}

// 删除公告
func DeleteAnn(id int) error {
	an, _ := FetchById(id)
	if an == nil {
		fmt.Println("delete FetchById 失败")
	}
	err := db.Delete(an).Error
	return err
}