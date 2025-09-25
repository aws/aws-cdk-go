package awsaps


// A component of a Amazon Managed Service for Prometheus scraper that can be configured for logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scraperComponentProperty := &ScraperComponentProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Config: &ComponentConfigProperty{
//   		Options: map[string]*string{
//   			"optionsKey": jsii.String("options"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scrapercomponent.html
//
type CfnScraper_ScraperComponentProperty struct {
	// The type of the scraper component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scrapercomponent.html#cfn-aps-scraper-scrapercomponent-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration settings for the scraper component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scrapercomponent.html#cfn-aps-scraper-scrapercomponent-config
	//
	Config interface{} `field:"optional" json:"config" yaml:"config"`
}

