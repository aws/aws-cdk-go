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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html
//
type CfnSuiteDefinition_SuiteDefinitionConfigurationProperty struct {
	// The device permission role arn of the test suite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-devicepermissionrolearn
	//
	DevicePermissionRoleArn *string `field:"required" json:"devicePermissionRoleArn" yaml:"devicePermissionRoleArn"`
	// The root group of the test suite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-rootgroup
	//
	RootGroup *string `field:"required" json:"rootGroup" yaml:"rootGroup"`
	// The devices being tested in the test suite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-devices
	//
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Whether the tests are intended for qualification in a suite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-intendedforqualification
	//
	IntendedForQualification interface{} `field:"optional" json:"intendedForQualification" yaml:"intendedForQualification"`
	// The Name of the suite definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-suitedefinitionname
	//
	SuiteDefinitionName *string `field:"optional" json:"suiteDefinitionName" yaml:"suiteDefinitionName"`
}

