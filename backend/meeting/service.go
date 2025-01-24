package meeting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hello-world/access_token"
	"hello-world/my_http"
	"net/http"
)

func UpdateMeeting(value UpdateMeetingRequest) error {

	body := my_http.StringToBuffer(fmt.Sprintf(`
	{
		"schedule_for": "harits.abdullah19@gmail.com",
		"agenda": "%s",
		"duration": 60,
		"password": "123456",
		"pre_schedule": false,
		"settings": {
			"value": "<Error: Too many levels of nesting to fake this schema>"
		},
		"start_time": "2022-03-25T07:29:29Z",
		"template_id": "5Cj3ceXoStO6TGOVvIOVPA==",
		"timezone": "America/Los_Angeles",
		"topic": "My Meeting",
		"type": 2
	}
	`, value.Agenda))

	_, err := my_http.Request("PATCH", "https://api.zoom.us/v2/meetings/"+value.MeetingId, body, http.StatusNoContent)

	if err != nil {

		return err
	}

	return nil

}

func GetAllMeeting() (*GetAllMeetingZoomApiResponse, error) {

	response, err := my_http.Request("GET", "https://api.zoom.us/v2/users/me/meetings?page_size=30&page_number=1", nil, http.StatusOK)

	if err != nil {
		return nil, err
	}

	post := &GetAllMeetingZoomApiResponse{}
	err = json.NewDecoder(response.Body).Decode(post)

	if err != nil {

		return nil, err
	}
	return post, nil
}

func CreateMeeting(value CreateMeetingRequest) (err error) {

	fmt.Println("CreateMeeting value", value.Agenda)

	// JSON body
	body := []byte(fmt.Sprintf(`
	{
    "agenda": "%s",
    "default_password": false,
    "duration": 60,
    "password": "123456",
    "pre_schedule": false,
    "recurrence": {
        "end_date_time": "2022-04-02T15:59:00Z",
        "end_times": 7,
        "monthly_day": 1,
        "monthly_week": 1,
        "monthly_week_day": 1,
        "repeat_interval": 1,
        "type": 1,
        "weekly_days": "1"
    },
    "schedule_for": "harits.abdullah19@gmail.com",
    "settings": {
        "additional_data_center_regions": [
            "TY"
        ],
        "allow_multiple_devices": true,
        "alternative_hosts_email_notification": true,
        "approval_type": 2,
        "approved_or_denied_countries_or_regions": {
            "approved_list": [
                "CX"
            ],
            "denied_list": [
                "CA"
            ],
            "enable": true,
            "method": "approve"
        },
        "audio": "telephony",
        "audio_conference_info": "test",
        "authentication_domains": "example.com",
        "authentication_exception": [
            {
                "email": "harits.abdullah19@gmail.com",
                "name": "Jill Chill"
            }
        ],
        "auto_recording": "cloud",
        "breakout_room": {
            "enable": true,
            "rooms": [
                {
                    "name": "room1",
                    "participants": [
                        "harits.abdullah19@gmail.com"
                    ]
                }
            ]
        },
        "calendar_type": 1,
        "close_registration": false,
        "contact_email": "harits.abdullah19@gmail.com",
        "contact_name": "Jill Chill",
        "email_notification": true,
        "encryption_type": "enhanced_encryption",
        "focus_mode": true,
        "global_dial_in_countries": [
        ],
        "host_video": true,
        "jbh_time": 0,
        "join_before_host": false,
        "language_interpretation": {
            "enable": true,
            "interpreters": [
                {
                    "email": "interpreter@example.com",
                    "languages": "ID,FR"
                }
            ]
        },
        "meeting_authentication": true,
        "meeting_invitees": [
            {
                "email": "harits.abdullah19@gmail.com"
            }
        ],
        "mute_upon_entry": false,
        "participant_video": false,
        "private_meeting": false,
        "registrants_confirmation_email": true,
        "registrants_email_notification": true,
        "registration_type": 1,
        "show_share_button": true,
        "use_pmi": false,
        "waiting_room": false,
        "watermark": false,
        "host_save_video_order": true,
        "alternative_host_update_polls": true
    },
    "start_time": "2022-03-25T07:32:55Z",
    "template_id": "Dv4YdINdTk+Z5RToadh5ug==",
    "timezone": "America/Los_Angeles",
    "topic": "My Meeting",
    "tracking_fields": [
        {
            "field": "field1",
            "value": "value1"
        }
    ],
    "type": 2
}`, value.Agenda))

	client := &http.Client{}

	request, err := http.NewRequest("POST", "https://api.zoom.us/v2/users/me/meetings", bytes.NewBuffer(body))

	if err != nil {

		return err
	}

	access_token.SetAccessToken(request)

	response, err := client.Do(request)

	if err != nil {

		return err
	}

	defer response.Body.Close()

	post := &CreateMeetingZoomApiResponse{}
	err = json.NewDecoder(response.Body).Decode(post)

	if err != nil {

		return err
	}

	fmt.Println("CreateMeeting response ", post.UUID)

	fmt.Println("CreateMeeting status response", response.Status)

	if response.StatusCode != http.StatusCreated {

		return (fmt.Errorf("Failed to create meeting with error code %d", response.StatusCode))
	}

	return nil

}
