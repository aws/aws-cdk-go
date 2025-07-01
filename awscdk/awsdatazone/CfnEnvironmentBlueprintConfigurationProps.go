package awsdatazone


// Properties for defining a `CfnEnvironmentBlueprintConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentBlueprintConfigurationProps := &CfnEnvironmentBlueprintConfigurationProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnabledRegions: []*string{
//   		jsii.String("enabledRegions"),
//   	},
//   	EnvironmentBlueprintIdentifier: jsii.String("environmentBlueprintIdentifier"),
//
//   	// the properties below are optional
//   	EnvironmentRolePermissionBoundary: jsii.String("environmentRolePermissionBoundary"),
//   	ManageAccessRoleArn: jsii.String("manageAccessRoleArn"),
//   	ProvisioningConfigurations: []interface{}{
//   		&ProvisioningConfigurationProperty{
//   			LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   				LocationRegistrationExcludeS3Locations: []*string{
//   					jsii.String("locationRegistrationExcludeS3Locations"),
//   				},
//   				LocationRegistrationRole: jsii.String("locationRegistrationRole"),
//   			},
//   		},
//   	},
//   	ProvisioningRoleArn: jsii.String("provisioningRoleArn"),
//   	RegionalParameters: []interface{}{
//   		&RegionalParameterProperty{
//   			Parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			Region: jsii.String("region"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html
//
type CfnEnvironmentBlueprintConfigurationProps struct {
	// The identifier of the Amazon DataZone domain in which an environment blueprint exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The enabled AWS Regions specified in a blueprint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-enabledregions
	//
	EnabledRegions *[]*string `field:"required" json:"enabledRegions" yaml:"enabledRegions"`
	// The identifier of the environment blueprint.
	//
	// In the current release, only the following values are supported: `DefaultDataLake` and `DefaultDataWarehouse` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-environmentblueprintidentifier
	//
	EnvironmentBlueprintIdentifier *string `field:"required" json:"environmentBlueprintIdentifier" yaml:"environmentBlueprintIdentifier"`
	// The environment role permission boundary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-environmentrolepermissionboundary
	//
	EnvironmentRolePermissionBoundary *string `field:"optional" json:"environmentRolePermissionBoundary" yaml:"environmentRolePermissionBoundary"`
	// The ARN of the manage access role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-manageaccessrolearn
	//
	ManageAccessRoleArn *string `field:"optional" json:"manageAccessRoleArn" yaml:"manageAccessRoleArn"`
	// The provisioning configuration of a blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-provisioningconfigurations
	//
	ProvisioningConfigurations interface{} `field:"optional" json:"provisioningConfigurations" yaml:"provisioningConfigurations"`
	// The ARN of the provisioning role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-provisioningrolearn
	//
	ProvisioningRoleArn *string `field:"optional" json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// The regional parameters of the environment blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-regionalparameters
	//
	RegionalParameters interface{} `field:"optional" json:"regionalParameters" yaml:"regionalParameters"`
}

