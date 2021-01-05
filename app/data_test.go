package app

import (
	"strconv"
	"testing"
)

// TODO: Add negative tests

func TestDataLayer(t *testing.T) {
	appointment1 := appointment{doctorID: 1, locationID: "address1", time: appointmentTime{day: "Monday", hour: 10}}
	appointment2 := appointment{doctorID: 1, locationID: "address2", time: appointmentTime{day: "Wednesday", hour: 2}}
	testDoctor := &doctor{id: 1, name: "Jow Smith", locations: []string{"address1", "address2"}, schedule: []*appointment{&appointment1, &appointment2}}

	// test create doctor
	ok, msg := createDoctor(testDoctor.id, testDoctor.name, testDoctor.locations, testDoctor.schedule)

	if !ok || msg != "" {
		t.Fatalf("Unexpected error: " + msg + "\n")
	}

	if len(doctorMap) != 1 {
		t.Fatalf("Unexpected number of doctors: " + strconv.Itoa(len(doctorMap)) + "\n")
	}

	// test get doctor
	ok, getDoctor := getDoctor(testDoctor.id)

	if !ok || (getDoctor == nil) {
		t.Fatalf("Could not retrieve doctor" + "\n")
	}

	if !compareDoctor(*testDoctor, *getDoctor) {
		t.Fatalf("Incorrect doctor information" + "\n")
	}

	if len(getDoctorAppointments(getDoctor)) != 0 {
		t.Fatalf("Incorrect appointments" + "\n")
	}

	// test createAppointment
	ok, msg = createAppointment(appointment1.doctorID, appointment1.locationID, appointment1.time)
	if !ok || msg != "" {
		t.Fatalf("Error creating appointment: " + msg + "\n")
	}

	ok, appointments := getAppointments(appointment1.doctorID)
	if !ok {
		t.Fatalf("Error getting appointments." + "\n")
	}

	if len(appointments) != 1 {
		t.Fatalf("Wrong number of appointments." + "\n")
	}

	if !compareAppointments(appointment1, *appointments[0]) {
		t.Fatalf("Appointments dont match." + "\n")
	}

	// test deleteAppointment
	ok, msg = deleteAppointment(appointment1.doctorID, appointment1.locationID, appointment1.time)
	if !ok || msg != "" {
		t.Fatalf("Error deleting appointment: " + msg + "\n")
	}

	ok, appointments = getAppointments(appointment1.doctorID)
	if !ok {
		t.Fatalf("Error getting appointments." + "\n")
	}

	if len(appointments) != 0 {
		t.Fatalf("Wrong number of appointments." + "\n")
	}

	// test delete doctor
	ok = deleteDoctor(testDoctor.id)

	if !ok {
		t.Fatalf("Unexpected error while deleting doctor." + "\n")
	}

	if len(doctorMap) != 0 {
		t.Fatalf("Unexpected number of doctors: " + strconv.Itoa(len(doctorMap)) + "\n")
	}
}

func TestCreateMultipleAppointments(t *testing.T) {
	appointment11 := appointment{doctorID: 1, locationID: "address1", time: appointmentTime{day: "Thursday", hour: 3}}
	appointment12 := appointment{doctorID: 1, locationID: "address2", time: appointmentTime{day: "Monday", hour: 3}}
	appointment13 := appointment{doctorID: 1, locationID: "address1", time: appointmentTime{day: "Monday", hour: 10}}
	appointment14 := appointment{doctorID: 1, locationID: "address2", time: appointmentTime{day: "Monday", hour: 10}}
	appointment15 := appointment{doctorID: 1, locationID: "address1", time: appointmentTime{day: "Monday", hour: 2}}
	appointment21 := appointment{doctorID: 2, locationID: "address3", time: appointmentTime{day: "Monday", hour: 10}}
	appointment22 := appointment{doctorID: 2, locationID: "address4", time: appointmentTime{day: "Monday", hour: 10}}
	testDoctor1 := &doctor{id: 1, name: "Jon Smith", locations: []string{"address1", "address2"}, schedule: []*appointment{&appointment11, &appointment12}}
	testDoctor2 := &doctor{id: 2, name: "Mike Baron", locations: []string{"address3", "address4"}, schedule: []*appointment{&appointment21, &appointment22}}

	// test create doctor
	ok, msg := createDoctor(testDoctor1.id, testDoctor1.name, testDoctor1.locations, testDoctor1.schedule)
	ok, msg = createDoctor(testDoctor2.id, testDoctor2.name, testDoctor2.locations, testDoctor2.schedule)

	ok, msg = createAppointment(appointment11.doctorID, appointment11.locationID, appointment11.time)
	if !ok || msg != "" {
		t.Fatalf("Error creating appointment11: " + msg + "\n")
	}

	ok, msg = createAppointment(appointment12.doctorID, appointment12.locationID, appointment12.time)
	if !ok || msg != "" {
		t.Fatalf("Error creating appointment12: " + msg + "\n")
	}

	ok, msg = createAppointment(appointment12.doctorID, appointment12.locationID, appointment12.time)
	if ok {
		t.Fatalf("Expected Error creating appointment12: " + msg + "\n")
	}

	ok, msg = createAppointment(appointment13.doctorID, appointment13.locationID, appointment13.time)
	if ok {
		t.Fatalf("Expected Error creating appointment13: " + msg + "\n")
	}

	ok, msg = createAppointment(appointment14.doctorID, appointment14.locationID, appointment14.time)
	if ok {
		t.Fatalf("Expected Error creating appointment14: " + msg + "\n")
	}

	ok, msg = createAppointment(appointment15.doctorID, appointment15.locationID, appointment15.time)
	if ok {
		t.Fatalf("Expected Error creating appointment15: " + msg + "\n")
	}

	ok, msg = createAppointment(appointment21.doctorID, appointment21.locationID, appointment21.time)
	if !ok || msg != "" {
		t.Fatalf("Error creating appointment21: " + msg + "\n")
	}

	ok, msg = createAppointment(appointment22.doctorID, appointment22.locationID, appointment22.time)
	if ok {
		t.Fatalf("Error creating appointment22: " + msg + "\n")
	}
}

func compareDoctor(doctor1 doctor, doctor2 doctor) bool {
	if doctor1.id != doctor2.id {
		return false
	}

	if doctor1.name != doctor2.name {
		return false
	}

	if len(doctor1.locations) != len(doctor2.locations) {
		return false
	}

	for _, loc1 := range doctor1.locations {
		var found = false
		for _, loc2 := range doctor2.locations {
			if loc2 == loc1 {
				found = true
			}
		}
		if !found {
			return false
		}
	}

	if len(doctor1.schedule) != len(doctor2.schedule) {
		return false
	}

	for _, app1 := range doctor1.schedule {
		var found = false
		for _, app2 := range doctor2.schedule {
			if compareAppointments(*app1, *app2) {
				found = true
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func compareAppointments(app1 appointment, app2 appointment) bool {
	if app1.doctorID == app2.doctorID && app1.locationID == app2.locationID && app1.time.day == app2.time.day && app1.time.hour == app2.time.hour {
		return true
	} else {
		return false
	}
}
