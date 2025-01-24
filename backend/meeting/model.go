package meeting

import "time"

type CreateMeetingRequest struct {
	Agenda string `json:"agenda" binding:"required"`
}

type UpdateMeetingRequest struct {
	MeetingId string `json:"meeting_id" binding:"required"`
	Agenda string `json:"agenda" binding:"required"`
}


type CreateMeetingZoomApiResponse struct {
	UUID string `json:"uuid"`
}

type GetAllMeetingZoomApiResponse struct {
	PageCount    int `json:"page_count"`
	PageNumber   int `json:"page_number"`
	PageSize     int `json:"page_size"`
	TotalRecords int `json:"total_records"`
	Meetings     []struct {
		UUID          string    `json:"uuid"`
		ID            int64     `json:"id"`
		HostID        string    `json:"host_id"`
		Topic         string    `json:"topic"`
		Type          int       `json:"type"`
		StartTime     time.Time `json:"start_time"`
		Duration      int       `json:"duration"`
		Timezone      string    `json:"timezone"`
		Agenda        string    `json:"agenda"`
		CreatedAt     time.Time `json:"created_at"`
		JoinURL       string    `json:"join_url"`
		SupportGoLive bool      `json:"supportGoLive"`
	} `json:"meetings"`
}
