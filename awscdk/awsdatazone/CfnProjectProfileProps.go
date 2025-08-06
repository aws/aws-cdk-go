package awsdatazone


// Properties for defining a `CfnProjectProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProfileProps := &CfnProjectProfileProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	DomainUnitIdentifier: jsii.String("domainUnitIdentifier"),
//   	EnvironmentConfigurations: []interface{}{
//   		&EnvironmentConfigurationProperty{
//   			AwsRegion: &RegionProperty{
//   				RegionName: jsii.String("regionName"),
//   			},
//   			EnvironmentBlueprintId: jsii.String("environmentBlueprintId"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			AwsAccount: &AwsAccountProperty{
//   				AwsAccountId: jsii.String("awsAccountId"),
//   			},
//   			ConfigurationParameters: &EnvironmentConfigurationParametersDetailsProperty{
//   				ParameterOverrides: []interface{}{
//   					&EnvironmentConfigurationParameterProperty{
//   						IsEditable: jsii.Boolean(false),
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				ResolvedParameters: []interface{}{
//   					&EnvironmentConfigurationParameterProperty{
//   						IsEditable: jsii.Boolean(false),
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				SsmPath: jsii.String("ssmPath"),
//   			},
//   			DeploymentMode: jsii.String("deploymentMode"),
//   			DeploymentOrder: jsii.Number(123),
//   			Description: jsii.String("description"),
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html
//
type CfnProjectProfileProps struct {
	// The name of a project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-domainunitidentifier
	//
	DomainUnitIdentifier *string `field:"optional" json:"domainUnitIdentifier" yaml:"domainUnitIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-environmentconfigurations
	//
	EnvironmentConfigurations interface{} `field:"optional" json:"environmentConfigurations" yaml:"environmentConfigurations"`
	// The status of a project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

