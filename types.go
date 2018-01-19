package dsb

// Account stores login information
type Account struct {
	username string
	password string
}

// Response contains all of the information
type Response struct {
	Resultcode  int        `json:"Resultcode"`
	StatusInfo  string     `json:"ResultStatusInfo"`
	StartIndex  int        `json:"StartIndex"`
	MenuItems   []MenuItem `json:"ResultMenuItems"`
	ChannelType int        `json:"ChannelType"`
	MandantID   string     `json:"MandantId"`
}

// MenuItem is a menu category
type MenuItem struct {
	Index         int             `json:"Index"`
	Icon          string          `json:"IconLink"`
	Title         string          `json:"Title"`
	Childs        []MenuItemChild `json:"Childs"`
	Method        string          `json:"MethodName"`
	NewCount      int             `json:"NewCount"`
	SaveLastState bool            `json:"SaveLastState"`
}

// MenuItemChild is a menu subcategory
type MenuItemChild struct {
	Index         int               `json:"Index"`
	Icon          string            `json:"IconLink"`
	Title         string            `json:"Title"`
	Root          MenuItemChildItem `json:"Root"`
	Childs        []MenuItemChild   `json:"Childs"`
	Method        string            `json:"MethodName"`
	NewCount      int               `json:"NewCount"`
	SaveLastState bool              `json:"SaveLastState"`
}

// MenuItemChildItem is part of MenuItemChild
type MenuItemChildItem struct {
	ID      string              `json:"Id"`
	Date    string              `json:"Date"`
	Title   string              `json:"Title"`
	Detail  string              `json:"Detail"`
	Tags    string              `json:"Tags"`
	ConType int                 `json:"ConType"`
	Prio    int                 `json:"Prio"`
	Index   int                 `json:"Index"`
	Childs  []MenuItemChildItem `json:"Childs"`
	Preview string              `json:"Preview"`
}
