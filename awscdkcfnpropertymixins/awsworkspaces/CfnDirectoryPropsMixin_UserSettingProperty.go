package awsworkspaces


// Information about the user's permission settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   userSettingProperty := &UserSettingProperty{
//   	Action: jsii.String("action"),
//   	MaximumLength: jsii.Number(123),
//   	Permission: jsii.String("permission"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-usersetting.html
//
type CfnDirectoryPropsMixin_UserSettingProperty struct {
	// Indicates the type of action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-usersetting.html#cfn-workspaces-directory-usersetting-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Indicates the maximum character length for the specified user setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-usersetting.html#cfn-workspaces-directory-usersetting-maximumlength
	//
	MaximumLength *float64 `field:"optional" json:"maximumLength" yaml:"maximumLength"`
	// Indicates if the setting is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-usersetting.html#cfn-workspaces-directory-usersetting-permission
	//
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

