package awsdatazone


// The details of the IAM User Profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamUserProfileDetailsProperty := &IamUserProfileDetailsProperty{
//   	Arn: jsii.String("arn"),
//   	GroupProfileId: jsii.String("groupProfileId"),
//   	SessionName: jsii.String("sessionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-iamuserprofiledetails.html
//
type CfnUserProfile_IamUserProfileDetailsProperty struct {
	// The ARN of the IAM User Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-iamuserprofiledetails.html#cfn-datazone-userprofile-iamuserprofiledetails-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The group profile ID of the IAM User Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-iamuserprofiledetails.html#cfn-datazone-userprofile-iamuserprofiledetails-groupprofileid
	//
	GroupProfileId *string `field:"optional" json:"groupProfileId" yaml:"groupProfileId"`
	// The session name of the IAM User Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-iamuserprofiledetails.html#cfn-datazone-userprofile-iamuserprofiledetails-sessionname
	//
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
}

