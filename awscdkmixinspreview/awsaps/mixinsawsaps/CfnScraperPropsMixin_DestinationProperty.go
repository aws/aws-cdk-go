package mixinsawsaps


// Where to send the metrics from a scraper.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationProperty := &DestinationProperty{
//   	AmpConfiguration: &AmpConfigurationProperty{
//   		WorkspaceArn: jsii.String("workspaceArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-destination.html
//
type CfnScraperPropsMixin_DestinationProperty struct {
	// The Amazon Managed Service for Prometheus workspace to send metrics to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-destination.html#cfn-aps-scraper-destination-ampconfiguration
	//
	AmpConfiguration interface{} `field:"optional" json:"ampConfiguration" yaml:"ampConfiguration"`
}

