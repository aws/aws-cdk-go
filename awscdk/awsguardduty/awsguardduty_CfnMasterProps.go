package awsguardduty


// Properties for defining a `CfnMaster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMasterProps := &cfnMasterProps{
//   	detectorId: jsii.String("detectorId"),
//   	masterId: jsii.String("masterId"),
//
//   	// the properties below are optional
//   	invitationId: jsii.String("invitationId"),
//   }
//
type CfnMasterProps struct {
	// The unique ID of the detector of the GuardDuty member account.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The AWS account ID of the account designated as the  administrator account.
	MasterId *string `field:"required" json:"masterId" yaml:"masterId"`
	// The ID of the invitation that is sent to the account designated as a member account.
	//
	// You can find the invitation ID by using the ListInvitation action of the  API.
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
}

