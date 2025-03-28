package awsaps


// The source of collected metrics for a scraper.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	EksConfiguration: &EksConfigurationProperty{
//   		ClusterArn: jsii.String("clusterArn"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-source.html
//
type CfnScraper_SourceProperty struct {
	// The Amazon EKS cluster from which a scraper collects metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-source.html#cfn-aps-scraper-source-eksconfiguration
	//
	EksConfiguration interface{} `field:"required" json:"eksConfiguration" yaml:"eksConfiguration"`
}

