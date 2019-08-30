package copilot

// Contact is a representation of the Autopilot API Contact data
type Contact struct {
	ID                 string `json:"contact_id,omitempty"`
	Email              string `json:"Email,omitempty"`
	FirstName          string `json:"FirstName,omitempty"`
	Twitter            string `json:"Twitter,omitempty"`
	LastName           string `json:"LastName,omitempty"`
	Salutation         string `json:"Salutation,omitempty"`
	Company            string `json:"Company,omitempty"`
	NumberOfEmployees  string `json:"NumberOfEmployees,omitempty"`
	Title              string `json:"Title,omitempty"`
	Industry           string `json:"Industry,omitempty"`
	Phone              string `json:"Phone,omitempty"`
	MobilePhone        string `json:"MobilePhone,omitempty"`
	Fax                string `json:"Fax,omitempty"`
	Website            string `json:"Website,omitempty"`
	MailingStreet      string `json:"MailingStreet,omitempty"`
	MailingCity        string `json:"MailingCity,omitempty"`
	MailingState       string `json:"MailingState,omitempty"`
	MailingPostalCode  string `json:"MailingPostalCode,omitempty"`
	MailingCountry     string `json:"MailingCountry,omitempty"`
	OwnerName          string `json:"owner_name,omitempty"`
	LeadSource         string `json:"LeadSource,omitempty"`
	Status             string `json:"Status,omitempty"`
	LinkedIn           string `json:"LinkedIn,omitempty"`
	Unsubscribed       string `json:"unsubscribed,omitempty"`
	Custom             string `json:"custom,omitempty"`
	AutopilotSessionID string `json:"_autopilot_session_id,omitempty"`
	AutopilotList      string `json:"_autopilot_list,omitempty"`
	Notify             string `json:"notify,omitempty"`
}
