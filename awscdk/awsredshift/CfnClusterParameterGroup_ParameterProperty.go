package awsredshift


// Describes a parameter in a cluster parameter group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterProperty := &ParameterProperty{
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
type CfnClusterParameterGroup_ParameterProperty struct {
	// The name of the parameter.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The value of the parameter.
	//
	// If `ParameterName` is `wlm_json_configuration` , then the maximum size of `ParameterValue` is 8000 characters.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

