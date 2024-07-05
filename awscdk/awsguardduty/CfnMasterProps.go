package awsguardduty


// Properties for defining a `CfnMaster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMasterProps := &CfnMasterProps{
//   	DetectorId: jsii.String("detectorId"),
//   	MasterId: jsii.String("masterId"),
//
//   	// the properties below are optional
//   	InvitationId: jsii.String("invitationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html
//
type CfnMasterProps struct {
	// The unique ID of the detector of the GuardDuty member account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-detectorid
	//
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The AWS account ID of the account designated as the GuardDuty administrator account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-masterid
	//
	MasterId *string `field:"required" json:"masterId" yaml:"masterId"`
	// The ID of the invitation that is sent to the account designated as a member account.
	//
	// You can find the invitation ID by running the [ListInvitations](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListInvitations.html) in the *GuardDuty API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-invitationid
	//
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
}

