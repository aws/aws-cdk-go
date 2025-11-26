package previewawsneptunemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEventSubscriptionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventSubscriptionMixinProps := &CfnEventSubscriptionMixinProps{
//   	Enabled: jsii.Boolean(false),
//   	EventCategories: []*string{
//   		jsii.String("eventCategories"),
//   	},
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   	SourceIds: []*string{
//   		jsii.String("sourceIds"),
//   	},
//   	SourceType: jsii.String("sourceType"),
//   	SubscriptionName: jsii.String("subscriptionName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html
//
type CfnEventSubscriptionMixinProps struct {
	// A Boolean value indicating if the subscription is enabled.
	//
	// True indicates the subscription is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-enabled
	//
	// Default: - true.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of event categories for a SourceType that you want to subscribe to.
	//
	// You can see a list of the categories for a given SourceType in the Events topic in the Amazon Neptune User Guide or by using the DescribeEventCategories action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-eventcategories
	//
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// The topic ARN of the event notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// The list of identifiers of the event sources for which events will be returned.
	//
	// If not specified, then all sources are included in the response. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it cannot end with a hyphen or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-sourceids
	//
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
	// The source type for the event notification subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// The name of the subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-subscriptionname
	//
	SubscriptionName *string `field:"optional" json:"subscriptionName" yaml:"subscriptionName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-eventsubscription.html#cfn-neptune-eventsubscription-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

