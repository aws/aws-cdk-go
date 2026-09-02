package awsaps


// Configuration for CloudWatch metrics destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchConfigurationProperty := &CloudWatchConfigurationProperty{
//   	DatasetArn: jsii.String("datasetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-cloudwatchconfiguration.html
//
type CfnScraper_CloudWatchConfigurationProperty struct {
	// ARN of a CloudWatch dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-cloudwatchconfiguration.html#cfn-aps-scraper-cloudwatchconfiguration-datasetarn
	//
	DatasetArn *string `field:"required" json:"datasetArn" yaml:"datasetArn"`
}

