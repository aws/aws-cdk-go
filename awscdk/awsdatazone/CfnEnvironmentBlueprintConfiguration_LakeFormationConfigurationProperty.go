package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lakeFormationConfigurationProperty := &LakeFormationConfigurationProperty{
//   	LocationRegistrationExcludeS3Locations: []*string{
//   		jsii.String("locationRegistrationExcludeS3Locations"),
//   	},
//   	LocationRegistrationRole: jsii.String("locationRegistrationRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentblueprintconfiguration-lakeformationconfiguration.html
//
type CfnEnvironmentBlueprintConfiguration_LakeFormationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentblueprintconfiguration-lakeformationconfiguration.html#cfn-datazone-environmentblueprintconfiguration-lakeformationconfiguration-locationregistrationexcludes3locations
	//
	LocationRegistrationExcludeS3Locations *[]*string `field:"optional" json:"locationRegistrationExcludeS3Locations" yaml:"locationRegistrationExcludeS3Locations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentblueprintconfiguration-lakeformationconfiguration.html#cfn-datazone-environmentblueprintconfiguration-lakeformationconfiguration-locationregistrationrole
	//
	LocationRegistrationRole *string `field:"optional" json:"locationRegistrationRole" yaml:"locationRegistrationRole"`
}

