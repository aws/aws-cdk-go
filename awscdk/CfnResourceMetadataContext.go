package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Manages resource-level CloudFormation Metadata Context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var constructSelector IConstructSelector
//
//   cfnResourceMetadataContext := cdk.CfnResourceMetadataContext_Of(this, &CfnResourceMetadataContextOptions{
//   	Selector: constructSelector,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnResourceMetadataContext interface {
	// Removes Context from the selected resources.
	Clear()
	// Replaces the complete Context value on the selected resources.
	//
	// Includes resources added after this call.
	Set(context *CfnResourceMetadataContextProps)
}

// The jsii proxy struct for CfnResourceMetadataContext
type jsiiProxy_CfnResourceMetadataContext struct {
	_ byte // padding
}

// Returns the resource context API for a construct scope.
func CfnResourceMetadataContext_Of(scope constructs.IConstruct, options *CfnResourceMetadataContextOptions) CfnResourceMetadataContext {
	_init_.Initialize()

	if err := validateCfnResourceMetadataContext_OfParameters(scope, options); err != nil {
		panic(err)
	}
	var returns CfnResourceMetadataContext

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnResourceMetadataContext",
		"of",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceMetadataContext) Clear() {
	_jsii_.InvokeVoid(
		c,
		"clear",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResourceMetadataContext) Set(context *CfnResourceMetadataContextProps) {
	if err := c.validateSetParameters(context); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"set",
		[]interface{}{context},
	)
}

