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
//   }
//
type CfnStack_UserSettingProperty struct {
	// The action that is enabled or disabled.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Indicates whether the action is enabled or disabled.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

