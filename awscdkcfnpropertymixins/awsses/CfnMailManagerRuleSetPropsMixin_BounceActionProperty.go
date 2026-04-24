package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   bounceActionProperty := &BounceActionProperty{
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	DiagnosticMessage: jsii.String("diagnosticMessage"),
//   	Message: jsii.String("message"),
//   	RoleArn: jsii.String("roleArn"),
//   	Sender: jsii.String("sender"),
//   	SmtpReplyCode: jsii.String("smtpReplyCode"),
//   	StatusCode: jsii.String("statusCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html
//
type CfnMailManagerRuleSetPropsMixin_BounceActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-diagnosticmessage
	//
	DiagnosticMessage *string `field:"optional" json:"diagnosticMessage" yaml:"diagnosticMessage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-sender
	//
	Sender *string `field:"optional" json:"sender" yaml:"sender"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-smtpreplycode
	//
	SmtpReplyCode *string `field:"optional" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-bounceaction.html#cfn-ses-mailmanagerruleset-bounceaction-statuscode
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

