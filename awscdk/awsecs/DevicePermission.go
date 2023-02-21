package awsecs


// Permissions for device access.
type DevicePermission string

const (
	// Read.
	DevicePermission_READ DevicePermission = "READ"
	// Write.
	DevicePermission_WRITE DevicePermission = "WRITE"
	// Make a node.
	DevicePermission_MKNOD DevicePermission = "MKNOD"
)

