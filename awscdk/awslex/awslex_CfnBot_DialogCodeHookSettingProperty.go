package awslex


// Specifies whether an intent uses the dialog code hook during conversations with a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dialogCodeHookSettingProperty := &dialogCodeHookSettingProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBot_DialogCodeHookSettingProperty struct {
	// Indicates whether an intent uses the dialog code hook during a conversation with a user.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

