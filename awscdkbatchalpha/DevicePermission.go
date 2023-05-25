package awscdkbatchalpha


// Permissions for device access.
// Experimental.
type DevicePermission string

const (
	// Read.
	// Experimental.
	DevicePermission_READ DevicePermission = "READ"
	// Write.
	// Experimental.
	DevicePermission_WRITE DevicePermission = "WRITE"
	// Make a node.
	// Experimental.
	DevicePermission_MKNOD DevicePermission = "MKNOD"
)

