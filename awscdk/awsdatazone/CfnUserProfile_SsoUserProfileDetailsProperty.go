package awsdatazone


// The details of the SSO User Profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssoUserProfileDetailsProperty := &SsoUserProfileDetailsProperty{
//   	FirstName: jsii.String("firstName"),
//   	LastName: jsii.String("lastName"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-ssouserprofiledetails.html
//
type CfnUserProfile_SsoUserProfileDetailsProperty struct {
	// The First Name of the IAM User Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-ssouserprofiledetails.html#cfn-datazone-userprofile-ssouserprofiledetails-firstname
	//
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The Last Name of the IAM User Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-ssouserprofiledetails.html#cfn-datazone-userprofile-ssouserprofiledetails-lastname
	//
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The username of the SSO User Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-ssouserprofiledetails.html#cfn-datazone-userprofile-ssouserprofiledetails-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

