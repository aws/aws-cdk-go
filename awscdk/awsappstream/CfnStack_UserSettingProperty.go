package awsappstream


// Specifies an action and whether the action is enabled or disabled for users during their streaming sessions.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html
//
type CfnStack_UserSettingProperty struct {
	// The action that is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html#cfn-appstream-stack-usersetting-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// Indicates whether the action is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html#cfn-appstream-stack-usersetting-permission
	//
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Specifies the number of characters that can be copied by end users from the local device to the remote session, and to the local device from the remote session.
	//
	// This can be specified only for the `CLIPBOARD_COPY_FROM_LOCAL_DEVICE` and `CLIPBOARD_COPY_TO_LOCAL_DEVICE` actions.
	//
	// This defaults to 20,971,520 (20 MB) when unspecified and the permission is `ENABLED` . This can't be specified when the permission is `DISABLED` .
	//
	// This can only be specified for AlwaysOn and OnDemand fleets. The attribute is not supported on Elastic fleets.
	//
	// The value can be between 1 and 20,971,520 (20 MB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html#cfn-appstream-stack-usersetting-maximumlength
	//
	MaximumLength *float64 `field:"optional" json:"maximumLength" yaml:"maximumLength"`
}

