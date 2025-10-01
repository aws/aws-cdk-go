package awslocation


// A reference to a Tracker resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trackerReference := &TrackerReference{
//   	TrackerArn: jsii.String("trackerArn"),
//   	TrackerName: jsii.String("trackerName"),
//   }
//
type TrackerReference struct {
	// The ARN of the Tracker resource.
	TrackerArn *string `field:"required" json:"trackerArn" yaml:"trackerArn"`
	// The TrackerName of the Tracker resource.
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
}

