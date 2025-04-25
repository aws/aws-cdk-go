package awsec2


// Describes a load balancer listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisLoadBalancerListenerProperty := &AnalysisLoadBalancerListenerProperty{
//   	InstancePort: jsii.Number(123),
//   	LoadBalancerPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancerlistener.html
//
type CfnNetworkInsightsAnalysis_AnalysisLoadBalancerListenerProperty struct {
	// [Classic Load Balancers] The back-end port for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancerlistener.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancerlistener-instanceport
	//
	InstancePort *float64 `field:"optional" json:"instancePort" yaml:"instancePort"`
	// The port on which the load balancer is listening.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisloadbalancerlistener.html#cfn-ec2-networkinsightsanalysis-analysisloadbalancerlistener-loadbalancerport
	//
	LoadBalancerPort *float64 `field:"optional" json:"loadBalancerPort" yaml:"loadBalancerPort"`
}

