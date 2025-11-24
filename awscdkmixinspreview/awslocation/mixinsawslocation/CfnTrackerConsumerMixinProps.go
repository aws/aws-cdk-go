package mixinsawslocation


// Properties for CfnTrackerConsumerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrackerConsumerMixinProps := &CfnTrackerConsumerMixinProps{
//   	ConsumerArn: jsii.String("consumerArn"),
//   	TrackerName: jsii.String("trackerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html
//
type CfnTrackerConsumerMixinProps struct {
	// The Amazon Resource Name (ARN) for the geofence collection to be associated to tracker resource.
	//
	// Used when you need to specify a resource across all AWS .
	//
	// - Format example: `arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollectionConsumer`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-consumerarn
	//
	ConsumerArn *string `field:"optional" json:"consumerArn" yaml:"consumerArn"`
	// The name for the tracker resource.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A-Z, a-z, 0-9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique tracker resource name.
	// - No spaces allowed. For example, `ExampleTracker` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-trackername
	//
	TrackerName *string `field:"optional" json:"trackerName" yaml:"trackerName"`
}

