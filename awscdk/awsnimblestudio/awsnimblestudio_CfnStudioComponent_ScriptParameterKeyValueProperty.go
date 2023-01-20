package awsnimblestudio


// A parameter for a studio component script, in the form of a key:value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptParameterKeyValueProperty := &scriptParameterKeyValueProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnStudioComponent_ScriptParameterKeyValueProperty struct {
	// A script parameter key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A script parameter value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

