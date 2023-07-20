package awsdetective


// Properties for defining a `CfnMemberInvitation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMemberInvitationProps := &CfnMemberInvitationProps{
//   	GraphArn: jsii.String("graphArn"),
//   	MemberEmailAddress: jsii.String("memberEmailAddress"),
//   	MemberId: jsii.String("memberId"),
//
//   	// the properties below are optional
//   	DisableEmailNotification: jsii.Boolean(false),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html
//
type CfnMemberInvitationProps struct {
	// The ARN of the behavior graph to invite the account to contribute data to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-grapharn
	//
	GraphArn *string `field:"required" json:"graphArn" yaml:"graphArn"`
	// The root user email address of the invited account.
	//
	// If the email address provided is not the root user email address for the provided account, the invitation creation fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberemailaddress
	//
	MemberEmailAddress *string `field:"required" json:"memberEmailAddress" yaml:"memberEmailAddress"`
	// The AWS account identifier of the invited account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberid
	//
	MemberId *string `field:"required" json:"memberId" yaml:"memberId"`
	// Whether to send an invitation email to the member account.
	//
	// If set to true, the member account does not receive an invitation email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-disableemailnotification
	//
	DisableEmailNotification interface{} `field:"optional" json:"disableEmailNotification" yaml:"disableEmailNotification"`
	// Customized text to include in the invitation email message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

