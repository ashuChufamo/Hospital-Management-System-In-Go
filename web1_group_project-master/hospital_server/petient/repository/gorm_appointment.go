package PetientRepository

import (
	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// AppointmentGormRepo Implements the menu.AppointmentRepository interface
type AppointmentGormRepo struct {
	conn *gorm.DB
}

// NewAppointmentGormRepo creates a new object of AppointmentGormRepo
func NewAppointmentGormRepo(db *gorm.DB) petient.AppointmentRepository {
	return &AppointmentGormRepo{conn: db}
}

// Appointments return all appointments from the database
func (appointmentRepo *AppointmentGormRepo) Appointments() ([]entity.Appointment, []error) {
	appointments := []entity.Appointment{}
	errs := appointmentRepo.conn.Find(&appointments).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return appointments, errs
}

// Appointment retrieves a appointment by its id from the database
func (appointmentRepo *AppointmentGormRepo) Appointment(id uint) (*entity.Appointment, []error) {
	appointment := entity.Appointment{}

	errs := appointmentRepo.conn.Debug().Find(&appointment, "id=?", id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &appointment, errs
}
