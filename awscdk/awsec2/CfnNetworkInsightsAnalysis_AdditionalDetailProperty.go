package awsec2


// Describes an additional detail for a path analysis.
//
// For more information, see [Reachability Analyzer additional detail codes](https://docs.aws.amazon.com/vpc/latest/reachability/additional-detail-codes.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   additionalDetailProperty := &AdditionalDetailProperty{
//   	AdditionalDetailType: jsii.String("additionalDetailType"),
//   	Component: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	LoadBalancers: []interface{}{
//   		&AnalysisComponentProperty{
//   			Arn: jsii.String("arn"),
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-additionaldetail.html
//
type CfnNetworkInsightsAnalysis_AdditionalDetailProperty struct {
	// The additional detail code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-additionaldetail.html#cfn-ec2-networkinsightsanalysis-additionaldetail-additionaldetailtype
	//
	AdditionalDetailType *string `field:"optional" json:"additionalDetailType" yaml:"additionalDetailType"`
	// The path component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-additionaldetail.html#cfn-ec2-networkinsightsanalysis-additionaldetail-component
	//
	Component interface{} `field:"optional" json:"component" yaml:"component"`
	// The load balancers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-additionaldetail.html#cfn-ec2-networkinsightsanalysis-additionaldetail-loadbalancers
	//
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// The name of the VPC endpoint service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-additionaldetail.html#cfn-ec2-networkinsightsanalysis-additionaldetail-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

