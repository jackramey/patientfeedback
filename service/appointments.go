package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"patientfeedback/api"
	. "patientfeedback/domain"
	"patientfeedback/storage"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type appointmentsHandler struct {
	store storage.ResourceStore
}

func (handler appointmentsHandler) GetAppointment(c echo.Context) error {
	appointmentId := c.Param("appointmentId")
	appointment, err := handler.store.GetAppointment(appointmentId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, appointment)
}

func (handler appointmentsHandler) PostFeedbackForAppointment(c echo.Context) error {
	req := c.Request()
	appointmentId := c.Param("appointmentId")
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return errors.Wrap(err, "unable to read request body")
	}

	var request api.CreateFeedbackRequest
	if err := json.Unmarshal(bytes, &request); err != nil {
		return errors.Wrap(err, "unable to unmarshal create feedback request")
	}

	feedback, err := handler.CreateFeedbackForAppointment(appointmentId, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, feedback)
}

// TODO separate business logic from handler
func (handler appointmentsHandler) CreateFeedbackForAppointment(appointmentId string, request api.CreateFeedbackRequest) (*Feedback, error) {
	appointment, err := handler.store.GetAppointment(appointmentId)
	if err != nil {
		return nil, err
	}

	if appointment == nil {
		return nil, errors.New("cannot create feedback for an appointment that does not exist")
	}

	feedback := Feedback{
		ResourceHeader: ResourceHeader{
			ID:   uuid.New().String(),
			Type: FeedbackResType,
		},
		Rating:     request.Rating,
		Understood: request.Understood,
		Comment:    request.Comment,
		Appointment: Reference{
			Type: AppointmentResType,
			ID:   appointmentId,
		},
	}

	// TODO transaction support would be nice here
	if err := handler.store.WriteAppointment(*appointment); err != nil {
		return nil, err
	}

	if err := handler.store.WriteFeedback(feedback); err != nil {
		return nil, err
	}

	return &feedback, nil
}
