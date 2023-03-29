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
type CfnNetworkInsightsAnalysis_AdditionalDetailProperty struct {
	// The additional detail code.
	AdditionalDetailType *string `field:"optional" json:"additionalDetailType" yaml:"additionalDetailType"`
	// The path component.
	Component interface{} `field:"optional" json:"component" yaml:"component"`
	// `CfnNetworkInsightsAnalysis.AdditionalDetailProperty.LoadBalancers`.
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// `CfnNetworkInsightsAnalysis.AdditionalDetailProperty.ServiceName`.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

