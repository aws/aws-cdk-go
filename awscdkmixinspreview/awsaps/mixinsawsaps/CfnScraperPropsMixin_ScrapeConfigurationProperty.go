package mixinsawsaps


// A scrape configuration for a scraper, base 64 encoded.
//
// For more information, see [Scraper configuration](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html#AMP-collector-configuration) in the *Amazon Managed Service for Prometheus User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scrapeConfigurationProperty := &ScrapeConfigurationProperty{
//   	ConfigurationBlob: jsii.String("configurationBlob"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scrapeconfiguration.html
//
type CfnScraperPropsMixin_ScrapeConfigurationProperty struct {
	// The base 64 encoded scrape configuration file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scrapeconfiguration.html#cfn-aps-scraper-scrapeconfiguration-configurationblob
	//
	ConfigurationBlob *string `field:"optional" json:"configurationBlob" yaml:"configurationBlob"`
}

