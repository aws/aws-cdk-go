package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentConfigurationParametersDetailsProperty := &EnvironmentConfigurationParametersDetailsProperty{
//   	ParameterOverrides: []interface{}{
//   		&EnvironmentConfigurationParameterProperty{
//   			IsEditable: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResolvedParameters: []interface{}{
//   		&EnvironmentConfigurationParameterProperty{
//   			IsEditable: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	SsmPath: jsii.String("ssmPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparametersdetails.html
//
type CfnProjectProfile_EnvironmentConfigurationParametersDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparametersdetails.html#cfn-datazone-projectprofile-environmentconfigurationparametersdetails-parameteroverrides
	//
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparametersdetails.html#cfn-datazone-projectprofile-environmentconfigurationparametersdetails-resolvedparameters
	//
	ResolvedParameters interface{} `field:"optional" json:"resolvedParameters" yaml:"resolvedParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparametersdetails.html#cfn-datazone-projectprofile-environmentconfigurationparametersdetails-ssmpath
	//
	SsmPath *string `field:"optional" json:"ssmPath" yaml:"ssmPath"`
}

