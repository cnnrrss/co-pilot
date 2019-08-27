package copilot

// Contact is a representation of the Autopilot API Contact data
type Contact struct {
	contact_id            string `json:"contact_id"`
	Email                 string `json:"Email"`
	FirstName             string `json:"FirstName"`
	Twitter               string `json:"Twitter"`
	LastName              string `json:"LastName"`
	Salutation            string `json:"Salutation"`
	Company               string `json:"Company"`
	NumberOfEmployees     string `json:"NumberOfEmployees"`
	Title                 string `json:"Title"`
	Industry              string `json:"Industry"`
	Phone                 string `json:"Phone"`
	MobilePhone           string `json:"MobilePhone"`
	Fax                   string `json:"Fax"`
	Website               string `json:"Website"`
	MailingStreet         string `json:"MailingStreet"`
	MailingCity           string `json:"MailingCity"`
	MailingState          string `json:"MailingState"`
	MailingPostalCode     string `json:"MailingPostalCode"`
	MailingCountry        string `json:"MailingCountry"`
	owner_name            string `json:"owner_name"`
	LeadSource            string `json:"LeadSource"`
	Status                string `json:"Status"`
	LinkedIn              string `json:"LinkedIn"`
	unsubscribed          string `json:unsubscribed""`
	custom                string `json:"custom"`
	_autopilot_session_id string `json:"_autopilot_session_id"`
	_autopilot_list       string `json:"_autopilot_list"`
	notify                string `json:"notify"`
}
