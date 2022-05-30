package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Kendra::DataSource`.
//
// Specifies a data source that you use to with an Amazon Kendra index.
//
// You specify a name, connector type and description for your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSource := awscdk.Aws_kendra.NewCfnDataSource(this, jsii.String("MyCfnDataSource"), &cfnDataSourceProps{
//   	indexId: jsii.String("indexId"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	customDocumentEnrichmentConfiguration: &customDocumentEnrichmentConfigurationProperty{
//   		inlineConfigurations: []interface{}{
//   			&inlineCustomDocumentEnrichmentConfigurationProperty{
//   				condition: &documentAttributeConditionProperty{
//   					conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   					operator: jsii.String("operator"),
//
//   					// the properties below are optional
//   					conditionOnValue: &documentAttributeValueProperty{
//   						dateValue: jsii.String("dateValue"),
//   						longValue: jsii.Number(123),
//   						stringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						stringValue: jsii.String("stringValue"),
//   					},
//   				},
//   				documentContentDeletion: jsii.Boolean(false),
//   				target: &documentAttributeTargetProperty{
//   					targetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   					// the properties below are optional
//   					targetDocumentAttributeValue: &documentAttributeValueProperty{
//   						dateValue: jsii.String("dateValue"),
//   						longValue: jsii.Number(123),
//   						stringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						stringValue: jsii.String("stringValue"),
//   					},
//   					targetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		postExtractionHookConfiguration: &hookConfigurationProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   			s3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			invocationCondition: &documentAttributeConditionProperty{
//   				conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				conditionOnValue: &documentAttributeValueProperty{
//   					dateValue: jsii.String("dateValue"),
//   					longValue: jsii.Number(123),
//   					stringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   		preExtractionHookConfiguration: &hookConfigurationProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   			s3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			invocationCondition: &documentAttributeConditionProperty{
//   				conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				conditionOnValue: &documentAttributeValueProperty{
//   					dateValue: jsii.String("dateValue"),
//   					longValue: jsii.Number(123),
//   					stringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	dataSourceConfiguration: &dataSourceConfigurationProperty{
//   		confluenceConfiguration: &confluenceConfigurationProperty{
//   			secretArn: jsii.String("secretArn"),
//   			serverUrl: jsii.String("serverUrl"),
//   			version: jsii.String("version"),
//
//   			// the properties below are optional
//   			attachmentConfiguration: &confluenceAttachmentConfigurationProperty{
//   				attachmentFieldMappings: []interface{}{
//   					&confluenceAttachmentToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				crawlAttachments: jsii.Boolean(false),
//   			},
//   			blogConfiguration: &confluenceBlogConfigurationProperty{
//   				blogFieldMappings: []interface{}{
//   					&confluenceBlogToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			pageConfiguration: &confluencePageConfigurationProperty{
//   				pageFieldMappings: []interface{}{
//   					&confluencePageToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			spaceConfiguration: &confluenceSpaceConfigurationProperty{
//   				crawlArchivedSpaces: jsii.Boolean(false),
//   				crawlPersonalSpaces: jsii.Boolean(false),
//   				excludeSpaces: []*string{
//   					jsii.String("excludeSpaces"),
//   				},
//   				includeSpaces: []*string{
//   					jsii.String("includeSpaces"),
//   				},
//   				spaceFieldMappings: []interface{}{
//   					&confluenceSpaceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		databaseConfiguration: &databaseConfigurationProperty{
//   			columnConfiguration: &columnConfigurationProperty{
//   				changeDetectingColumns: []*string{
//   					jsii.String("changeDetectingColumns"),
//   				},
//   				documentDataColumnName: jsii.String("documentDataColumnName"),
//   				documentIdColumnName: jsii.String("documentIdColumnName"),
//
//   				// the properties below are optional
//   				documentTitleColumnName: jsii.String("documentTitleColumnName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			connectionConfiguration: &connectionConfigurationProperty{
//   				databaseHost: jsii.String("databaseHost"),
//   				databaseName: jsii.String("databaseName"),
//   				databasePort: jsii.Number(123),
//   				secretArn: jsii.String("secretArn"),
//   				tableName: jsii.String("tableName"),
//   			},
//   			databaseEngineType: jsii.String("databaseEngineType"),
//
//   			// the properties below are optional
//   			aclConfiguration: &aclConfigurationProperty{
//   				allowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   			},
//   			sqlConfiguration: &sqlConfigurationProperty{
//   				queryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   			},
//   			vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		googleDriveConfiguration: &googleDriveConfigurationProperty{
//   			secretArn: jsii.String("secretArn"),
//
//   			// the properties below are optional
//   			excludeMimeTypes: []*string{
//   				jsii.String("excludeMimeTypes"),
//   			},
//   			excludeSharedDrives: []*string{
//   				jsii.String("excludeSharedDrives"),
//   			},
//   			excludeUserAccounts: []*string{
//   				jsii.String("excludeUserAccounts"),
//   			},
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   		},
//   		oneDriveConfiguration: &oneDriveConfigurationProperty{
//   			oneDriveUsers: &oneDriveUsersProperty{
//   				oneDriveUserList: []*string{
//   					jsii.String("oneDriveUserList"),
//   				},
//   				oneDriveUserS3Path: &s3PathProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			secretArn: jsii.String("secretArn"),
//   			tenantDomain: jsii.String("tenantDomain"),
//
//   			// the properties below are optional
//   			disableLocalGroups: jsii.Boolean(false),
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   		},
//   		s3Configuration: &s3DataSourceConfigurationProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			accessControlListConfiguration: &accessControlListConfigurationProperty{
//   				keyPath: jsii.String("keyPath"),
//   			},
//   			documentsMetadataConfiguration: &documentsMetadataConfigurationProperty{
//   				s3Prefix: jsii.String("s3Prefix"),
//   			},
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			inclusionPrefixes: []*string{
//   				jsii.String("inclusionPrefixes"),
//   			},
//   		},
//   		salesforceConfiguration: &salesforceConfigurationProperty{
//   			secretArn: jsii.String("secretArn"),
//   			serverUrl: jsii.String("serverUrl"),
//
//   			// the properties below are optional
//   			chatterFeedConfiguration: &salesforceChatterFeedConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				includeFilterTypes: []*string{
//   					jsii.String("includeFilterTypes"),
//   				},
//   			},
//   			crawlAttachments: jsii.Boolean(false),
//   			excludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
//   			},
//   			includeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   			knowledgeArticleConfiguration: &salesforceKnowledgeArticleConfigurationProperty{
//   				includedStates: []*string{
//   					jsii.String("includedStates"),
//   				},
//
//   				// the properties below are optional
//   				customKnowledgeArticleTypeConfigurations: []interface{}{
//   					&salesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   						documentDataFieldName: jsii.String("documentDataFieldName"),
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   						fieldMappings: []interface{}{
//   							&dataSourceToIndexFieldMappingProperty{
//   								dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   								indexFieldName: jsii.String("indexFieldName"),
//
//   								// the properties below are optional
//   								dateFieldFormat: jsii.String("dateFieldFormat"),
//   							},
//   						},
//   					},
//   				},
//   				standardKnowledgeArticleTypeConfiguration: &salesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   					documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   					// the properties below are optional
//   					documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					fieldMappings: []interface{}{
//   						&dataSourceToIndexFieldMappingProperty{
//   							dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							indexFieldName: jsii.String("indexFieldName"),
//
//   							// the properties below are optional
//   							dateFieldFormat: jsii.String("dateFieldFormat"),
//   						},
//   					},
//   				},
//   			},
//   			standardObjectAttachmentConfiguration: &salesforceStandardObjectAttachmentConfigurationProperty{
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			standardObjectConfigurations: []interface{}{
//   				&salesforceStandardObjectConfigurationProperty{
//   					documentDataFieldName: jsii.String("documentDataFieldName"),
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					fieldMappings: []interface{}{
//   						&dataSourceToIndexFieldMappingProperty{
//   							dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							indexFieldName: jsii.String("indexFieldName"),
//
//   							// the properties below are optional
//   							dateFieldFormat: jsii.String("dateFieldFormat"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		serviceNowConfiguration: &serviceNowConfigurationProperty{
//   			hostUrl: jsii.String("hostUrl"),
//   			secretArn: jsii.String("secretArn"),
//   			serviceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//
//   			// the properties below are optional
//   			authenticationType: jsii.String("authenticationType"),
//   			knowledgeArticleConfiguration: &serviceNowKnowledgeArticleConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				crawlAttachments: jsii.Boolean(false),
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				excludeAttachmentFilePatterns: []*string{
//   					jsii.String("excludeAttachmentFilePatterns"),
//   				},
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				filterQuery: jsii.String("filterQuery"),
//   				includeAttachmentFilePatterns: []*string{
//   					jsii.String("includeAttachmentFilePatterns"),
//   				},
//   			},
//   			serviceCatalogConfiguration: &serviceNowServiceCatalogConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				crawlAttachments: jsii.Boolean(false),
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				excludeAttachmentFilePatterns: []*string{
//   					jsii.String("excludeAttachmentFilePatterns"),
//   				},
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				includeAttachmentFilePatterns: []*string{
//   					jsii.String("includeAttachmentFilePatterns"),
//   				},
//   			},
//   		},
//   		sharePointConfiguration: &sharePointConfigurationProperty{
//   			secretArn: jsii.String("secretArn"),
//   			sharePointVersion: jsii.String("sharePointVersion"),
//   			urls: []*string{
//   				jsii.String("urls"),
//   			},
//
//   			// the properties below are optional
//   			crawlAttachments: jsii.Boolean(false),
//   			disableLocalGroups: jsii.Boolean(false),
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			sslCertificateS3Path: &s3PathProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//   			},
//   			useChangeLog: jsii.Boolean(false),
//   			vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		webCrawlerConfiguration: &webCrawlerConfigurationProperty{
//   			urls: &webCrawlerUrlsProperty{
//   				seedUrlConfiguration: &webCrawlerSeedUrlConfigurationProperty{
//   					seedUrls: []*string{
//   						jsii.String("seedUrls"),
//   					},
//
//   					// the properties below are optional
//   					webCrawlerMode: jsii.String("webCrawlerMode"),
//   				},
//   				siteMapsConfiguration: &webCrawlerSiteMapsConfigurationProperty{
//   					siteMaps: []*string{
//   						jsii.String("siteMaps"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			authenticationConfiguration: &webCrawlerAuthenticationConfigurationProperty{
//   				basicAuthentication: []interface{}{
//   					&webCrawlerBasicAuthenticationProperty{
//   						credentials: jsii.String("credentials"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			crawlDepth: jsii.Number(123),
//   			maxContentSizePerPageInMegaBytes: jsii.Number(123),
//   			maxLinksPerPage: jsii.Number(123),
//   			maxUrlsPerMinuteCrawlRate: jsii.Number(123),
//   			proxyConfiguration: &proxyConfigurationProperty{
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//
//   				// the properties below are optional
//   				credentials: jsii.String("credentials"),
//   			},
//   			urlExclusionPatterns: []*string{
//   				jsii.String("urlExclusionPatterns"),
//   			},
//   			urlInclusionPatterns: []*string{
//   				jsii.String("urlInclusionPatterns"),
//   			},
//   		},
//   		workDocsConfiguration: &workDocsConfigurationProperty{
//   			organizationId: jsii.String("organizationId"),
//
//   			// the properties below are optional
//   			crawlComments: jsii.Boolean(false),
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			useChangeLog: jsii.Boolean(false),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	roleArn: jsii.String("roleArn"),
//   	schedule: jsii.String("schedule"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
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
	//
	// The contents of the configuration depend on the type of data source. You can only specify one type of data source in the configuration. Choose from one of the following data sources.
	//
	// - Amazon S3
	// - Confluence
	// - Custom
	// - Database
	// - Microsoft OneDrive
	// - Microsoft SharePoint
	// - Salesforce
	// - ServiceNow
	//
	// You can't specify the `Configuration` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `Configuration` parameter is required for all other data sources.
	DataSourceConfiguration() interface{}
	SetDataSourceConfiguration(val interface{})
	// A description of the data source.
	Description() *string
	SetDescription(val *string)
	// The identifier of the index that should be associated with this data source.
	IndexId() *string
	SetIndexId(val *string)
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
	//
	// You can't specify the `RoleArn` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `RoleArn` parameter is required for all other data sources.
	RoleArn() *string
	SetRoleArn(val *string)
	// Sets the frequency that Amazon Kendra checks the documents in your data source and updates the index.
	//
	// If you don't set a schedule, Amazon Kendra doesn't periodically update the index.
	Schedule() *string
	SetSchedule(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// The type of the data source.
	Type() *string
	SetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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
	internal.Type__awscdkIInspectable
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


// Create a new `AWS::Kendra::DataSource`.
func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::DataSource`.
func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource) SetCustomDocumentEnrichmentConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"customDocumentEnrichmentConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDataSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"dataSourceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetIndexId(val *string) {
	_jsii_.Set(
		j,
		"indexId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetType(val *string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies access control list files for the documents in a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlListConfigurationProperty := &accessControlListConfigurationProperty{
//   	keyPath: jsii.String("keyPath"),
//   }
//
type CfnDataSource_AccessControlListConfigurationProperty struct {
	// Path to the AWS S3 bucket that contains the access control list files.
	KeyPath *string `field:"optional" json:"keyPath" yaml:"keyPath"`
}

// Provides information about the column that should be used for filtering the query response by groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aclConfigurationProperty := &aclConfigurationProperty{
//   	allowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   }
//
type CfnDataSource_AclConfigurationProperty struct {
	// A list of groups, separated by semi-colons, that filters a query response based on user context.
	//
	// The document is only returned to users that are in one of the groups specified in the `UserContext` field of the [Query](https://docs.aws.amazon.com/kendra/latest/dg/API_Query.html) operation.
	AllowedGroupsColumnName *string `field:"required" json:"allowedGroupsColumnName" yaml:"allowedGroupsColumnName"`
}

// Provides information about how Amazon Kendra should use the columns of a database in an index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnConfigurationProperty := &columnConfigurationProperty{
//   	changeDetectingColumns: []*string{
//   		jsii.String("changeDetectingColumns"),
//   	},
//   	documentDataColumnName: jsii.String("documentDataColumnName"),
//   	documentIdColumnName: jsii.String("documentIdColumnName"),
//
//   	// the properties below are optional
//   	documentTitleColumnName: jsii.String("documentTitleColumnName"),
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_ColumnConfigurationProperty struct {
	// One to five columns that indicate when a document in the database has changed.
	ChangeDetectingColumns *[]*string `field:"required" json:"changeDetectingColumns" yaml:"changeDetectingColumns"`
	// The column that contains the contents of the document.
	DocumentDataColumnName *string `field:"required" json:"documentDataColumnName" yaml:"documentDataColumnName"`
	// The column that provides the document's unique identifier.
	DocumentIdColumnName *string `field:"required" json:"documentIdColumnName" yaml:"documentIdColumnName"`
	// The column that contains the title of the document.
	DocumentTitleColumnName *string `field:"optional" json:"documentTitleColumnName" yaml:"documentTitleColumnName"`
	// An array of objects that map database column names to the corresponding fields in an index.
	//
	// You must first create the fields in the index using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

// Configuration of attachment settings for the Confluence data source.
//
// Attachment settings are optional, if you don't specify settings attachments, Amazon Kendra won't index them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceAttachmentConfigurationProperty := &confluenceAttachmentConfigurationProperty{
//   	attachmentFieldMappings: []interface{}{
//   		&confluenceAttachmentToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	crawlAttachments: jsii.Boolean(false),
//   }
//
type CfnDataSource_ConfluenceAttachmentConfigurationProperty struct {
	// Maps attributes or field names of Confluence attachments to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `AttachentFieldMappings` parameter, you must specify at least one field mapping.
	AttachmentFieldMappings interface{} `field:"optional" json:"attachmentFieldMappings" yaml:"attachmentFieldMappings"`
	// Indicates whether Amazon Kendra indexes attachments to the pages and blogs in the Confluence data source.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
}

// Maps attributes or field names of Confluence attachments to Amazon Kendra index field names.
//
// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confuence data source field names must exist in your Confluence custom metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceAttachmentToIndexFieldMappingProperty := &confluenceAttachmentToIndexFieldMappingProperty{
//   	dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	indexFieldName: jsii.String("indexFieldName"),
//
//   	// the properties below are optional
//   	dateFieldFormat: jsii.String("dateFieldFormat"),
//   }
//
type CfnDataSource_ConfluenceAttachmentToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	//
	// You must first create the index field using the `UpdateIndex` API.
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Configuration of blog settings for the Confluence data source.
//
// Blogs are always indexed unless filtered from the index by the `ExclusionPatterns` or `InclusionPatterns` fields in the `ConfluenceConfiguration` object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceBlogConfigurationProperty := &confluenceBlogConfigurationProperty{
//   	blogFieldMappings: []interface{}{
//   		&confluenceBlogToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_ConfluenceBlogConfigurationProperty struct {
	// Maps attributes or field names of Confluence blogs to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `BlogFieldMappings` parameter, you must specify at least one field mapping.
	BlogFieldMappings interface{} `field:"optional" json:"blogFieldMappings" yaml:"blogFieldMappings"`
}

// Maps attributes or field names of Confluence blog to Amazon Kendra index field names.
//
// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceBlogToIndexFieldMappingProperty := &confluenceBlogToIndexFieldMappingProperty{
//   	dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	indexFieldName: jsii.String("indexFieldName"),
//
//   	// the properties below are optional
//   	dateFieldFormat: jsii.String("dateFieldFormat"),
//   }
//
type CfnDataSource_ConfluenceBlogToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Provides the configuration information to connect to Confluence as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceConfigurationProperty := &confluenceConfigurationProperty{
//   	secretArn: jsii.String("secretArn"),
//   	serverUrl: jsii.String("serverUrl"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	attachmentConfiguration: &confluenceAttachmentConfigurationProperty{
//   		attachmentFieldMappings: []interface{}{
//   			&confluenceAttachmentToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		crawlAttachments: jsii.Boolean(false),
//   	},
//   	blogConfiguration: &confluenceBlogConfigurationProperty{
//   		blogFieldMappings: []interface{}{
//   			&confluenceBlogToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	pageConfiguration: &confluencePageConfigurationProperty{
//   		pageFieldMappings: []interface{}{
//   			&confluencePageToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	spaceConfiguration: &confluenceSpaceConfigurationProperty{
//   		crawlArchivedSpaces: jsii.Boolean(false),
//   		crawlPersonalSpaces: jsii.Boolean(false),
//   		excludeSpaces: []*string{
//   			jsii.String("excludeSpaces"),
//   		},
//   		includeSpaces: []*string{
//   			jsii.String("includeSpaces"),
//   		},
//   		spaceFieldMappings: []interface{}{
//   			&confluenceSpaceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDataSource_ConfluenceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the key-value pairs required to connect to your Confluence server.
	//
	// The secret must contain a JSON structure with the following keys:
	//
	// - username—The user name or email address of a user with administrative privileges for the Confluence server.
	// - password—The password associated with the user logging in to the Confluence server.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The URL of your Confluence instance.
	//
	// Use the full URL of the server. For example, *https://server.example.com:port/* . You can also use an IP address, for example, *https://192.168.1.113/* .
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Specifies the version of the Confluence installation that you are connecting to.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Configuration information for indexing attachments to Confluence blogs and pages.
	AttachmentConfiguration interface{} `field:"optional" json:"attachmentConfiguration" yaml:"attachmentConfiguration"`
	// Configuration information for indexing Confluence blogs.
	BlogConfiguration interface{} `field:"optional" json:"blogConfiguration" yaml:"blogConfiguration"`
	// >A list of regular expression patterns to exclude certain blog posts, pages, spaces, or attachments in your Confluence.
	//
	// Content that matches the patterns are excluded from the index. Content that doesn't match the patterns is included in the index. If content matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the content isn't included in the index.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of regular expression patterns to include certain blog posts, pages, spaces, or attachments in your Confluence.
	//
	// Content that matches the patterns are included in the index. Content that doesn't match the patterns is excluded from the index. If content matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the content isn't included in the index.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Configuration information for indexing Confluence pages.
	PageConfiguration interface{} `field:"optional" json:"pageConfiguration" yaml:"pageConfiguration"`
	// Configuration information for indexing Confluence spaces.
	SpaceConfiguration interface{} `field:"optional" json:"spaceConfiguration" yaml:"spaceConfiguration"`
	// Configuration information for an Amazon Virtual Private Cloud to connect to your Confluence.
	//
	// For more information, see [Configuring a VPC](https://docs.aws.amazon.com/kendra/latest/dg/vpc-configuration.html) .
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// Configuration of the page settings for the Confluence data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluencePageConfigurationProperty := &confluencePageConfigurationProperty{
//   	pageFieldMappings: []interface{}{
//   		&confluencePageToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_ConfluencePageConfigurationProperty struct {
	// >Maps attributes or field names of Confluence pages to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `PageFieldMappings` parameter, you must specify at least one field mapping.
	PageFieldMappings interface{} `field:"optional" json:"pageFieldMappings" yaml:"pageFieldMappings"`
}

// >Maps attributes or field names of Confluence pages to Amazon Kendra index field names.
//
// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluencePageToIndexFieldMappingProperty := &confluencePageToIndexFieldMappingProperty{
//   	dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	indexFieldName: jsii.String("indexFieldName"),
//
//   	// the properties below are optional
//   	dateFieldFormat: jsii.String("dateFieldFormat"),
//   }
//
type CfnDataSource_ConfluencePageToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Configuration information for indexing Confluence spaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceSpaceConfigurationProperty := &confluenceSpaceConfigurationProperty{
//   	crawlArchivedSpaces: jsii.Boolean(false),
//   	crawlPersonalSpaces: jsii.Boolean(false),
//   	excludeSpaces: []*string{
//   		jsii.String("excludeSpaces"),
//   	},
//   	includeSpaces: []*string{
//   		jsii.String("includeSpaces"),
//   	},
//   	spaceFieldMappings: []interface{}{
//   		&confluenceSpaceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_ConfluenceSpaceConfigurationProperty struct {
	// Specifies whether Amazon Kendra should index archived spaces.
	CrawlArchivedSpaces interface{} `field:"optional" json:"crawlArchivedSpaces" yaml:"crawlArchivedSpaces"`
	// Specifies whether Amazon Kendra should index personal spaces.
	//
	// Users can add restrictions to items in personal spaces. If personal spaces are indexed, queries without user context information may return restricted items from a personal space in their results. For more information, see [Filtering on user context](https://docs.aws.amazon.com/kendra/latest/dg/user-context-filter.html) .
	CrawlPersonalSpaces interface{} `field:"optional" json:"crawlPersonalSpaces" yaml:"crawlPersonalSpaces"`
	// A list of space keys of Confluence spaces.
	//
	// If you include a key, the blogs, documents, and attachments in the space are not indexed. If a space is in both the `ExcludeSpaces` and the `IncludeSpaces` list, the space is excluded.
	ExcludeSpaces *[]*string `field:"optional" json:"excludeSpaces" yaml:"excludeSpaces"`
	// A list of space keys for Confluence spaces.
	//
	// If you include a key, the blogs, documents, and attachments in the space are indexed. Spaces that aren't in the list aren't indexed. A space in the list must exist. Otherwise, Amazon Kendra logs an error when the data source is synchronized. If a space is in both the `IncludeSpaces` and the `ExcludeSpaces` list, the space is excluded.
	IncludeSpaces *[]*string `field:"optional" json:"includeSpaces" yaml:"includeSpaces"`
	// Maps attributes or field names of Confluence spaces to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `SpaceFieldMappings` parameter, you must specify at least one field mapping.
	SpaceFieldMappings interface{} `field:"optional" json:"spaceFieldMappings" yaml:"spaceFieldMappings"`
}

// >Maps attributes or field names of Confluence spaces to Amazon Kendra index field names.
//
// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceSpaceToIndexFieldMappingProperty := &confluenceSpaceToIndexFieldMappingProperty{
//   	dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	indexFieldName: jsii.String("indexFieldName"),
//
//   	// the properties below are optional
//   	dateFieldFormat: jsii.String("dateFieldFormat"),
//   }
//
type CfnDataSource_ConfluenceSpaceToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Provides the configuration information that's required to connect to a database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionConfigurationProperty := &connectionConfigurationProperty{
//   	databaseHost: jsii.String("databaseHost"),
//   	databaseName: jsii.String("databaseName"),
//   	databasePort: jsii.Number(123),
//   	secretArn: jsii.String("secretArn"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnDataSource_ConnectionConfigurationProperty struct {
	// The name of the host for the database.
	//
	// Can be either a string (host.subdomain.domain.tld) or an IPv4 or IPv6 address.
	DatabaseHost *string `field:"required" json:"databaseHost" yaml:"databaseHost"`
	// The name of the database containing the document data.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The port that the database uses for connections.
	DatabasePort *float64 `field:"required" json:"databasePort" yaml:"databasePort"`
	// The Amazon Resource Name (ARN) of credentials stored in AWS Secrets Manager .
	//
	// The credentials should be a user/password pair. For more information, see [Using a Database Data Source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-database.html) . For more information about AWS Secrets Manager , see [What Is AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) in the *AWS Secrets Manager* user guide.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The name of the table that contains the document data.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

// Provides the configuration information for altering document metadata and content during the document ingestion process.
//
// For more information, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDocumentEnrichmentConfigurationProperty := &customDocumentEnrichmentConfigurationProperty{
//   	inlineConfigurations: []interface{}{
//   		&inlineCustomDocumentEnrichmentConfigurationProperty{
//   			condition: &documentAttributeConditionProperty{
//   				conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				conditionOnValue: &documentAttributeValueProperty{
//   					dateValue: jsii.String("dateValue"),
//   					longValue: jsii.Number(123),
//   					stringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			documentContentDeletion: jsii.Boolean(false),
//   			target: &documentAttributeTargetProperty{
//   				targetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   				// the properties below are optional
//   				targetDocumentAttributeValue: &documentAttributeValueProperty{
//   					dateValue: jsii.String("dateValue"),
//   					longValue: jsii.Number(123),
//   					stringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   				targetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	postExtractionHookConfiguration: &hookConfigurationProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		invocationCondition: &documentAttributeConditionProperty{
//   			conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   			operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			conditionOnValue: &documentAttributeValueProperty{
//   				dateValue: jsii.String("dateValue"),
//   				longValue: jsii.Number(123),
//   				stringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	preExtractionHookConfiguration: &hookConfigurationProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		invocationCondition: &documentAttributeConditionProperty{
//   			conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   			operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			conditionOnValue: &documentAttributeValueProperty{
//   				dateValue: jsii.String("dateValue"),
//   				longValue: jsii.Number(123),
//   				stringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnDataSource_CustomDocumentEnrichmentConfigurationProperty struct {
	// Configuration information to alter document attributes or metadata fields and content when ingesting documents into Amazon Kendra.
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// Configuration information for invoking a Lambda function in AWS Lambda on the structured documents with their metadata and text extracted.
	//
	// You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Advanced data manipulation](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#advanced-data-manipulation) .
	PostExtractionHookConfiguration interface{} `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// Configuration information for invoking a Lambda function in AWS Lambda on the original or raw documents before extracting their metadata and text.
	//
	// You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Advanced data manipulation](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#advanced-data-manipulation) .
	PreExtractionHookConfiguration interface{} `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
	// The Amazon Resource Name (ARN) of a role with permission to run `PreExtractionHookConfiguration` and `PostExtractionHookConfiguration` for altering document metadata and content during the document ingestion process.
	//
	// For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html) .
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

// Provides the configuration information for an Amazon Kendra data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationProperty := &dataSourceConfigurationProperty{
//   	confluenceConfiguration: &confluenceConfigurationProperty{
//   		secretArn: jsii.String("secretArn"),
//   		serverUrl: jsii.String("serverUrl"),
//   		version: jsii.String("version"),
//
//   		// the properties below are optional
//   		attachmentConfiguration: &confluenceAttachmentConfigurationProperty{
//   			attachmentFieldMappings: []interface{}{
//   				&confluenceAttachmentToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			crawlAttachments: jsii.Boolean(false),
//   		},
//   		blogConfiguration: &confluenceBlogConfigurationProperty{
//   			blogFieldMappings: []interface{}{
//   				&confluenceBlogToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   		exclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		inclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		pageConfiguration: &confluencePageConfigurationProperty{
//   			pageFieldMappings: []interface{}{
//   				&confluencePageToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   		spaceConfiguration: &confluenceSpaceConfigurationProperty{
//   			crawlArchivedSpaces: jsii.Boolean(false),
//   			crawlPersonalSpaces: jsii.Boolean(false),
//   			excludeSpaces: []*string{
//   				jsii.String("excludeSpaces"),
//   			},
//   			includeSpaces: []*string{
//   				jsii.String("includeSpaces"),
//   			},
//   			spaceFieldMappings: []interface{}{
//   				&confluenceSpaceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   		vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	databaseConfiguration: &databaseConfigurationProperty{
//   		columnConfiguration: &columnConfigurationProperty{
//   			changeDetectingColumns: []*string{
//   				jsii.String("changeDetectingColumns"),
//   			},
//   			documentDataColumnName: jsii.String("documentDataColumnName"),
//   			documentIdColumnName: jsii.String("documentIdColumnName"),
//
//   			// the properties below are optional
//   			documentTitleColumnName: jsii.String("documentTitleColumnName"),
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   		connectionConfiguration: &connectionConfigurationProperty{
//   			databaseHost: jsii.String("databaseHost"),
//   			databaseName: jsii.String("databaseName"),
//   			databasePort: jsii.Number(123),
//   			secretArn: jsii.String("secretArn"),
//   			tableName: jsii.String("tableName"),
//   		},
//   		databaseEngineType: jsii.String("databaseEngineType"),
//
//   		// the properties below are optional
//   		aclConfiguration: &aclConfigurationProperty{
//   			allowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   		},
//   		sqlConfiguration: &sqlConfigurationProperty{
//   			queryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   		},
//   		vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	googleDriveConfiguration: &googleDriveConfigurationProperty{
//   		secretArn: jsii.String("secretArn"),
//
//   		// the properties below are optional
//   		excludeMimeTypes: []*string{
//   			jsii.String("excludeMimeTypes"),
//   		},
//   		excludeSharedDrives: []*string{
//   			jsii.String("excludeSharedDrives"),
//   		},
//   		excludeUserAccounts: []*string{
//   			jsii.String("excludeUserAccounts"),
//   		},
//   		exclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		inclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   	},
//   	oneDriveConfiguration: &oneDriveConfigurationProperty{
//   		oneDriveUsers: &oneDriveUsersProperty{
//   			oneDriveUserList: []*string{
//   				jsii.String("oneDriveUserList"),
//   			},
//   			oneDriveUserS3Path: &s3PathProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//   			},
//   		},
//   		secretArn: jsii.String("secretArn"),
//   		tenantDomain: jsii.String("tenantDomain"),
//
//   		// the properties below are optional
//   		disableLocalGroups: jsii.Boolean(false),
//   		exclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		inclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   	},
//   	s3Configuration: &s3DataSourceConfigurationProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		accessControlListConfiguration: &accessControlListConfigurationProperty{
//   			keyPath: jsii.String("keyPath"),
//   		},
//   		documentsMetadataConfiguration: &documentsMetadataConfigurationProperty{
//   			s3Prefix: jsii.String("s3Prefix"),
//   		},
//   		exclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		inclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		inclusionPrefixes: []*string{
//   			jsii.String("inclusionPrefixes"),
//   		},
//   	},
//   	salesforceConfiguration: &salesforceConfigurationProperty{
//   		secretArn: jsii.String("secretArn"),
//   		serverUrl: jsii.String("serverUrl"),
//
//   		// the properties below are optional
//   		chatterFeedConfiguration: &salesforceChatterFeedConfigurationProperty{
//   			documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   			// the properties below are optional
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			includeFilterTypes: []*string{
//   				jsii.String("includeFilterTypes"),
//   			},
//   		},
//   		crawlAttachments: jsii.Boolean(false),
//   		excludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		includeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   		knowledgeArticleConfiguration: &salesforceKnowledgeArticleConfigurationProperty{
//   			includedStates: []*string{
//   				jsii.String("includedStates"),
//   			},
//
//   			// the properties below are optional
//   			customKnowledgeArticleTypeConfigurations: []interface{}{
//   				&salesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   					documentDataFieldName: jsii.String("documentDataFieldName"),
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					fieldMappings: []interface{}{
//   						&dataSourceToIndexFieldMappingProperty{
//   							dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							indexFieldName: jsii.String("indexFieldName"),
//
//   							// the properties below are optional
//   							dateFieldFormat: jsii.String("dateFieldFormat"),
//   						},
//   					},
//   				},
//   			},
//   			standardKnowledgeArticleTypeConfiguration: &salesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   		},
//   		standardObjectAttachmentConfiguration: &salesforceStandardObjectAttachmentConfigurationProperty{
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   		standardObjectConfigurations: []interface{}{
//   			&salesforceStandardObjectConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	serviceNowConfiguration: &serviceNowConfigurationProperty{
//   		hostUrl: jsii.String("hostUrl"),
//   		secretArn: jsii.String("secretArn"),
//   		serviceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//
//   		// the properties below are optional
//   		authenticationType: jsii.String("authenticationType"),
//   		knowledgeArticleConfiguration: &serviceNowKnowledgeArticleConfigurationProperty{
//   			documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   			// the properties below are optional
//   			crawlAttachments: jsii.Boolean(false),
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			excludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			filterQuery: jsii.String("filterQuery"),
//   			includeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   		},
//   		serviceCatalogConfiguration: &serviceNowServiceCatalogConfigurationProperty{
//   			documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   			// the properties below are optional
//   			crawlAttachments: jsii.Boolean(false),
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			excludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			includeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   		},
//   	},
//   	sharePointConfiguration: &sharePointConfigurationProperty{
//   		secretArn: jsii.String("secretArn"),
//   		sharePointVersion: jsii.String("sharePointVersion"),
//   		urls: []*string{
//   			jsii.String("urls"),
//   		},
//
//   		// the properties below are optional
//   		crawlAttachments: jsii.Boolean(false),
//   		disableLocalGroups: jsii.Boolean(false),
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		exclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		inclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		sslCertificateS3Path: &s3PathProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//   		},
//   		useChangeLog: jsii.Boolean(false),
//   		vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	webCrawlerConfiguration: &webCrawlerConfigurationProperty{
//   		urls: &webCrawlerUrlsProperty{
//   			seedUrlConfiguration: &webCrawlerSeedUrlConfigurationProperty{
//   				seedUrls: []*string{
//   					jsii.String("seedUrls"),
//   				},
//
//   				// the properties below are optional
//   				webCrawlerMode: jsii.String("webCrawlerMode"),
//   			},
//   			siteMapsConfiguration: &webCrawlerSiteMapsConfigurationProperty{
//   				siteMaps: []*string{
//   					jsii.String("siteMaps"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		authenticationConfiguration: &webCrawlerAuthenticationConfigurationProperty{
//   			basicAuthentication: []interface{}{
//   				&webCrawlerBasicAuthenticationProperty{
//   					credentials: jsii.String("credentials"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		crawlDepth: jsii.Number(123),
//   		maxContentSizePerPageInMegaBytes: jsii.Number(123),
//   		maxLinksPerPage: jsii.Number(123),
//   		maxUrlsPerMinuteCrawlRate: jsii.Number(123),
//   		proxyConfiguration: &proxyConfigurationProperty{
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//
//   			// the properties below are optional
//   			credentials: jsii.String("credentials"),
//   		},
//   		urlExclusionPatterns: []*string{
//   			jsii.String("urlExclusionPatterns"),
//   		},
//   		urlInclusionPatterns: []*string{
//   			jsii.String("urlInclusionPatterns"),
//   		},
//   	},
//   	workDocsConfiguration: &workDocsConfigurationProperty{
//   		organizationId: jsii.String("organizationId"),
//
//   		// the properties below are optional
//   		crawlComments: jsii.Boolean(false),
//   		exclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		inclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		useChangeLog: jsii.Boolean(false),
//   	},
//   }
//
type CfnDataSource_DataSourceConfigurationProperty struct {
	// Provides the configuration information to connect to Confluence as your data source.
	ConfluenceConfiguration interface{} `field:"optional" json:"confluenceConfiguration" yaml:"confluenceConfiguration"`
	// Provides the configuration information to connect to a database as your data source.
	DatabaseConfiguration interface{} `field:"optional" json:"databaseConfiguration" yaml:"databaseConfiguration"`
	// Provides the configuration information to connect to Google Drive as your data source.
	GoogleDriveConfiguration interface{} `field:"optional" json:"googleDriveConfiguration" yaml:"googleDriveConfiguration"`
	// Provides the configuration information to connect to Microsoft OneDrive as your data source.
	OneDriveConfiguration interface{} `field:"optional" json:"oneDriveConfiguration" yaml:"oneDriveConfiguration"`
	// Provides the configuration information to connect to an Amazon S3 bucket as your data source.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// Provides the configuration information to connect to Salesforce as your data source.
	SalesforceConfiguration interface{} `field:"optional" json:"salesforceConfiguration" yaml:"salesforceConfiguration"`
	// Provides the configuration information to connect to ServiceNow as your data source.
	ServiceNowConfiguration interface{} `field:"optional" json:"serviceNowConfiguration" yaml:"serviceNowConfiguration"`
	// Provides the configuration information to connect to Microsoft SharePoint as your data source.
	SharePointConfiguration interface{} `field:"optional" json:"sharePointConfiguration" yaml:"sharePointConfiguration"`
	// Provides the configuration information required for Amazon Kendra Web Crawler.
	WebCrawlerConfiguration interface{} `field:"optional" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
	// Provides the configuration information to connect to Amazon WorkDocs as your data source.
	WorkDocsConfiguration interface{} `field:"optional" json:"workDocsConfiguration" yaml:"workDocsConfiguration"`
}

// Maps a column or attribute in the data source to an index field.
//
// You must first create the fields in the index using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceToIndexFieldMappingProperty := &dataSourceToIndexFieldMappingProperty{
//   	dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   	indexFieldName: jsii.String("indexFieldName"),
//
//   	// the properties below are optional
//   	dateFieldFormat: jsii.String("dateFieldFormat"),
//   }
//
type CfnDataSource_DataSourceToIndexFieldMappingProperty struct {
	// The name of the column or attribute in the data source.
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the field in the index.
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// The type of data stored in the column or attribute.
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Provides the configuration information to connect to an Amazon VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceVpcConfigurationProperty := &dataSourceVpcConfigurationProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnDataSource_DataSourceVpcConfigurationProperty struct {
	// A list of identifiers of security groups within your Amazon VPC.
	//
	// The security groups should enable Amazon Kendra to connect to the data source.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of identifiers for subnets within your Amazon VPC.
	//
	// The subnets should be able to connect to each other in the VPC, and they should have outgoing access to the Internet through a NAT device.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

// Provides the configuration information to connect to a index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseConfigurationProperty := &databaseConfigurationProperty{
//   	columnConfiguration: &columnConfigurationProperty{
//   		changeDetectingColumns: []*string{
//   			jsii.String("changeDetectingColumns"),
//   		},
//   		documentDataColumnName: jsii.String("documentDataColumnName"),
//   		documentIdColumnName: jsii.String("documentIdColumnName"),
//
//   		// the properties below are optional
//   		documentTitleColumnName: jsii.String("documentTitleColumnName"),
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	connectionConfiguration: &connectionConfigurationProperty{
//   		databaseHost: jsii.String("databaseHost"),
//   		databaseName: jsii.String("databaseName"),
//   		databasePort: jsii.Number(123),
//   		secretArn: jsii.String("secretArn"),
//   		tableName: jsii.String("tableName"),
//   	},
//   	databaseEngineType: jsii.String("databaseEngineType"),
//
//   	// the properties below are optional
//   	aclConfiguration: &aclConfigurationProperty{
//   		allowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   	},
//   	sqlConfiguration: &sqlConfigurationProperty{
//   		queryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   	},
//   	vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDataSource_DatabaseConfigurationProperty struct {
	// Information about where the index should get the document information from the database.
	ColumnConfiguration interface{} `field:"required" json:"columnConfiguration" yaml:"columnConfiguration"`
	// Configuration information that's required to connect to a database.
	ConnectionConfiguration interface{} `field:"required" json:"connectionConfiguration" yaml:"connectionConfiguration"`
	// The type of database engine that runs the database.
	DatabaseEngineType *string `field:"required" json:"databaseEngineType" yaml:"databaseEngineType"`
	// Information about the database column that provides information for user context filtering.
	AclConfiguration interface{} `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// Provides information about how Amazon Kendra uses quote marks around SQL identifiers when querying a database data source.
	SqlConfiguration interface{} `field:"optional" json:"sqlConfiguration" yaml:"sqlConfiguration"`
	// Provides information for connecting to an Amazon VPC.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// The condition used for the target document attribute or metadata field when ingesting documents into Amazon Kendra.
//
// You use this with [DocumentAttributeTarget to apply the condition](https://docs.aws.amazon.com/kendra/latest/dg/API_DocumentAttributeTarget.html) .
//
// For example, you can create the 'Department' target field and have it prefill department names associated with the documents based on information in the 'Source_URI' field. Set the condition that if the 'Source_URI' field contains 'financial' in its URI value, then prefill the target field 'Department' with the target value 'Finance' for the document.
//
// Amazon Kendra cannot create a target field if it has not already been created as an index field. After you create your index field, you can create a document metadata field using `DocumentAttributeTarget` . Amazon Kendra then will map your newly created metadata field to your index field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentAttributeConditionProperty := &documentAttributeConditionProperty{
//   	conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   	operator: jsii.String("operator"),
//
//   	// the properties below are optional
//   	conditionOnValue: &documentAttributeValueProperty{
//   		dateValue: jsii.String("dateValue"),
//   		longValue: jsii.Number(123),
//   		stringListValue: []*string{
//   			jsii.String("stringListValue"),
//   		},
//   		stringValue: jsii.String("stringValue"),
//   	},
//   }
//
type CfnDataSource_DocumentAttributeConditionProperty struct {
	// The identifier of the document attribute used for the condition.
	//
	// For example, 'Source_URI' could be an identifier for the attribute or metadata field that contains source URIs associated with the documents.
	//
	// Amazon Kendra currently does not support `_document_body` as an attribute key used for the condition.
	ConditionDocumentAttributeKey *string `field:"required" json:"conditionDocumentAttributeKey" yaml:"conditionDocumentAttributeKey"`
	// The condition operator.
	//
	// For example, you can use 'Contains' to partially match a string.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value used by the operator.
	//
	// For example, you can specify the value 'financial' for strings in the 'Source_URI' field that partially match or contain this value.
	ConditionOnValue interface{} `field:"optional" json:"conditionOnValue" yaml:"conditionOnValue"`
}

// The target document attribute or metadata field you want to alter when ingesting documents into Amazon Kendra.
//
// For example, you can delete customer identification numbers associated with the documents, stored in the document metadata field called 'Customer_ID'. You set the target key as 'Customer_ID' and the deletion flag to `TRUE` . This removes all customer ID values in the field 'Customer_ID'. This would scrub personally identifiable information from each document's metadata.
//
// Amazon Kendra cannot create a target field if it has not already been created as an index field. After you create your index field, you can create a document metadata field using `DocumentAttributeTarget` . Amazon Kendra then will map your newly created metadata field to your index field.
//
// You can also use this with [DocumentAttributeCondition](https://docs.aws.amazon.com/kendra/latest/dg/API_DocumentAttributeCondition.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentAttributeTargetProperty := &documentAttributeTargetProperty{
//   	targetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   	// the properties below are optional
//   	targetDocumentAttributeValue: &documentAttributeValueProperty{
//   		dateValue: jsii.String("dateValue"),
//   		longValue: jsii.Number(123),
//   		stringListValue: []*string{
//   			jsii.String("stringListValue"),
//   		},
//   		stringValue: jsii.String("stringValue"),
//   	},
//   	targetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   }
//
type CfnDataSource_DocumentAttributeTargetProperty struct {
	// The identifier of the target document attribute or metadata field.
	//
	// For example, 'Department' could be an identifier for the target attribute or metadata field that includes the department names associated with the documents.
	TargetDocumentAttributeKey *string `field:"required" json:"targetDocumentAttributeKey" yaml:"targetDocumentAttributeKey"`
	// The target value you want to create for the target attribute.
	//
	// For example, 'Finance' could be the target value for the target attribute key 'Department'.
	TargetDocumentAttributeValue interface{} `field:"optional" json:"targetDocumentAttributeValue" yaml:"targetDocumentAttributeValue"`
	// `TRUE` to delete the existing target value for your specified target attribute key.
	//
	// You cannot create a target value and set this to `TRUE` . To create a target value ( `TargetDocumentAttributeValue` ), set this to `FALSE` .
	TargetDocumentAttributeValueDeletion interface{} `field:"optional" json:"targetDocumentAttributeValueDeletion" yaml:"targetDocumentAttributeValueDeletion"`
}

// The value of a document attribute.
//
// You can only provide one value for a document attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentAttributeValueProperty := &documentAttributeValueProperty{
//   	dateValue: jsii.String("dateValue"),
//   	longValue: jsii.Number(123),
//   	stringListValue: []*string{
//   		jsii.String("stringListValue"),
//   	},
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnDataSource_DocumentAttributeValueProperty struct {
	// A date expressed as an ISO 8601 string.
	//
	// It is important for the time zone to be included in the ISO 8601 date-time format. For example, 2012-03-25T12:30:10+01:00 is the ISO 8601 date-time format for March 25th 2012 at 12:30PM (plus 10 seconds) in Central European Time.
	DateValue *string `field:"optional" json:"dateValue" yaml:"dateValue"`
	// A long integer value.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// A list of strings.
	StringListValue *[]*string `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// A string, such as "department".
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// Document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes.
//
// Each metadata file contains metadata about a single document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentsMetadataConfigurationProperty := &documentsMetadataConfigurationProperty{
//   	s3Prefix: jsii.String("s3Prefix"),
//   }
//
type CfnDataSource_DocumentsMetadataConfigurationProperty struct {
	// A prefix used to filter metadata configuration files in the AWS S3 bucket.
	//
	// The S3 bucket might contain multiple metadata files. Use `S3Prefix` to include only the desired metadata files.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

// Provides the configuration information to connect to Google Drive as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   googleDriveConfigurationProperty := &googleDriveConfigurationProperty{
//   	secretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	excludeMimeTypes: []*string{
//   		jsii.String("excludeMimeTypes"),
//   	},
//   	excludeSharedDrives: []*string{
//   		jsii.String("excludeSharedDrives"),
//   	},
//   	excludeUserAccounts: []*string{
//   		jsii.String("excludeUserAccounts"),
//   	},
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   }
//
type CfnDataSource_GoogleDriveConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of a AWS Secrets Manager secret that contains the credentials required to connect to Google Drive.
	//
	// For more information, see [Using a Google Workspace Drive data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-google-drive.html) .
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// A list of MIME types to exclude from the index. All documents matching the specified MIME type are excluded.
	//
	// For a list of MIME types, see [Using a Google Workspace Drive data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-google-drive.html) .
	ExcludeMimeTypes *[]*string `field:"optional" json:"excludeMimeTypes" yaml:"excludeMimeTypes"`
	// A list of identifiers or shared drives to exclude from the index.
	//
	// All files and folders stored on the shared drive are excluded.
	ExcludeSharedDrives *[]*string `field:"optional" json:"excludeSharedDrives" yaml:"excludeSharedDrives"`
	// A list of email addresses of the users.
	//
	// Documents owned by these users are excluded from the index. Documents shared with excluded users are indexed unless they are excluded in another way.
	ExcludeUserAccounts *[]*string `field:"optional" json:"excludeUserAccounts" yaml:"excludeUserAccounts"`
	// A list of regular expression patterns to exclude certain items in your Google Drive, including shared drives and users' My Drives.
	//
	// Items that match the patterns are excluded from the index. Items that don't match the patterns are included in the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Maps Google Drive data source attributes or field names to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Google Drive fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Google Drive data source field names must exist in your Google Drive custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain items in your Google Drive, including shared drives and users' My Drives.
	//
	// Items that match the patterns are included in the index. Items that don't match the patterns are excluded from the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
}

// Provides the configuration information for invoking a Lambda function in AWS Lambda to alter document metadata and content when ingesting documents into Amazon Kendra.
//
// You can configure your Lambda function using [PreExtractionHookConfiguration](https://docs.aws.amazon.com/kendra/latest/dg/API_CustomDocumentEnrichmentConfiguration.html) if you want to apply advanced alterations on the original or raw documents. If you want to apply advanced alterations on the Amazon Kendra structured documents, you must configure your Lambda function using [PostExtractionHookConfiguration](https://docs.aws.amazon.com/kendra/latest/dg/API_CustomDocumentEnrichmentConfiguration.html) . You can only invoke one Lambda function. However, this function can invoke other functions it requires.
//
// For more information, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hookConfigurationProperty := &hookConfigurationProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   	s3Bucket: jsii.String("s3Bucket"),
//
//   	// the properties below are optional
//   	invocationCondition: &documentAttributeConditionProperty{
//   		conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   		operator: jsii.String("operator"),
//
//   		// the properties below are optional
//   		conditionOnValue: &documentAttributeValueProperty{
//   			dateValue: jsii.String("dateValue"),
//   			longValue: jsii.Number(123),
//   			stringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   }
//
type CfnDataSource_HookConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of a role with permission to run a Lambda function during ingestion.
	//
	// For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html) .
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Stores the original, raw documents or the structured, parsed documents before and after altering them.
	//
	// For more information, see [Data contracts for Lambda functions](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#cde-data-contracts-lambda) .
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The condition used for when a Lambda function should be invoked.
	//
	// For example, you can specify a condition that if there are empty date-time values, then Amazon Kendra should invoke a function that inserts the current date-time.
	InvocationCondition interface{} `field:"optional" json:"invocationCondition" yaml:"invocationCondition"`
}

// Provides the configuration information for applying basic logic to alter document metadata and content when ingesting documents into Amazon Kendra.
//
// To apply advanced logic, to go beyond what you can do with basic logic, see [HookConfiguration](https://docs.aws.amazon.com/kendra/latest/dg/API_HookConfiguration.html) .
//
// For more information, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inlineCustomDocumentEnrichmentConfigurationProperty := &inlineCustomDocumentEnrichmentConfigurationProperty{
//   	condition: &documentAttributeConditionProperty{
//   		conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   		operator: jsii.String("operator"),
//
//   		// the properties below are optional
//   		conditionOnValue: &documentAttributeValueProperty{
//   			dateValue: jsii.String("dateValue"),
//   			longValue: jsii.Number(123),
//   			stringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	documentContentDeletion: jsii.Boolean(false),
//   	target: &documentAttributeTargetProperty{
//   		targetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   		// the properties below are optional
//   		targetDocumentAttributeValue: &documentAttributeValueProperty{
//   			dateValue: jsii.String("dateValue"),
//   			longValue: jsii.Number(123),
//   			stringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			stringValue: jsii.String("stringValue"),
//   		},
//   		targetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   	},
//   }
//
type CfnDataSource_InlineCustomDocumentEnrichmentConfigurationProperty struct {
	// Configuration of the condition used for the target document attribute or metadata field when ingesting documents into Amazon Kendra.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// `TRUE` to delete content if the condition used for the target attribute is met.
	DocumentContentDeletion interface{} `field:"optional" json:"documentContentDeletion" yaml:"documentContentDeletion"`
	// Configuration of the target document attribute or metadata field when ingesting documents into Amazon Kendra.
	//
	// You can also include a value.
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

// Provides the configuration information to connect to OneDrive as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oneDriveConfigurationProperty := &oneDriveConfigurationProperty{
//   	oneDriveUsers: &oneDriveUsersProperty{
//   		oneDriveUserList: []*string{
//   			jsii.String("oneDriveUserList"),
//   		},
//   		oneDriveUserS3Path: &s3PathProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//   		},
//   	},
//   	secretArn: jsii.String("secretArn"),
//   	tenantDomain: jsii.String("tenantDomain"),
//
//   	// the properties below are optional
//   	disableLocalGroups: jsii.Boolean(false),
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   }
//
type CfnDataSource_OneDriveConfigurationProperty struct {
	// A list of user accounts whose documents should be indexed.
	OneDriveUsers interface{} `field:"required" json:"oneDriveUsers" yaml:"oneDriveUsers"`
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the user name and password to connect to OneDrive.
	//
	// The user named should be the application ID for the OneDrive application, and the password is the application key for the OneDrive application.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The Azure Active Directory domain of the organization.
	TenantDomain *string `field:"required" json:"tenantDomain" yaml:"tenantDomain"`
	// A Boolean value that specifies whether local groups are disabled ( `True` ) or enabled ( `False` ).
	DisableLocalGroups interface{} `field:"optional" json:"disableLocalGroups" yaml:"disableLocalGroups"`
	// A list of regular expression patterns to exclude certain documents in your OneDrive.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the file name.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map OneDrive data source attributes or field names to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to OneDrive fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The OneDrive data source field names must exist in your OneDrive custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain documents in your OneDrive.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the file name.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
}

// User accounts whose documents should be indexed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oneDriveUsersProperty := &oneDriveUsersProperty{
//   	oneDriveUserList: []*string{
//   		jsii.String("oneDriveUserList"),
//   	},
//   	oneDriveUserS3Path: &s3PathProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataSource_OneDriveUsersProperty struct {
	// A list of users whose documents should be indexed.
	//
	// Specify the user names in email format, for example, `username@tenantdomain` . If you need to index the documents of more than 100 users, use the `OneDriveUserS3Path` field to specify the location of a file containing a list of users.
	OneDriveUserList *[]*string `field:"optional" json:"oneDriveUserList" yaml:"oneDriveUserList"`
	// The S3 bucket location of a file containing a list of users whose documents should be indexed.
	OneDriveUserS3Path interface{} `field:"optional" json:"oneDriveUserS3Path" yaml:"oneDriveUserS3Path"`
}

// Provides the configuration information for a web proxy to connect to website hosts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proxyConfigurationProperty := &proxyConfigurationProperty{
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//
//   	// the properties below are optional
//   	credentials: jsii.String("credentials"),
//   }
//
type CfnDataSource_ProxyConfigurationProperty struct {
	// The name of the website host you want to connect to via a web proxy server.
	//
	// For example, the host name of https://a.example.com/page1.html is "a.example.com".
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port number of the website host you want to connect to via a web proxy server.
	//
	// For example, the port for https://a.example.com/page1.html is 443, the standard port for HTTPS.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Your secret ARN, which you can create in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html).
	//
	// The credentials are optional. You use a secret if web proxy credentials are required to connect to a website host. Amazon Kendra currently support basic authentication to connect to a web proxy server. The secret stores your credentials.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
}

// Provides the configuration information to connect to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DataSourceConfigurationProperty := &s3DataSourceConfigurationProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	accessControlListConfiguration: &accessControlListConfigurationProperty{
//   		keyPath: jsii.String("keyPath"),
//   	},
//   	documentsMetadataConfiguration: &documentsMetadataConfigurationProperty{
//   		s3Prefix: jsii.String("s3Prefix"),
//   	},
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	inclusionPrefixes: []*string{
//   		jsii.String("inclusionPrefixes"),
//   	},
//   }
//
type CfnDataSource_S3DataSourceConfigurationProperty struct {
	// The name of the bucket that contains the documents.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Provides the path to the S3 bucket that contains the user context filtering files for the data source.
	//
	// For the format of the file, see [Access control for S3 data sources](https://docs.aws.amazon.com/kendra/latest/dg/s3-acl.html) .
	AccessControlListConfiguration interface{} `field:"optional" json:"accessControlListConfiguration" yaml:"accessControlListConfiguration"`
	// Specifies document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes.
	//
	// Each metadata file contains metadata about a single document.
	DocumentsMetadataConfiguration interface{} `field:"optional" json:"documentsMetadataConfiguration" yaml:"documentsMetadataConfiguration"`
	// A list of glob patterns for documents that should not be indexed.
	//
	// If a document that matches an inclusion prefix or inclusion pattern also matches an exclusion pattern, the document is not indexed.
	//
	// Some [examples](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) are:
	//
	// - **.png , *.jpg* will exclude all PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
	// - **internal** will exclude all files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
	// - *** /*internal** will exclude all internal-related files in a directory and its subdirectories.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of glob patterns for documents that should be indexed.
	//
	// If a document that matches an inclusion pattern also matches an exclusion pattern, the document is not indexed.
	//
	// Some [examples](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) are:
	//
	// - **.txt* will include all text files in a directory (files with the extension .txt).
	// - *** /*.txt* will include all text files in a directory and its subdirectories.
	// - **tax** will include all files in a directory that contain 'tax' in the file name, such as 'tax', 'taxes', 'income_tax'.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// A list of S3 prefixes for the documents that should be included in the index.
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

// Information required to find a specific file in an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3PathProperty := &s3PathProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   }
//
type CfnDataSource_S3PathProperty struct {
	// The name of the S3 bucket that contains the file.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the file.
	Key *string `field:"required" json:"key" yaml:"key"`
}

// The configuration information for syncing a Salesforce chatter feed.
//
// The contents of the object comes from the Salesforce FeedItem table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceChatterFeedConfigurationProperty := &salesforceChatterFeedConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   	// the properties below are optional
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	includeFilterTypes: []*string{
//   		jsii.String("includeFilterTypes"),
//   	},
//   }
//
type CfnDataSource_SalesforceChatterFeedConfigurationProperty struct {
	// The name of the column in the Salesforce FeedItem table that contains the content to index.
	//
	// Typically this is the `Body` column.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the column in the Salesforce FeedItem table that contains the title of the document.
	//
	// This is typically the `Title` column.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps fields from a Salesforce chatter feed into Amazon Kendra index fields.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Filters the documents in the feed based on status of the user.
	//
	// When you specify `ACTIVE_USERS` only documents from users who have an active account are indexed. When you specify `STANDARD_USER` only documents for Salesforce standard users are documented. You can specify both.
	IncludeFilterTypes *[]*string `field:"optional" json:"includeFilterTypes" yaml:"includeFilterTypes"`
}

// Provides the configuration information to connect to Salesforce as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceConfigurationProperty := &salesforceConfigurationProperty{
//   	secretArn: jsii.String("secretArn"),
//   	serverUrl: jsii.String("serverUrl"),
//
//   	// the properties below are optional
//   	chatterFeedConfiguration: &salesforceChatterFeedConfigurationProperty{
//   		documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		includeFilterTypes: []*string{
//   			jsii.String("includeFilterTypes"),
//   		},
//   	},
//   	crawlAttachments: jsii.Boolean(false),
//   	excludeAttachmentFilePatterns: []*string{
//   		jsii.String("excludeAttachmentFilePatterns"),
//   	},
//   	includeAttachmentFilePatterns: []*string{
//   		jsii.String("includeAttachmentFilePatterns"),
//   	},
//   	knowledgeArticleConfiguration: &salesforceKnowledgeArticleConfigurationProperty{
//   		includedStates: []*string{
//   			jsii.String("includedStates"),
//   		},
//
//   		// the properties below are optional
//   		customKnowledgeArticleTypeConfigurations: []interface{}{
//   			&salesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   		},
//   		standardKnowledgeArticleTypeConfiguration: &salesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   			documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   			// the properties below are optional
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   	},
//   	standardObjectAttachmentConfiguration: &salesforceStandardObjectAttachmentConfigurationProperty{
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	standardObjectConfigurations: []interface{}{
//   		&salesforceStandardObjectConfigurationProperty{
//   			documentDataFieldName: jsii.String("documentDataFieldName"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the key/value pairs required to connect to your Salesforce instance.
	//
	// The secret must contain a JSON structure with the following keys:
	//
	// - authenticationUrl - The OAUTH endpoint that Amazon Kendra connects to get an OAUTH token.
	// - consumerKey - The application public key generated when you created your Salesforce application.
	// - consumerSecret - The application private key generated when you created your Salesforce application.
	// - password - The password associated with the user logging in to the Salesforce instance.
	// - securityToken - The token associated with the user account logging in to the Salesforce instance.
	// - username - The user name of the user logging in to the Salesforce instance.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The instance URL for the Salesforce site that you want to index.
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Configuration information for Salesforce chatter feeds.
	ChatterFeedConfiguration interface{} `field:"optional" json:"chatterFeedConfiguration" yaml:"chatterFeedConfiguration"`
	// Indicates whether Amazon Kendra should index attachments to Salesforce objects.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// A list of regular expression patterns to exclude certain documents in your Salesforce.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the name of the attached file.
	ExcludeAttachmentFilePatterns *[]*string `field:"optional" json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// A list of regular expression patterns to include certain documents in your Salesforce.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the name of the attached file.
	IncludeAttachmentFilePatterns *[]*string `field:"optional" json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
	// Configuration information for the knowledge article types that Amazon Kendra indexes.
	//
	// Amazon Kendra indexes standard knowledge articles and the standard fields of knowledge articles, or the custom fields of custom knowledge articles, but not both.
	KnowledgeArticleConfiguration interface{} `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Configuration information for processing attachments to Salesforce standard objects.
	StandardObjectAttachmentConfiguration interface{} `field:"optional" json:"standardObjectAttachmentConfiguration" yaml:"standardObjectAttachmentConfiguration"`
	// Configuration of the Salesforce standard objects that Amazon Kendra indexes.
	StandardObjectConfigurations interface{} `field:"optional" json:"standardObjectConfigurations" yaml:"standardObjectConfigurations"`
}

// Provides the configuration information for indexing Salesforce custom articles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceCustomKnowledgeArticleTypeConfigurationProperty := &salesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceCustomKnowledgeArticleTypeConfigurationProperty struct {
	// The name of the field in the custom knowledge article that contains the document data to index.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the field in the custom knowledge article that contains the document title.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps attributes or field names of the custom knowledge article to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Salesforce fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Salesforce data source field names must exist in your Salesforce custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

// Provides the configuration information for the knowledge article types that Amazon Kendra indexes.
//
// Amazon Kendra indexes standard knowledge articles and the standard fields of knowledge articles, or the custom fields of custom knowledge articles, but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceKnowledgeArticleConfigurationProperty := &salesforceKnowledgeArticleConfigurationProperty{
//   	includedStates: []*string{
//   		jsii.String("includedStates"),
//   	},
//
//   	// the properties below are optional
//   	customKnowledgeArticleTypeConfigurations: []interface{}{
//   		&salesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   			documentDataFieldName: jsii.String("documentDataFieldName"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   	},
//   	standardKnowledgeArticleTypeConfiguration: &salesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   		documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceKnowledgeArticleConfigurationProperty struct {
	// Specifies the document states that should be included when Amazon Kendra indexes knowledge articles.
	//
	// You must specify at least one state.
	IncludedStates *[]*string `field:"required" json:"includedStates" yaml:"includedStates"`
	// Configuration information for custom Salesforce knowledge articles.
	CustomKnowledgeArticleTypeConfigurations interface{} `field:"optional" json:"customKnowledgeArticleTypeConfigurations" yaml:"customKnowledgeArticleTypeConfigurations"`
	// Configuration information for standard Salesforce knowledge articles.
	StandardKnowledgeArticleTypeConfiguration interface{} `field:"optional" json:"standardKnowledgeArticleTypeConfiguration" yaml:"standardKnowledgeArticleTypeConfiguration"`
}

// Provides the configuration information for standard Salesforce knowledge articles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceStandardKnowledgeArticleTypeConfigurationProperty := &salesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   	// the properties below are optional
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceStandardKnowledgeArticleTypeConfigurationProperty struct {
	// The name of the field that contains the document data to index.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the field that contains the document title.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps attributes or field names of the knowledge article to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Salesforce fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Salesforce data source field names must exist in your Salesforce custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

// Provides the configuration information for processing attachments to Salesforce standard objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceStandardObjectAttachmentConfigurationProperty := &salesforceStandardObjectAttachmentConfigurationProperty{
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceStandardObjectAttachmentConfigurationProperty struct {
	// The name of the field used for the document title.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// One or more objects that map fields in attachments to Amazon Kendra index fields.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

// Specifies configuration information for indexing a single standard object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceStandardObjectConfigurationProperty := &salesforceStandardObjectConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceStandardObjectConfigurationProperty struct {
	// The name of the field in the standard object table that contains the document contents.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the standard object.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the field in the standard object table that contains the document title.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps attributes or field names of the standard object to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Salesforce fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Salesforce data source field names must exist in your Salesforce custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

// Provides the configuration information to connect to ServiceNow as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowConfigurationProperty := &serviceNowConfigurationProperty{
//   	hostUrl: jsii.String("hostUrl"),
//   	secretArn: jsii.String("secretArn"),
//   	serviceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//
//   	// the properties below are optional
//   	authenticationType: jsii.String("authenticationType"),
//   	knowledgeArticleConfiguration: &serviceNowKnowledgeArticleConfigurationProperty{
//   		documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		crawlAttachments: jsii.Boolean(false),
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		excludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		filterQuery: jsii.String("filterQuery"),
//   		includeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   	},
//   	serviceCatalogConfiguration: &serviceNowServiceCatalogConfigurationProperty{
//   		documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		crawlAttachments: jsii.Boolean(false),
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		excludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		includeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   	},
//   }
//
type CfnDataSource_ServiceNowConfigurationProperty struct {
	// The ServiceNow instance that the data source connects to.
	//
	// The host endpoint should look like the following: *{instance}.service-now.com.*
	HostUrl *string `field:"required" json:"hostUrl" yaml:"hostUrl"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the user name and password required to connect to the ServiceNow instance.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The identifier of the release that the ServiceNow host is running.
	//
	// If the host is not running the `LONDON` release, use `OTHERS` .
	ServiceNowBuildVersion *string `field:"required" json:"serviceNowBuildVersion" yaml:"serviceNowBuildVersion"`
	// The type of authentication used to connect to the ServiceNow instance.
	//
	// If you choose `HTTP_BASIC` , Amazon Kendra is authenticated using the user name and password provided in the AWS Secrets Manager secret in the `SecretArn` field. When you choose `OAUTH2` , Amazon Kendra is authenticated using the OAuth token and secret provided in the Secrets Manager secret, and the user name and password are used to determine which information Amazon Kendra has access to.
	//
	// When you use `OAUTH2` authentication, you must generate a token and a client secret using the ServiceNow console. For more information, see [Using a ServiceNow data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-servicenow.html) .
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Configuration information for crawling knowledge articles in the ServiceNow site.
	KnowledgeArticleConfiguration interface{} `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Configuration information for crawling service catalogs in the ServiceNow site.
	ServiceCatalogConfiguration interface{} `field:"optional" json:"serviceCatalogConfiguration" yaml:"serviceCatalogConfiguration"`
}

// Provides the configuration information for crawling knowledge articles in the ServiceNow site.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowKnowledgeArticleConfigurationProperty := &serviceNowKnowledgeArticleConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   	// the properties below are optional
//   	crawlAttachments: jsii.Boolean(false),
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	excludeAttachmentFilePatterns: []*string{
//   		jsii.String("excludeAttachmentFilePatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	filterQuery: jsii.String("filterQuery"),
//   	includeAttachmentFilePatterns: []*string{
//   		jsii.String("includeAttachmentFilePatterns"),
//   	},
//   }
//
type CfnDataSource_ServiceNowKnowledgeArticleConfigurationProperty struct {
	// The name of the ServiceNow field that is mapped to the index document contents field in the Amazon Kendra index.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// Indicates whether Amazon Kendra should index attachments to knowledge articles.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// The name of the ServiceNow field that is mapped to the index document title field.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// A list of regular expression patterns to exclude certain attachments of knowledge articles in your ServiceNow.
	//
	// Item that match the patterns are excluded from the index. Items that don't match the patterns are included in the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	//
	// The regex is applied to the field specified in the `PatternTargetField` .
	ExcludeAttachmentFilePatterns *[]*string `field:"optional" json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// Maps attributes or field names of knoweldge articles to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to ServiceNow fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The ServiceNow data source field names must exist in your ServiceNow custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A query that selects the knowledge articles to index.
	//
	// The query can return articles from multiple knowledge bases, and the knowledge bases can be public or private.
	//
	// The query string must be one generated by the ServiceNow console. For more information, see [Specifying documents to index with a query](https://docs.aws.amazon.com/kendra/latest/dg/servicenow-query.html) .
	FilterQuery *string `field:"optional" json:"filterQuery" yaml:"filterQuery"`
	// A list of regular expression patterns to include certain attachments of knowledge articles in your ServiceNow.
	//
	// Item that match the patterns are included in the index. Items that don't match the patterns are excluded from the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	//
	// The regex is applied to the field specified in the `PatternTargetField` .
	IncludeAttachmentFilePatterns *[]*string `field:"optional" json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
}

// Provides the configuration information for crawling service catalog items in the ServiceNow site.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowServiceCatalogConfigurationProperty := &serviceNowServiceCatalogConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   	// the properties below are optional
//   	crawlAttachments: jsii.Boolean(false),
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	excludeAttachmentFilePatterns: []*string{
//   		jsii.String("excludeAttachmentFilePatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	includeAttachmentFilePatterns: []*string{
//   		jsii.String("includeAttachmentFilePatterns"),
//   	},
//   }
//
type CfnDataSource_ServiceNowServiceCatalogConfigurationProperty struct {
	// The name of the ServiceNow field that is mapped to the index document contents field in the Amazon Kendra index.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// Indicates whether Amazon Kendra should crawl attachments to the service catalog items.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// The name of the ServiceNow field that is mapped to the index document title field.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// A list of regular expression patterns to exclude certain attachments of catalogs in your ServiceNow.
	//
	// Item that match the patterns are excluded from the index. Items that don't match the patterns are included in the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	//
	// The regex is applied to the file name of the attachment.
	ExcludeAttachmentFilePatterns *[]*string `field:"optional" json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// Maps attributes or field names of catalogs to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to ServiceNow fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The ServiceNow data source field names must exist in your ServiceNow custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain attachments of catalogs in your ServiceNow.
	//
	// Item that match the patterns are included in the index. Items that don't match the patterns are excluded from the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	//
	// The regex is applied to the file name of the attachment.
	IncludeAttachmentFilePatterns *[]*string `field:"optional" json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
}

// Provides the configuration information to connect to Microsoft SharePoint as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharePointConfigurationProperty := &sharePointConfigurationProperty{
//   	secretArn: jsii.String("secretArn"),
//   	sharePointVersion: jsii.String("sharePointVersion"),
//   	urls: []*string{
//   		jsii.String("urls"),
//   	},
//
//   	// the properties below are optional
//   	crawlAttachments: jsii.Boolean(false),
//   	disableLocalGroups: jsii.Boolean(false),
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	sslCertificateS3Path: &s3PathProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//   	useChangeLog: jsii.Boolean(false),
//   	vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDataSource_SharePointConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of credentials stored in AWS Secrets Manager .
	//
	// The credentials should be a user/password pair. If you use SharePoint Server, you also need to provide the sever domain name as part of the credentials. For more information, see [Using a Microsoft SharePoint Data Source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-sharepoint.html) . For more information about AWS Secrets Manager see [What Is AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) in the *AWS Secrets Manager* user guide.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The version of Microsoft SharePoint that you are using as a data source.
	SharePointVersion *string `field:"required" json:"sharePointVersion" yaml:"sharePointVersion"`
	// The URLs of the Microsoft SharePoint site that contains the documents that should be indexed.
	Urls *[]*string `field:"required" json:"urls" yaml:"urls"`
	// `TRUE` to include attachments to documents stored in your Microsoft SharePoint site in the index;
	//
	// otherwise, `FALSE` .
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// A Boolean value that specifies whether local groups are disabled ( `True` ) or enabled ( `False` ).
	DisableLocalGroups interface{} `field:"optional" json:"disableLocalGroups" yaml:"disableLocalGroups"`
	// The Microsoft SharePoint attribute field that contains the title of the document.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// A list of regular expression patterns.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an exclusion pattern and an inclusion pattern, the document is not included in the index.
	//
	// The regex is applied to the display URL of the SharePoint document.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map Microsoft SharePoint attributes to custom fields in the Amazon Kendra index.
	//
	// You must first create the index fields using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation before you map SharePoint attributes. For more information, see [Mapping Data Source Fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) .
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain documents in your SharePoint.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The regex is applied to the display URL of the SharePoint document.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Information required to find a specific file in an Amazon S3 bucket.
	SslCertificateS3Path interface{} `field:"optional" json:"sslCertificateS3Path" yaml:"sslCertificateS3Path"`
	// `TRUE` to use the SharePoint change log to determine which documents require updating in the index.
	//
	// Depending on the change log's size, it may take longer for Amazon Kendra to use the change log than to scan all of your documents in SharePoint.
	UseChangeLog interface{} `field:"optional" json:"useChangeLog" yaml:"useChangeLog"`
	// Provides information for connecting to an Amazon VPC.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// Provides information that configures Amazon Kendra to use a SQL database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqlConfigurationProperty := &sqlConfigurationProperty{
//   	queryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   }
//
type CfnDataSource_SqlConfigurationProperty struct {
	// Determines whether Amazon Kendra encloses SQL identifiers for tables and column names in double quotes (") when making a database query.
	//
	// You can set the value to `DOUBLE_QUOTES` or `NONE` .
	//
	// By default, Amazon Kendra passes SQL identifiers the way that they are entered into the data source configuration. It does not change the case of identifiers or enclose them in quotes.
	//
	// PostgreSQL internally converts uppercase characters to lower case characters in identifiers unless they are quoted. Choosing this option encloses identifiers in quotes so that PostgreSQL does not convert the character's case.
	//
	// For MySQL databases, you must enable the ansi_quotes option when you set this field to `DOUBLE_QUOTES` .
	QueryIdentifiersEnclosingOption *string `field:"optional" json:"queryIdentifiersEnclosingOption" yaml:"queryIdentifiersEnclosingOption"`
}

// Provides the configuration information to connect to websites that require user authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerAuthenticationConfigurationProperty := &webCrawlerAuthenticationConfigurationProperty{
//   	basicAuthentication: []interface{}{
//   		&webCrawlerBasicAuthenticationProperty{
//   			credentials: jsii.String("credentials"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnDataSource_WebCrawlerAuthenticationConfigurationProperty struct {
	// The list of configuration information that's required to connect to and crawl a website host using basic authentication credentials.
	//
	// The list includes the name and port number of the website host.
	BasicAuthentication interface{} `field:"optional" json:"basicAuthentication" yaml:"basicAuthentication"`
}

// Provides the configuration information to connect to websites that require basic user authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerBasicAuthenticationProperty := &webCrawlerBasicAuthenticationProperty{
//   	credentials: jsii.String("credentials"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_WebCrawlerBasicAuthenticationProperty struct {
	// Your secret ARN, which you can create in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html).
	//
	// You use a secret if basic authentication credentials are required to connect to a website. The secret stores your credentials of user name and password.
	Credentials *string `field:"required" json:"credentials" yaml:"credentials"`
	// The name of the website host you want to connect to using authentication credentials.
	//
	// For example, the host name of https://a.example.com/page1.html is "a.example.com".
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port number of the website host you want to connect to using authentication credentials.
	//
	// For example, the port for https://a.example.com/page1.html is 443, the standard port for HTTPS.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// Provides the configuration information required for Amazon Kendra Web Crawler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerConfigurationProperty := &webCrawlerConfigurationProperty{
//   	urls: &webCrawlerUrlsProperty{
//   		seedUrlConfiguration: &webCrawlerSeedUrlConfigurationProperty{
//   			seedUrls: []*string{
//   				jsii.String("seedUrls"),
//   			},
//
//   			// the properties below are optional
//   			webCrawlerMode: jsii.String("webCrawlerMode"),
//   		},
//   		siteMapsConfiguration: &webCrawlerSiteMapsConfigurationProperty{
//   			siteMaps: []*string{
//   				jsii.String("siteMaps"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	authenticationConfiguration: &webCrawlerAuthenticationConfigurationProperty{
//   		basicAuthentication: []interface{}{
//   			&webCrawlerBasicAuthenticationProperty{
//   				credentials: jsii.String("credentials"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	crawlDepth: jsii.Number(123),
//   	maxContentSizePerPageInMegaBytes: jsii.Number(123),
//   	maxLinksPerPage: jsii.Number(123),
//   	maxUrlsPerMinuteCrawlRate: jsii.Number(123),
//   	proxyConfiguration: &proxyConfigurationProperty{
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//
//   		// the properties below are optional
//   		credentials: jsii.String("credentials"),
//   	},
//   	urlExclusionPatterns: []*string{
//   		jsii.String("urlExclusionPatterns"),
//   	},
//   	urlInclusionPatterns: []*string{
//   		jsii.String("urlInclusionPatterns"),
//   	},
//   }
//
type CfnDataSource_WebCrawlerConfigurationProperty struct {
	// Specifies the seed or starting point URLs of the websites or the sitemap URLs of the websites you want to crawl.
	//
	// You can include website subdomains. You can list up to 100 seed URLs and up to three sitemap URLs.
	//
	// You can only crawl websites that use the secure communication protocol, Hypertext Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it could be that the website is blocked from crawling.
	//
	// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use Amazon Kendra Web Crawler to index your own webpages, or webpages that you have authorization to index.*
	Urls interface{} `field:"required" json:"urls" yaml:"urls"`
	// Configuration information required to connect to websites using authentication.
	//
	// You can connect to websites using basic authentication of user name and password.
	//
	// You must provide the website host name and port number. For example, the host name of https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard port for HTTPS. You use a secret in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) to store your authentication credentials.
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Specifies the number of levels in a website that you want to crawl.
	//
	// The first level begins from the website seed or starting point URL. For example, if a website has 3 levels – index level (i.e. seed in this example), sections level, and subsections level – and you are only interested in crawling information up to the sections level (i.e. levels 0-1), you can set your depth to 1.
	//
	// The default crawl depth is set to 2.
	CrawlDepth *float64 `field:"optional" json:"crawlDepth" yaml:"crawlDepth"`
	// The maximum size (in MB) of a webpage or attachment to crawl.
	//
	// Files larger than this size (in MB) are skipped/not crawled.
	//
	// The default maximum size of a webpage or attachment is set to 50 MB.
	MaxContentSizePerPageInMegaBytes *float64 `field:"optional" json:"maxContentSizePerPageInMegaBytes" yaml:"maxContentSizePerPageInMegaBytes"`
	// The maximum number of URLs on a webpage to include when crawling a website. This number is per webpage.
	//
	// As a website’s webpages are crawled, any URLs the webpages link to are also crawled. URLs on a webpage are crawled in order of appearance.
	//
	// The default maximum links per page is 100.
	MaxLinksPerPage *float64 `field:"optional" json:"maxLinksPerPage" yaml:"maxLinksPerPage"`
	// The maximum number of URLs crawled per website host per minute.
	//
	// A minimum of one URL is required.
	//
	// The default maximum number of URLs crawled per website host per minute is 300.
	MaxUrlsPerMinuteCrawlRate *float64 `field:"optional" json:"maxUrlsPerMinuteCrawlRate" yaml:"maxUrlsPerMinuteCrawlRate"`
	// Configuration information required to connect to your internal websites via a web proxy.
	//
	// You must provide the website host name and port number. For example, the host name of https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard port for HTTPS.
	//
	// Web proxy credentials are optional and you can use them to connect to a web proxy server that requires basic authentication. To store web proxy credentials, you use a secret in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) .
	ProxyConfiguration interface{} `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// A list of regular expression patterns to exclude certain URLs to crawl.
	//
	// URLs that match the patterns are excluded from the index. URLs that don't match the patterns are included in the index. If a URL matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the URL file isn't included in the index.
	UrlExclusionPatterns *[]*string `field:"optional" json:"urlExclusionPatterns" yaml:"urlExclusionPatterns"`
	// A list of regular expression patterns to include certain URLs to crawl.
	//
	// URLs that match the patterns are included in the index. URLs that don't match the patterns are excluded from the index. If a URL matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the URL file isn't included in the index.
	UrlInclusionPatterns *[]*string `field:"optional" json:"urlInclusionPatterns" yaml:"urlInclusionPatterns"`
}

// Provides the configuration information of the seed or starting point URLs to crawl.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerSeedUrlConfigurationProperty := &webCrawlerSeedUrlConfigurationProperty{
//   	seedUrls: []*string{
//   		jsii.String("seedUrls"),
//   	},
//
//   	// the properties below are optional
//   	webCrawlerMode: jsii.String("webCrawlerMode"),
//   }
//
type CfnDataSource_WebCrawlerSeedUrlConfigurationProperty struct {
	// The list of seed or starting point URLs of the websites you want to crawl.
	//
	// The list can include a maximum of 100 seed URLs.
	SeedUrls *[]*string `field:"required" json:"seedUrls" yaml:"seedUrls"`
	// You can choose one of the following modes:.
	//
	// - `HOST_ONLY` – crawl only the website host names. For example, if the seed URL is "abc.example.com", then only URLs with host name "abc.example.com" are crawled.
	// - `SUBDOMAINS` – crawl the website host names with subdomains. For example, if the seed URL is "abc.example.com", then "a.abc.example.com" and "b.abc.example.com" are also crawled.
	// - `EVERYTHING` – crawl the website host names with subdomains and other domains that the webpages link to.
	//
	// The default mode is set to `HOST_ONLY` .
	WebCrawlerMode *string `field:"optional" json:"webCrawlerMode" yaml:"webCrawlerMode"`
}

// Provides the configuration information of the sitemap URLs to crawl.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerSiteMapsConfigurationProperty := &webCrawlerSiteMapsConfigurationProperty{
//   	siteMaps: []*string{
//   		jsii.String("siteMaps"),
//   	},
//   }
//
type CfnDataSource_WebCrawlerSiteMapsConfigurationProperty struct {
	// The list of sitemap URLs of the websites you want to crawl.
	//
	// The list can include a maximum of three sitemap URLs.
	SiteMaps *[]*string `field:"required" json:"siteMaps" yaml:"siteMaps"`
}

// Specifies the seed or starting point URLs of the websites or the sitemap URLs of the websites you want to crawl.
//
// You can include website subdomains. You can list up to 100 seed URLs and up to three sitemap URLs.
//
// You can only crawl websites that use the secure communication protocol, Hypertext Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it could be that the website is blocked from crawling.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerUrlsProperty := &webCrawlerUrlsProperty{
//   	seedUrlConfiguration: &webCrawlerSeedUrlConfigurationProperty{
//   		seedUrls: []*string{
//   			jsii.String("seedUrls"),
//   		},
//
//   		// the properties below are optional
//   		webCrawlerMode: jsii.String("webCrawlerMode"),
//   	},
//   	siteMapsConfiguration: &webCrawlerSiteMapsConfigurationProperty{
//   		siteMaps: []*string{
//   			jsii.String("siteMaps"),
//   		},
//   	},
//   }
//
type CfnDataSource_WebCrawlerUrlsProperty struct {
	// Configuration of the seed or starting point URLs of the websites you want to crawl.
	//
	// You can choose to crawl only the website host names, or the website host names with subdomains, or the website host names with subdomains and other domains that the webpages link to.
	//
	// You can list up to 100 seed URLs.
	SeedUrlConfiguration interface{} `field:"optional" json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// Configuration of the sitemap URLs of the websites you want to crawl.
	//
	// Only URLs belonging to the same website host names are crawled. You can list up to three sitemap URLs.
	SiteMapsConfiguration interface{} `field:"optional" json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

// Provides the configuration information to connect to Amazon WorkDocs as your data source.
//
// Amazon WorkDocs connector is available in Oregon, North Virginia, Sydney, Singapore and Ireland regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workDocsConfigurationProperty := &workDocsConfigurationProperty{
//   	organizationId: jsii.String("organizationId"),
//
//   	// the properties below are optional
//   	crawlComments: jsii.Boolean(false),
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	useChangeLog: jsii.Boolean(false),
//   }
//
type CfnDataSource_WorkDocsConfigurationProperty struct {
	// The identifier of the directory corresponding to your Amazon WorkDocs site repository.
	//
	// You can find the organization ID in the [AWS Directory Service](https://docs.aws.amazon.com/directoryservicev2/) by going to *Active Directory* , then *Directories* . Your Amazon WorkDocs site directory has an ID, which is the organization ID. You can also set up a new Amazon WorkDocs directory in the AWS Directory Service console and enable a Amazon WorkDocs site for the directory in the Amazon WorkDocs console.
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// `TRUE` to include comments on documents in your index.
	//
	// Including comments in your index means each comment is a document that can be searched on.
	//
	// The default is set to `FALSE` .
	CrawlComments interface{} `field:"optional" json:"crawlComments" yaml:"crawlComments"`
	// A list of regular expression patterns to exclude certain files in your Amazon WorkDocs site repository.
	//
	// Files that match the patterns are excluded from the index. Files that don’t match the patterns are included in the index. If a file matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the file isn't included in the index.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map Amazon WorkDocs data source attributes or field names to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Amazon WorkDocs fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Amazon WorkDocs data source field names must exist in your Amazon WorkDocs custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain files in your Amazon WorkDocs site repository.
	//
	// Files that match the patterns are included in the index. Files that don't match the patterns are excluded from the index. If a file matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the file isn't included in the index.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// `TRUE` to use the Amazon WorkDocs change log to determine which documents require updating in the index.
	//
	// Depending on the change log's size, it may take longer for Amazon Kendra to use the change log than to scan all of your documents in Amazon WorkDocs.
	UseChangeLog interface{} `field:"optional" json:"useChangeLog" yaml:"useChangeLog"`
}

// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &cfnDataSourceProps{
//   	indexId: jsii.String("indexId"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	customDocumentEnrichmentConfiguration: &customDocumentEnrichmentConfigurationProperty{
//   		inlineConfigurations: []interface{}{
//   			&inlineCustomDocumentEnrichmentConfigurationProperty{
//   				condition: &documentAttributeConditionProperty{
//   					conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   					operator: jsii.String("operator"),
//
//   					// the properties below are optional
//   					conditionOnValue: &documentAttributeValueProperty{
//   						dateValue: jsii.String("dateValue"),
//   						longValue: jsii.Number(123),
//   						stringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						stringValue: jsii.String("stringValue"),
//   					},
//   				},
//   				documentContentDeletion: jsii.Boolean(false),
//   				target: &documentAttributeTargetProperty{
//   					targetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   					// the properties below are optional
//   					targetDocumentAttributeValue: &documentAttributeValueProperty{
//   						dateValue: jsii.String("dateValue"),
//   						longValue: jsii.Number(123),
//   						stringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						stringValue: jsii.String("stringValue"),
//   					},
//   					targetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		postExtractionHookConfiguration: &hookConfigurationProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   			s3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			invocationCondition: &documentAttributeConditionProperty{
//   				conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				conditionOnValue: &documentAttributeValueProperty{
//   					dateValue: jsii.String("dateValue"),
//   					longValue: jsii.Number(123),
//   					stringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   		preExtractionHookConfiguration: &hookConfigurationProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   			s3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			invocationCondition: &documentAttributeConditionProperty{
//   				conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				conditionOnValue: &documentAttributeValueProperty{
//   					dateValue: jsii.String("dateValue"),
//   					longValue: jsii.Number(123),
//   					stringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	dataSourceConfiguration: &dataSourceConfigurationProperty{
//   		confluenceConfiguration: &confluenceConfigurationProperty{
//   			secretArn: jsii.String("secretArn"),
//   			serverUrl: jsii.String("serverUrl"),
//   			version: jsii.String("version"),
//
//   			// the properties below are optional
//   			attachmentConfiguration: &confluenceAttachmentConfigurationProperty{
//   				attachmentFieldMappings: []interface{}{
//   					&confluenceAttachmentToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				crawlAttachments: jsii.Boolean(false),
//   			},
//   			blogConfiguration: &confluenceBlogConfigurationProperty{
//   				blogFieldMappings: []interface{}{
//   					&confluenceBlogToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			pageConfiguration: &confluencePageConfigurationProperty{
//   				pageFieldMappings: []interface{}{
//   					&confluencePageToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			spaceConfiguration: &confluenceSpaceConfigurationProperty{
//   				crawlArchivedSpaces: jsii.Boolean(false),
//   				crawlPersonalSpaces: jsii.Boolean(false),
//   				excludeSpaces: []*string{
//   					jsii.String("excludeSpaces"),
//   				},
//   				includeSpaces: []*string{
//   					jsii.String("includeSpaces"),
//   				},
//   				spaceFieldMappings: []interface{}{
//   					&confluenceSpaceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		databaseConfiguration: &databaseConfigurationProperty{
//   			columnConfiguration: &columnConfigurationProperty{
//   				changeDetectingColumns: []*string{
//   					jsii.String("changeDetectingColumns"),
//   				},
//   				documentDataColumnName: jsii.String("documentDataColumnName"),
//   				documentIdColumnName: jsii.String("documentIdColumnName"),
//
//   				// the properties below are optional
//   				documentTitleColumnName: jsii.String("documentTitleColumnName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			connectionConfiguration: &connectionConfigurationProperty{
//   				databaseHost: jsii.String("databaseHost"),
//   				databaseName: jsii.String("databaseName"),
//   				databasePort: jsii.Number(123),
//   				secretArn: jsii.String("secretArn"),
//   				tableName: jsii.String("tableName"),
//   			},
//   			databaseEngineType: jsii.String("databaseEngineType"),
//
//   			// the properties below are optional
//   			aclConfiguration: &aclConfigurationProperty{
//   				allowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   			},
//   			sqlConfiguration: &sqlConfigurationProperty{
//   				queryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   			},
//   			vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		googleDriveConfiguration: &googleDriveConfigurationProperty{
//   			secretArn: jsii.String("secretArn"),
//
//   			// the properties below are optional
//   			excludeMimeTypes: []*string{
//   				jsii.String("excludeMimeTypes"),
//   			},
//   			excludeSharedDrives: []*string{
//   				jsii.String("excludeSharedDrives"),
//   			},
//   			excludeUserAccounts: []*string{
//   				jsii.String("excludeUserAccounts"),
//   			},
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   		},
//   		oneDriveConfiguration: &oneDriveConfigurationProperty{
//   			oneDriveUsers: &oneDriveUsersProperty{
//   				oneDriveUserList: []*string{
//   					jsii.String("oneDriveUserList"),
//   				},
//   				oneDriveUserS3Path: &s3PathProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			secretArn: jsii.String("secretArn"),
//   			tenantDomain: jsii.String("tenantDomain"),
//
//   			// the properties below are optional
//   			disableLocalGroups: jsii.Boolean(false),
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   		},
//   		s3Configuration: &s3DataSourceConfigurationProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			accessControlListConfiguration: &accessControlListConfigurationProperty{
//   				keyPath: jsii.String("keyPath"),
//   			},
//   			documentsMetadataConfiguration: &documentsMetadataConfigurationProperty{
//   				s3Prefix: jsii.String("s3Prefix"),
//   			},
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			inclusionPrefixes: []*string{
//   				jsii.String("inclusionPrefixes"),
//   			},
//   		},
//   		salesforceConfiguration: &salesforceConfigurationProperty{
//   			secretArn: jsii.String("secretArn"),
//   			serverUrl: jsii.String("serverUrl"),
//
//   			// the properties below are optional
//   			chatterFeedConfiguration: &salesforceChatterFeedConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				includeFilterTypes: []*string{
//   					jsii.String("includeFilterTypes"),
//   				},
//   			},
//   			crawlAttachments: jsii.Boolean(false),
//   			excludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
//   			},
//   			includeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   			knowledgeArticleConfiguration: &salesforceKnowledgeArticleConfigurationProperty{
//   				includedStates: []*string{
//   					jsii.String("includedStates"),
//   				},
//
//   				// the properties below are optional
//   				customKnowledgeArticleTypeConfigurations: []interface{}{
//   					&salesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   						documentDataFieldName: jsii.String("documentDataFieldName"),
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   						fieldMappings: []interface{}{
//   							&dataSourceToIndexFieldMappingProperty{
//   								dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   								indexFieldName: jsii.String("indexFieldName"),
//
//   								// the properties below are optional
//   								dateFieldFormat: jsii.String("dateFieldFormat"),
//   							},
//   						},
//   					},
//   				},
//   				standardKnowledgeArticleTypeConfiguration: &salesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   					documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   					// the properties below are optional
//   					documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					fieldMappings: []interface{}{
//   						&dataSourceToIndexFieldMappingProperty{
//   							dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							indexFieldName: jsii.String("indexFieldName"),
//
//   							// the properties below are optional
//   							dateFieldFormat: jsii.String("dateFieldFormat"),
//   						},
//   					},
//   				},
//   			},
//   			standardObjectAttachmentConfiguration: &salesforceStandardObjectAttachmentConfigurationProperty{
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   			standardObjectConfigurations: []interface{}{
//   				&salesforceStandardObjectConfigurationProperty{
//   					documentDataFieldName: jsii.String("documentDataFieldName"),
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   					fieldMappings: []interface{}{
//   						&dataSourceToIndexFieldMappingProperty{
//   							dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   							indexFieldName: jsii.String("indexFieldName"),
//
//   							// the properties below are optional
//   							dateFieldFormat: jsii.String("dateFieldFormat"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		serviceNowConfiguration: &serviceNowConfigurationProperty{
//   			hostUrl: jsii.String("hostUrl"),
//   			secretArn: jsii.String("secretArn"),
//   			serviceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//
//   			// the properties below are optional
//   			authenticationType: jsii.String("authenticationType"),
//   			knowledgeArticleConfiguration: &serviceNowKnowledgeArticleConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				crawlAttachments: jsii.Boolean(false),
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				excludeAttachmentFilePatterns: []*string{
//   					jsii.String("excludeAttachmentFilePatterns"),
//   				},
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				filterQuery: jsii.String("filterQuery"),
//   				includeAttachmentFilePatterns: []*string{
//   					jsii.String("includeAttachmentFilePatterns"),
//   				},
//   			},
//   			serviceCatalogConfiguration: &serviceNowServiceCatalogConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   				// the properties below are optional
//   				crawlAttachments: jsii.Boolean(false),
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				excludeAttachmentFilePatterns: []*string{
//   					jsii.String("excludeAttachmentFilePatterns"),
//   				},
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   				includeAttachmentFilePatterns: []*string{
//   					jsii.String("includeAttachmentFilePatterns"),
//   				},
//   			},
//   		},
//   		sharePointConfiguration: &sharePointConfigurationProperty{
//   			secretArn: jsii.String("secretArn"),
//   			sharePointVersion: jsii.String("sharePointVersion"),
//   			urls: []*string{
//   				jsii.String("urls"),
//   			},
//
//   			// the properties below are optional
//   			crawlAttachments: jsii.Boolean(false),
//   			disableLocalGroups: jsii.Boolean(false),
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			sslCertificateS3Path: &s3PathProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//   			},
//   			useChangeLog: jsii.Boolean(false),
//   			vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		webCrawlerConfiguration: &webCrawlerConfigurationProperty{
//   			urls: &webCrawlerUrlsProperty{
//   				seedUrlConfiguration: &webCrawlerSeedUrlConfigurationProperty{
//   					seedUrls: []*string{
//   						jsii.String("seedUrls"),
//   					},
//
//   					// the properties below are optional
//   					webCrawlerMode: jsii.String("webCrawlerMode"),
//   				},
//   				siteMapsConfiguration: &webCrawlerSiteMapsConfigurationProperty{
//   					siteMaps: []*string{
//   						jsii.String("siteMaps"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			authenticationConfiguration: &webCrawlerAuthenticationConfigurationProperty{
//   				basicAuthentication: []interface{}{
//   					&webCrawlerBasicAuthenticationProperty{
//   						credentials: jsii.String("credentials"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			crawlDepth: jsii.Number(123),
//   			maxContentSizePerPageInMegaBytes: jsii.Number(123),
//   			maxLinksPerPage: jsii.Number(123),
//   			maxUrlsPerMinuteCrawlRate: jsii.Number(123),
//   			proxyConfiguration: &proxyConfigurationProperty{
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//
//   				// the properties below are optional
//   				credentials: jsii.String("credentials"),
//   			},
//   			urlExclusionPatterns: []*string{
//   				jsii.String("urlExclusionPatterns"),
//   			},
//   			urlInclusionPatterns: []*string{
//   				jsii.String("urlInclusionPatterns"),
//   			},
//   		},
//   		workDocsConfiguration: &workDocsConfigurationProperty{
//   			organizationId: jsii.String("organizationId"),
//
//   			// the properties below are optional
//   			crawlComments: jsii.Boolean(false),
//   			exclusionPatterns: []*string{
//   				jsii.String("exclusionPatterns"),
//   			},
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			inclusionPatterns: []*string{
//   				jsii.String("inclusionPatterns"),
//   			},
//   			useChangeLog: jsii.Boolean(false),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	roleArn: jsii.String("roleArn"),
//   	schedule: jsii.String("schedule"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataSourceProps struct {
	// The identifier of the index that should be associated with this data source.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// The name of the data source.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the data source.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Configuration information for altering document metadata and content during the document ingestion process.
	CustomDocumentEnrichmentConfiguration interface{} `field:"optional" json:"customDocumentEnrichmentConfiguration" yaml:"customDocumentEnrichmentConfiguration"`
	// Configuration information for an Amazon Kendra data source.
	//
	// The contents of the configuration depend on the type of data source. You can only specify one type of data source in the configuration. Choose from one of the following data sources.
	//
	// - Amazon S3
	// - Confluence
	// - Custom
	// - Database
	// - Microsoft OneDrive
	// - Microsoft SharePoint
	// - Salesforce
	// - ServiceNow
	//
	// You can't specify the `Configuration` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `Configuration` parameter is required for all other data sources.
	DataSourceConfiguration interface{} `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// A description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of a role with permission to access the data source.
	//
	// You can't specify the `RoleArn` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `RoleArn` parameter is required for all other data sources.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Sets the frequency that Amazon Kendra checks the documents in your data source and updates the index.
	//
	// If you don't set a schedule, Amazon Kendra doesn't periodically update the index.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Kendra::Faq`.
//
// Specifies an new set of frequently asked question (FAQ) questions and answers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFaq := awscdk.Aws_kendra.NewCfnFaq(this, jsii.String("MyCfnFaq"), &cfnFaqProps{
//   	indexId: jsii.String("indexId"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	s3Path: &s3PathProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	fileFormat: jsii.String("fileFormat"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFaq interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `arn:aws:kendra:us-west-2:111122223333:index/335c3741-41df-46a6-b5d3-61f85b787884/faq/f61995a6-cd5c-4e99-9cfc-58816d8bfaa7`.
	AttrArn() *string
	// The identifier for the FAQ. For example:.
	//
	// `f61995a6-cd5c-4e99-9cfc-58816d8bfaa7`.
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
	// A description of the FAQ.
	Description() *string
	SetDescription(val *string)
	// The format of the input file.
	//
	// You can choose between a basic CSV format, a CSV format that includes customs attributes in a header, and a JSON format that includes custom attributes.
	//
	// The format must match the format of the file stored in the S3 bucket identified in the S3Path parameter.
	//
	// Valid values are:
	//
	// - `CSV`
	// - `CSV_WITH_HEADER`
	// - `JSON`.
	FileFormat() *string
	SetFileFormat(val *string)
	// The identifier of the index that contains the FAQ.
	IndexId() *string
	SetIndexId(val *string)
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
	// The name that you assigned the FAQ when you created or updated the FAQ.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that contains the FAQ.
	RoleArn() *string
	SetRoleArn(val *string)
	// The Amazon Simple Storage Service (Amazon S3) location of the FAQ input data.
	S3Path() interface{}
	SetS3Path(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnFaq
type jsiiProxy_CfnFaq struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFaq) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) FileFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) IndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) S3Path() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Kendra::Faq`.
func NewCfnFaq(scope constructs.Construct, id *string, props *CfnFaqProps) CfnFaq {
	_init_.Initialize()

	j := jsiiProxy_CfnFaq{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::Faq`.
func NewCfnFaq_Override(c CfnFaq, scope constructs.Construct, id *string, props *CfnFaqProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFaq) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetFileFormat(val *string) {
	_jsii_.Set(
		j,
		"fileFormat",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetIndexId(val *string) {
	_jsii_.Set(
		j,
		"indexId",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetS3Path(val interface{}) {
	_jsii_.Set(
		j,
		"s3Path",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFaq_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFaq_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnFaq_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFaq_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFaq) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFaq) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFaq) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFaq) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFaq) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFaq) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFaq) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFaq) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFaq) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFaq) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFaq) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFaq) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFaq) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFaq) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFaq) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information required to find a specific file in an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3PathProperty := &s3PathProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   }
