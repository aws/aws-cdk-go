package previewawskendramixins


// Provides the configuration information for an Amazon Kendra data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	ConfluenceConfiguration: &ConfluenceConfigurationProperty{
//   		AttachmentConfiguration: &ConfluenceAttachmentConfigurationProperty{
//   			AttachmentFieldMappings: []interface{}{
//   				&ConfluenceAttachmentToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			CrawlAttachments: jsii.Boolean(false),
//   		},
//   		BlogConfiguration: &ConfluenceBlogConfigurationProperty{
//   			BlogFieldMappings: []interface{}{
//   				&ConfluenceBlogToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   		},
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		PageConfiguration: &ConfluencePageConfigurationProperty{
//   			PageFieldMappings: []interface{}{
//   				&ConfluencePageToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   		ServerUrl: jsii.String("serverUrl"),
//   		SpaceConfiguration: &ConfluenceSpaceConfigurationProperty{
//   			CrawlArchivedSpaces: jsii.Boolean(false),
//   			CrawlPersonalSpaces: jsii.Boolean(false),
//   			ExcludeSpaces: []*string{
//   				jsii.String("excludeSpaces"),
//   			},
//   			IncludeSpaces: []*string{
//   				jsii.String("includeSpaces"),
//   			},
//   			SpaceFieldMappings: []interface{}{
//   				&ConfluenceSpaceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   		},
//   		Version: jsii.String("version"),
//   		VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	DatabaseConfiguration: &DatabaseConfigurationProperty{
//   		AclConfiguration: &AclConfigurationProperty{
//   			AllowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   		},
//   		ColumnConfiguration: &ColumnConfigurationProperty{
//   			ChangeDetectingColumns: []*string{
//   				jsii.String("changeDetectingColumns"),
//   			},
//   			DocumentDataColumnName: jsii.String("documentDataColumnName"),
//   			DocumentIdColumnName: jsii.String("documentIdColumnName"),
//   			DocumentTitleColumnName: jsii.String("documentTitleColumnName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   		},
//   		ConnectionConfiguration: &ConnectionConfigurationProperty{
//   			DatabaseHost: jsii.String("databaseHost"),
//   			DatabaseName: jsii.String("databaseName"),
//   			DatabasePort: jsii.Number(123),
//   			SecretArn: jsii.String("secretArn"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		DatabaseEngineType: jsii.String("databaseEngineType"),
//   		SqlConfiguration: &SqlConfigurationProperty{
//   			QueryIdentifiersEnclosingOption: jsii.String("queryIdentifiersEnclosingOption"),
//   		},
//   		VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	GoogleDriveConfiguration: &GoogleDriveConfigurationProperty{
//   		ExcludeMimeTypes: []*string{
//   			jsii.String("excludeMimeTypes"),
//   		},
//   		ExcludeSharedDrives: []*string{
//   			jsii.String("excludeSharedDrives"),
//   		},
//   		ExcludeUserAccounts: []*string{
//   			jsii.String("excludeUserAccounts"),
//   		},
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	OneDriveConfiguration: &OneDriveConfigurationProperty{
//   		DisableLocalGroups: jsii.Boolean(false),
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		OneDriveUsers: &OneDriveUsersProperty{
//   			OneDriveUserList: []*string{
//   				jsii.String("oneDriveUserList"),
//   			},
//   			OneDriveUserS3Path: &S3PathProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   			},
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   		TenantDomain: jsii.String("tenantDomain"),
//   	},
//   	S3Configuration: &S3DataSourceConfigurationProperty{
//   		AccessControlListConfiguration: &AccessControlListConfigurationProperty{
//   			KeyPath: jsii.String("keyPath"),
//   		},
//   		BucketName: jsii.String("bucketName"),
//   		DocumentsMetadataConfiguration: &DocumentsMetadataConfigurationProperty{
//   			S3Prefix: jsii.String("s3Prefix"),
//   		},
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		InclusionPrefixes: []*string{
//   			jsii.String("inclusionPrefixes"),
//   		},
//   	},
//   	SalesforceConfiguration: &SalesforceConfigurationProperty{
//   		ChatterFeedConfiguration: &SalesforceChatterFeedConfigurationProperty{
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			IncludeFilterTypes: []*string{
//   				jsii.String("includeFilterTypes"),
//   			},
//   		},
//   		CrawlAttachments: jsii.Boolean(false),
//   		ExcludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		IncludeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   		KnowledgeArticleConfiguration: &SalesforceKnowledgeArticleConfigurationProperty{
//   			CustomKnowledgeArticleTypeConfigurations: []interface{}{
//   				&SalesforceCustomKnowledgeArticleTypeConfigurationProperty{
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
//   			IncludedStates: []*string{
//   				jsii.String("includedStates"),
//   			},
//   			StandardKnowledgeArticleTypeConfiguration: &SalesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   			},
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   		ServerUrl: jsii.String("serverUrl"),
//   		StandardObjectAttachmentConfiguration: &SalesforceStandardObjectAttachmentConfigurationProperty{
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   		},
//   		StandardObjectConfigurations: []interface{}{
//   			&SalesforceStandardObjectConfigurationProperty{
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	ServiceNowConfiguration: &ServiceNowConfigurationProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//   		HostUrl: jsii.String("hostUrl"),
//   		KnowledgeArticleConfiguration: &ServiceNowKnowledgeArticleConfigurationProperty{
//   			CrawlAttachments: jsii.Boolean(false),
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			ExcludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			FilterQuery: jsii.String("filterQuery"),
//   			IncludeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   		ServiceCatalogConfiguration: &ServiceNowServiceCatalogConfigurationProperty{
//   			CrawlAttachments: jsii.Boolean(false),
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			ExcludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
//   			},
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			IncludeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   		},
//   		ServiceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//   	},
//   	SharePointConfiguration: &SharePointConfigurationProperty{
//   		CrawlAttachments: jsii.Boolean(false),
//   		DisableLocalGroups: jsii.Boolean(false),
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		SecretArn: jsii.String("secretArn"),
//   		SharePointVersion: jsii.String("sharePointVersion"),
//   		SslCertificateS3Path: &S3PathProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   		Urls: []*string{
//   			jsii.String("urls"),
//   		},
//   		UseChangeLog: jsii.Boolean(false),
//   		VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	TemplateConfiguration: &TemplateConfigurationProperty{
//   		Template: jsii.String("template"),
//   	},
//   	WebCrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   		AuthenticationConfiguration: &WebCrawlerAuthenticationConfigurationProperty{
//   			BasicAuthentication: []interface{}{
//   				&WebCrawlerBasicAuthenticationProperty{
//   					Credentials: jsii.String("credentials"),
//   					Host: jsii.String("host"),
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		CrawlDepth: jsii.Number(123),
//   		MaxContentSizePerPageInMegaBytes: jsii.Number(123),
//   		MaxLinksPerPage: jsii.Number(123),
//   		MaxUrlsPerMinuteCrawlRate: jsii.Number(123),
//   		ProxyConfiguration: &ProxyConfigurationProperty{
//   			Credentials: jsii.String("credentials"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   		UrlExclusionPatterns: []*string{
//   			jsii.String("urlExclusionPatterns"),
//   		},
//   		UrlInclusionPatterns: []*string{
//   			jsii.String("urlInclusionPatterns"),
//   		},
//   		Urls: &WebCrawlerUrlsProperty{
//   			SeedUrlConfiguration: &WebCrawlerSeedUrlConfigurationProperty{
//   				SeedUrls: []*string{
//   					jsii.String("seedUrls"),
//   				},
//   				WebCrawlerMode: jsii.String("webCrawlerMode"),
//   			},
//   			SiteMapsConfiguration: &WebCrawlerSiteMapsConfigurationProperty{
//   				SiteMaps: []*string{
//   					jsii.String("siteMaps"),
//   				},
//   			},
//   		},
//   	},
//   	WorkDocsConfiguration: &WorkDocsConfigurationProperty{
//   		CrawlComments: jsii.Boolean(false),
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		OrganizationId: jsii.String("organizationId"),
//   		UseChangeLog: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html
//
type CfnDataSourcePropsMixin_DataSourceConfigurationProperty struct {
	// Provides the configuration information to connect to Confluence as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-confluenceconfiguration
	//
	ConfluenceConfiguration interface{} `field:"optional" json:"confluenceConfiguration" yaml:"confluenceConfiguration"`
	// Provides the configuration information to connect to a database as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-databaseconfiguration
	//
	DatabaseConfiguration interface{} `field:"optional" json:"databaseConfiguration" yaml:"databaseConfiguration"`
	// Provides the configuration information to connect to Google Drive as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-googledriveconfiguration
	//
	GoogleDriveConfiguration interface{} `field:"optional" json:"googleDriveConfiguration" yaml:"googleDriveConfiguration"`
	// Provides the configuration information to connect to Microsoft OneDrive as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-onedriveconfiguration
	//
	OneDriveConfiguration interface{} `field:"optional" json:"oneDriveConfiguration" yaml:"oneDriveConfiguration"`
	// Provides the configuration information to connect to an Amazon S3 bucket as your data source.
	//
	// > Amazon Kendra now supports an upgraded Amazon S3 connector.
	// >
	// > You must now use the [TemplateConfiguration](https://docs.aws.amazon.com/kendra/latest/APIReference/API_TemplateConfiguration.html) object instead of the `S3DataSourceConfiguration` object to configure your connector.
	// >
	// > Connectors configured using the older console and API architecture will continue to function as configured. However, you won't be able to edit or update them. If you want to edit or update your connector configuration, you must create a new connector.
	// >
	// > We recommended migrating your connector workflow to the upgraded version. Support for connectors configured using the older architecture is scheduled to end by June 2024.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// Provides the configuration information to connect to Salesforce as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-salesforceconfiguration
	//
	SalesforceConfiguration interface{} `field:"optional" json:"salesforceConfiguration" yaml:"salesforceConfiguration"`
	// Provides the configuration information to connect to ServiceNow as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-servicenowconfiguration
	//
	ServiceNowConfiguration interface{} `field:"optional" json:"serviceNowConfiguration" yaml:"serviceNowConfiguration"`
	// Provides the configuration information to connect to Microsoft SharePoint as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-sharepointconfiguration
	//
	SharePointConfiguration interface{} `field:"optional" json:"sharePointConfiguration" yaml:"sharePointConfiguration"`
	// Provides a template for the configuration information to connect to your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// Provides the configuration information required for Amazon Kendra Web Crawler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-webcrawlerconfiguration
	//
	WebCrawlerConfiguration interface{} `field:"optional" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
	// Provides the configuration information to connect to WorkDocs as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-workdocsconfiguration
	//
	WorkDocsConfiguration interface{} `field:"optional" json:"workDocsConfiguration" yaml:"workDocsConfiguration"`
}

