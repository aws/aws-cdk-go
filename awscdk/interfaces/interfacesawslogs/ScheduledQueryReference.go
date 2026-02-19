package interfacesawslogs


// A reference to a ScheduledQuery resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledQueryReference := &ScheduledQueryReference{
//   	ScheduledQueryArn: jsii.String("scheduledQueryArn"),
//   }
//
type ScheduledQueryReference struct {
	// The ScheduledQueryArn of the ScheduledQuery resource.
	ScheduledQueryArn *string `field:"required" json:"scheduledQueryArn" yaml:"scheduledQueryArn"`
}

