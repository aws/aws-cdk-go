// Version 2 of the AWS Cloud Development Kit library
package awscdk


// The Parameter data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterProperty := &parameterProperty{
//   	parameterKey: jsii.String("parameterKey"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnStackSet_ParameterProperty struct {
	// The key associated with the parameter.
	//
	// If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that's specified in your template.
	ParameterKey *string `field:"required" json:"parameterKey" yaml:"parameterKey"`
	// The input value associated with the parameter.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

