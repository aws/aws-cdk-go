package mixinsawsosis

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsosis/mixinsawsosis/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::OSIS::Pipeline resource creates an Amazon OpenSearch Ingestion pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnPipelinePropsMixin := awscdkmixinspreview.Mixins.NewCfnPipelinePropsMixin(&CfnPipelineMixinProps{
//   	BufferOptions: &BufferOptionsProperty{
//   		PersistentBufferEnabled: jsii.Boolean(false),
//   	},
//   	EncryptionAtRestOptions: &EncryptionAtRestOptionsProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	LogPublishingOptions: &LogPublishingOptionsProperty{
//   		CloudWatchLogDestination: &CloudWatchLogDestinationProperty{
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		IsLoggingEnabled: jsii.Boolean(false),
//   	},
//   	MaxUnits: jsii.Number(123),
//   	MinUnits: jsii.Number(123),
//   	PipelineConfigurationBody: jsii.String("pipelineConfigurationBody"),
//   	PipelineName: jsii.String("pipelineName"),
//   	PipelineRoleArn: jsii.String("pipelineRoleArn"),
//   	ResourcePolicy: &ResourcePolicyProperty{
//   		Policy: policy,
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcOptions: &VpcOptionsProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcAttachmentOptions: &VpcAttachmentOptionsProperty{
//   			AttachToVpc: jsii.Boolean(false),
//   			CidrBlock: jsii.String("cidrBlock"),
//   		},
//   		VpcEndpointManagement: jsii.String("vpcEndpointManagement"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html
//
type CfnPipelinePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPipelineMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPipelinePropsMixin
type jsiiProxy_CfnPipelinePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Props() *CfnPipelineMixinProps {
	var returns *CfnPipelineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OSIS::Pipeline`.
func NewCfnPipelinePropsMixin(props *CfnPipelineMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPipelinePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPipelinePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipelinePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OSIS::Pipeline`.
func NewCfnPipelinePropsMixin_Override(c CfnPipelinePropsMixin, props *CfnPipelineMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPipelinePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipelinePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipelinePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_osis.mixins.CfnPipelinePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipelinePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPipelinePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

