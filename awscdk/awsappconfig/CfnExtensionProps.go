package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnExtension`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var actions interface{}
//
//   cfnExtensionProps := &CfnExtensionProps{
//   	Actions: actions,
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	LatestVersionNumber: jsii.Number(123),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &ParameterProperty{
//   			"required": jsii.Boolean(false),
//
//   			// the properties below are optional
//   			"description": jsii.String("description"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extension.html
//
type CfnExtensionProps struct {
	// The actions defined in the extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extension.html#cfn-appconfig-extension-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// A name for the extension.
	//
	// Each extension name in your account must be unique. Extension versions use the same name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extension.html#cfn-appconfig-extension-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Information about the extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extension.html#cfn-appconfig-extension-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// You can omit this field when you create an extension.
	//
	// When you create a new version, specify the most recent current version number. For example, you create version 3, enter 2 for this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extension.html#cfn-appconfig-extension-latestversionnumber
	//
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// The parameters accepted by the extension.
	//
	// You specify parameter values when you associate the extension to an AWS AppConfig resource by using the `CreateExtensionAssociation` API action. For AWS Lambda extension actions, these parameters are included in the Lambda request object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extension.html#cfn-appconfig-extension-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Adds one or more tags for the specified extension.
	//
	// Tags are metadata that help you categorize resources in different ways, for example, by purpose, owner, or environment. Each tag consists of a key and an optional value, both of which you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extension.html#cfn-appconfig-extension-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