//
type CfnFaq_S3PathProperty struct {
	// The name of the S3 bucket that contains the file.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the file.
	Key *string `field:"required" json:"key" yaml:"key"`
}

// Properties for defining a `CfnFaq`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFaqProps := &cfnFaqProps{
//   	indexId: jsii.String("indexId"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	s3Path: &s3PathProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	fileFormat: jsii.String("fileFormat"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFaqProps struct {
	// The identifier of the index that contains the FAQ.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// The name that you assigned the FAQ when you created or updated the FAQ.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that contains the FAQ.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Simple Storage Service (Amazon S3) location of the FAQ input data.
	S3Path interface{} `field:"required" json:"s3Path" yaml:"s3Path"`
	// A description of the FAQ.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The format of the input file.
	//
	// You can choose between a basic CSV format, a CSV format that includes customs attributes in a header, and a JSON format that includes custom attributes.
	//
	// The format must match the format of the file stored in the S3 bucket identified in the S3Path parameter.
	//
	// Valid values are:
	//
	// - `CSV`
	// - `CSV_WITH_HEADER`
	// - `JSON`.
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Kendra::Index`.
//
// Specifies a new Amazon Kendra index. And index is a collection of documents and associated metadata that you want to search for relevant documents.
//
// Once the index is active you can add documents to your index using the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation or using one of the supported data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndex := awscdk.Aws_kendra.NewCfnIndex(this, jsii.String("MyCfnIndex"), &cfnIndexProps{
//   	edition: jsii.String("edition"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	capacityUnits: &capacityUnitsConfigurationProperty{
//   		queryCapacityUnits: jsii.Number(123),
//   		storageCapacityUnits: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	documentMetadataConfigurations: []interface{}{
//   		&documentMetadataConfigurationProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			relevance: &relevanceProperty{
//   				duration: jsii.String("duration"),
//   				freshness: jsii.Boolean(false),
//   				importance: jsii.Number(123),
//   				rankOrder: jsii.String("rankOrder"),
//   				valueImportanceItems: []interface{}{
//   					&valueImportanceItemProperty{
//   						key: jsii.String("key"),
//   						value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			search: &searchProperty{
//   				displayable: jsii.Boolean(false),
//   				facetable: jsii.Boolean(false),
//   				searchable: jsii.Boolean(false),
//   				sortable: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	serverSideEncryptionConfiguration: &serverSideEncryptionConfigurationProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userContextPolicy: jsii.String("userContextPolicy"),
//   	userTokenConfigurations: []interface{}{
//   		&userTokenConfigurationProperty{
//   			jsonTokenTypeConfiguration: &jsonTokenTypeConfigurationProperty{
//   				groupAttributeField: jsii.String("groupAttributeField"),
//   				userNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   			jwtTokenTypeConfiguration: &jwtTokenTypeConfigurationProperty{
//   				keyLocation: jsii.String("keyLocation"),
//
//   				// the properties below are optional
//   				claimRegex: jsii.String("claimRegex"),
//   				groupAttributeField: jsii.String("groupAttributeField"),
//   				issuer: jsii.String("issuer"),
//   				secretManagerArn: jsii.String("secretManagerArn"),
//   				url: jsii.String("url"),
//   				userNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   		},
//   	},
//   })
//
type CfnIndex interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the index.
	//
	// For example: `arn:aws:kendra:us-west-2:111122223333:index/0123456789abcdef` .
	AttrArn() *string
	// The identifier for the index.
	//
	// For example: `f4aeaa10-8056-4b2c-a343-522ca0f41234` .
	AttrId() *string
	// Specifies capacity units configured for your index.
	//
	// You can add and remove capacity units to tune an index to your requirements. You can set capacity units only for Enterprise edition indexes.
	CapacityUnits() interface{}
	SetCapacityUnits(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the index.
	Description() *string
	SetDescription(val *string)
	// Specifies the properties of an index field.
	//
	// You can add either a custom or a built-in field. You can add and remove built-in fields at any time. When a built-in field is removed it's configuration reverts to the default for the field. Custom fields can't be removed from an index after they are added.
	DocumentMetadataConfigurations() interface{}
	SetDocumentMetadataConfigurations(val interface{})
	// Indicates whether the index is a enterprise edition index or a developer edition index.
	//
	// Valid values are `DEVELOPER_EDITION` and `ENTERPRISE_EDITION` .
	Edition() *string
	SetEdition(val *string)
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
	// The identifier of the index.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics.
	//
	// This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	RoleArn() *string
	SetRoleArn(val *string)
	// The identifier of the AWS KMS customer managed key (CMK) to use to encrypt data indexed by Amazon Kendra.
	//
	// Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration() interface{}
	SetServerSideEncryptionConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If there is an access control list, it is ignored. You can filter on user and group attributes.
	//
	// USER_TOKEN
	//
	// - Enables SSO and token-based user access control. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	UserContextPolicy() *string
	SetUserContextPolicy(val *string)
	// Defines the type of user token used for the index.
	UserTokenConfigurations() interface{}
	SetUserTokenConfigurations(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnIndex
type jsiiProxy_CfnIndex struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIndex) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CapacityUnits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) DocumentMetadataConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentMetadataConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) ServerSideEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UserContextPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userContextPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UserTokenConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userTokenConfigurations",
		&returns,
	)
	return returns
}


// Create a new `AWS::Kendra::Index`.
func NewCfnIndex(scope constructs.Construct, id *string, props *CfnIndexProps) CfnIndex {
	_init_.Initialize()

	j := jsiiProxy_CfnIndex{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::Index`.
