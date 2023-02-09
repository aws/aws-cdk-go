package awsiotcoredeviceadvisor

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSuiteDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var suiteDefinitionConfiguration interface{}
//
//   cfnSuiteDefinitionProps := &cfnSuiteDefinitionProps{
//   	suiteDefinitionConfiguration: suiteDefinitionConfiguration,
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSuiteDefinitionProps struct {
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
	SuiteDefinitionConfiguration interface{} `field:"required" json:"suiteDefinitionConfiguration" yaml:"suiteDefinitionConfiguration"`
	// Metadata that can be used to manage the the Suite Definition.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

