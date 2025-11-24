package mixinsawsdetective


// Properties for CfnMemberInvitationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemberInvitationMixinProps := &CfnMemberInvitationMixinProps{
//   	DisableEmailNotification: jsii.Boolean(false),
//   	GraphArn: jsii.String("graphArn"),
//   	MemberEmailAddress: jsii.String("memberEmailAddress"),
//   	MemberId: jsii.String("memberId"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html
//
type CfnMemberInvitationMixinProps struct {
	// Whether to send an invitation email to the member account.
	//
	// If set to true, the member account does not receive an invitation email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-disableemailnotification
	//
	// Default: - false.
	//
	DisableEmailNotification interface{} `field:"optional" json:"disableEmailNotification" yaml:"disableEmailNotification"`
	// The ARN of the behavior graph to invite the account to contribute data to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-grapharn
	//
	GraphArn *string `field:"optional" json:"graphArn" yaml:"graphArn"`
	// The root user email address of the invited account.
	//
	// If the email address provided is not the root user email address for the provided account, the invitation creation fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberemailaddress
	//
	MemberEmailAddress *string `field:"optional" json:"memberEmailAddress" yaml:"memberEmailAddress"`
	// The AWS account identifier of the invited account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberid
	//
	MemberId *string `field:"optional" json:"memberId" yaml:"memberId"`
	// Customized text to include in the invitation email message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

