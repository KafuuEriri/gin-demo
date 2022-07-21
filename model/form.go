package model

type Form struct {
	Id        int64  `json:"id" gorm:"primary_key"`                        // 主键
	Name      string `json:"name" gorm:"type:varchar(50);not null;unique"` //名称
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`             // 创建时间
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`             // 修改时间
}
