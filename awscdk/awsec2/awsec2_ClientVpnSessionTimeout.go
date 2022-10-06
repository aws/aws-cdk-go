package awsec2


// Maximum VPN session duration time.
// Experimental.
type ClientVpnSessionTimeout string

const (
	// 8 hours.
	// Experimental.
	ClientVpnSessionTimeout_EIGHT_HOURS ClientVpnSessionTimeout = "EIGHT_HOURS"
	// 10 hours.
	// Experimental.
	ClientVpnSessionTimeout_TEN_HOURS ClientVpnSessionTimeout = "TEN_HOURS"
	// 12 hours.
	// Experimental.
	ClientVpnSessionTimeout_TWELVE_HOURS ClientVpnSessionTimeout = "TWELVE_HOURS"
	// 24 hours.
	// Experimental.
	ClientVpnSessionTimeout_TWENTY_FOUR_HOURS ClientVpnSessionTimeout = "TWENTY_FOUR_HOURS"
)

