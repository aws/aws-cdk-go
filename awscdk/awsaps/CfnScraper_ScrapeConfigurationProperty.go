package awsaps


// Scraper configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scrapeConfigurationProperty := &ScrapeConfigurationProperty{
//   	ConfigurationBlob: jsii.String("configurationBlob"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scrapeconfiguration.html
//
type CfnScraper_ScrapeConfigurationProperty struct {
	// Prometheus compatible scrape configuration in base64 encoded blob format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scrapeconfiguration.html#cfn-aps-scraper-scrapeconfiguration-configurationblob
	//
	ConfigurationBlob *string `field:"required" json:"configurationBlob" yaml:"configurationBlob"`
}

