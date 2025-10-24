package awsglobalaccelerator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Basic options for creating a new EndpointGroup.
//
// Example:
//   var alb ApplicationLoadBalancer
//   var listener Listener
//
//
//   listener.AddEndpointGroup(jsii.String("Group"), &EndpointGroupOptions{
//   	Endpoints: []IEndpoint{
//   		ga_endpoints.NewApplicationLoadBalancerEndpoint(alb, &ApplicationLoadBalancerEndpointOptions{
//   			Weight: jsii.Number(128),
//   			PreserveClientIp: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
type EndpointGroupOptions struct {
	// Name of the endpoint group.
	// Default: - logical ID of the resource.
	//
	EndpointGroupName *string `field:"optional" json:"endpointGroupName" yaml:"endpointGroupName"`
	// Initial list of endpoints for this group.
	// Default: - Group is initially empty.
	//
	Endpoints *[]IEndpoint `field:"optional" json:"endpoints" yaml:"endpoints"`
	// The time between health checks for each endpoint.
	//
	// Must be either 10 or 30 seconds.
	// Default: Duration.seconds(30)
	//
	HealthCheckInterval awscdk.Duration `field:"optional" json:"healthCheckInterval" yaml:"healthCheckInterval"`
	// The ping path for health checks (if the protocol is HTTP(S)).
	// Default: '/'.
	//
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port used to perform health checks.
	// Default: - The listener's port.
	//
	HealthCheckPort *float64 `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// The protocol used to perform health checks.
	// Default: HealthCheckProtocol.TCP
	//
	HealthCheckProtocol HealthCheckProtocol `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.
	// Default: 3.
	//
	HealthCheckThreshold *float64 `field:"optional" json:"healthCheckThreshold" yaml:"healthCheckThreshold"`
	// Override the destination ports used to route traffic to an endpoint.
	//
	// Unless overridden, the port used to hit the endpoint will be the same as the port
	// that traffic arrives on at the listener.
	// Default: - No overrides.
	//
	PortOverrides *[]*PortOverride `field:"optional" json:"portOverrides" yaml:"portOverrides"`
	// The AWS Region where the endpoint group is located.
	// Default: - region of the first endpoint in this group, or the stack region if that region can't be determined.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The percentage of traffic to send to this AWS Region.
	//
	// The percentage is applied to the traffic that would otherwise have been
	// routed to the Region based on optimal routing. Additional traffic is
	// distributed to other endpoint groups for this listener.
	// Default: 100.
	//
	TrafficDialPercentage *float64 `field:"optional" json:"trafficDialPercentage" yaml:"trafficDialPercentage"`
}

