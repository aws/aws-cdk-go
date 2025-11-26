package previewawsresiliencehubmixins


// Indicates an event you would like to subscribe and get notification for.
//
// Currently, AWS Resilience Hub supports notifications only for *Drift detected* and *Scheduled assessment failure* events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventSubscriptionProperty := &EventSubscriptionProperty{
//   	EventType: jsii.String("eventType"),
//   	Name: jsii.String("name"),
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-eventsubscription.html
//
type CfnAppPropsMixin_EventSubscriptionProperty struct {
	// The type of event you would like to subscribe and get notification for.
	//
	// Currently, AWS Resilience Hub supports notifications only for *Drift detected* ( `DriftDetected` ) and *Scheduled assessment failure* ( `ScheduledAssessmentFailure` ) events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-eventsubscription.html#cfn-resiliencehub-app-eventsubscription-eventtype
	//
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// Unique name to identify an event subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-eventsubscription.html#cfn-resiliencehub-app-eventsubscription-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic.
	//
	// The format for this ARN is: `arn:partition:sns:region:account:topic-name` . For more information about ARNs, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-eventsubscription.html#cfn-resiliencehub-app-eventsubscription-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

