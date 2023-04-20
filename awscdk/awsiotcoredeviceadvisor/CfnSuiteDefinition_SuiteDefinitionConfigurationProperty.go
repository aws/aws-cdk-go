package awsiotcoredeviceadvisor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   suiteDefinitionConfigurationProperty := &SuiteDefinitionConfigurationProperty{
//   	DevicePermissionRoleArn: jsii.String("devicePermissionRoleArn"),
//   	RootGroup: jsii.String("rootGroup"),
//
//   	// the properties below are optional
//   	Devices: []interface{}{
//   		&DeviceUnderTestProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			ThingArn: jsii.String("thingArn"),
//   		},
//   	},
//   	IntendedForQualification: jsii.Boolean(false),
//   	SuiteDefinitionName: jsii.String("suiteDefinitionName"),
//   }
//
type CfnSuiteDefinition_SuiteDefinitionConfigurationProperty struct {
	// `CfnSuiteDefinition.SuiteDefinitionConfigurationProperty.DevicePermissionRoleArn`.
	DevicePermissionRoleArn *string `field:"required" json:"devicePermissionRoleArn" yaml:"devicePermissionRoleArn"`
	// `CfnSuiteDefinition.SuiteDefinitionConfigurationProperty.RootGroup`.
	RootGroup *string `field:"required" json:"rootGroup" yaml:"rootGroup"`
	// `CfnSuiteDefinition.SuiteDefinitionConfigurationProperty.Devices`.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// `CfnSuiteDefinition.SuiteDefinitionConfigurationProperty.IntendedForQualification`.
	IntendedForQualification interface{} `field:"optional" json:"intendedForQualification" yaml:"intendedForQualification"`
	// `CfnSuiteDefinition.SuiteDefinitionConfigurationProperty.SuiteDefinitionName`.
	SuiteDefinitionName *string `field:"optional" json:"suiteDefinitionName" yaml:"suiteDefinitionName"`
}

