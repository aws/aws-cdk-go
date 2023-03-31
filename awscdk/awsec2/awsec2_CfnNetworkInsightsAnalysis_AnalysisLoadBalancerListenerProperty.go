package awsec2


// Describes a load balancer listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisLoadBalancerListenerProperty := &analysisLoadBalancerListenerProperty{
//   	instancePort: jsii.Number(123),
//   	loadBalancerPort: jsii.Number(123),
//   }
//
type CfnNetworkInsightsAnalysis_AnalysisLoadBalancerListenerProperty struct {
	// [Classic Load Balancers] The back-end port for the listener.
	InstancePort *float64 `field:"optional" json:"instancePort" yaml:"instancePort"`
	// The port on which the load balancer is listening.
	LoadBalancerPort *float64 `field:"optional" json:"loadBalancerPort" yaml:"loadBalancerPort"`
}

