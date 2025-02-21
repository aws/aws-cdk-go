package awsopsworks


// Properties for defining a `CfnUserProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProfileProps := &CfnUserProfileProps{
//   	IamUserArn: jsii.String("iamUserArn"),
//
//   	// the properties below are optional
//   	AllowSelfManagement: jsii.Boolean(false),
//   	SshPublicKey: jsii.String("sshPublicKey"),
//   	SshUsername: jsii.String("sshUsername"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html
//
type CfnUserProfileProps struct {
	// The user's IAM ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-iamuserarn
	//
	IamUserArn *string `field:"required" json:"iamUserArn" yaml:"iamUserArn"`
	// Whether users can specify their own SSH public key through the My Settings page.
	//
	// For more information, see [Managing User Permissions](https://docs.aws.amazon.com/opsworks/latest/userguide/security-settingsshkey.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-allowselfmanagement
	//
	AllowSelfManagement interface{} `field:"optional" json:"allowSelfManagement" yaml:"allowSelfManagement"`
	// The user's SSH public key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-sshpublickey
	//
	SshPublicKey *string `field:"optional" json:"sshPublicKey" yaml:"sshPublicKey"`
	// The user's SSH user name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-sshusername
	//
	SshUsername *string `field:"optional" json:"sshUsername" yaml:"sshUsername"`
}

