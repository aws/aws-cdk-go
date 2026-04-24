package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bounceActionProperty := &BounceActionProperty{
//   	DiagnosticMessage: jsii.String("diagnosticMessage"),
//   	RoleArn: jsii.String("roleArn"),
//   	Sender: jsii.String("sender"),
//   	SmtpReplyCode: jsii.String("smtpReplyCode"),
//   	StatusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html
//
type CfnMailManagerRuleSet_BounceActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-diagnosticmessage
	//
	DiagnosticMessage *string `field:"required" json:"diagnosticMessage" yaml:"diagnosticMessage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-sender
	//
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-smtpreplycode
	//
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-statuscode
	//
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

