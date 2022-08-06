package model

import (
	"database/sql"
	"time"
)

/* 建表语句
CREATE TABLE `blacklists`  (
	`id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`email` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
	`type` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
	`created_at` timestamp NULL DEFAULT NULL,
	`updated_at` timestamp NULL DEFAULT NULL,
	PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = DYNAMIC;
*/

type Blacklists struct {
	Id        uint         `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Email     string       `gorm:"column:email;type:varchar(50);NOT NULL" json:"email"`
	Type      string       `gorm:"column:type;type:varchar(20);NOT NULL" json:"type"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

/*
CREATE DATABASE jacquesyu_user_info_db;
CREATE TABLE IF NOT EXISTS `jacquesyu_user_info_db`.`user_info_tbl`(
  `id` BIGINT(20) UNSIGNED AUTO_INCREMENT COMMENT '用户 ID',
  `real_name` VARCHAR(64) NOT NULL COMMENT '用户真实姓名',
  `nick_name` VARCHAR(64) NOT NULL COMMENT '用户昵称',
  `phone_num` CHAR(20) NOT NULL COMMENT '用户手机',
  `birth_date` DATE NOT NULL COMMENT '用户生日',
  `create_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` CHAR(1) DEFAULT "0" COMMENT '是否已删除',
  PRIMARY KEY ( `id` ),
  INDEX `idx_phone_num_available` (`phone_num`,`is_deleted`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/

type UserInfoTbl struct {
	Id        uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:用户 ID" json:"id"`
	RealName  string    `gorm:"column:real_name;type:varchar(64);comment:用户真实姓名;NOT NULL" json:"real_name"`
	NickName  string    `gorm:"column:nick_name;type:varchar(64);comment:用户昵称;NOT NULL" json:"nick_name"`
	PhoneNum  string    `gorm:"column:phone_num;type:char(20);comment:用户手机;NOT NULL" json:"phone_num"`
	BirthDate time.Time `gorm:"column:birth_date;type:date;comment:用户生日;NOT NULL" json:"birth_date"`
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_at"`
	IsDeleted string    `gorm:"column:is_deleted;type:char(1);default:0;comment:是否已删除" json:"is_deleted"`
}

func (m *UserInfoTbl) TableName() string {
	return "user_info_tbl"
}