func NewCfnIndex_Override(c CfnIndex, scope constructs.Construct, id *string, props *CfnIndexProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIndex) SetCapacityUnits(val interface{}) {
	_jsii_.Set(
		j,
		"capacityUnits",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetDocumentMetadataConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"documentMetadataConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetEdition(val *string) {
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetServerSideEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"serverSideEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetUserContextPolicy(val *string) {
	_jsii_.Set(
		j,
		"userContextPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetUserTokenConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"userTokenConfigurations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnIndex_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnIndex_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIndex_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIndex) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIndex) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIndex) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIndex) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIndex) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIndex) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIndex) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIndex) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIndex) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIndex) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIndex) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIndex) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIndex) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIndex) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies additional capacity units configured for your Enterprise Edition index.
//
// You can add and remove capacity units to fit your usage requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityUnitsConfigurationProperty := &capacityUnitsConfigurationProperty{
//   	queryCapacityUnits: jsii.Number(123),
//   	storageCapacityUnits: jsii.Number(123),
//   }
//
type CfnIndex_CapacityUnitsConfigurationProperty struct {
	// The amount of extra query capacity for an index and [GetQuerySuggestions](https://docs.aws.amazon.com/kendra/latest/dg/API_GetQuerySuggestions.html) capacity.
	//
	// A single extra capacity unit for an index provides 0.1 queries per second or approximately 8,000 queries per day.
	//
	// `GetQuerySuggestions` capacity is five times the provisioned query capacity for an index, or the base capacity of 2.5 calls per second, whichever is higher. For example, the base capacity for an index is 0.1 queries per second, and `GetQuerySuggestions` capacity has a base of 2.5 calls per second. If you add another 0.1 queries per second to total 0.2 queries per second for an index, the `GetQuerySuggestions` capacity is 2.5 calls per second (higher than five times 0.2 queries per second).
	QueryCapacityUnits *float64 `field:"required" json:"queryCapacityUnits" yaml:"queryCapacityUnits"`
	// The amount of extra storage capacity for an index.
	//
	// A single capacity unit provides 30 GB of storage space or 100,000 documents, whichever is reached first.
	StorageCapacityUnits *float64 `field:"required" json:"storageCapacityUnits" yaml:"storageCapacityUnits"`
}

