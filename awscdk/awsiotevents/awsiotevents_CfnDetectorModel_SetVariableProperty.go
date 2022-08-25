package awsiotevents


// Information about the variable and its new value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   setVariableProperty := &setVariableProperty{
//   	value: jsii.String("value"),
//   	variableName: jsii.String("variableName"),
//   }
//
type CfnDetectorModel_SetVariableProperty struct {
	// The new value of the variable.
	Value *string `field:"required" json:"value" yaml:"value"`
	// The name of the variable.
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
}

