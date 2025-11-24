package mixinsawsiotcoredeviceadvisor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSuiteDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var suiteDefinitionConfiguration interface{}
//
//   cfnSuiteDefinitionMixinProps := &CfnSuiteDefinitionMixinProps{
//   	SuiteDefinitionConfiguration: suiteDefinitionConfiguration,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotcoredeviceadvisor-suitedefinition.html
//
type CfnSuiteDefinitionMixinProps struct {
	// Gets the suite definition configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotcoredeviceadvisor-suitedefinition.html#cfn-iotcoredeviceadvisor-suitedefinition-suitedefinitionconfiguration
	//
	SuiteDefinitionConfiguration interface{} `field:"optional" json:"suiteDefinitionConfiguration" yaml:"suiteDefinitionConfiguration"`
	// Metadata that can be used to manage the the Suite Definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotcoredeviceadvisor-suitedefinition.html#cfn-iotcoredeviceadvisor-suitedefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

