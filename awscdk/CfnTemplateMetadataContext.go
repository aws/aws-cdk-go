package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Manages template-level CloudFormation Metadata Context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack Stack
//
//   cfnTemplateMetadataContext := cdk.CfnTemplateMetadataContext_Of(stack)
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnTemplateMetadataContext interface {
	// Removes the template Context value.
	Clear()
	// Replaces the complete template Context value.
	//
	// Values must be concrete because CloudFormation does not resolve intrinsic
	// functions in template metadata.
	Set(context *CfnTemplateMetadataContextProps)
}

// The jsii proxy struct for CfnTemplateMetadataContext
type jsiiProxy_CfnTemplateMetadataContext struct {
	_ byte // padding
}

// Returns the template context API for a stack.
func CfnTemplateMetadataContext_Of(stack Stack) CfnTemplateMetadataContext {
	_init_.Initialize()

	if err := validateCfnTemplateMetadataContext_OfParameters(stack); err != nil {
		panic(err)
	}
	var returns CfnTemplateMetadataContext

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnTemplateMetadataContext",
		"of",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplateMetadataContext) Clear() {
	_jsii_.InvokeVoid(
		c,
		"clear",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTemplateMetadataContext) Set(context *CfnTemplateMetadataContextProps) {
	if err := c.validateSetParameters(context); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"set",
		[]interface{}{context},
	)
}

