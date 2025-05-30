package models

import "time"

type Inscripcion struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"` // ID único de la inscripción
	UsuarioID   uint      `gorm:"not null"`                 // ID del usuario que se inscribe
	ActividadID uint      `gorm:"not null"`                 // ID de la actividad a la que se inscribe
	Fecha       time.Time `gorm:"not null"`                 // Fecha de inscripción

	//Relaciones
	Usuario   Usuario   `gorm:"foreignKey:UsuarioID;references:ID"`   // Relación con el usuario que se inscribe
	Actividad Actividad `gorm:"foreignKey:ActividadID;references:ID"` // Relación con la actividad a la que se inscribe
}
