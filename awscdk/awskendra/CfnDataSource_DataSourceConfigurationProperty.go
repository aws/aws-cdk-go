package awskendra


// Provides the configuration information for an Amazon Kendra data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	ConfluenceConfiguration: &ConfluenceConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   		ServerUrl: jsii.String("serverUrl"),
//   		Version: jsii.String("version"),
//
//   		// the properties below are optional
//   		AttachmentConfiguration: &ConfluenceAttachmentConfigurationProperty{
//   			AttachmentFieldMappings: []interface{}{
//   				&ConfluenceAttachmentToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   			CrawlAttachments: jsii.Boolean(false),
//   		},
//   		BlogConfiguration: &ConfluenceBlogConfigurationProperty{
//   			BlogFieldMappings: []interface{}{
//   				&ConfluenceBlogToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
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
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
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
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
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
//   	DatabaseConfiguration: &DatabaseConfigurationProperty{
//   		ColumnConfiguration: &ColumnConfigurationProperty{
//   			ChangeDetectingColumns: []*string{
//   				jsii.String("changeDetectingColumns"),
//   			},
//   			DocumentDataColumnName: jsii.String("documentDataColumnName"),
//   			DocumentIdColumnName: jsii.String("documentIdColumnName"),
//
//   			// the properties below are optional
//   			DocumentTitleColumnName: jsii.String("documentTitleColumnName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
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
//
//   		// the properties below are optional
//   		AclConfiguration: &AclConfigurationProperty{
//   			AllowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   		},
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
//   		SecretArn: jsii.String("secretArn"),
//
//   		// the properties below are optional
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
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   	},
//   	OneDriveConfiguration: &OneDriveConfigurationProperty{
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
//
//   		// the properties below are optional
//   		DisableLocalGroups: jsii.Boolean(false),
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   	},
//   	S3Configuration: &S3DataSourceConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		AccessControlListConfiguration: &AccessControlListConfigurationProperty{
//   			KeyPath: jsii.String("keyPath"),
//   		},
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
//   		SecretArn: jsii.String("secretArn"),
//   		ServerUrl: jsii.String("serverUrl"),
//
//   		// the properties below are optional
//   		ChatterFeedConfiguration: &SalesforceChatterFeedConfigurationProperty{
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   			// the properties below are optional
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
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
//   			IncludedStates: []*string{
//   				jsii.String("includedStates"),
//   			},
//
//   			// the properties below are optional
//   			CustomKnowledgeArticleTypeConfigurations: []interface{}{
//   				&SalesforceCustomKnowledgeArticleTypeConfigurationProperty{
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
//   			StandardKnowledgeArticleTypeConfiguration: &SalesforceStandardKnowledgeArticleTypeConfigurationProperty{
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
//   			},
//   		},
//   		StandardObjectAttachmentConfiguration: &SalesforceStandardObjectAttachmentConfigurationProperty{
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   		StandardObjectConfigurations: []interface{}{
//   			&SalesforceStandardObjectConfigurationProperty{
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   				Name: jsii.String("name"),
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
//   			},
//   		},
//   	},
//   	ServiceNowConfiguration: &ServiceNowConfigurationProperty{
//   		HostUrl: jsii.String("hostUrl"),
//   		SecretArn: jsii.String("secretArn"),
//   		ServiceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//
//   		// the properties below are optional
//   		AuthenticationType: jsii.String("authenticationType"),
//   		KnowledgeArticleConfiguration: &ServiceNowKnowledgeArticleConfigurationProperty{
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   			// the properties below are optional
//   			CrawlAttachments: jsii.Boolean(false),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			ExcludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
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
//   			FilterQuery: jsii.String("filterQuery"),
//   			IncludeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   		},
//   		ServiceCatalogConfiguration: &ServiceNowServiceCatalogConfigurationProperty{
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   			// the properties below are optional
//   			CrawlAttachments: jsii.Boolean(false),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			ExcludeAttachmentFilePatterns: []*string{
//   				jsii.String("excludeAttachmentFilePatterns"),
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
//   			IncludeAttachmentFilePatterns: []*string{
//   				jsii.String("includeAttachmentFilePatterns"),
//   			},
//   		},
//   	},
//   	SharePointConfiguration: &SharePointConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   		SharePointVersion: jsii.String("sharePointVersion"),
//   		Urls: []*string{
//   			jsii.String("urls"),
//   		},
//
//   		// the properties below are optional
//   		CrawlAttachments: jsii.Boolean(false),
//   		DisableLocalGroups: jsii.Boolean(false),
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		SslCertificateS3Path: &S3PathProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
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
//   		Urls: &WebCrawlerUrlsProperty{
//   			SeedUrlConfiguration: &WebCrawlerSeedUrlConfigurationProperty{
//   				SeedUrls: []*string{
//   					jsii.String("seedUrls"),
//   				},
//
//   				// the properties below are optional
//   				WebCrawlerMode: jsii.String("webCrawlerMode"),
//   			},
//   			SiteMapsConfiguration: &WebCrawlerSiteMapsConfigurationProperty{
//   				SiteMaps: []*string{
//   					jsii.String("siteMaps"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
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
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//
//   			// the properties below are optional
//   			Credentials: jsii.String("credentials"),
//   		},
//   		UrlExclusionPatterns: []*string{
//   			jsii.String("urlExclusionPatterns"),
//   		},
//   		UrlInclusionPatterns: []*string{
//   			jsii.String("urlInclusionPatterns"),
//   		},
//   	},
//   	WorkDocsConfiguration: &WorkDocsConfigurationProperty{
//   		OrganizationId: jsii.String("organizationId"),
//
//   		// the properties below are optional
//   		CrawlComments: jsii.Boolean(false),
//   		ExclusionPatterns: []*string{
//   			jsii.String("exclusionPatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		InclusionPatterns: []*string{
//   			jsii.String("inclusionPatterns"),
//   		},
//   		UseChangeLog: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html
//
type CfnDataSource_DataSourceConfigurationProperty struct {
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// Provides the configuration information required for Amazon Kendra Web Crawler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-webcrawlerconfiguration
	//
	WebCrawlerConfiguration interface{} `field:"optional" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
	// Provides the configuration information to connect to Amazon WorkDocs as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourceconfiguration.html#cfn-kendra-datasource-datasourceconfiguration-workdocsconfiguration
	//
	WorkDocsConfiguration interface{} `field:"optional" json:"workDocsConfiguration" yaml:"workDocsConfiguration"`
}

