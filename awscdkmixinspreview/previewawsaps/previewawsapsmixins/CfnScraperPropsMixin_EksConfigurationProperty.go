package previewawsapsmixins


// The `EksConfiguration` structure describes the connection to the Amazon EKS cluster from which a scraper collects metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eksConfigurationProperty := &EksConfigurationProperty{
//   	ClusterArn: jsii.String("clusterArn"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-eksconfiguration.html
//
type CfnScraperPropsMixin_EksConfigurationProperty struct {
	// ARN of the Amazon EKS cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-eksconfiguration.html#cfn-aps-scraper-eksconfiguration-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// A list of the security group IDs for the Amazon EKS cluster VPC configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-eksconfiguration.html#cfn-aps-scraper-eksconfiguration-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs for the Amazon EKS cluster VPC configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-eksconfiguration.html#cfn-aps-scraper-eksconfiguration-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

