package awsec2


// Maximum VPN session duration time.
type ClientVpnSessionTimeout string

const (
	// 8 hours.
	ClientVpnSessionTimeout_EIGHT_HOURS ClientVpnSessionTimeout = "EIGHT_HOURS"
	// 10 hours.
	ClientVpnSessionTimeout_TEN_HOURS ClientVpnSessionTimeout = "TEN_HOURS"
	// 12 hours.
	ClientVpnSessionTimeout_TWELVE_HOURS ClientVpnSessionTimeout = "TWELVE_HOURS"
	// 24 hours.
	ClientVpnSessionTimeout_TWENTY_FOUR_HOURS ClientVpnSessionTimeout = "TWENTY_FOUR_HOURS"
)

