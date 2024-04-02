package awsaps


// Configuration for Amazon Managed Prometheus metrics destination.
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
	// ARN of an Amazon Managed Prometheus workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-ampconfiguration.html#cfn-aps-scraper-ampconfiguration-workspacearn
	//
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
}

