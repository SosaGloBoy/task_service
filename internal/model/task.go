package model

import "time"

type Task struct {
	Id          uint          `gorm:"primaryKey;autoIncrement" json:"id"`     // Указание, что это первичный ключ с автоинкрементом
	Title       string        `gorm:"not null" json:"title"`                  // Заголовок задания, обязательное поле
	Description string        `gorm:"type:text" json:"description"`           // Описание задания
	VMImagePath string        `gorm:"type:varchar(255)" json:"vm_image_path"` // Путь к образу виртуальной машины
	CreatedAt   time.Time     `gorm:"autoCreateTime" json:"created_at"`       // Автоматическое создание времени
	UpdatedAt   time.Time     `gorm:"autoUpdateTime" json:"updated_at"`       // Автоматическое обновление времени
	ArticleID   int           `json:"article_id"`
	Duration    time.Duration `json:"duration"`
	Steps       []Step        `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE" json:"steps"` // Связь с шагами (One-to-Many)
}

type Step struct {
	Id                uint   `gorm:"primaryKey;autoIncrement" json:"id"` // Уникальный идентификатор шага
	TaskID            uint   `gorm:"not null" json:"task_id"`            // Внешний ключ для связи с задачей
	Description       string `gorm:"type:text" json:"description"`       // Описание шага
	Command           string `gorm:"type:text" json:"command"`           // Команда для выполнения
	ExpectedOutput    string `gorm:"type:text" json:"expected_output"`
	ValidationCommand string `gorm:"type:text" json:"validation_command"`
	Status            string `gorm:"type:varchar(50)" json:"status"` // Статус выполнения
}
