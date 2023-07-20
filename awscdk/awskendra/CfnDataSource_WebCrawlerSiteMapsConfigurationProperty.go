package awskendra


// Provides the configuration information of the sitemap URLs to crawl.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerSiteMapsConfigurationProperty := &WebCrawlerSiteMapsConfigurationProperty{
//   	SiteMaps: []*string{
//   		jsii.String("siteMaps"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlersitemapsconfiguration.html
//
type CfnDataSource_WebCrawlerSiteMapsConfigurationProperty struct {
	// The list of sitemap URLs of the websites you want to crawl.
	//
	// The list can include a maximum of three sitemap URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlersitemapsconfiguration.html#cfn-kendra-datasource-webcrawlersitemapsconfiguration-sitemaps
	//
	SiteMaps *[]*string `field:"required" json:"siteMaps" yaml:"siteMaps"`
}

