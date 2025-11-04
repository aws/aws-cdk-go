package awscdkeksv2alpha


// Represents the different types of access entries that can be used in an Amazon EKS cluster.
// Experimental.
type AccessEntryType string

const (
	// Represents a standard access entry.
	// Experimental.
	AccessEntryType_STANDARD AccessEntryType = "STANDARD"
	// Represents a Fargate Linux access entry.
	// Experimental.
	AccessEntryType_FARGATE_LINUX AccessEntryType = "FARGATE_LINUX"
	// Represents an EC2 Linux access entry.
	// Experimental.
	AccessEntryType_EC2_LINUX AccessEntryType = "EC2_LINUX"
	// Represents an EC2 Windows access entry.
	// Experimental.
	AccessEntryType_EC2_WINDOWS AccessEntryType = "EC2_WINDOWS"
)

