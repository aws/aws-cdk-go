package awsaps


// Scraper metrics destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &DestinationProperty{
//   	AmpConfiguration: &AmpConfigurationProperty{
//   		WorkspaceArn: jsii.String("workspaceArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-destination.html
//
type CfnScraper_DestinationProperty struct {
	// Configuration for Amazon Managed Prometheus metrics destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-destination.html#cfn-aps-scraper-destination-ampconfiguration
	//
	AmpConfiguration interface{} `field:"required" json:"ampConfiguration" yaml:"ampConfiguration"`
}

