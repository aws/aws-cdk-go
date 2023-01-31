package awscloudtrail


// Types of events that CloudTrail can log.
//
// Example:
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &trailProps{
//   	// ...
//   	managementEvents: cloudtrail.readWriteType_READ_ONLY,
//   })
//
// Experimental.
type ReadWriteType string

const (
	// Read-only events include API operations that read your resources, but don't make changes.
	//
	// For example, read-only events include the Amazon EC2 DescribeSecurityGroups
	// and DescribeSubnets API operations.
	// Experimental.
	ReadWriteType_READ_ONLY ReadWriteType = "READ_ONLY"
	// Write-only events include API operations that modify (or might modify) your resources.
	//
	// For example, the Amazon EC2 RunInstances and TerminateInstances API
	// operations modify your instances.
	// Experimental.
	ReadWriteType_WRITE_ONLY ReadWriteType = "WRITE_ONLY"
	// All events.
	// Experimental.
	ReadWriteType_ALL ReadWriteType = "ALL"
	// No events.
	// Experimental.
	ReadWriteType_NONE ReadWriteType = "NONE"
)

