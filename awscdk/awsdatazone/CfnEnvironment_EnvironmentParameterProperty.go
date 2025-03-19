package awsdatazone


// The parameter details of the environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentParameterProperty := &EnvironmentParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environment-environmentparameter.html
//
type CfnEnvironment_EnvironmentParameterProperty struct {
	// The name of the environment parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environment-environmentparameter.html#cfn-datazone-environment-environmentparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the environment parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environment-environmentparameter.html#cfn-datazone-environment-environmentparameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