// Specifies the properties of a custom index field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentMetadataConfigurationProperty := &documentMetadataConfigurationProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	relevance: &relevanceProperty{
//   		duration: jsii.String("duration"),
//   		freshness: jsii.Boolean(false),
//   		importance: jsii.Number(123),
//   		rankOrder: jsii.String("rankOrder"),
//   		valueImportanceItems: []interface{}{
//   			&valueImportanceItemProperty{
//   				key: jsii.String("key"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	search: &searchProperty{
//   		displayable: jsii.Boolean(false),
//   		facetable: jsii.Boolean(false),
//   		searchable: jsii.Boolean(false),
//   		sortable: jsii.Boolean(false),
//   	},
//   }
//
type CfnIndex_DocumentMetadataConfigurationProperty struct {
	// The name of the index field.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the index field.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Provides manual tuning parameters to determine how the field affects the search results.
	Relevance interface{} `field:"optional" json:"relevance" yaml:"relevance"`
	// Provides information about how the field is used during a search.
	Search interface{} `field:"optional" json:"search" yaml:"search"`
}

// Provides the configuration information for the JSON token type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonTokenTypeConfigurationProperty := &jsonTokenTypeConfigurationProperty{
//   	groupAttributeField: jsii.String("groupAttributeField"),
//   	userNameAttributeField: jsii.String("userNameAttributeField"),
//   }
//
type CfnIndex_JsonTokenTypeConfigurationProperty struct {
	// The group attribute field.
	GroupAttributeField *string `field:"required" json:"groupAttributeField" yaml:"groupAttributeField"`
	// The user name attribute field.
	UserNameAttributeField *string `field:"required" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

// Provides the configuration information for the JWT token type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jwtTokenTypeConfigurationProperty := &jwtTokenTypeConfigurationProperty{
//   	keyLocation: jsii.String("keyLocation"),
//
//   	// the properties below are optional
//   	claimRegex: jsii.String("claimRegex"),
//   	groupAttributeField: jsii.String("groupAttributeField"),
//   	issuer: jsii.String("issuer"),
//   	secretManagerArn: jsii.String("secretManagerArn"),
//   	url: jsii.String("url"),
//   	userNameAttributeField: jsii.String("userNameAttributeField"),
//   }
//
type CfnIndex_JwtTokenTypeConfigurationProperty struct {
	// The location of the key.
	KeyLocation *string `field:"required" json:"keyLocation" yaml:"keyLocation"`
	// The regular expression that identifies the claim.
	ClaimRegex *string `field:"optional" json:"claimRegex" yaml:"claimRegex"`
	// The group attribute field.
	GroupAttributeField *string `field:"optional" json:"groupAttributeField" yaml:"groupAttributeField"`
	// The issuer of the token.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The Amazon Resource Name (arn) of the secret.
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The signing key URL.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name attribute field.
	UserNameAttributeField *string `field:"optional" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

// Provides information for manually tuning the relevance of a field in a search.
//
// When a query includes terms that match the field, the results are given a boost in the response based on these tuning parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relevanceProperty := &relevanceProperty{
//   	duration: jsii.String("duration"),
//   	freshness: jsii.Boolean(false),
//   	importance: jsii.Number(123),
//   	rankOrder: jsii.String("rankOrder"),
//   	valueImportanceItems: []interface{}{
//   		&valueImportanceItemProperty{
//   			key: jsii.String("key"),
//   			value: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnIndex_RelevanceProperty struct {
	// Specifies the time period that the boost applies to.
	//
	// For example, to make the boost apply to documents with the field value within the last month, you would use "2628000s". Once the field value is beyond the specified range, the effect of the boost drops off. The higher the importance, the faster the effect drops off. If you don't specify a value, the default is 3 months. The value of the field is a numeric string followed by the character "s", for example "86400s" for one day, or "604800s" for one week.
	//
	// Only applies to `DATE` fields.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Indicates that this field determines how "fresh" a document is.
	//
	// For example, if document 1 was created on November 5, and document 2 was created on October 31, document 1 is "fresher" than document 2. You can only set the `Freshness` field on one `DATE` type field. Only applies to `DATE` fields.
	Freshness interface{} `field:"optional" json:"freshness" yaml:"freshness"`
	// The relative importance of the field in the search.
	//
	// Larger numbers provide more of a boost than smaller numbers.
	Importance *float64 `field:"optional" json:"importance" yaml:"importance"`
	// Determines how values should be interpreted.
	//
	// When the `RankOrder` field is `ASCENDING` , higher numbers are better. For example, a document with a rating score of 10 is higher ranking than a document with a rating score of 1.
	//
	// When the `RankOrder` field is `DESCENDING` , lower numbers are better. For example, in a task tracking application, a priority 1 task is more important than a priority 5 task.
	//
	// Only applies to `LONG` and `DOUBLE` fields.
	RankOrder *string `field:"optional" json:"rankOrder" yaml:"rankOrder"`
	// An array of key-value pairs that contains an array of values that should be given a different boost when they appear in the search result list.
	//
	// For example, if you are boosting query terms that match the department field in the result, query terms that match the department field are boosted in the result. You can add entries from the department field to boost documents with those values higher.
	//
	// For example, you can add entries to the map with names of departments. If you add "HR", 5 and "Legal",3 those departments are given special attention when they appear in the metadata of a document.
	ValueImportanceItems interface{} `field:"optional" json:"valueImportanceItems" yaml:"valueImportanceItems"`
}

