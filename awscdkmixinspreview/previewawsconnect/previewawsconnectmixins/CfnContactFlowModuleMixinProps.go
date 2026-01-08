package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnContactFlowModulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContactFlowModuleMixinProps := &CfnContactFlowModuleMixinProps{
//   	Content: jsii.String("content"),
//   	Description: jsii.String("description"),
//   	ExternalInvocationConfiguration: &ExternalInvocationConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	Settings: jsii.String("settings"),
//   	State: jsii.String("state"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html
//
type CfnContactFlowModuleMixinProps struct {
	// The content of the flow module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html#cfn-connect-contactflowmodule-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The description of the flow module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html#cfn-connect-contactflowmodule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The external invocation configuration for the flow module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html#cfn-connect-contactflowmodule-externalinvocationconfiguration
	//
	ExternalInvocationConfiguration interface{} `field:"optional" json:"externalInvocationConfiguration" yaml:"externalInvocationConfiguration"`
	// The Amazon Resource Name (ARN) of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html#cfn-connect-contactflowmodule-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The name of the flow module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html#cfn-connect-contactflowmodule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The configuration settings for the flow module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html#cfn-connect-contactflowmodule-settings
	//
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
	// The state of the flow module.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html#cfn-connect-contactflowmodule-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodule.html#cfn-connect-contactflowmodule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

