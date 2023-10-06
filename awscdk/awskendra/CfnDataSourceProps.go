package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &CfnDataSourceProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html
//
type CfnDataSourceProps struct {
	// The identifier of the index you want to use with the data source connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-indexid
	//
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Configuration information for altering document metadata and content during the document ingestion process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration
	//
	CustomDocumentEnrichmentConfiguration interface{} `field:"optional" json:"customDocumentEnrichmentConfiguration" yaml:"customDocumentEnrichmentConfiguration"`
	// Configuration information for an Amazon Kendra data source.
	//
	// The contents of the configuration depend on the type of data source. You can only specify one type of data source in the configuration.
	//
	// You can't specify the `Configuration` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `Configuration` parameter is required for all other data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-datasourceconfiguration
	//
	DataSourceConfiguration interface{} `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// A description for the data source connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The code for a language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The Amazon Resource Name (ARN) of a role with permission to access the data source.
	//
	// You can't specify the `RoleArn` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `RoleArn` parameter is required for all other data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Sets the frequency that Amazon Kendra checks the documents in your data source and updates the index.
	//
	// If you don't set a schedule, Amazon Kendra doesn't periodically update the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-schedule
	//
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

