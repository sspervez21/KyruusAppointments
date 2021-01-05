package app

import (
	"KyruusAppointments/models"
	"KyruusAppointments/restapi/operations"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAppointment function
func CreateAppointment(params operations.CreateAppointmentParams) middleware.Responder {

	ok, errorMessage := createAppointment(*params.CreateAppointment.DoctorID, *params.CreateAppointment.LocationID, appointmentTime{day: *params.CreateAppointment.Day, hour: *params.CreateAppointment.Time})

	if !ok {
		return operations.NewCreateAppointmentConflict().WithPayload(&models.BadRequest{
			Code:    int64(operations.CreateAppointmentConflictCode),
			Message: errorMessage,
		})
	}

	return operations.NewCreateAppointmentOK()
}

// GetAppointments function
func GetAppointments(params operations.GetAppointmentsParams) middleware.Responder {

	ok, appointments := getAppointments(params.DoctorID)

	if !ok {
		return operations.NewGetAppointmentsNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.GetAppointmentsNotFoundCode),
			Message: "Doctor not found",
		})
	}

	var appointmentsModel = make([]*models.Appointment, len(appointments))

	for _, appointment := range appointments {
		appointmentsModel = append(appointmentsModel, &models.Appointment{DoctorID: &appointment.doctorID, LocationID: &appointment.locationID, Day: &appointment.time.day, Time: &appointment.time.hour})
	}

	return operations.NewGetAppointmentsOK().WithPayload(appointmentsModel)
}

// DeleteAppointment function
func DeleteAppointment(params operations.DeleteAppointmentParams) middleware.Responder {

	ok, errorMessage := deleteAppointment(*params.DeleteAppointment.DoctorID, *params.DeleteAppointment.LocationID, appointmentTime{day: *params.DeleteAppointment.Day, hour: *params.DeleteAppointment.Time})

	if !ok {
		return operations.NewDeleteAppointmentNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.DeleteAppointmentNotFoundCode),
			Message: errorMessage,
		})
	}

	return operations.NewDeleteAppointmentOK()
}

// CreateDoctor function
func CreateDoctor(params operations.CreateDoctorParams) middleware.Responder {

	var schedule []*appointment

	for _, appointmentModel := range params.CreateDoctorInput.Schedule {
		appointment := appointment{doctorID: *appointmentModel.DoctorID, locationID: *appointmentModel.LocationID, time: appointmentTime{day: *appointmentModel.Day, hour: *appointmentModel.Time}}
		schedule = append(schedule, &appointment)
	}

	ok, errorMessage := createDoctor(*params.CreateDoctorInput.ID, *params.CreateDoctorInput.Name, params.CreateDoctorInput.Locations, schedule)

	if !ok {
		return operations.NewCreateDoctorConflict().WithPayload(&models.BadRequest{
			Code:    int64(operations.CreateDoctorConflictCode),
			Message: errorMessage,
		})
	}

	return operations.NewCreateDoctorOK()
}

// GetDoctor function
func GetDoctor(params operations.GetDoctorParams) middleware.Responder {

	ok, doctor := getDoctor(params.DoctorID)

	if !ok {
		return operations.NewGetDoctorNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.GetDoctorNotFoundCode),
			Message: "Doctor with ID " + strconv.FormatInt(params.DoctorID, 10) + " does not exist.",
		})
	}

	var scheduleModel []*models.Appointment

	for _, appointment := range doctor.schedule {
		appointmentModel := models.Appointment{DoctorID: &appointment.doctorID, LocationID: &appointment.locationID, Day: &appointment.time.day, Time: &appointment.time.hour}
		scheduleModel = append(scheduleModel, &appointmentModel)
	}

	doctorModel := models.Doctor{Name: &doctor.name, ID: &doctor.id, Locations: doctor.locations, Schedule: scheduleModel}

	return operations.NewGetDoctorOK().WithPayload(&doctorModel)
}

// DeleteDoctor function
func DeleteDoctor(params operations.DeleteDoctorParams) middleware.Responder {

	ok := deleteDoctor(params.DoctorID)

	if !ok {
		return operations.NewDeleteDoctorNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.DeleteDoctorNotFoundCode),
			Message: "Doctor with ID " + strconv.FormatInt(params.DoctorID, 10) + " does not exist.",
		})
	}

	return operations.NewDeleteDoctorOK()
}

// UpdateDoctor function
func UpdateDoctor(params operations.UpdateDoctorParams) middleware.Responder {

	var schedule []*appointment

	for _, appointmentModel := range params.UpdateDoctorInput.Schedule {
		appointment := appointment{doctorID: *appointmentModel.DoctorID, locationID: *appointmentModel.LocationID, time: appointmentTime{day: *appointmentModel.Day, hour: *appointmentModel.Time}}
		schedule = append(schedule, &appointment)
	}

	ok := updateDoctor(params.DoctorID, *params.UpdateDoctorInput.Name, params.UpdateDoctorInput.Locations, schedule)

	if !ok {
		return operations.NewUpdateDoctorNotFound().WithPayload(&models.BadRequest{
			Code:    int64(operations.UpdateDoctorNotFoundCode),
			Message: "Doctor with ID " + strconv.FormatInt(params.DoctorID, 10) + " does not exist.",
		})
	}

	return operations.NewUpdateDoctorOK()
}
