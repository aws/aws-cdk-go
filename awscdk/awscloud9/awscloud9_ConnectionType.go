package awscloud9


// The connection type used for connecting to an Amazon EC2 environment.
// Experimental.
type ConnectionType string

const (
	// Conect through SSH.
	// Experimental.
	ConnectionType_CONNECT_SSH ConnectionType = "CONNECT_SSH"
	// Connect through AWS Systems Manager.
	// Experimental.
	ConnectionType_CONNECT_SSM ConnectionType = "CONNECT_SSM"
)

