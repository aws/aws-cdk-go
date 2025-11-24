package mixinsawsbedrock


// The rate limits for the URLs that you want to crawl.
//
// You should be authorized to crawl the URLs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webCrawlerLimitsProperty := &WebCrawlerLimitsProperty{
//   	MaxPages: jsii.Number(123),
//   	RateLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerlimits.html
//
type CfnDataSourcePropsMixin_WebCrawlerLimitsProperty struct {
	// The max number of web pages crawled from your source URLs, up to 25,000 pages.
	//
	// If the web pages exceed this limit, the data source sync will fail and no web pages will be ingested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerlimits.html#cfn-bedrock-datasource-webcrawlerlimits-maxpages
	//
	MaxPages *float64 `field:"optional" json:"maxPages" yaml:"maxPages"`
	// The max rate at which pages are crawled, up to 300 per minute per host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerlimits.html#cfn-bedrock-datasource-webcrawlerlimits-ratelimit
	//
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

