package mixinsawskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awskendra/mixinsawskendra/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a data source connector that you want to use with an Amazon Kendra index.
//
// You specify a name, data source connector type and description for your data source. You also specify configuration information for the data source connector.
//
// > `CreateDataSource` does *not* support connectors which [require a `TemplateConfiguration` object](https://docs.aws.amazon.com/kendra/latest/dg/ds-schemas.html) for connecting to Amazon Kendra .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataSourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnDataSourcePropsMixin(&CfnDataSourceMixinProps{
//   	CustomDocumentEnrichmentConfiguration: &CustomDocumentEnrichmentConfigurationProperty{
//   		InlineConfigurations: []interface{}{
//   			&InlineCustomDocumentEnrichmentConfigurationProperty{
//   				Condition: &DocumentAttributeConditionProperty{
//   					ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   					ConditionOnValue: &DocumentAttributeValueProperty{
//   						DateValue: jsii.String("dateValue"),
//   						LongValue: jsii.Number(123),
//   						StringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   					},
//   					Operator: jsii.String("operator"),
//   				},
//   				DocumentContentDeletion: jsii.Boolean(false),
//   				Target: &DocumentAttributeTargetProperty{
//   					TargetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//   					TargetDocumentAttributeValue: &DocumentAttributeValueProperty{
//   						DateValue: jsii.String("dateValue"),
//   						LongValue: jsii.Number(123),
//   						StringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   					},
//   					TargetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		PostExtractionHookConfiguration: &HookConfigurationProperty{
//   			InvocationCondition: &DocumentAttributeConditionProperty{
//   				ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				ConditionOnValue: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   				Operator: jsii.String("operator"),
//   			},
//   			LambdaArn: jsii.String("lambdaArn"),
//   			S3Bucket: jsii.String("s3Bucket"),
//   		},
//   		PreExtractionHookConfiguration: &HookConfigurationProperty{
//   			InvocationCondition: &DocumentAttributeConditionProperty{
//   				ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				ConditionOnValue: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   				Operator: jsii.String("operator"),
//   			},
//   			LambdaArn: jsii.String("lambdaArn"),
//   			S3Bucket: jsii.String("s3Bucket"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	DataSourceConfiguration: &DataSourceConfigurationProperty{
//   		ConfluenceConfiguration: &ConfluenceConfigurationProperty{
//   			AttachmentConfiguration: &ConfluenceAttachmentConfigurationProperty{
//   				AttachmentFieldMappings: []interface{}{
//   					&ConfluenceAttachmentToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   				CrawlAttachments: jsii.Boolean(false),
//   			},
//   			BlogConfiguration: &ConfluenceBlogConfigurationProperty{
//   				BlogFieldMappings: []interface{}{
//   					&ConfluenceBlogToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   			},
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			PageConfiguration: &ConfluencePageConfigurationProperty{
//   				PageFieldMappings: []interface{}{
//   					&ConfluencePageToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   			ServerUrl: jsii.String("serverUrl"),
//   			SpaceConfiguration: &ConfluenceSpaceConfigurationProperty{
//   				CrawlArchivedSpaces: jsii.Boolean(false),
//   				CrawlPersonalSpaces: jsii.Boolean(false),
//   				ExcludeSpaces: []*string{
//   					jsii.String("excludeSpaces"),
//   				},
//   				IncludeSpaces: []*string{
//   					jsii.String("includeSpaces"),
//   				},
//   				SpaceFieldMappings: []interface{}{
//   					&ConfluenceSpaceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   			},
//   			Version: jsii.String("version"),
//   			VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		DatabaseConfiguration: &DatabaseConfigurationProperty{
//   			AclConfiguration: &AclConfigurationProperty{
//   				AllowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   			},
//   			ColumnConfiguration: &ColumnConfigurationProperty{
//   				ChangeDetectingColumns: []*string{
//   					jsii.String("changeDetectingColumns"),
//   				},
//   				DocumentDataColumnName: jsii.String("documentDataColumnName"),
//   				DocumentIdColumnName: jsii.String("documentIdColumnName"),
//   				DocumentTitleColumnName: jsii.String("documentTitleColumnName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   			},
//   			ConnectionConfiguration: &ConnectionConfigurationProperty{
//   				DatabaseHost: jsii.String("databaseHost"),
//   				DatabaseName: jsii.String("databaseName"),
//   				DatabasePort: jsii.Number(123),
//   				SecretArn: jsii.String("secretArn"),
//   				TableName: jsii.String("tableName"),
//   			},
//   			DatabaseEngineType: jsii.String("databaseEngineType"),
//   			SqlConfiguration: &SqlConfigurationProperty{
//   				QueryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   			},
//   			VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		GoogleDriveConfiguration: &GoogleDriveConfigurationProperty{
//   			ExcludeMimeTypes: []*string{
//   				jsii.String("excludeMimeTypes"),
//   			},
//   			ExcludeSharedDrives: []*string{
//   				jsii.String("excludeSharedDrives"),
//   			},
//   			ExcludeUserAccounts: []*string{
//   				jsii.String("excludeUserAccounts"),
//   			},
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		OneDriveConfiguration: &OneDriveConfigurationProperty{
//   			DisableLocalGroups: jsii.Boolean(false),
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			OneDriveUsers: &OneDriveUsersProperty{
//   				OneDriveUserList: []*string{
//   					jsii.String("oneDriveUserList"),
//   				},
//   				OneDriveUserS3Path: &S3PathProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   			TenantDomain: jsii.String("tenantDomain"),
//   		},
//   		S3Configuration: &S3DataSourceConfigurationProperty{
//   			AccessControlListConfiguration: &AccessControlListConfigurationProperty{
//   				KeyPath: jsii.String("keyPath"),
//   			},
//   			BucketName: jsii.String("bucketName"),
//   			DocumentsMetadataConfiguration: &DocumentsMetadataConfigurationProperty{
//   				S3Prefix: jsii.String("s3Prefix"),
//   			},
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			InclusionPrefixes: []*string{
//   				jsii.String("inclusionPrefixes"),
//   			},
//   		},
//   		SalesforceConfiguration: &SalesforceConfigurationProperty{
//   			ChatterFeedConfiguration: &SalesforceChatterFeedConfigurationProperty{
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   				IncludeFilterTypes: []*string{
//   					jsii.String("includeFilterTypes"),
//   				},
//   			},
//   			CrawlAttachments: jsii.Boolean(false),
//   			ExcludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
//   			},
//   			IncludeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   			KnowledgeArticleConfiguration: &SalesforceKnowledgeArticleConfigurationProperty{
//   				CustomKnowledgeArticleTypeConfigurations: []interface{}{
//   					&SalesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   						DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   						DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   						FieldMappings: []interface{}{
//   							&DataSourceToIndexFieldMappingProperty{
//   								DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   								DateFieldFormat: jsii.String("dateFieldFormat"),
//   								IndexFieldName: jsii.String("indexFieldName"),
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				IncludedStates: []*string{
//   					jsii.String("includedStates"),
//   				},
//   				StandardKnowledgeArticleTypeConfiguration: &SalesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   					DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   					DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					FieldMappings: []interface{}{
//   						&DataSourceToIndexFieldMappingProperty{
//   							DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							DateFieldFormat: jsii.String("dateFieldFormat"),
//   							IndexFieldName: jsii.String("indexFieldName"),
//   						},
//   					},
//   				},
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   			ServerUrl: jsii.String("serverUrl"),
//   			StandardObjectAttachmentConfiguration: &SalesforceStandardObjectAttachmentConfigurationProperty{
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   			},
//   			StandardObjectConfigurations: []interface{}{
//   				&SalesforceStandardObjectConfigurationProperty{
//   					DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   					DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					FieldMappings: []interface{}{
//   						&DataSourceToIndexFieldMappingProperty{
//   							DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							DateFieldFormat: jsii.String("dateFieldFormat"),
//   							IndexFieldName: jsii.String("indexFieldName"),
//   						},
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		ServiceNowConfiguration: &ServiceNowConfigurationProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
//   			HostUrl: jsii.String("hostUrl"),
//   			KnowledgeArticleConfiguration: &ServiceNowKnowledgeArticleConfigurationProperty{
//   				CrawlAttachments: jsii.Boolean(false),
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				ExcludeAttachmentFilePatterns: []*string{
//   					jsii.String("excludeAttachmentFilePatterns"),
//   				},
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   				FilterQuery: jsii.String("filterQuery"),
//   				IncludeAttachmentFilePatterns: []*string{
//   					jsii.String("includeAttachmentFilePatterns"),
//   				},
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   			ServiceCatalogConfiguration: &ServiceNowServiceCatalogConfigurationProperty{
//   				CrawlAttachments: jsii.Boolean(false),
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				ExcludeAttachmentFilePatterns: []*string{
//   					jsii.String("excludeAttachmentFilePatterns"),
//   				},
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   				IncludeAttachmentFilePatterns: []*string{
//   					jsii.String("includeAttachmentFilePatterns"),
//   				},
//   			},
//   			ServiceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//   		},
//   		SharePointConfiguration: &SharePointConfigurationProperty{
//   			CrawlAttachments: jsii.Boolean(false),
//   			DisableLocalGroups: jsii.Boolean(false),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			SecretArn: jsii.String("secretArn"),
//   			SharePointVersion: jsii.String("sharePointVersion"),
//   			SslCertificateS3Path: &S3PathProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//   			Urls: []*string{
//   				jsii.String("urls"),
//   			},
//   			UseChangeLog: jsii.Boolean(false),
//   			VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		TemplateConfiguration: &TemplateConfigurationProperty{
//   			Template: jsii.String("template"),
//   		},
//   		WebCrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   			AuthenticationConfiguration: &WebCrawlerAuthenticationConfigurationProperty{
//   				BasicAuthentication: []interface{}{
//   					&WebCrawlerBasicAuthenticationProperty{
//   						Credentials: jsii.String("credentials"),
//   						Host: jsii.String("host"),
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			CrawlDepth: jsii.Number(123),
//   			MaxContentSizePerPageInMegaBytes: jsii.Number(123),
//   			MaxLinksPerPage: jsii.Number(123),
//   			MaxUrlsPerMinuteCrawlRate: jsii.Number(123),
//   			ProxyConfiguration: &ProxyConfigurationProperty{
//   				Credentials: jsii.String("credentials"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   			UrlExclusionPatterns: []*string{
//   				jsii.String("urlExclusionPatterns"),
//   			},
//   			UrlInclusionPatterns: []*string{
//   				jsii.String("urlInclusionPatterns"),
//   			},
//   			Urls: &WebCrawlerUrlsProperty{
//   				SeedUrlConfiguration: &WebCrawlerSeedUrlConfigurationProperty{
//   					SeedUrls: []*string{
//   						jsii.String("seedUrls"),
//   					},
//   					WebCrawlerMode: jsii.String("webCrawlerMode"),
//   				},
//   				SiteMapsConfiguration: &WebCrawlerSiteMapsConfigurationProperty{
//   					SiteMaps: []*string{
//   						jsii.String("siteMaps"),
//   					},
//   				},
//   			},
//   		},
//   		WorkDocsConfiguration: &WorkDocsConfigurationProperty{
//   			CrawlComments: jsii.Boolean(false),
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			OrganizationId: jsii.String("organizationId"),
//   			UseChangeLog: jsii.Boolean(false),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IndexId: jsii.String("indexId"),
//   	LanguageCode: jsii.String("languageCode"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Schedule: jsii.String("schedule"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html
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


// Create a mixin to apply properties to `AWS::Kendra::DataSource`.
func NewCfnDataSourcePropsMixin(props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kendra.mixins.CfnDataSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Kendra::DataSource`.
func NewCfnDataSourcePropsMixin_Override(c CfnDataSourcePropsMixin, props *CfnDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kendra.mixins.CfnDataSourcePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_kendra.mixins.CfnDataSourcePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_kendra.mixins.CfnDataSourcePropsMixin",
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

