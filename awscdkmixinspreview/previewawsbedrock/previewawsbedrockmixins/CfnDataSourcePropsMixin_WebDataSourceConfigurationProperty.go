package previewawsbedrockmixins


// The configuration details for the web data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webDataSourceConfigurationProperty := &WebDataSourceConfigurationProperty{
//   	CrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   		CrawlerLimits: &WebCrawlerLimitsProperty{
//   			MaxPages: jsii.Number(123),
//   			RateLimit: jsii.Number(123),
//   		},
//   		ExclusionFilters: []*string{
//   			jsii.String("exclusionFilters"),
//   		},
//   		InclusionFilters: []*string{
//   			jsii.String("inclusionFilters"),
//   		},
//   		Scope: jsii.String("scope"),
//   		UserAgent: jsii.String("userAgent"),
//   		UserAgentHeader: jsii.String("userAgentHeader"),
//   	},
//   	SourceConfiguration: &WebSourceConfigurationProperty{
//   		UrlConfiguration: &UrlConfigurationProperty{
//   			SeedUrls: []interface{}{
//   				&SeedUrlProperty{
//   					Url: jsii.String("url"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webdatasourceconfiguration.html
//
type CfnDataSourcePropsMixin_WebDataSourceConfigurationProperty struct {
	// The Web Crawler configuration details for the web data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webdatasourceconfiguration.html#cfn-bedrock-datasource-webdatasourceconfiguration-crawlerconfiguration
	//
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// The source configuration details for the web data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webdatasourceconfiguration.html#cfn-bedrock-datasource-webdatasourceconfiguration-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

