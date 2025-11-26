package previewawsdatazonemixins


// The environment configuration user parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environmentConfigurationUserParameterProperty := &EnvironmentConfigurationUserParameterProperty{
//   	EnvironmentConfigurationName: jsii.String("environmentConfigurationName"),
//   	EnvironmentId: jsii.String("environmentId"),
//   	EnvironmentParameters: []interface{}{
//   		&EnvironmentParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentconfigurationuserparameter.html
//
type CfnProjectPropsMixin_EnvironmentConfigurationUserParameterProperty struct {
	// The environment configuration name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentconfigurationuserparameter.html#cfn-datazone-project-environmentconfigurationuserparameter-environmentconfigurationname
	//
	EnvironmentConfigurationName *string `field:"optional" json:"environmentConfigurationName" yaml:"environmentConfigurationName"`
	// The ID of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentconfigurationuserparameter.html#cfn-datazone-project-environmentconfigurationuserparameter-environmentid
	//
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// The environment parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-project-environmentconfigurationuserparameter.html#cfn-datazone-project-environmentconfigurationuserparameter-environmentparameters
	//
	EnvironmentParameters interface{} `field:"optional" json:"environmentParameters" yaml:"environmentParameters"`
}

