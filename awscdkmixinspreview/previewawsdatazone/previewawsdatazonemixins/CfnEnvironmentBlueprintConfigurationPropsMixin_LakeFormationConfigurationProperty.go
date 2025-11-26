package previewawsdatazonemixins


// The Lake Formation configuration of the Data Lake blueprint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnEnvironmentBlueprintConfigurationPropsMixin_LakeFormationConfigurationProperty struct {
	// Specifies certain Amazon S3 locations if you do not want Amazon DataZone to automatically register them in hybrid mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentblueprintconfiguration-lakeformationconfiguration.html#cfn-datazone-environmentblueprintconfiguration-lakeformationconfiguration-locationregistrationexcludes3locations
	//
	LocationRegistrationExcludeS3Locations *[]*string `field:"optional" json:"locationRegistrationExcludeS3Locations" yaml:"locationRegistrationExcludeS3Locations"`
	// The role that is used to manage read/write access to the chosen Amazon S3 bucket(s) for Data Lake using AWS Lake Formation hybrid access mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentblueprintconfiguration-lakeformationconfiguration.html#cfn-datazone-environmentblueprintconfiguration-lakeformationconfiguration-locationregistrationrole
	//
	LocationRegistrationRole *string `field:"optional" json:"locationRegistrationRole" yaml:"locationRegistrationRole"`
}

