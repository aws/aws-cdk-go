package awsworkspaces


// Information about the user's permission settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userSettingProperty := &UserSettingProperty{
//   	Action: jsii.String("action"),
//   	Permission: jsii.String("permission"),
//
//   	// the properties below are optional
//   	MaximumLength: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-usersetting.html
//
type CfnDirectory_UserSettingProperty struct {
	// Indicates the type of action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-usersetting.html#cfn-workspaces-directory-usersetting-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// Indicates if the setting is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-usersetting.html#cfn-workspaces-directory-usersetting-permission
	//
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Indicates the maximum character length for the specified user setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-usersetting.html#cfn-workspaces-directory-usersetting-maximumlength
	//
	MaximumLength *float64 `field:"optional" json:"maximumLength" yaml:"maximumLength"`
}

