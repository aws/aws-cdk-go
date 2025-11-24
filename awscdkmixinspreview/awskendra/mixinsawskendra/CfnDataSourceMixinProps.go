package mixinsawskendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataSourceMixinProps := &CfnDataSourceMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html
//
type CfnDataSourceMixinProps struct {
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
	// The identifier of the index you want to use with the data source connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-indexid
	//
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
	// The code for a language.
	//
	// This shows a supported language for all documents in the data source. English is supported by default. For more information on supported languages, including their codes, see [Adding documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
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
	// The type of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-datasource.html#cfn-kendra-datasource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

