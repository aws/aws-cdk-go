package awsbedrock


// The configuration details for the web data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webDataSourceConfigurationProperty := &WebDataSourceConfigurationProperty{
//   	SourceConfiguration: &WebSourceConfigurationProperty{
//   		UrlConfiguration: &UrlConfigurationProperty{
//   			SeedUrls: []interface{}{
//   				&SeedUrlProperty{
//   					Url: jsii.String("url"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	CrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   		CrawlerLimits: &WebCrawlerLimitsProperty{
//   			RateLimit: jsii.Number(123),
//   		},
//   		ExclusionFilters: []*string{
//   			jsii.String("exclusionFilters"),
//   		},
//   		InclusionFilters: []*string{
//   			jsii.String("inclusionFilters"),
//   		},
//   		Scope: jsii.String("scope"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webdatasourceconfiguration.html
//
type CfnDataSource_WebDataSourceConfigurationProperty struct {
	// The source configuration details for the web data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webdatasourceconfiguration.html#cfn-bedrock-datasource-webdatasourceconfiguration-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"required" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The Web Crawler configuration details for the web data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webdatasourceconfiguration.html#cfn-bedrock-datasource-webdatasourceconfiguration-crawlerconfiguration
	//
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
}

