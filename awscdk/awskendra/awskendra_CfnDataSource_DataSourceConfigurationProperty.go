package awskendra


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

