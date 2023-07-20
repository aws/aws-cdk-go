package awscdk


// The Parameter data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterProperty := &ParameterProperty{
//   	ParameterKey: jsii.String("parameterKey"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html
//
type CfnStackSet_ParameterProperty struct {
	// The key associated with the parameter.
	//
	// If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that's specified in your template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html#cfn-cloudformation-stackset-parameter-parameterkey
	//
	ParameterKey *string `field:"required" json:"parameterKey" yaml:"parameterKey"`
	// The input value associated with the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html#cfn-cloudformation-stackset-parameter-parametervalue
	//
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

