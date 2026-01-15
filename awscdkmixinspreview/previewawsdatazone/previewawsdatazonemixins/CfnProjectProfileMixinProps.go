package previewawsdatazonemixins


// Properties for CfnProjectProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProjectProfileMixinProps := &CfnProjectProfileMixinProps{
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	DomainUnitIdentifier: jsii.String("domainUnitIdentifier"),
//   	EnvironmentConfigurations: []interface{}{
//   		&EnvironmentConfigurationProperty{
//   			AwsAccount: &AwsAccountProperty{
//   				AwsAccountId: jsii.String("awsAccountId"),
//   			},
//   			AwsRegion: &RegionProperty{
//   				RegionName: jsii.String("regionName"),
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
//   			EnvironmentBlueprintId: jsii.String("environmentBlueprintId"),
//   			EnvironmentConfigurationId: jsii.String("environmentConfigurationId"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//   	UseDefaultConfigurations: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html
//
type CfnProjectProfileMixinProps struct {
	// The description of the project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A domain ID of the project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// A domain unit ID of the project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-domainunitidentifier
	//
	DomainUnitIdentifier *string `field:"optional" json:"domainUnitIdentifier" yaml:"domainUnitIdentifier"`
	// Environment configurations of a project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-environmentconfigurations
	//
	EnvironmentConfigurations interface{} `field:"optional" json:"environmentConfigurations" yaml:"environmentConfigurations"`
	// The name of a project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The status of a project profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-usedefaultconfigurations
	//
	UseDefaultConfigurations interface{} `field:"optional" json:"useDefaultConfigurations" yaml:"useDefaultConfigurations"`
}

