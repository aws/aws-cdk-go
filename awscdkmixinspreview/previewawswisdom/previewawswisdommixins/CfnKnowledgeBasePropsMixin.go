package previewawswisdommixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswisdom/previewawswisdommixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBasePropsMixin := awscdkmixinspreview.Mixins.NewCfnKnowledgeBasePropsMixin(&CfnKnowledgeBaseMixinProps{
//   	Description: jsii.String("description"),
//   	KnowledgeBaseType: jsii.String("knowledgeBaseType"),
//   	Name: jsii.String("name"),
//   	RenderingConfiguration: &RenderingConfigurationProperty{
//   		TemplateUri: jsii.String("templateUri"),
//   	},
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	SourceConfiguration: &SourceConfigurationProperty{
//   		AppIntegrations: &AppIntegrationsConfigurationProperty{
//   			AppIntegrationArn: jsii.String("appIntegrationArn"),
//   			ObjectFields: []*string{
//   				jsii.String("objectFields"),
//   			},
//   		},
//   		ManagedSourceConfiguration: &ManagedSourceConfigurationProperty{
//   			WebCrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   				CrawlerLimits: &CrawlerLimitsProperty{
//   					RateLimit: jsii.Number(123),
//   				},
//   				ExclusionFilters: []*string{
//   					jsii.String("exclusionFilters"),
//   				},
//   				InclusionFilters: []*string{
//   					jsii.String("inclusionFilters"),
//   				},
//   				Scope: jsii.String("scope"),
//   				UrlConfiguration: &UrlConfigurationProperty{
//   					SeedUrls: []interface{}{
//   						&SeedUrlProperty{
//   							Url: jsii.String("url"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VectorIngestionConfiguration: &VectorIngestionConfigurationProperty{
//   		ChunkingConfiguration: &ChunkingConfigurationProperty{
//   			ChunkingStrategy: jsii.String("chunkingStrategy"),
//   			FixedSizeChunkingConfiguration: &FixedSizeChunkingConfigurationProperty{
//   				MaxTokens: jsii.Number(123),
//   				OverlapPercentage: jsii.Number(123),
//   			},
//   			HierarchicalChunkingConfiguration: &HierarchicalChunkingConfigurationProperty{
//   				LevelConfigurations: []interface{}{
//   					&HierarchicalChunkingLevelConfigurationProperty{
//   						MaxTokens: jsii.Number(123),
//   					},
//   				},
//   				OverlapTokens: jsii.Number(123),
//   			},
//   			SemanticChunkingConfiguration: &SemanticChunkingConfigurationProperty{
//   				BreakpointPercentileThreshold: jsii.Number(123),
//   				BufferSize: jsii.Number(123),
//   				MaxTokens: jsii.Number(123),
//   			},
//   		},
//   		ParsingConfiguration: &ParsingConfigurationProperty{
//   			BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   				ModelArn: jsii.String("modelArn"),
//   				ParsingPrompt: &ParsingPromptProperty{
//   					ParsingPromptText: jsii.String("parsingPromptText"),
//   				},
//   			},
//   			ParsingStrategy: jsii.String("parsingStrategy"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html
//
type CfnKnowledgeBasePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnKnowledgeBaseMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnKnowledgeBasePropsMixin
type jsiiProxy_CfnKnowledgeBasePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnKnowledgeBasePropsMixin) Props() *CfnKnowledgeBaseMixinProps {
	var returns *CfnKnowledgeBaseMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKnowledgeBasePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Wisdom::KnowledgeBase`.
func NewCfnKnowledgeBasePropsMixin(props *CfnKnowledgeBaseMixinProps, options *mixins.CfnPropertyMixinOptions) CfnKnowledgeBasePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnKnowledgeBasePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnKnowledgeBasePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnKnowledgeBasePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Wisdom::KnowledgeBase`.
func NewCfnKnowledgeBasePropsMixin_Override(c CfnKnowledgeBasePropsMixin, props *CfnKnowledgeBaseMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnKnowledgeBasePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnKnowledgeBasePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKnowledgeBasePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnKnowledgeBasePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKnowledgeBasePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnKnowledgeBasePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnKnowledgeBasePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnKnowledgeBasePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

