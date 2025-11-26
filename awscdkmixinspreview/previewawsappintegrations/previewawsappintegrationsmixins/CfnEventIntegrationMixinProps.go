package previewawsappintegrationsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEventIntegrationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventIntegrationMixinProps := &CfnEventIntegrationMixinProps{
//   	Description: jsii.String("description"),
//   	EventBridgeBus: jsii.String("eventBridgeBus"),
//   	EventFilter: &EventFilterProperty{
//   		Source: jsii.String("source"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-eventintegration.html
//
type CfnEventIntegrationMixinProps struct {
	// The event integration description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-eventintegration.html#cfn-appintegrations-eventintegration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon EventBridge bus for the event integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-eventintegration.html#cfn-appintegrations-eventintegration-eventbridgebus
	//
	EventBridgeBus *string `field:"optional" json:"eventBridgeBus" yaml:"eventBridgeBus"`
	// The event integration filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-eventintegration.html#cfn-appintegrations-eventintegration-eventfilter
	//
	EventFilter interface{} `field:"optional" json:"eventFilter" yaml:"eventFilter"`
	// The name of the event integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-eventintegration.html#cfn-appintegrations-eventintegration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-eventintegration.html#cfn-appintegrations-eventintegration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

