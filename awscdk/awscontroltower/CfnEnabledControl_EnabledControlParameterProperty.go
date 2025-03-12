package awscontroltower


// A set of parameters that configure the behavior of the enabled control.
//
// Expressed as a key/value pair, where `Key` is of type `String` and `Value` is of type `Document` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   enabledControlParameterProperty := &EnabledControlParameterProperty{
//   	Key: jsii.String("key"),
//   	Value: value,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controltower-enabledcontrol-enabledcontrolparameter.html
//
type CfnEnabledControl_EnabledControlParameterProperty struct {
	// The key of a key/value pair.
	//
	// It is of type `string` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controltower-enabledcontrol-enabledcontrolparameter.html#cfn-controltower-enabledcontrol-enabledcontrolparameter-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of a key/value pair.
	//
	// It can be of type `array` , `string` , `number` , `object` , or `boolean` . [Note: The *Type* field that follows may show a single type such as Number, which is only one possible type.]
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controltower-enabledcontrol-enabledcontrolparameter.html#cfn-controltower-enabledcontrol-enabledcontrolparameter-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

