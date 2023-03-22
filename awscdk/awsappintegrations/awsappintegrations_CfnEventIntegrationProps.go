package awsappintegrations

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEventIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventIntegrationProps := &cfnEventIntegrationProps{
//   	eventBridgeBus: jsii.String("eventBridgeBus"),
//   	eventFilter: &eventFilterProperty{
//   		source: jsii.String("source"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEventIntegrationProps struct {
	// The Amazon EventBridge bus for the event integration.
	EventBridgeBus *string `field:"required" json:"eventBridgeBus" yaml:"eventBridgeBus"`
	// The event integration filter.
	EventFilter interface{} `field:"required" json:"eventFilter" yaml:"eventFilter"`
	// The name of the event integration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The event integration description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

