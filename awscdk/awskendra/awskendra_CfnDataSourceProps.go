package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

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

