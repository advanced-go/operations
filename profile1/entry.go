package profile1

const (
	CreatedTSName = "created_ts"
	RegionName    = "region"
	ZoneName      = "zone"
	SubZoneName   = "sub_zone"
	HostName      = "host"
)

var (
	//safeEntry = common.NewSafe()
	//entryData = Entry{Days: [7][24]Window{[{Hour: 1, Tag:"Off-Peak", Rate: 10}]}}
	win = [24]Window{
		{Hour: 0, Tag: "Scale-Down", Rate: 5},
		{Hour: 1, Tag: "Scale-Down", Rate: 5},
		{Hour: 2, Tag: "Off-Peak", Rate: 5},
		{Hour: 3, Tag: "Off-Peak", Rate: 5},
		{Hour: 4, Tag: "Off-Peak", Rate: 5},
		{Hour: 5, Tag: "Off-Peak", Rate: 5},
		{Hour: 6, Tag: "Off-Peak", Rate: 5},
		{Hour: 7, Tag: "Scale-Up", Rate: 5},
		{Hour: 8, Tag: "Scale-Up", Rate: 5},
		{Hour: 9, Tag: "Scale-Up", Rate: 5},
		{Hour: 10, Tag: "Scale-Up", Rate: 5},
		{Hour: 11, Tag: "Scale-Up", Rate: 5},
		{Hour: 12, Tag: "Peak", Rate: 5},
		{Hour: 13, Tag: "Peak", Rate: 5},
		{Hour: 14, Tag: "Peak", Rate: 5},
		{Hour: 15, Tag: "Peak", Rate: 5},
		{Hour: 16, Tag: "Scale-Down", Rate: 5},
		{Hour: 17, Tag: "Scale-Down", Rate: 5},
		{Hour: 18, Tag: "Scale-Down", Rate: 5},
		{Hour: 19, Tag: "Scale-Down", Rate: 5},
		{Hour: 20, Tag: "Scale-Down", Rate: 5},
		{Hour: 21, Tag: "Scale-Down", Rate: 5},
		{Hour: 22, Tag: "Scale-Down", Rate: 5},
		{Hour: 23, Tag: "Scale-Down", Rate: 5},
	}

	entry = Entry{Days: [7][24]Window{win}}
)

/*
{Window: {Hour: 1, Tag:"Off-Peak", Rate: 10}],
//{Region: "us-west1", Zone: "a", Host: "www.host1.com", AgentId: "test-agent", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
*/
type Window struct {
	Hour int
	Tag  string // Peak,Off-Peak,Scale-Up,Scale-Down
	Rate int
}

// Entry - host
type Entry struct {
	Days [7][24]Window
}
