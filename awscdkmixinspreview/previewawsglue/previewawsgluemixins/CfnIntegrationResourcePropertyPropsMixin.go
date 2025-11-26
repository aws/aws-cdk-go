package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::IntegrationResourceProperty` resource type can be used to setup `ResourceProperty` of the AWS Glue connection (for the SaaS source), DynamoDB Database (for DynamoDB source), or AWS Glue database ARN (for the target).
//
// ResourceProperty is used to define the properties requires to setup the integration, including the role to access the connection or database, KMS keys, event bus for event notifications and VPC connection. To set both source and target properties the same API needs to be invoked twice, once with the AWS Glue connection ARN as ResourceArn with SourceProcessingProperties and next, with the AWS Glue database ARN as ResourceArn with TargetProcessingProperties respectively.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIntegrationResourcePropertyPropsMixin := awscdkmixinspreview.Mixins.NewCfnIntegrationResourcePropertyPropsMixin(&CfnIntegrationResourcePropertyMixinProps{
//   	ResourceArn: jsii.String("resourceArn"),
//   	SourceProcessingProperties: &SourceProcessingPropertiesProperty{
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetProcessingProperties: &TargetProcessingPropertiesProperty{
//   		ConnectionName: jsii.String("connectionName"),
//   		EventBusArn: jsii.String("eventBusArn"),
//   		KmsArn: jsii.String("kmsArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integrationresourceproperty.html
//
type CfnIntegrationResourcePropertyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIntegrationResourcePropertyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIntegrationResourcePropertyPropsMixin
type jsiiProxy_CfnIntegrationResourcePropertyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIntegrationResourcePropertyPropsMixin) Props() *CfnIntegrationResourcePropertyMixinProps {
	var returns *CfnIntegrationResourcePropertyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResourcePropertyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::IntegrationResourceProperty`.
func NewCfnIntegrationResourcePropertyPropsMixin(props *CfnIntegrationResourcePropertyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIntegrationResourcePropertyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIntegrationResourcePropertyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIntegrationResourcePropertyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationResourcePropertyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::IntegrationResourceProperty`.
func NewCfnIntegrationResourcePropertyPropsMixin_Override(c CfnIntegrationResourcePropertyPropsMixin, props *CfnIntegrationResourcePropertyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationResourcePropertyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIntegrationResourcePropertyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegrationResourcePropertyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationResourcePropertyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegrationResourcePropertyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnIntegrationResourcePropertyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIntegrationResourcePropertyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegrationResourcePropertyPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

