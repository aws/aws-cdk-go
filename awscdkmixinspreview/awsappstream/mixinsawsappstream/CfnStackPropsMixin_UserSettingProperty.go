package mixinsawsappstream


// Specifies an action and whether the action is enabled or disabled for users during their streaming sessions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userSettingProperty := &UserSettingProperty{
//   	Action: jsii.String("action"),
//   	MaximumLength: jsii.Number(123),
//   	Permission: jsii.String("permission"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html
//
type CfnStackPropsMixin_UserSettingProperty struct {
	// The action that is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html#cfn-appstream-stack-usersetting-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Specifies the number of characters that can be copied by end users from the local device to the remote session, and to the local device from the remote session.
	//
	// This can be specified only for the `CLIPBOARD_COPY_FROM_LOCAL_DEVICE` and `CLIPBOARD_COPY_TO_LOCAL_DEVICE` actions.
	//
	// This defaults to 20,971,520 (20 MB) when unspecified and the permission is `ENABLED` . This can't be specified when the permission is `DISABLED` .
	//
	// The value can be between 1 and 20,971,520 (20 MB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html#cfn-appstream-stack-usersetting-maximumlength
	//
	MaximumLength *float64 `field:"optional" json:"maximumLength" yaml:"maximumLength"`
	// Indicates whether the action is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html#cfn-appstream-stack-usersetting-permission
	//
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