// Provides information about how a custom index field is used during a search.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   searchProperty := &searchProperty{
//   	displayable: jsii.Boolean(false),
//   	facetable: jsii.Boolean(false),
//   	searchable: jsii.Boolean(false),
//   	sortable: jsii.Boolean(false),
//   }
//
type CfnIndex_SearchProperty struct {
	// Determines whether the field is returned in the query response.
	//
	// The default is `true` .
	Displayable interface{} `field:"optional" json:"displayable" yaml:"displayable"`
	// Indicates that the field can be used to create search facets, a count of results for each value in the field.
	//
	// The default is `false` .
	Facetable interface{} `field:"optional" json:"facetable" yaml:"facetable"`
	// Determines whether the field is used in the search.
	//
	// If the `Searchable` field is `true` , you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is `true` for string fields and `false` for number and date fields.
	Searchable interface{} `field:"optional" json:"searchable" yaml:"searchable"`
	// Indicates that the field can be used to sort the search results.
	//
	// The default is `false` .
	Sortable interface{} `field:"optional" json:"sortable" yaml:"sortable"`
}

// Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt data indexed by Amazon Kendra.
//
// We suggest that you use a CMK from your account to help secure your index. Amazon Kendra doesn't support asymmetric CMKs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionConfigurationProperty := &serverSideEncryptionConfigurationProperty{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnIndex_ServerSideEncryptionConfigurationProperty struct {
	// The identifier of the AWS KMS key .
	//
	// Amazon Kendra doesn't support asymmetric keys.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

