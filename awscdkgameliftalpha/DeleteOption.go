package awscdkgameliftalpha


// The type of delete to perform.
//
// To delete a game server group, specify the DeleteOption.
// Experimental.
type DeleteOption string

const (
	// Terminates the game server group and Amazon EC2 Auto Scaling group only when it has no game servers that are in UTILIZED status.
	// Experimental.
	DeleteOption_SAFE_DELETE DeleteOption = "SAFE_DELETE"
	// Terminates the game server group, including all active game servers regardless of their utilization status, and the Amazon EC2 Auto Scaling group.
	// Experimental.
	DeleteOption_FORCE_DELETE DeleteOption = "FORCE_DELETE"
	// Does a safe delete of the game server group but retains the Amazon EC2 Auto Scaling group as is.
	// Experimental.
	DeleteOption_RETAIN DeleteOption = "RETAIN"
)

