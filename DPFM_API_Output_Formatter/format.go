package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-authenticator-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-authenticator-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToUserCreates(subfuncSDC *sub_func_complementer.SDC) (*User, error) {
	data := subfuncSDC.Message.User

	user, err := TypeConverter[*User](data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func ConvertToSMSAuthCreates(subfuncSDC *sub_func_complementer.SDC) (*SMSAuth, error) {
	data := subfuncSDC.Message.SMSAuth

	sMSAuth, err := TypeConverter[*SMSAuth](data)
	if err != nil {
		return nil, err
	}

	return sMSAuth, nil
}

func ConvertToGoogleAccountAuthCreates(subfuncSDC *sub_func_complementer.SDC) (*GoogleAccountAuth, error) {
	data := subfuncSDC.Message.GoogleAccountAuth

	googleAccountAuth, err := TypeConverter[*GoogleAccountAuth](data)
	if err != nil {
		return nil, err
	}

	return googleAccountAuth, nil
}

func ConvertToInstagramAuthCreates(subfuncSDC *sub_func_complementer.SDC) (*InstagramAuth, error) {
	data := subfuncSDC.Message.InstagramAuth

	instagramAuth, err := TypeConverter[*InstagramAuth](data)
	if err != nil {
		return nil, err
	}

	return instagramAuth, nil
}

func ConvertToUserUpdates(userData dpfm_api_input_reader.User) (*User, error) {
	data := userData

	user, err := TypeConverter[*User](data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func ConvertToSMSAuthUpdates(sMSAuthData dpfm_api_input_reader.SMSAuth) (*SMSAuth, error) {
	data := sMSAuthData

	sMSAuth, err := TypeConverter[*SMSAuth](data)
	if err != nil {
		return nil, err
	}

	return sMSAuth, nil
}

func ConvertToGoogleAccountAuthUpdates(googleAccountAuthData dpfm_api_input_reader.GoogleAccountAuth) (*GoogleAccountAuth, error) {
	data := googleAccountAuthData

	googleAccountAuth, err := TypeConverter[*GoogleAccountAuth](data)
	if err != nil {
		return nil, err
	}

	return googleAccountAuth, nil
}

func ConvertToInstagramAuthUpdates(instagramAuthData dpfm_api_input_reader.InstagramAuth) (*InstagramAuth, error) {
	data := instagramAuthData

	instagramAuth, err := TypeConverter[*InstagramAuth](data)
	if err != nil {
		return nil, err
	}

	return instagramAuth, nil
}

func ConvertToUser(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.User = &sub_func_complementer.User{
		UserID:              input.User.UserID,
		BusinessPartner:     input.User.BusinessPartner,
		Password:            input.User.Password,
		Qos:                 input.User.Qos,
		IsEncrypt:           input.User.IsEncrypt,
		Language:            input.User.Language,
		LastLoginDate:       input.User.LastLoginDate,
		LastLoginTime:       input.User.LastLoginTime,
		CreationDate:        input.User.CreationDate,
		CreationTime:        input.User.CreationTime,
		LastChangeDate:      input.User.LastChangeDate,
		LastChangeTime:      input.User.LastChangeTime,
		IsMarkedForDeletion: input.User.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToSMSAuth(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.SMSAuth = &sub_func_complementer.SMSAuth{
		UserID:              input.User.UserID,
		MobilePhoneNumber:   input.User.SMSAuth.MobilePhoneNumber,
		AuthenticationCode:  input.User.SMSAuth.AuthenticationCode,
		CreationDate:        input.User.SMSAuth.CreationDate,
		CreationTime:        input.User.SMSAuth.CreationTime,
		LastChangeDate:      input.User.SMSAuth.LastChangeDate,
		LastChangeTime:      input.User.SMSAuth.LastChangeTime,
		IsMarkedForDeletion: input.User.SMSAuth.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToGoogleAccountAuth(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.GoogleAccountAuth = &sub_func_complementer.GoogleAccountAuth{
		UserID:              input.User.UserID,
		EmailAddress:        input.User.GoogleAccountAuth.EmailAddress,
		GoogleID:            input.User.GoogleAccountAuth.GoogleID,
		AccessToken:         input.User.GoogleAccountAuth.AccessToken,
		CreationDate:        input.User.GoogleAccountAuth.CreationDate,
		CreationTime:        input.User.GoogleAccountAuth.CreationTime,
		LastChangeDate:      input.User.GoogleAccountAuth.LastChangeDate,
		LastChangeTime:      input.User.GoogleAccountAuth.LastChangeTime,
		IsMarkedForDeletion: input.User.GoogleAccountAuth.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToInstagramAuth(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.InstagramAuth = &sub_func_complementer.InstagramAuth{
		UserID:              input.User.UserID,
		InstagramID:         input.User.InstagramAuth.InstagramID,
		AccessToken:         input.User.InstagramAuth.AccessToken,
		CreationDate:        input.User.InstagramAuth.CreationDate,
		CreationTime:        input.User.InstagramAuth.CreationTime,
		LastChangeDate:      input.User.InstagramAuth.LastChangeDate,
		LastChangeTime:      input.User.InstagramAuth.LastChangeTime,
		IsMarkedForDeletion: input.User.InstagramAuth.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
