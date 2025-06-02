package awsneptune


// Properties for defining a `CfnEventSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSubscriptionProps := &CfnEventSubscriptionProps{
//   	Enabled: jsii.Boolean(false),
//   	EventCategories: []*string{
//   		jsii.String("eventCategories"),
//   	},
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   	SourceIds: []*string{
//   		jsii.String("sourceIds"),
//   	},
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html
//
type CfnEventSubscriptionProps struct {
	// A Boolean value indicating if the subscription is enabled.
	//
	// True indicates the subscription is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-eventcategories
	//
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// The topic ARN of the event notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-sourceids
	//
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
	// The source type for the event notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

