package awsaps


// Configuration for scraper logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scraperLoggingConfigurationProperty := &ScraperLoggingConfigurationProperty{
//   	LoggingDestination: &ScraperLoggingDestinationProperty{
//   		CloudWatchLogs: &CloudWatchLogDestinationProperty{
//   			LogGroupArn: jsii.String("logGroupArn"),
//   		},
//   	},
//   	ScraperComponents: []interface{}{
//   		&ScraperComponentProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Config: &ComponentConfigProperty{
//   				Options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scraperloggingconfiguration.html
//
type CfnScraper_ScraperLoggingConfigurationProperty struct {
	// Destination for scraper logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scraperloggingconfiguration.html#cfn-aps-scraper-scraperloggingconfiguration-loggingdestination
	//
	LoggingDestination interface{} `field:"required" json:"loggingDestination" yaml:"loggingDestination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scraperloggingconfiguration.html#cfn-aps-scraper-scraperloggingconfiguration-scrapercomponents
	//
	ScraperComponents interface{} `field:"required" json:"scraperComponents" yaml:"scraperComponents"`
}

