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
//   	ManageAccessRoleArn: jsii.String("manageAccessRoleArn"),
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-enabledregions
	//
	EnabledRegions *[]*string `field:"required" json:"enabledRegions" yaml:"enabledRegions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-environmentblueprintidentifier
	//
	EnvironmentBlueprintIdentifier *string `field:"required" json:"environmentBlueprintIdentifier" yaml:"environmentBlueprintIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-manageaccessrolearn
	//
	ManageAccessRoleArn *string `field:"optional" json:"manageAccessRoleArn" yaml:"manageAccessRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-provisioningrolearn
	//
	ProvisioningRoleArn *string `field:"optional" json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-environmentblueprintconfiguration.html#cfn-datazone-environmentblueprintconfiguration-regionalparameters
	//
	RegionalParameters interface{} `field:"optional" json:"regionalParameters" yaml:"regionalParameters"`
}

