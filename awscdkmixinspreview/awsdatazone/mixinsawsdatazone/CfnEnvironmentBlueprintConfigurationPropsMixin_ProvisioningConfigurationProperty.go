package mixinsawsdatazone


// The provisioning configuration of the blueprint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnEnvironmentBlueprintConfigurationPropsMixin_ProvisioningConfigurationProperty struct {
	// The Lake Formation configuration of the Data Lake blueprint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentblueprintconfiguration-provisioningconfiguration.html#cfn-datazone-environmentblueprintconfiguration-provisioningconfiguration-lakeformationconfiguration
	//
	LakeFormationConfiguration interface{} `field:"optional" json:"lakeFormationConfiguration" yaml:"lakeFormationConfiguration"`
}

