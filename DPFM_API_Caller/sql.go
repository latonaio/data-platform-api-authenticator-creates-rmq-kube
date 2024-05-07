package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-authenticator-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-authenticator-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-authenticator-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-authenticator-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var user *dpfm_api_output_formatter.User
	var sMSAuth *dpfm_api_output_formatter.SMSAuth
	var googleAccountAuth *dpfm_api_output_formatter.GoogleAccountAuth
	var instagramAuth *dpfm_api_output_formatter.InstagramAuth
	for _, fn := range accepter {
		switch fn {
		case "User":
			user = c.userCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "SMSAuth":
			sMSAuth = c.sMSAuthCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "GoogleAccountAuth":
			googleAccountAuth = c.googleAccountAuthCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "InstagramAuth":
			instagramAuth = c.instagramAuthCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		User:              user,
		SMSAuth:           sMSAuth,
		GoogleAccountAuth: googleAccountAuth,
		InstagramAuth:     instagramAuth,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var user *dpfm_api_output_formatter.User
	var sMSAuth *dpfm_api_output_formatter.SMSAuth
	var googleAccountAuth *dpfm_api_output_formatter.GoogleAccountAuth
	var instagramAuth *dpfm_api_output_formatter.InstagramAuth
	for _, fn := range accepter {
		switch fn {
		case "User":
			user = c.userUpdateSql(mtx, input, output, errs, log)
		case "SMSAuth":
			sMSAuth = c.sMSAuthUpdateSql(mtx, input, output, errs, log)
		case "GoogleAccountAuth":
			googleAccountAuth = c.googleAccountAuthUpdateSql(mtx, input, output, errs, log)
		case "InstagramAuth":
			instagramAuth = c.instagramAuthUpdateSql(mtx, input, output, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		User:              user,
		SMSAuth:           sMSAuth,
		GoogleAccountAuth: googleAccountAuth,
		InstagramAuth:     instagramAuth,
	}

	return data
}

func (c *DPFMAPICaller) userCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.User {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID

	dpfm_api_output_formatter.ConvertToUser(input, subfuncSDC)

	userData := subfuncSDC.Message.User
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": userData, "function": "AuthenticatorUser", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "User Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToUserCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) sMSAuthCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.SMSAuth {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID

	dpfm_api_output_formatter.ConvertToSMSAuth(input, subfuncSDC)

	sMSAuthData := subfuncSDC.Message.SMSAuth
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": sMSAuthData, "function": "AuthenticatorSMSAuth", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "SMSAuth Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToSMSAuthCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) googleAccountAuthCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.GoogleAccountAuth {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID

	dpfm_api_output_formatter.ConvertToGoogleAccountAuth(input, subfuncSDC)

	googleAccountAuthData := subfuncSDC.Message.GoogleAccountAuth
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": googleAccountAuthData, "function": "AuthenticatorGoogleAccountAuth", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "GoogleAccountAuth Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGoogleAccountAuthCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) instagramAuthCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.InstagramAuth {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID

	dpfm_api_output_formatter.ConvertToInstagramAuth(input, subfuncSDC)

	instagramAuthData := subfuncSDC.Message.InstagramAuth
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": instagramAuthData, "function": "AuthenticatorInstagramAuth", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "InstagramAuth Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToInstagramAuthCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) userUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.User {
	user := input.User
	userData := dpfm_api_processing_formatter.ConvertToUserUpdates(user)

	sessionID := input.RuntimeSessionID
	if userIsUpdate(userData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": userData, "function": "AuthenticatorUser", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "User Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToUserUpdates(user)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) sMSAuthUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.SMSAuth {
	sMSAuth := input.User.SMSAuth
	sMSAuthData := dpfm_api_processing_formatter.ConvertToSMSAuthUpdates(sMSAuth)

	sessionID := input.RuntimeSessionID
	if sMSAuthIsUpdate(sMSAuthData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": sMSAuthData, "function": "AuthenticatorSMSAuth", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "SMSAuth Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToSMSAuthUpdates(sMSAuth)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) googleAccountAuthUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.GoogleAccountAuth {
	googleAccountAuth := input.User.GoogleAccountAuth
	googleAccountAuthData := dpfm_api_processing_formatter.ConvertToGoogleAccountAuthUpdates(googleAccountAuth)

	sessionID := input.RuntimeSessionID
	if googleAccountAuthIsUpdate(googleAccountAuthData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": googleAccountAuthData, "function": "AuthenticatorGoogleAccountAuth", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "GoogleAccountAuth Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGoogleAccountAuthUpdates(googleAccountAuth)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) instagramAuthUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.InstagramAuth {
	instagramAuth := input.User.InstagramAuth
	instagramAuthData := dpfm_api_processing_formatter.ConvertToInstagramAuthUpdates(instagramAuth)

	sessionID := input.RuntimeSessionID
	if instagramAuthIsUpdate(instagramAuthData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": instagramAuthData, "function": "AuthenticatorInstagramAuth", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "InstagramAuth Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToInstagramAuthUpdates(instagramAuth)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func userIsUpdate(user *dpfm_api_processing_formatter.UserUpdates) bool {
	userID := user.UserID

	return !(userID == "")
}

func sMSAuthIsUpdate(sMSAuth *dpfm_api_processing_formatter.SMSAuthUpdates) bool {
	userID := sMSAuth.UserID

	return !(userID == "")
}

func googleAccountAuthIsUpdate(googleAccountAuth *dpfm_api_processing_formatter.GoogleAccountAuthUpdates) bool {
	userID := googleAccountAuth.UserID

	return !(userID == "")
}

func instagramAuthIsUpdate(instagramAuth *dpfm_api_processing_formatter.InstagramAuthUpdates) bool {
	userID := instagramAuth.UserID

	return !(userID == "")
}
