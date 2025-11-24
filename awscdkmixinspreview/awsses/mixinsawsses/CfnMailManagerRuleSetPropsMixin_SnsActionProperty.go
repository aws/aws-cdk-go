package mixinsawsses


// The action to publish the email content to an Amazon SNS topic.
//
// When executed, this action will send the email as a notification to the specified SNS topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snsActionProperty := &SnsActionProperty{
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	Encoding: jsii.String("encoding"),
//   	PayloadType: jsii.String("payloadType"),
//   	RoleArn: jsii.String("roleArn"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html
//
type CfnMailManagerRuleSetPropsMixin_SnsActionProperty struct {
	// A policy that states what to do in the case of failure.
	//
	// The action will fail if there are configuration errors. For example, specified SNS topic has been deleted or the role lacks necessary permissions to call the `sns:Publish` API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// The encoding to use for the email within the Amazon SNS notification.
	//
	// The default value is `UTF-8` . Use `BASE64` if you need to preserve all special characters, especially when the original message uses a different encoding format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-encoding
	//
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The expected payload type within the Amazon SNS notification.
	//
	// `CONTENT` attempts to publish the full email content with 20KB of headers content. `HEADERS` extracts up to 100KB of header content to include in the notification, email content will not be included to the notification. The default value is `CONTENT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-payloadtype
	//
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// The Amazon Resource Name (ARN) of the IAM Role to use while writing to Amazon SNS.
	//
	// This role must have access to the `sns:Publish` API for the given topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the Amazon SNS Topic to which notification for the email received will be published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-snsaction.html#cfn-ses-mailmanagerruleset-snsaction-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

