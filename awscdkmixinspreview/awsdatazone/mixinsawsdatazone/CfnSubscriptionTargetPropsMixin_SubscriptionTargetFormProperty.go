package mixinsawsdatazone


// The details of the subscription target configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subscriptionTargetFormProperty := &SubscriptionTargetFormProperty{
//   	Content: jsii.String("content"),
//   	FormName: jsii.String("formName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-subscriptiontarget-subscriptiontargetform.html
//
type CfnSubscriptionTargetPropsMixin_SubscriptionTargetFormProperty struct {
	// The content of the subscription target configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-subscriptiontarget-subscriptiontargetform.html#cfn-datazone-subscriptiontarget-subscriptiontargetform-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The form name included in the subscription target configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-subscriptiontarget-subscriptiontargetform.html#cfn-datazone-subscriptiontarget-subscriptiontargetform-formname
	//
	FormName *string `field:"optional" json:"formName" yaml:"formName"`
}

