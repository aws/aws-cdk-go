package previewawsdatazonemixins


// The configuration of an environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environmentConfigurationProperty := &EnvironmentConfigurationProperty{
//   	AwsAccount: &AwsAccountProperty{
//   		AwsAccountId: jsii.String("awsAccountId"),
//   	},
//   	AwsRegion: &RegionProperty{
//   		RegionName: jsii.String("regionName"),
//   	},
//   	ConfigurationParameters: &EnvironmentConfigurationParametersDetailsProperty{
//   		ParameterOverrides: []interface{}{
//   			&EnvironmentConfigurationParameterProperty{
//   				IsEditable: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResolvedParameters: []interface{}{
//   			&EnvironmentConfigurationParameterProperty{
//   				IsEditable: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		SsmPath: jsii.String("ssmPath"),
//   	},
//   	DeploymentMode: jsii.String("deploymentMode"),
//   	DeploymentOrder: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	EnvironmentBlueprintId: jsii.String("environmentBlueprintId"),
//   	EnvironmentConfigurationId: jsii.String("environmentConfigurationId"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html
//
type CfnProjectProfilePropsMixin_EnvironmentConfigurationProperty struct {
	// The AWS account of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-awsaccount
	//
	AwsAccount interface{} `field:"optional" json:"awsAccount" yaml:"awsAccount"`
	// The AWS Region of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-awsregion
	//
	AwsRegion interface{} `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The configuration parameters of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-configurationparameters
	//
	ConfigurationParameters interface{} `field:"optional" json:"configurationParameters" yaml:"configurationParameters"`
	// The deployment mode of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-deploymentmode
	//
	DeploymentMode *string `field:"optional" json:"deploymentMode" yaml:"deploymentMode"`
	// The deployment order of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-deploymentorder
	//
	DeploymentOrder *float64 `field:"optional" json:"deploymentOrder" yaml:"deploymentOrder"`
	// The environment description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The environment blueprint ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-environmentblueprintid
	//
	EnvironmentBlueprintId *string `field:"optional" json:"environmentBlueprintId" yaml:"environmentBlueprintId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-environmentconfigurationid
	//
	EnvironmentConfigurationId *string `field:"optional" json:"environmentConfigurationId" yaml:"environmentConfigurationId"`
	// The environment name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

