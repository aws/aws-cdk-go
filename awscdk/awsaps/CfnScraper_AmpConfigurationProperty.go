package awsaps


// The `AmpConfiguration` structure defines the Amazon Managed Service for Prometheus instance a scraper should send metrics to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ampConfigurationProperty := &AmpConfigurationProperty{
//   	WorkspaceArn: jsii.String("workspaceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-ampconfiguration.html
//
type CfnScraper_AmpConfigurationProperty struct {
	// ARN of the Amazon Managed Service for Prometheus workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-ampconfiguration.html#cfn-aps-scraper-ampconfiguration-workspacearn
	//
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
}

