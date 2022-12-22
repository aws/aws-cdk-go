package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Kendra::DataSource`.
//
// Creates a data source connector that you want to use with an Amazon Kendra index.
//
// You specify a name, data source connector type and description for your data source. You also specify configuration information for the data source connector.
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
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Configuration information for altering document metadata and content during the document ingestion process.
	CustomDocumentEnrichmentConfiguration() interface{}
	SetCustomDocumentEnrichmentConfiguration(val interface{})
	// Configuration information for an Amazon Kendra data source.
	//
	// The contents of the configuration depend on the type of data source. You can only specify one type of data source in the configuration.
	//
	// You can't specify the `Configuration` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `Configuration` parameter is required for all other data sources.
	DataSourceConfiguration() interface{}
	SetDataSourceConfiguration(val interface{})
	// A description for the data source connector.
	Description() *string
	SetDescription(val *string)
	// The identifier of the index you want to use with the data source connector.
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
	// Experimental.
	LogicalId() *string
	// The name of the data source.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
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
	// Experimental.
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
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnDataSource(scope awscdk.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	if err := validateNewCfnDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"monocdk.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::DataSource`.
func NewCfnDataSource_Override(c CfnDataSource, scope awscdk.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kendra.CfnDataSource",
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
// Experimental.
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kendra.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kendra.CfnDataSource",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kendra.CfnDataSource",
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
		"monocdk.aws_kendra.CfnDataSource",
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

func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
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

func (c *jsiiProxy_CfnDataSource) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnDataSource) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

