package mixinsawsguardduty


// Properties for CfnMasterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMasterMixinProps := &CfnMasterMixinProps{
//   	DetectorId: jsii.String("detectorId"),
//   	InvitationId: jsii.String("invitationId"),
//   	MasterId: jsii.String("masterId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html
//
type CfnMasterMixinProps struct {
	// The unique ID of the detector of the GuardDuty member account.
	//
	// To find the `detectorId` in the current Region, see the
	// Settings page in the GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html) API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-detectorid
	//
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// The ID of the invitation that is sent to the account designated as a member account.
	//
	// You can find the invitation ID by running the [ListInvitations](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListInvitations.html) in the *GuardDuty API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-invitationid
	//
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
	// The AWS account ID of the account designated as the GuardDuty administrator account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-masterid
	//
	MasterId *string `field:"optional" json:"masterId" yaml:"masterId"`
}

