package mixinsawsdatazone


// The details of the environment configuration parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnProjectProfilePropsMixin_EnvironmentConfigurationParametersDetailsProperty struct {
	// The parameter overrides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparametersdetails.html#cfn-datazone-projectprofile-environmentconfigurationparametersdetails-parameteroverrides
	//
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// The resolved environment configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparametersdetails.html#cfn-datazone-projectprofile-environmentconfigurationparametersdetails-resolvedparameters
	//
	ResolvedParameters interface{} `field:"optional" json:"resolvedParameters" yaml:"resolvedParameters"`
	// Ssm path environment configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfigurationparametersdetails.html#cfn-datazone-projectprofile-environmentconfigurationparametersdetails-ssmpath
	//
	SsmPath *string `field:"optional" json:"ssmPath" yaml:"ssmPath"`
}

