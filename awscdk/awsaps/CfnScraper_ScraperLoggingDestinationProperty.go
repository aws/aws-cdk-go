package awsaps


// The destination where scraper logs are sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scraperLoggingDestinationProperty := &ScraperLoggingDestinationProperty{
//   	CloudWatchLogs: &CloudWatchLogDestinationProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scraperloggingdestination.html
//
type CfnScraper_ScraperLoggingDestinationProperty struct {
	// The CloudWatch Logs configuration for the scraper logging destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-scraperloggingdestination.html#cfn-aps-scraper-scraperloggingdestination-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
}

