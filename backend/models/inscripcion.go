package models

import "time"

type Usuario struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"` // Cambiado de UsuarioID a ID para mayor consistencia
	UsuarioID   uint      `gorm:"primaryKey;autoIncrement"`
	ActividadID uint      `gorm:"not null"` // ID de la actividad a la que se inscribe
	Fecha       time.Time `gorm:"not null"` // Fecha de inscripción
	//Relaciones
	Usuario   Usuario   `gorm:"foreignKey:UsuarioID;references:ID"`   // Relación con el usuario que se inscribe
	Actividad Actividad `gorm:"foreignKey:ActividadID;references:ID"` // Relación con la actividad a la que se inscribe
}
