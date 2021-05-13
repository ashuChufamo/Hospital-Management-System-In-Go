package Doctor

import "github.com/web1_group_project/hospital_server/entity"

// UserService specifies application user related services
type PatientService interface {
	Patientes() ([]entity.Petient, []error)
	Patient(id uint) (*entity.Petient, []error)
	UpdatePatient(user *entity.Petient) (*entity.Petient, []error)
	DeletePatient(id uint) (*entity.Petient, []error)
	StorePatient(user *entity.Petient) (*entity.Petient, []error)
}

// RoleService speifies application user role related services

type AppointmentService interface {
	Appointments() ([]entity.Doctor, []error)
	Appointment(id uint) (*entity.Doctor, []error)
	UpdateAppointment(user *entity.Doctor) (*entity.Doctor, []error)
	AppUpdateAppointment(appointment *entity.Appointment) (*entity.Appointment, []error)
	UpdatePrescription(prescription *entity.Prescription) (*entity.Prescription, []error)

	DeleteAppointment(id uint) (*entity.Doctor, []error)
	AppAppointment(id uint) (*entity.Appointment, []error)
	Prescribtion(id uint) (*entity.Prescription, []error)
}
