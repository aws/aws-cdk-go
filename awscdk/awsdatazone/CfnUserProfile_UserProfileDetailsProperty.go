package awsdatazone


// The details of the user profile in Amazon DataZone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userProfileDetailsProperty := &UserProfileDetailsProperty{
//   	Iam: &IamUserProfileDetailsProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	Sso: &SsoUserProfileDetailsProperty{
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   		Username: jsii.String("username"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-userprofiledetails.html
//
type CfnUserProfile_UserProfileDetailsProperty struct {
	// The IAM details included in the user profile details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-userprofiledetails.html#cfn-datazone-userprofile-userprofiledetails-iam
	//
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// The single sign-on details included in the user profile details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-userprofiledetails.html#cfn-datazone-userprofile-userprofiledetails-sso
	//
	Sso interface{} `field:"optional" json:"sso" yaml:"sso"`
}

