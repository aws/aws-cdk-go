package awsec2


// Describes a load balancer target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisLoadBalancerTargetProperty := &AnalysisLoadBalancerTargetProperty{
//   	Address: jsii.String("address"),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Instance: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html
//
type CfnNetworkInsightsAnalysis_AnalysisLoadBalancerTargetProperty struct {
	// The IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancertarget-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancertarget-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Information about the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancertarget-instance
	//
	Instance interface{} `field:"optional" json:"instance" yaml:"instance"`
	// The port on which the target is listening.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancertarget.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancertarget-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

