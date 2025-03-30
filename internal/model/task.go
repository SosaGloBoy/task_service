package model

import "time"

type Task struct {
	Id          uint      `gorm:"primaryKey;autoIncrement" json:"id"`     // Указание, что это первичный ключ с автоинкрементом
	Title       string    `gorm:"not null" json:"title"`                  // Заголовок задания, обязательное поле
	Description string    `gorm:"type:text" json:"description"`           // Описание задания
	VMImagePath string    `gorm:"type:varchar(255)" json:"vm_image_path"` // Путь к образу виртуальной машины
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`       // Автоматическое создание времени
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`       // Автоматическое обновление времени
}
