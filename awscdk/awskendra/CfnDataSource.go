package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskendra/internal"
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSource := awscdk.Aws_kendra.NewCfnDataSource(this, jsii.String("MyCfnDataSource"), &CfnDataSourceProps{
//   	IndexId: jsii.String("indexId"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	CustomDocumentEnrichmentConfiguration: &CustomDocumentEnrichmentConfigurationProperty{
//   		InlineConfigurations: []interface{}{
//   			&InlineCustomDocumentEnrichmentConfigurationProperty{
//   				Condition: &DocumentAttributeConditionProperty{
//   					ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   					Operator: jsii.String("operator"),
//
//   					// the properties below are optional
//   					ConditionOnValue: &DocumentAttributeValueProperty{
//   						DateValue: jsii.String("dateValue"),
//   						LongValue: jsii.Number(123),
//   						StringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   				DocumentContentDeletion: jsii.Boolean(false),
//   				Target: &DocumentAttributeTargetProperty{
//   					TargetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   					// the properties below are optional
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
//   			LambdaArn: jsii.String("lambdaArn"),
//   			S3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			InvocationCondition: &DocumentAttributeConditionProperty{
//   				ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				Operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				ConditionOnValue: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   		PreExtractionHookConfiguration: &HookConfigurationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			S3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			InvocationCondition: &DocumentAttributeConditionProperty{
//   				ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				Operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				ConditionOnValue: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	DataSourceConfiguration: &DataSourceConfigurationProperty{
//   		ConfluenceConfiguration: &ConfluenceConfigurationProperty{
//   			SecretArn: jsii.String("secretArn"),
//   			ServerUrl: jsii.String("serverUrl"),
//   			Version: jsii.String("version"),
//
//   			// the properties below are optional
//   			AttachmentConfiguration: &ConfluenceAttachmentConfigurationProperty{
//   				AttachmentFieldMappings: []interface{}{
//   					&ConfluenceAttachmentToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				CrawlAttachments: jsii.Boolean(false),
//   			},
//   			BlogConfiguration: &ConfluenceBlogConfigurationProperty{
//   				BlogFieldMappings: []interface{}{
//   					&ConfluenceBlogToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
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
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
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
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
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
//   		DatabaseConfiguration: &DatabaseConfigurationProperty{
//   			ColumnConfiguration: &ColumnConfigurationProperty{
//   				ChangeDetectingColumns: []*string{
//   					jsii.String("changeDetectingColumns"),
//   				},
//   				DocumentDataColumnName: jsii.String("documentDataColumnName"),
//   				DocumentIdColumnName: jsii.String("documentIdColumnName"),
//
//   				// the properties below are optional
//   				DocumentTitleColumnName: jsii.String("documentTitleColumnName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
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
//
//   			// the properties below are optional
//   			AclConfiguration: &AclConfigurationProperty{
//   				AllowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   			},
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
//   			SecretArn: jsii.String("secretArn"),
//
//   			// the properties below are optional
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
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   		},
//   		OneDriveConfiguration: &OneDriveConfigurationProperty{
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
//
//   			// the properties below are optional
//   			DisableLocalGroups: jsii.Boolean(false),
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   		},
//   		S3Configuration: &S3DataSourceConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			AccessControlListConfiguration: &AccessControlListConfigurationProperty{
//   				KeyPath: jsii.String("keyPath"),
//   			},
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
//   			SecretArn: jsii.String("secretArn"),
//   			ServerUrl: jsii.String("serverUrl"),
//
//   			// the properties below are optional
//   			ChatterFeedConfiguration: &SalesforceChatterFeedConfigurationProperty{
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
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
//   				IncludedStates: []*string{
//   					jsii.String("includedStates"),
//   				},
//
//   				// the properties below are optional
//   				CustomKnowledgeArticleTypeConfigurations: []interface{}{
//   					&SalesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   						DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   						FieldMappings: []interface{}{
//   							&DataSourceToIndexFieldMappingProperty{
//   								DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   								IndexFieldName: jsii.String("indexFieldName"),
//
//   								// the properties below are optional
//   								DateFieldFormat: jsii.String("dateFieldFormat"),
//   							},
//   						},
//   					},
//   				},
//   				StandardKnowledgeArticleTypeConfiguration: &SalesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   					DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   					// the properties below are optional
//   					DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					FieldMappings: []interface{}{
//   						&DataSourceToIndexFieldMappingProperty{
//   							DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							IndexFieldName: jsii.String("indexFieldName"),
//
//   							// the properties below are optional
//   							DateFieldFormat: jsii.String("dateFieldFormat"),
//   						},
//   					},
//   				},
//   			},
//   			StandardObjectAttachmentConfiguration: &SalesforceStandardObjectAttachmentConfigurationProperty{
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			StandardObjectConfigurations: []interface{}{
//   				&SalesforceStandardObjectConfigurationProperty{
//   					DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					FieldMappings: []interface{}{
//   						&DataSourceToIndexFieldMappingProperty{
//   							DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							IndexFieldName: jsii.String("indexFieldName"),
//
//   							// the properties below are optional
//   							DateFieldFormat: jsii.String("dateFieldFormat"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		ServiceNowConfiguration: &ServiceNowConfigurationProperty{
//   			HostUrl: jsii.String("hostUrl"),
//   			SecretArn: jsii.String("secretArn"),
//   			ServiceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//
//   			// the properties below are optional
//   			AuthenticationType: jsii.String("authenticationType"),
//   			KnowledgeArticleConfiguration: &ServiceNowKnowledgeArticleConfigurationProperty{
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				CrawlAttachments: jsii.Boolean(false),
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				ExcludeAttachmentFilePatterns: []*string{
//   					jsii.String("excludeAttachmentFilePatterns"),
//   				},
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				FilterQuery: jsii.String("filterQuery"),
//   				IncludeAttachmentFilePatterns: []*string{
//   					jsii.String("includeAttachmentFilePatterns"),
//   				},
//   			},
//   			ServiceCatalogConfiguration: &ServiceNowServiceCatalogConfigurationProperty{
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				CrawlAttachments: jsii.Boolean(false),
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				ExcludeAttachmentFilePatterns: []*string{
//   					jsii.String("excludeAttachmentFilePatterns"),
//   				},
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				IncludeAttachmentFilePatterns: []*string{
//   					jsii.String("includeAttachmentFilePatterns"),
//   				},
//   			},
//   		},
//   		SharePointConfiguration: &SharePointConfigurationProperty{
//   			SecretArn: jsii.String("secretArn"),
//   			SharePointVersion: jsii.String("sharePointVersion"),
//   			Urls: []*string{
//   				jsii.String("urls"),
//   			},
//
//   			// the properties below are optional
//   			CrawlAttachments: jsii.Boolean(false),
//   			DisableLocalGroups: jsii.Boolean(false),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			SslCertificateS3Path: &S3PathProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
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
//   			Urls: &WebCrawlerUrlsProperty{
//   				SeedUrlConfiguration: &WebCrawlerSeedUrlConfigurationProperty{
//   					SeedUrls: []*string{
//   						jsii.String("seedUrls"),
//   					},
//
//   					// the properties below are optional
//   					WebCrawlerMode: jsii.String("webCrawlerMode"),
//   				},
//   				SiteMapsConfiguration: &WebCrawlerSiteMapsConfigurationProperty{
//   					SiteMaps: []*string{
//   						jsii.String("siteMaps"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
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
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//
//   				// the properties below are optional
//   				Credentials: jsii.String("credentials"),
//   			},
//   			UrlExclusionPatterns: []*string{
//   				jsii.String("urlExclusionPatterns"),
//   			},
//   			UrlInclusionPatterns: []*string{
//   				jsii.String("urlInclusionPatterns"),
//   			},
//   		},
//   		WorkDocsConfiguration: &WorkDocsConfigurationProperty{
//   			OrganizationId: jsii.String("organizationId"),
//
//   			// the properties below are optional
//   			CrawlComments: jsii.Boolean(false),
//   			ExclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			InclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			UseChangeLog: jsii.Boolean(false),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	LanguageCode: jsii.String("languageCode"),
//   	RoleArn: jsii.String("roleArn"),
//   	Schedule: jsii.String("schedule"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html
//
type CfnDataSource interface {
	awscdk.CfnResource
	IDataSourceRef
	awscdk.IInspectable
	awscdk.ITaggable
	// The Amazon Resource Name (ARN) of the data source. For example:.
	//
	// `arn:aws:kendra:us-west-2:111122223333:index/335c3741-41df-46a6-b5d3-61f85b787884/data-source/b8cae438-6787-4091-8897-684a652bbb0a`.
	AttrArn() *string
	// The identifier for the data source. For example:.
	//
	// `b8cae438-6787-4091-8897-684a652bbb0a` .
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Configuration information for altering document metadata and content during the document ingestion process.
	CustomDocumentEnrichmentConfiguration() interface{}
	SetCustomDocumentEnrichmentConfiguration(val interface{})
	// Configuration information for an Amazon Kendra data source.
	DataSourceConfiguration() interface{}
	SetDataSourceConfiguration(val interface{})
	// A reference to a DataSource resource.
	DataSourceRef() *DataSourceReference
	// A description for the data source connector.
	Description() *string
	SetDescription(val *string)
	// The identifier of the index you want to use with the data source connector.
	IndexId() *string
	SetIndexId(val *string)
	// The code for a language.
	LanguageCode() *string
	SetLanguageCode(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the data source.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of a role with permission to access the data source.
	RoleArn() *string
	SetRoleArn(val *string)
	// Sets the frequency that Amazon Kendra checks the documents in your data source and updates the index.
	Schedule() *string
	SetSchedule(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An array of key-value pairs to apply to this resource.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The type of the data source.
	Type() *string
	SetType(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_IDataSourceRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnDataSource) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CustomDocumentEnrichmentConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDocumentEnrichmentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceRef() *DataSourceReference {
	var returns *DataSourceReference
	_jsii_.Get(
		j,
		"dataSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) IndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	if err := validateNewCfnDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource)SetCustomDocumentEnrichmentConfiguration(val interface{}) {
	if err := j.validateSetCustomDocumentEnrichmentConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDocumentEnrichmentConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetDataSourceConfiguration(val interface{}) {
	if err := j.validateSetDataSourceConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetIndexId(val *string) {
	if err := j.validateSetIndexIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetLanguageCode(val *string) {
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDataSource_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSource) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDataSource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

