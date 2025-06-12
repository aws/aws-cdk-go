package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisioningConfigurationProperty := &ProvisioningConfigurationProperty{
//   	LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   		LocationRegistrationExcludeS3Locations: []*string{
//   			jsii.String("locationRegistrationExcludeS3Locations"),
//   		},
//   		LocationRegistrationRole: jsii.String("locationRegistrationRole"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentblueprintconfiguration-provisioningconfiguration.html
//
type CfnEnvironmentBlueprintConfiguration_ProvisioningConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentblueprintconfiguration-provisioningconfiguration.html#cfn-datazone-environmentblueprintconfiguration-provisioningconfiguration-lakeformationconfiguration
	//
	LakeFormationConfiguration interface{} `field:"required" json:"lakeFormationConfiguration" yaml:"lakeFormationConfiguration"`
}