// Provides the configuration information for a token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userTokenConfigurationProperty := &userTokenConfigurationProperty{
//   	jsonTokenTypeConfiguration: &jsonTokenTypeConfigurationProperty{
//   		groupAttributeField: jsii.String("groupAttributeField"),
//   		userNameAttributeField: jsii.String("userNameAttributeField"),
//   	},
//   	jwtTokenTypeConfiguration: &jwtTokenTypeConfigurationProperty{
//   		keyLocation: jsii.String("keyLocation"),
//
//   		// the properties below are optional
//   		claimRegex: jsii.String("claimRegex"),
//   		groupAttributeField: jsii.String("groupAttributeField"),
//   		issuer: jsii.String("issuer"),
//   		secretManagerArn: jsii.String("secretManagerArn"),
//   		url: jsii.String("url"),
//   		userNameAttributeField: jsii.String("userNameAttributeField"),
//   	},
//   }
//
type CfnIndex_UserTokenConfigurationProperty struct {
	// Information about the JSON token type configuration.
	JsonTokenTypeConfiguration interface{} `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// Information about the JWT token type configuration.
	JwtTokenTypeConfiguration interface{} `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

// Specifies a key-value pair that determines the search boost value that a document receives when the key is part of the metadata of a document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueImportanceItemProperty := &valueImportanceItemProperty{
//   	key: jsii.String("key"),
//   	value: jsii.Number(123),
//   }
//
type CfnIndex_ValueImportanceItemProperty struct {
	// The document metadata value that receives the search boost.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The boost value that a document receives when the key is part of the metadata of a document.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

// Properties for defining a `CfnIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIndexProps := &cfnIndexProps{
//   	edition: jsii.String("edition"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	capacityUnits: &capacityUnitsConfigurationProperty{
//   		queryCapacityUnits: jsii.Number(123),
//   		storageCapacityUnits: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	documentMetadataConfigurations: []interface{}{
//   		&documentMetadataConfigurationProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			relevance: &relevanceProperty{
//   				duration: jsii.String("duration"),
//   				freshness: jsii.Boolean(false),
//   				importance: jsii.Number(123),
//   				rankOrder: jsii.String("rankOrder"),
//   				valueImportanceItems: []interface{}{
//   					&valueImportanceItemProperty{
//   						key: jsii.String("key"),
//   						value: jsii.Number(123),
//   					},
//   				},
//   			},
//   			search: &searchProperty{
//   				displayable: jsii.Boolean(false),
//   				facetable: jsii.Boolean(false),
//   				searchable: jsii.Boolean(false),
//   				sortable: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	serverSideEncryptionConfiguration: &serverSideEncryptionConfigurationProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userContextPolicy: jsii.String("userContextPolicy"),
//   	userTokenConfigurations: []interface{}{
//   		&userTokenConfigurationProperty{
//   			jsonTokenTypeConfiguration: &jsonTokenTypeConfigurationProperty{
//   				groupAttributeField: jsii.String("groupAttributeField"),
//   				userNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   			jwtTokenTypeConfiguration: &jwtTokenTypeConfigurationProperty{
//   				keyLocation: jsii.String("keyLocation"),
//
//   				// the properties below are optional
//   				claimRegex: jsii.String("claimRegex"),
//   				groupAttributeField: jsii.String("groupAttributeField"),
//   				issuer: jsii.String("issuer"),
//   				secretManagerArn: jsii.String("secretManagerArn"),
//   				url: jsii.String("url"),
//   				userNameAttributeField: jsii.String("userNameAttributeField"),
//   			},
//   		},
//   	},
//   }
//
type CfnIndexProps struct {
	// Indicates whether the index is a enterprise edition index or a developer edition index.
	//
	// Valid values are `DEVELOPER_EDITION` and `ENTERPRISE_EDITION` .
	Edition *string `field:"required" json:"edition" yaml:"edition"`
	// The identifier of the index.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics.
	//
	// This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Specifies capacity units configured for your index.
	//
	// You can add and remove capacity units to tune an index to your requirements. You can set capacity units only for Enterprise edition indexes.
	CapacityUnits interface{} `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// A description of the index.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the properties of an index field.
	//
	// You can add either a custom or a built-in field. You can add and remove built-in fields at any time. When a built-in field is removed it's configuration reverts to the default for the field. Custom fields can't be removed from an index after they are added.
	DocumentMetadataConfigurations interface{} `field:"optional" json:"documentMetadataConfigurations" yaml:"documentMetadataConfigurations"`
	// The identifier of the AWS KMS customer managed key (CMK) to use to encrypt data indexed by Amazon Kendra.
	//
	// Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If there is an access control list, it is ignored. You can filter on user and group attributes.
	//
	// USER_TOKEN
	//
	// - Enables SSO and token-based user access control. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	UserContextPolicy *string `field:"optional" json:"userContextPolicy" yaml:"userContextPolicy"`
	// Defines the type of user token used for the index.
	UserTokenConfigurations interface{} `field:"optional" json:"userTokenConfigurations" yaml:"userTokenConfigurations"`
}

