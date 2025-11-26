package previewawsiotcoredeviceadvisormixins


// The configuration of the Suite Definition. Listed below are the required elements of the `SuiteDefinitionConfiguration` .
//
// - ***devicePermissionRoleArn*** - The device permission arn.
//
// This is a required element.
//
// *Type:* String
// - ***devices*** - The list of configured devices under test. For more information on devices under test, see [DeviceUnderTest](https://docs.aws.amazon.com/iot/latest/apireference/API_iotdeviceadvisor_DeviceUnderTest.html)
//
// Not a required element.
//
// *Type:* List of devices under test
// - ***intendedForQualification*** - The tests intended for qualification in a suite.
//
// Not a required element.
//
// *Type:* Boolean
// - ***rootGroup*** - The test suite root group. For more information on creating and using root groups see the [Device Advisor workflow](https://docs.aws.amazon.com/iot/latest/developerguide/device-advisor-workflow.html) .
//
// This is a required element.
//
// *Type:* String
// - ***suiteDefinitionName*** - The Suite Definition Configuration name.
//
// This is a required element.
//
// *Type:* String.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   suiteDefinitionConfigurationProperty := &SuiteDefinitionConfigurationProperty{
//   	DevicePermissionRoleArn: jsii.String("devicePermissionRoleArn"),
//   	Devices: []interface{}{
//   		&DeviceUnderTestProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			ThingArn: jsii.String("thingArn"),
//   		},
//   	},
//   	IntendedForQualification: jsii.Boolean(false),
//   	RootGroup: jsii.String("rootGroup"),
//   	SuiteDefinitionName: jsii.String("suiteDefinitionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html
//
type CfnSuiteDefinitionPropsMixin_SuiteDefinitionConfigurationProperty struct {
	// Gets the device permission ARN.
	//
	// This is a required parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-devicepermissionrolearn
	//
	DevicePermissionRoleArn *string `field:"optional" json:"devicePermissionRoleArn" yaml:"devicePermissionRoleArn"`
	// Gets the devices configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-devices
	//
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Gets the tests intended for qualification in a suite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-intendedforqualification
	//
	IntendedForQualification interface{} `field:"optional" json:"intendedForQualification" yaml:"intendedForQualification"`
	// Gets the test suite root group.
	//
	// This is a required parameter. For updating or creating the latest qualification suite, if `intendedForQualification` is set to true, `rootGroup` can be an empty string. If `intendedForQualification` is false, `rootGroup` cannot be an empty string. If `rootGroup` is empty, and `intendedForQualification` is set to true, all the qualification tests are included, and the configuration is default.
	//
	// For a qualification suite, the minimum length is 0, and the maximum is 2048. For a non-qualification suite, the minimum length is 1, and the maximum is 2048.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-rootgroup
	//
	RootGroup *string `field:"optional" json:"rootGroup" yaml:"rootGroup"`
	// Gets the suite definition name.
	//
	// This is a required parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration-suitedefinitionname
	//
	SuiteDefinitionName *string `field:"optional" json:"suiteDefinitionName" yaml:"suiteDefinitionName"`
}

