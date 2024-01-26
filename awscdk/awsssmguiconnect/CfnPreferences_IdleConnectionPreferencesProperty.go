package awsssmguiconnect


// Idle Connection Preferences.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idleConnectionPreferencesProperty := &IdleConnectionPreferencesProperty{
//   	Alert: &IdleConnectionAlertProperty{
//   		Value: jsii.Number(123),
//
//   		// the properties below are optional
//   		Type: jsii.String("type"),
//   	},
//   	Timeout: &IdleConnectionTimeoutProperty{
//   		Value: jsii.Number(123),
//
//   		// the properties below are optional
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-idleconnectionpreferences.html
//
type CfnPreferences_IdleConnectionPreferencesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-idleconnectionpreferences.html#cfn-ssmguiconnect-preferences-idleconnectionpreferences-alert
	//
	Alert interface{} `field:"optional" json:"alert" yaml:"alert"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-idleconnectionpreferences.html#cfn-ssmguiconnect-preferences-idleconnectionpreferences-timeout
	//
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

