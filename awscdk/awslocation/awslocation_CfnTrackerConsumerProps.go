package awslocation


// Properties for defining a `CfnTrackerConsumer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrackerConsumerProps := &cfnTrackerConsumerProps{
//   	consumerArn: jsii.String("consumerArn"),
//   	trackerName: jsii.String("trackerName"),
//   }
//
type CfnTrackerConsumerProps struct {
	// The Amazon Resource Name (ARN) for the geofence collection that consumes the tracker resource updates.
	//
	// - Format example: `arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollectionConsumer`.
	ConsumerArn *string `field:"required" json:"consumerArn" yaml:"consumerArn"`
	// The name for the tracker resource.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A-Z, a-z, 0-9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique tracker resource name.
	// - No spaces allowed. For example, `ExampleTracker` .
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
}

