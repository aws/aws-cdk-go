package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::DevEndpoint` resource specifies a development endpoint where a developer can remotely debug ETL scripts for AWS Glue .
//
// For more information, see [DevEndpoint Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-dev-endpoint.html#aws-glue-api-jobs-dev-endpoint-DevEndpoint) in the AWS Glue Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var arguments_ interface{}
//   var tags interface{}
//
//   cfnDevEndpointPropsMixin := awscdkmixinspreview.Mixins.NewCfnDevEndpointPropsMixin(&CfnDevEndpointMixinProps{
//   	Arguments: arguments_,
//   	EndpointName: jsii.String("endpointName"),
//   	ExtraJarsS3Path: jsii.String("extraJarsS3Path"),
//   	ExtraPythonLibsS3Path: jsii.String("extraPythonLibsS3Path"),
//   	GlueVersion: jsii.String("glueVersion"),
//   	NumberOfNodes: jsii.Number(123),
//   	NumberOfWorkers: jsii.Number(123),
//   	PublicKey: jsii.String("publicKey"),
//   	PublicKeys: []*string{
//   		jsii.String("publicKeys"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: tags,
//   	WorkerType: jsii.String("workerType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html
//
type CfnDevEndpointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDevEndpointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDevEndpointPropsMixin
type jsiiProxy_CfnDevEndpointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDevEndpointPropsMixin) Props() *CfnDevEndpointMixinProps {
	var returns *CfnDevEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::DevEndpoint`.
func NewCfnDevEndpointPropsMixin(props *CfnDevEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDevEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDevEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDevEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDevEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::DevEndpoint`.
func NewCfnDevEndpointPropsMixin_Override(c CfnDevEndpointPropsMixin, props *CfnDevEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDevEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDevEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDevEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDevEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDevEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnDevEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDevEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDevEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

