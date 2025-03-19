package awsbedrock


// The connection configuration for the data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ConfluenceConfiguration: &ConfluenceDataSourceConfigurationProperty{
//   		SourceConfiguration: &ConfluenceSourceConfigurationProperty{
//   			AuthType: jsii.String("authType"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			HostType: jsii.String("hostType"),
//   			HostUrl: jsii.String("hostUrl"),
//   		},
//
//   		// the properties below are optional
//   		CrawlerConfiguration: &ConfluenceCrawlerConfigurationProperty{
//   			FilterConfiguration: &CrawlFilterConfigurationProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   					Filters: []interface{}{
//   						&PatternObjectFilterProperty{
//   							ObjectType: jsii.String("objectType"),
//
//   							// the properties below are optional
//   							ExclusionFilters: []*string{
//   								jsii.String("exclusionFilters"),
//   							},
//   							InclusionFilters: []*string{
//   								jsii.String("inclusionFilters"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	S3Configuration: &S3DataSourceConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//
//   		// the properties below are optional
//   		BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   		InclusionPrefixes: []*string{
//   			jsii.String("inclusionPrefixes"),
//   		},
//   	},
//   	SalesforceConfiguration: &SalesforceDataSourceConfigurationProperty{
//   		SourceConfiguration: &SalesforceSourceConfigurationProperty{
//   			AuthType: jsii.String("authType"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			HostUrl: jsii.String("hostUrl"),
//   		},
//
//   		// the properties below are optional
//   		CrawlerConfiguration: &SalesforceCrawlerConfigurationProperty{
//   			FilterConfiguration: &CrawlFilterConfigurationProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   					Filters: []interface{}{
//   						&PatternObjectFilterProperty{
//   							ObjectType: jsii.String("objectType"),
//
//   							// the properties below are optional
//   							ExclusionFilters: []*string{
//   								jsii.String("exclusionFilters"),
//   							},
//   							InclusionFilters: []*string{
//   								jsii.String("inclusionFilters"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SharePointConfiguration: &SharePointDataSourceConfigurationProperty{
//   		SourceConfiguration: &SharePointSourceConfigurationProperty{
//   			AuthType: jsii.String("authType"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			Domain: jsii.String("domain"),
//   			HostType: jsii.String("hostType"),
//   			SiteUrls: []*string{
//   				jsii.String("siteUrls"),
//   			},
//
//   			// the properties below are optional
//   			TenantId: jsii.String("tenantId"),
//   		},
//
//   		// the properties below are optional
//   		CrawlerConfiguration: &SharePointCrawlerConfigurationProperty{
//   			FilterConfiguration: &CrawlFilterConfigurationProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   					Filters: []interface{}{
//   						&PatternObjectFilterProperty{
//   							ObjectType: jsii.String("objectType"),
//
//   							// the properties below are optional
//   							ExclusionFilters: []*string{
//   								jsii.String("exclusionFilters"),
//   							},
//   							InclusionFilters: []*string{
//   								jsii.String("inclusionFilters"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	WebConfiguration: &WebDataSourceConfigurationProperty{
//   		SourceConfiguration: &WebSourceConfigurationProperty{
//   			UrlConfiguration: &UrlConfigurationProperty{
//   				SeedUrls: []interface{}{
//   					&SeedUrlProperty{
//   						Url: jsii.String("url"),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		CrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   			CrawlerLimits: &WebCrawlerLimitsProperty{
//   				RateLimit: jsii.Number(123),
//   			},
//   			ExclusionFilters: []*string{
//   				jsii.String("exclusionFilters"),
//   			},
//   			InclusionFilters: []*string{
//   				jsii.String("inclusionFilters"),
//   			},
//   			Scope: jsii.String("scope"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html
//
type CfnDataSource_DataSourceConfigurationProperty struct {
	// The type of data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration information to connect to Confluence as your data source.
	//
	// > Confluence data source connector is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-confluenceconfiguration
	//
	ConfluenceConfiguration interface{} `field:"optional" json:"confluenceConfiguration" yaml:"confluenceConfiguration"`
	// The configuration information to connect to Amazon S3 as your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// The configuration information to connect to Salesforce as your data source.
	//
	// > Salesforce data source connector is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-salesforceconfiguration
	//
	SalesforceConfiguration interface{} `field:"optional" json:"salesforceConfiguration" yaml:"salesforceConfiguration"`
	// The configuration information to connect to SharePoint as your data source.
	//
	// > SharePoint data source connector is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-sharepointconfiguration
	//
	SharePointConfiguration interface{} `field:"optional" json:"sharePointConfiguration" yaml:"sharePointConfiguration"`
	// The configuration of web URLs to crawl for your data source. You should be authorized to crawl the URLs.
	//
	// > Crawling web URLs as your data source is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-webconfiguration
	//
	WebConfiguration interface{} `field:"optional" json:"webConfiguration" yaml:"webConfiguration"`
}

