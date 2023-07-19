package models

type Category struct {
	Model

	Name string `gorm:"type:character varying;not null;unique"`
}

func (c Category) TableName() string {
	return "categories"
}
