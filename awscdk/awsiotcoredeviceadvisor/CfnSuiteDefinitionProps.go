package awsiotcoredeviceadvisor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   cfnSuiteDefinitionProps := &CfnSuiteDefinitionProps{
//   	SuiteDefinitionConfiguration: suiteDefinitionConfiguration,
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotcoredeviceadvisor-suitedefinition.html
//
type CfnSuiteDefinitionProps struct {
	// Gets the suite definition configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotcoredeviceadvisor-suitedefinition.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration
	//
	SuiteDefinitionConfiguration interface{} `field:"required" json:"suiteDefinitionConfiguration" yaml:"suiteDefinitionConfiguration"`
	// Metadata that can be used to manage the the Suite Definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotcoredeviceadvisor-suitedefinition.html#cfn-iotcoredeviceadvisor-suitedefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

