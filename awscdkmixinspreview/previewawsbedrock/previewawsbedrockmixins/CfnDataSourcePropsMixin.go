package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > Properties with `__Update requires: Replacement__` can result in the creation of a new data source and deletion of the old one.
//
// This can happen if you also change the Name of the data source.
//
// Specifies a data source as a resource in a top-level template. Minimally, you must specify the following properties:
//
// - Name – Specify a name for the data source.
// - KnowledgeBaseId – Specify the ID of the knowledge base for the data source to belong to.
// - DataSourceConfiguration – Specify information about the Amazon S3 bucket containing the data source. The following sub-properties are required:
//
// - Type – Specify the value `S3` .
//
// For more information about setting up data sources in Amazon Bedrock , see [Set up a data source for your knowledge base](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-ds.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataSourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnDataSourcePropsMixin(&CfnDataSourceMixinProps{
//   	DataDeletionPolicy: jsii.String("dataDeletionPolicy"),
//   	DataSourceConfiguration: &DataSourceConfigurationProperty{
//   		ConfluenceConfiguration: &ConfluenceDataSourceConfigurationProperty{
//   			CrawlerConfiguration: &ConfluenceCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   								ObjectType: jsii.String("objectType"),
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			SourceConfiguration: &ConfluenceSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				HostType: jsii.String("hostType"),
//   				HostUrl: jsii.String("hostUrl"),
//   			},
//   		},
//   		S3Configuration: &S3DataSourceConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   			InclusionPrefixes: []*string{
//   				jsii.String("inclusionPrefixes"),
//   			},
//   		},
//   		SalesforceConfiguration: &SalesforceDataSourceConfigurationProperty{
//   			CrawlerConfiguration: &SalesforceCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   								ObjectType: jsii.String("objectType"),
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			SourceConfiguration: &SalesforceSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				HostUrl: jsii.String("hostUrl"),
//   			},
//   		},
//   		SharePointConfiguration: &SharePointDataSourceConfigurationProperty{
//   			CrawlerConfiguration: &SharePointCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   								ObjectType: jsii.String("objectType"),
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			SourceConfiguration: &SharePointSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				Domain: jsii.String("domain"),
//   				HostType: jsii.String("hostType"),
//   				SiteUrls: []*string{
//   					jsii.String("siteUrls"),
//   				},
//   				TenantId: jsii.String("tenantId"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   		WebConfiguration: &WebDataSourceConfigurationProperty{
//   			CrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   				CrawlerLimits: &WebCrawlerLimitsProperty{
//   					MaxPages: jsii.Number(123),
//   					RateLimit: jsii.Number(123),
//   				},
//   				ExclusionFilters: []*string{
//   					jsii.String("exclusionFilters"),
//   				},
//   				InclusionFilters: []*string{
//   					jsii.String("inclusionFilters"),
//   				},
//   				Scope: jsii.String("scope"),
//   				UserAgent: jsii.String("userAgent"),
//   				UserAgentHeader: jsii.String("userAgentHeader"),
//   			},
//   			SourceConfiguration: &WebSourceConfigurationProperty{
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
//   	Description: jsii.String("description"),
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	Name: jsii.String("name"),
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
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
//   		ContextEnrichmentConfiguration: &ContextEnrichmentConfigurationProperty{
//   			BedrockFoundationModelConfiguration: &BedrockFoundationModelContextEnrichmentConfigurationProperty{
//   				EnrichmentStrategyConfiguration: &EnrichmentStrategyConfigurationProperty{
//   					Method: jsii.String("method"),
//   				},
//   				ModelArn: jsii.String("modelArn"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		CustomTransformationConfiguration: &CustomTransformationConfigurationProperty{
//   			IntermediateStorage: &IntermediateStorageProperty{
//   				S3Location: &S3LocationProperty{
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   			Transformations: []interface{}{
//   				&TransformationProperty{
//   					StepToApply: jsii.String("stepToApply"),
//   					TransformationFunction: &TransformationFunctionProperty{
//   						TransformationLambdaConfiguration: &TransformationLambdaConfigurationProperty{
//   							LambdaArn: jsii.String("lambdaArn"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		ParsingConfiguration: &ParsingConfigurationProperty{
//   			BedrockDataAutomationConfiguration: &BedrockDataAutomationConfigurationProperty{
//   				ParsingModality: jsii.String("parsingModality"),
//   			},
//   			BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   				ModelArn: jsii.String("modelArn"),
//   				ParsingModality: jsii.String("parsingModality"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html
//
type CfnDataSourcePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataSourceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataSourcePropsMixin
type jsiiProxy_CfnDataSourcePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataSourcePropsMixin) Props() *CfnDataSourceMixinProps {
	var returns *CfnDataSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSourcePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::DataSource`.
func NewCfnDataSourcePropsMixin(props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::DataSource`.
func NewCfnDataSourcePropsMixin_Override(c CfnDataSourcePropsMixin, props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnDataSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnDataSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSourcePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDataSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

