package awsdatazone


// The configuration of an environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentConfigurationProperty := &EnvironmentConfigurationProperty{
//   	AwsRegion: &RegionProperty{
//   		RegionName: jsii.String("regionName"),
//   	},
//   	EnvironmentBlueprintId: jsii.String("environmentBlueprintId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AwsAccount: &AwsAccountProperty{
//   		AwsAccountId: jsii.String("awsAccountId"),
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
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html
//
type CfnProjectProfile_EnvironmentConfigurationProperty struct {
	// The AWS Region of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-awsregion
	//
	AwsRegion interface{} `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The environment blueprint ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-environmentblueprintid
	//
	EnvironmentBlueprintId *string `field:"required" json:"environmentBlueprintId" yaml:"environmentBlueprintId"`
	// The environment name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS account of the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-awsaccount
	//
	AwsAccount interface{} `field:"optional" json:"awsAccount" yaml:"awsAccount"`
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
	// The environment ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-environmentconfiguration.html#cfn-datazone-projectprofile-environmentconfiguration-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

