package awsconfig


// Input parameters in the form of key-value pairs for the conformance pack, both of which you define.
//
// Keys can have a maximum character length of 255 characters, and values can have a maximum length of 4096 characters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conformancePackInputParameterProperty := &conformancePackInputParameterProperty{
//   	parameterName: jsii.String("parameterName"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnOrganizationConformancePack_ConformancePackInputParameterProperty struct {
	// One part of a key-value pair.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// One part of a key-value pair.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

