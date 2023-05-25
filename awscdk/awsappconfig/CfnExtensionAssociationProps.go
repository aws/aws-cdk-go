package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnExtensionAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExtensionAssociationProps := &CfnExtensionAssociationProps{
//   	ExtensionIdentifier: jsii.String("extensionIdentifier"),
//   	ExtensionVersionNumber: jsii.Number(123),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnExtensionAssociationProps struct {
	// The name, the ID, or the Amazon Resource Name (ARN) of the extension.
	ExtensionIdentifier *string `field:"optional" json:"extensionIdentifier" yaml:"extensionIdentifier"`
	// The version number of the extension.
	//
	// If not specified, AWS AppConfig uses the maximum version of the extension.
	ExtensionVersionNumber *float64 `field:"optional" json:"extensionVersionNumber" yaml:"extensionVersionNumber"`
	// The parameter names and values defined in the extensions.
	//
	// Extension parameters marked `Required` must be entered for this field.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The ARN of an application, configuration profile, or environment.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Adds one or more tags for the specified extension association.
	//
	// Tags are metadata that help you categorize resources in different ways, for example, by purpose, owner, or environment. Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

