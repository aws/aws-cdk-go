package awsses


// This action to delivers an email to a mailbox.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliverToMailboxActionProperty := &DeliverToMailboxActionProperty{
//   	MailboxArn: jsii.String("mailboxArn"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertomailboxaction.html
//
type CfnMailManagerRuleSet_DeliverToMailboxActionProperty struct {
	// The Amazon Resource Name (ARN) of a WorkMail organization to deliver the email to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertomailboxaction.html#cfn-ses-mailmanagerruleset-delivertomailboxaction-mailboxarn
	//
	MailboxArn *string `field:"required" json:"mailboxArn" yaml:"mailboxArn"`
	// The Amazon Resource Name (ARN) of an IAM role to use to execute this action.
	//
	// The role must have access to the workmail:DeliverToMailbox API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertomailboxaction.html#cfn-ses-mailmanagerruleset-delivertomailboxaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A policy that states what to do in the case of failure.
	//
	// The action will fail if there are configuration errors. For example, the mailbox ARN is no longer valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-delivertomailboxaction.html#cfn-ses-mailmanagerruleset-delivertomailboxaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
}

