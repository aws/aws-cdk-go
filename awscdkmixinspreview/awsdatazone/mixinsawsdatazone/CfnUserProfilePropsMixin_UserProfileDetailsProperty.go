package mixinsawsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnUserProfilePropsMixin_UserProfileDetailsProperty struct {
	// The details of the IAM User Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-userprofiledetails.html#cfn-datazone-userprofile-userprofiledetails-iam
	//
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// The details of the SSO User Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-userprofiledetails.html#cfn-datazone-userprofile-userprofiledetails-sso
	//
	Sso interface{} `field:"optional" json:"sso" yaml:"sso"`
}

