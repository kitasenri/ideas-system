package models

import (
	"regexp"
	"time"

	"github.com/wcl48/valval"
)

var VALIDATE_KEY = valval.String(
	valval.MaxLength(256),
	valval.Regexp(regexp.MustCompile(`^[0-9a-z ]+$`)),
)

var VALIDATE_NAME = valval.String(
	valval.MaxLength(256),
	valval.Regexp(regexp.MustCompile(`^[0-9a-z ]+$`)),
)

type CommonPartsHeader struct {
	Id   string `sql:"type:varchar(255);" gorm:"primary_key;"`
	Name string `sql:"type:varchar(255);" gorm:"not null;"`
}

type CommonPartsMessage1 struct {
	Message1 string `sql:"type:mediumtext;" gorm:"not null;"`
	Message2 string `sql:"type:mediumtext;" gorm:"not null;"`
}

type CommonPartsMessage2 struct {
	Message3 string `sql:"type:mediumtext;" gorm:"not null;"`
	Message4 string `sql:"type:mediumtext;" gorm:"not null;"`
	Message5 string `sql:"type:mediumtext;" gorm:"not null;"`
}

type CommonPartsFooter struct {
	Version   int       `gorm:"not null;"`
	Deleted   bool      `sql:"DEFAULT:false;" gorm:"not null;"`
	UpdatedBy string    `sql:"type:varchar(255);"`
	UpdateAt  time.Time `gorm:"type:timestamp;"`
	CreatedBy string    `sql:"type:varchar(255);"`
	CreateAt  time.Time `sql:"DEFAULT:current_timestamp;"`
}
