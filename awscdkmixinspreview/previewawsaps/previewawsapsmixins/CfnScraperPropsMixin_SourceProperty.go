package previewawsapsmixins


// The source of collected metrics for a scraper.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceProperty := &SourceProperty{
//   	EksConfiguration: &EksConfigurationProperty{
//   		ClusterArn: jsii.String("clusterArn"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-source.html
//
type CfnScraperPropsMixin_SourceProperty struct {
	// The Amazon EKS cluster from which a scraper collects metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-source.html#cfn-aps-scraper-source-eksconfiguration
	//
	EksConfiguration interface{} `field:"optional" json:"eksConfiguration" yaml:"eksConfiguration"`
	// Configuration for VPC metrics source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-source.html#cfn-aps-scraper-source-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

