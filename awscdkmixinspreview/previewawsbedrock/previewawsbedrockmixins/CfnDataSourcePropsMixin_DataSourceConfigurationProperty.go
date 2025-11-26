package previewawsbedrockmixins


// The connection configuration for the data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	ConfluenceConfiguration: &ConfluenceDataSourceConfigurationProperty{
//   		CrawlerConfiguration: &ConfluenceCrawlerConfigurationProperty{
//   			FilterConfiguration: &CrawlFilterConfigurationProperty{
//   				PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   					Filters: []interface{}{
//   						&PatternObjectFilterProperty{
//   							ExclusionFilters: []*string{
//   								jsii.String("exclusionFilters"),
//   							},
//   							InclusionFilters: []*string{
//   								jsii.String("inclusionFilters"),
//   							},
//   							ObjectType: jsii.String("objectType"),
//   						},
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		SourceConfiguration: &ConfluenceSourceConfigurationProperty{
//   			AuthType: jsii.String("authType"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			HostType: jsii.String("hostType"),
//   			HostUrl: jsii.String("hostUrl"),
//   		},
//   	},
//   	S3Configuration: &S3DataSourceConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   		InclusionPrefixes: []*string{
//   			jsii.String("inclusionPrefixes"),
//   		},
//   	},
//   	SalesforceConfiguration: &SalesforceDataSourceConfigurationProperty{
//   		CrawlerConfiguration: &SalesforceCrawlerConfigurationProperty{
//   			FilterConfiguration: &CrawlFilterConfigurationProperty{
//   				PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   					Filters: []interface{}{
//   						&PatternObjectFilterProperty{
//   							ExclusionFilters: []*string{
//   								jsii.String("exclusionFilters"),
//   							},
//   							InclusionFilters: []*string{
//   								jsii.String("inclusionFilters"),
//   							},
//   							ObjectType: jsii.String("objectType"),
//   						},
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		SourceConfiguration: &SalesforceSourceConfigurationProperty{
//   			AuthType: jsii.String("authType"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			HostUrl: jsii.String("hostUrl"),
//   		},
//   	},
//   	SharePointConfiguration: &SharePointDataSourceConfigurationProperty{
//   		CrawlerConfiguration: &SharePointCrawlerConfigurationProperty{
//   			FilterConfiguration: &CrawlFilterConfigurationProperty{
//   				PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   					Filters: []interface{}{
//   						&PatternObjectFilterProperty{
//   							ExclusionFilters: []*string{
//   								jsii.String("exclusionFilters"),
//   							},
//   							InclusionFilters: []*string{
//   								jsii.String("inclusionFilters"),
//   							},
//   							ObjectType: jsii.String("objectType"),
//   						},
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		SourceConfiguration: &SharePointSourceConfigurationProperty{
//   			AuthType: jsii.String("authType"),
//   			CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   			Domain: jsii.String("domain"),
//   			HostType: jsii.String("hostType"),
//   			SiteUrls: []*string{
//   				jsii.String("siteUrls"),
//   			},
//   			TenantId: jsii.String("tenantId"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	WebConfiguration: &WebDataSourceConfigurationProperty{
//   		CrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   			CrawlerLimits: &WebCrawlerLimitsProperty{
//   				MaxPages: jsii.Number(123),
//   				RateLimit: jsii.Number(123),
//   			},
//   			ExclusionFilters: []*string{
//   				jsii.String("exclusionFilters"),
//   			},
//   			InclusionFilters: []*string{
//   				jsii.String("inclusionFilters"),
//   			},
//   			Scope: jsii.String("scope"),
//   			UserAgent: jsii.String("userAgent"),
//   			UserAgentHeader: jsii.String("userAgentHeader"),
//   		},
//   		SourceConfiguration: &WebSourceConfigurationProperty{
//   			UrlConfiguration: &UrlConfigurationProperty{
//   				SeedUrls: []interface{}{
//   					&SeedUrlProperty{
//   						Url: jsii.String("url"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html
//
type CfnDataSourcePropsMixin_DataSourceConfigurationProperty struct {
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
	// The type of data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The configuration of web URLs to crawl for your data source. You should be authorized to crawl the URLs.
	//
	// > Crawling web URLs as your data source is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-datasourceconfiguration.html#cfn-bedrock-datasource-datasourceconfiguration-webconfiguration
	//
	WebConfiguration interface{} `field:"optional" json:"webConfiguration" yaml:"webConfiguration"`
}

