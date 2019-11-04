package dsb

// Account stores login information
type Account struct {
	username string
	password string

	BundleID   string
	Webservice string
	AppVersion string
	Lang       string
	Device     string
	OsVersion  string
	UserAgent  string
}

// Response contains all of the information
type Response struct {
	StatusCode  int        `json:"Resultcode"`
	Status      string     `json:"ResultStatusInfo"`
	StartIndex  int        `json:"StartIndex"`
	Categorys   []Category `json:"ResultMenuItems"`
	ChannelType int        `json:"ChannelType"`
	MandantID   string     `json:"MandantId"`
}

// Category contains Menus
type Category struct {
	Index         int    `json:"Index"`
	Icon          string `json:"IconLink"`
	Title         string `json:"Title"`
	Menus         []Menu `json:"Childs"`
	Method        string `json:"MethodName"`
	NewCount      int    `json:"NewCount"`
	SaveLastState bool   `json:"SaveLastState"`
}

// Menu contains MenusItems
type Menu struct {
	Index int      `json:"Index"`
	Icon  string   `json:"IconLink"`
	Title string   `json:"Title"`
	Root  MenuItem `json:"Root"`
	//Childs        []Menu   `json:"Childs"`
	Method        string `json:"MethodName"`
	NewCount      int    `json:"NewCount"`
	SaveLastState bool   `json:"SaveLastState"`
}

// MenuItem can contain more MenuItems
type MenuItem struct {
	ID      string     `json:"Id"`
	Date    string     `json:"Date"`
	Title   string     `json:"Title"`
	Detail  string     `json:"Detail"`
	Tags    string     `json:"Tags"`
	ConType int        `json:"ConType"`
	Prio    int        `json:"Prio"`
	Index   int        `json:"Index"`
	Childs  []MenuItem `json:"Childs"`
	Preview string     `json:"Preview"`
}
