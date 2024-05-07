package dpfm_api_processing_formatter

type UserUpdates struct {
	UserID				string	`json:"UserID"`
	Password			string	`json:"Password"`
	LastLoginDate		*string	`json:"LastLoginDate"`
	LastLoginTime		*string	`json:"LastLoginTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
}

type SMSAuthUpdates struct {
	UserID				string	`json:"UserID"`
	AuthenticationCode	int		`json:"AuthenticationCode"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
}

type GoogleAccountAuthUpdates struct {
	UserID				string	`json:"UserID"`
	AccessToken			string	`json:"AccessToken"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
}

type InstagramAuthUpdates struct {
	UserID				string	`json:"UserID"`
	InstagramID			string	`json:"InstagramID"`
	AccessToken			string	`json:"AccessToken"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
}
