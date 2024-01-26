package awsssmguiconnect


// Properties for defining a `CfnPreferences`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPreferencesProps := &CfnPreferencesProps{
//   	IdleConnection: []interface{}{
//   		&IdleConnectionPreferencesProperty{
//   			Alert: &IdleConnectionAlertProperty{
//   				Value: jsii.Number(123),
//
//   				// the properties below are optional
//   				Type: jsii.String("type"),
//   			},
//   			Timeout: &IdleConnectionTimeoutProperty{
//   				Value: jsii.Number(123),
//
//   				// the properties below are optional
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmguiconnect-preferences.html
//
type CfnPreferencesProps struct {
	// A map for Idle Connection Preferences.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmguiconnect-preferences.html#cfn-ssmguiconnect-preferences-idleconnection
	//
	IdleConnection interface{} `field:"optional" json:"idleConnection" yaml:"idleConnection"`
}

