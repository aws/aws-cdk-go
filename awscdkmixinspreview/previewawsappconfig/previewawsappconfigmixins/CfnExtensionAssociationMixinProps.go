package previewawsappconfigmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnExtensionAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnExtensionAssociationMixinProps := &CfnExtensionAssociationMixinProps{
//   	ExtensionIdentifier: jsii.String("extensionIdentifier"),
//   	ExtensionVersionNumber: jsii.Number(123),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extensionassociation.html
//
type CfnExtensionAssociationMixinProps struct {
	// The name, the ID, or the Amazon Resource Name (ARN) of the extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extensionassociation.html#cfn-appconfig-extensionassociation-extensionidentifier
	//
	ExtensionIdentifier *string `field:"optional" json:"extensionIdentifier" yaml:"extensionIdentifier"`
	// The version number of the extension.
	//
	// If not specified, AWS AppConfig uses the maximum version of the extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extensionassociation.html#cfn-appconfig-extensionassociation-extensionversionnumber
	//
	ExtensionVersionNumber *float64 `field:"optional" json:"extensionVersionNumber" yaml:"extensionVersionNumber"`
	// The parameter names and values defined in the extensions.
	//
	// Extension parameters marked `Required` must be entered for this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extensionassociation.html#cfn-appconfig-extensionassociation-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The ARN of an application, configuration profile, or environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extensionassociation.html#cfn-appconfig-extensionassociation-resourceidentifier
	//
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Adds one or more tags for the specified extension association.
	//
	// Tags are metadata that help you categorize resources in different ways, for example, by purpose, owner, or environment. Each tag consists of a key and an optional value, both of which you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extensionassociation.html#cfn-appconfig-extensionassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

