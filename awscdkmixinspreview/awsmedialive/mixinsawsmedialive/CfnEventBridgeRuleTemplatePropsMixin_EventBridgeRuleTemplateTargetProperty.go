package mixinsawsmedialive


// The target to which to send matching events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventBridgeRuleTemplateTargetProperty := &EventBridgeRuleTemplateTargetProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-eventbridgeruletemplate-eventbridgeruletemplatetarget.html
//
type CfnEventBridgeRuleTemplatePropsMixin_EventBridgeRuleTemplateTargetProperty struct {
	// Target ARNs must be either an SNS topic or CloudWatch log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-eventbridgeruletemplate-eventbridgeruletemplatetarget.html#cfn-medialive-eventbridgeruletemplate-eventbridgeruletemplatetarget-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

