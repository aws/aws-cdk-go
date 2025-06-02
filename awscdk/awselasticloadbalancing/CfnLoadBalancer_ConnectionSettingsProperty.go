package awselasticloadbalancing


// Specifies the idle timeout value for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionSettingsProperty := &ConnectionSettingsProperty{
//   	IdleTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-connectionsettings.html
//
type CfnLoadBalancer_ConnectionSettingsProperty struct {
	// The time, in seconds, that the connection is allowed to be idle (no data has been sent over the connection) before it is closed by the load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancing-loadbalancer-connectionsettings.html#cfn-elasticloadbalancing-loadbalancer-connectionsettings-idletimeout
	//
	IdleTimeout *float64 `field:"required" json:"idleTimeout" yaml:"idleTimeout"`
}

