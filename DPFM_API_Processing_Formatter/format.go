package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-authenticator-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToUserUpdates(user dpfm_api_input_reader.User) *UserUpdates {
	data := user

	return &UserUpdates{
			UserID:					data.UserID,
			Password:				data.Password,
			LastLoginDate:			data.LastLoginDate,
			LastLoginTime:			data.LastLoginTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
	}
}

func ConvertToSMSAuthUpdates(sMSAuth dpfm_api_input_reader.SMSAuth) *SMSAuthUpdates {
	data := sMSAuth

	return &SMSAuthUpdates{
			UserID:					data.UserID,
			AuthenticationCode:		data.AuthenticationCode,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
	}
}

func ConvertToGoogleAccountAuthUpdates(googleAccountAuth dpfm_api_input_reader.GoogleAccountAuth) *GoogleAccountAuthUpdates {
	data := googleAccountAuth

	return &GoogleAccountAuthUpdates{
			UserID:					data.UserID,
			AccessToken:			data.AccessToken,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
	}
}

func ConvertToInstagramAuthUpdates(instagramAuth dpfm_api_input_reader.InstagramAuth) *InstagramAuthUpdates {
	data := instagramAuth

	return &InstagramAuthUpdates{
			UserID:					data.UserID,
			InstagramID:			data.InstagramID,
			AccessToken:			data.AccessToken,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
	}
}
