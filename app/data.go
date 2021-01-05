package app

type doctor struct {
	name         string
	id           int64
	locations    []string
	appointments map[string]map[string]map[int64]bool // doctor->location->day->hour->isAvailable
	schedule     []*appointment
}

type appointment struct {
	id         int64
	doctorID   int64
	locationID string
	time       appointmentTime
}

type appointmentTime struct {
	// TODO: Turn these into an enum perhaps
	day  string // accepted values [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	hour int64  // accepted values [8 am - 6 pm]
}

// A collection of doctors
var doctorMap = make(map[int64]*doctor)

func createAppointment(doctorID int64, locationID string, time appointmentTime) (bool, string) {
	doctor, ok := doctorMap[doctorID]

	if !ok {
		return false, "The doctor does not exist."
	}

	locationFound := false
	for _, location := range doctor.locations {
		if location == locationID {
			locationFound = true
		}
	}

	if !locationFound {
		return false, "The doctor does not practice in this location."
	}

	timeOk, msg := validateTime(time)

	if !timeOk {
		return false, msg
	}

	for _, location := range doctor.locations {
		appointment := getAppointment(doctorID, location, time)
		if appointment != nil {
			return false, "This doctor already has an appointment at this time."
		}
	}

	if doctor.appointments[locationID][time.day] == nil {
		return false, "The doctor is not available at this time."
	}

	doctor.appointments[locationID][time.day][time.hour] = false

	return true, ""
}

func validateTime(time appointmentTime) (bool, string) {
	if time.day != "Monday" && time.day != "Tuesday" && time.day != "Wednesday" && time.day != "Thursday" && time.day != "Friday" && time.day != "Saturday" && time.day != "Sunday" {
		return false, "Incorrect day specified."
	}

	if time.hour < 1 || time.hour > 11 || (time.hour >= 5 && time.hour <= 7) { // noon hour not included
		return false, "Incorrect time specified."
	}

	return true, ""
}

func getAppointment(doctorID int64, locationID string, time appointmentTime) *appointment {
	doctor, ok := doctorMap[doctorID]

	if !ok {
		return nil
	}

	if doctor.appointments[locationID] != nil && doctor.appointments[locationID][time.day] != nil && doctor.appointments[locationID][time.day][time.hour] == false {
		appointment := appointment{doctorID: doctorID, locationID: locationID, time: appointmentTime{day: time.day, hour: time.hour}}
		return &appointment
	}

	return nil
}

func getAppointments(doctorID int64) (bool, []*appointment) {
	doctor, ok := doctorMap[doctorID]

	if !ok {
		return false, nil
	}

	return true, getDoctorAppointments(doctor)
}

func getDoctorAppointments(doctor *doctor) []*appointment {
	var appointments []*appointment

	for _, location := range doctor.locations {
		mapDayHour := doctor.appointments[location]
		for day, mapHour := range mapDayHour {
			for hour, available := range mapHour {
				if !available {
					appointment := appointment{doctorID: doctor.id, locationID: location, time: appointmentTime{day: day, hour: hour}}
					appointments = append(appointments, &appointment)
				}
			}
		}
	}

	return appointments
}

func deleteAppointment(doctorID int64, locationID string, time appointmentTime) (bool, string) {
	doctor, ok := doctorMap[doctorID]

	if !ok {
		return false, "The doctor does not exist."
	}

	locationFound := false
	for _, location := range doctor.locations {
		if location == locationID {
			locationFound = true
		}
	}

	if !locationFound {
		return false, "The doctor does not practice in this location."
	}

	timeOk, msg := validateTime(time)

	if !timeOk {
		return false, msg
	}

	if doctor.appointments[locationID][time.day] == nil || doctor.appointments[locationID][time.day][time.hour] == true {
		return false, "This appointment does not exist."
	}

	// delete the appointment
	doctor.appointments[locationID][time.day][time.hour] = true

	return true, ""
}

func createDoctor(id int64, name string, locations []string, schedule []*appointment) (bool, string) {
	_, ok := doctorMap[id]

	if ok {
		return false, "The doctor already exists, perhaps you meant to update instead."
	}

	doctor := doctor{name: name, id: id, locations: locations}
	doctor.appointments = make(map[string]map[string]map[int64]bool)
	doctor.schedule = schedule

	for _, location := range locations {
		doctor.appointments[location] = make(map[string]map[int64]bool)
	}

	for _, appointment := range schedule {
		if doctor.appointments[appointment.locationID][appointment.time.day] == nil {
			doctor.appointments[appointment.locationID][appointment.time.day] = make(map[int64]bool)
		}

		// Indicate this time is available
		doctor.appointments[appointment.locationID][appointment.time.day][appointment.time.hour] = true
	}

	doctorMap[id] = &doctor
	return true, ""
}

func getDoctor(id int64) (bool, *doctor) {
	doctor, ok := doctorMap[id]

	if !ok {
		return false, nil
	}

	return true, doctor
}

func deleteDoctor(id int64) bool {
	_, ok := doctorMap[id]

	if !ok {
		return false
	}

	delete(doctorMap, id)
	return true
}

// TODO: Add ability to retain existing appointments
// Semantics today are that all existing appointments are cancelled when a doctor is updated
func updateDoctor(id int64, name string, locations []string, schedule []*appointment) bool {
	_, ok := doctorMap[id]

	if !ok {
		return false
	}

	doctor := doctor{name: name, id: id, locations: locations}
	doctor.appointments = make(map[string]map[string]map[int64]bool)
	doctor.schedule = schedule

	for _, location := range locations {
		doctor.appointments[location] = make(map[string]map[int64]bool)
	}

	// existingAppointments := getDoctorAppointments(existingDoctor)

	for _, appointment := range schedule {
		if doctor.appointments[appointment.locationID][appointment.time.day] == nil {
			doctor.appointments[appointment.locationID][appointment.time.day] = make(map[int64]bool)
		}

		// Indicate this time is available
		doctor.appointments[appointment.locationID][appointment.time.day][appointment.time.hour] = true

		// TODO: if existing appointment, set it here
	}

	doctorMap[id] = &doctor
	return true
}
