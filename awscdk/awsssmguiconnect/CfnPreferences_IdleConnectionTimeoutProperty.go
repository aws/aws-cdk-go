package awsssmguiconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idleConnectionTimeoutProperty := &IdleConnectionTimeoutProperty{
//   	Value: jsii.Number(123),
//
//   	// the properties below are optional
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-idleconnectiontimeout.html
//
type CfnPreferences_IdleConnectionTimeoutProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-idleconnectiontimeout.html#cfn-ssmguiconnect-preferences-idleconnectiontimeout-value
	//
	// Default: - 10.
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-idleconnectiontimeout.html#cfn-ssmguiconnect-preferences-idleconnectiontimeout-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

