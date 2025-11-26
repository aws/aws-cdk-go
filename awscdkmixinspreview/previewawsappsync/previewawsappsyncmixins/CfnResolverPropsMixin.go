package previewawsappsyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappsync/previewawsappsyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppSync::Resolver` resource defines the logical GraphQL resolver that you attach to fields in a schema.
//
// Request and response templates for resolvers are written in Apache Velocity Template Language (VTL) format. For more information about resolvers, see [Resolver Mapping Template Reference](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html) .
//
// > When you submit an update, AWS CloudFormation updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CloudFormation template. Changing the Amazon S3 file content without changing a property value will not result in an update operation.
// >
// > See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *AWS CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResolverPropsMixin := awscdkmixinspreview.Mixins.NewCfnResolverPropsMixin(&CfnResolverMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	CachingConfig: &CachingConfigProperty{
//   		CachingKeys: []*string{
//   			jsii.String("cachingKeys"),
//   		},
//   		Ttl: jsii.Number(123),
//   	},
//   	Code: jsii.String("code"),
//   	CodeS3Location: jsii.String("codeS3Location"),
//   	DataSourceName: jsii.String("dataSourceName"),
//   	FieldName: jsii.String("fieldName"),
//   	Kind: jsii.String("kind"),
//   	MaxBatchSize: jsii.Number(123),
//   	MetricsConfig: jsii.String("metricsConfig"),
//   	PipelineConfig: &PipelineConfigProperty{
//   		Functions: []*string{
//   			jsii.String("functions"),
//   		},
//   	},
//   	RequestMappingTemplate: jsii.String("requestMappingTemplate"),
//   	RequestMappingTemplateS3Location: jsii.String("requestMappingTemplateS3Location"),
//   	ResponseMappingTemplate: jsii.String("responseMappingTemplate"),
//   	ResponseMappingTemplateS3Location: jsii.String("responseMappingTemplateS3Location"),
//   	Runtime: &AppSyncRuntimeProperty{
//   		Name: jsii.String("name"),
//   		RuntimeVersion: jsii.String("runtimeVersion"),
//   	},
//   	SyncConfig: &SyncConfigProperty{
//   		ConflictDetection: jsii.String("conflictDetection"),
//   		ConflictHandler: jsii.String("conflictHandler"),
//   		LambdaConflictHandlerConfig: &LambdaConflictHandlerConfigProperty{
//   			LambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   		},
//   	},
//   	TypeName: jsii.String("typeName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html
//
type CfnResolverPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResolverMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResolverPropsMixin
type jsiiProxy_CfnResolverPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResolverPropsMixin) Props() *CfnResolverMixinProps {
	var returns *CfnResolverMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppSync::Resolver`.
func NewCfnResolverPropsMixin(props *CfnResolverMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResolverPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResolverPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResolverPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnResolverPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppSync::Resolver`.
func NewCfnResolverPropsMixin_Override(c CfnResolverPropsMixin, props *CfnResolverMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnResolverPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResolverPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResolverPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnResolverPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnResolverPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnResolverPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

