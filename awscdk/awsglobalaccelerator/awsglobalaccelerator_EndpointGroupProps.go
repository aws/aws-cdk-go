package awsglobalaccelerator

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Property of the EndpointGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var endpoint iEndpoint
//   var listener listener
//
//   endpointGroupProps := &endpointGroupProps{
//   	listener: listener,
//
//   	// the properties below are optional
//   	endpointGroupName: jsii.String("endpointGroupName"),
//   	endpoints: []*iEndpoint{
//   		endpoint,
//   	},
//   	healthCheckInterval: duration,
//   	healthCheckPath: jsii.String("healthCheckPath"),
//   	healthCheckPort: jsii.Number(123),
//   	healthCheckProtocol: awscdk.Aws_globalaccelerator.healthCheckProtocol_TCP,
//   	healthCheckThreshold: jsii.Number(123),
//   	portOverrides: []portOverride{
//   		&portOverride{
//   			endpointPort: jsii.Number(123),
//   			listenerPort: jsii.Number(123),
//   		},
//   	},
//   	region: jsii.String("region"),
//   	trafficDialPercentage: jsii.Number(123),
//   }
//
// Experimental.
type EndpointGroupProps struct {
	// Name of the endpoint group.
	// Experimental.
	EndpointGroupName *string `field:"optional" json:"endpointGroupName" yaml:"endpointGroupName"`
	// Initial list of endpoints for this group.
	// Experimental.
	Endpoints *[]IEndpoint `field:"optional" json:"endpoints" yaml:"endpoints"`
	// The time between health checks for each endpoint.
	//
	// Must be either 10 or 30 seconds.
	// Experimental.
	HealthCheckInterval awscdk.Duration `field:"optional" json:"healthCheckInterval" yaml:"healthCheckInterval"`
	// The ping path for health checks (if the protocol is HTTP(S)).
	// Experimental.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port used to perform health checks.
	// Experimental.
	HealthCheckPort *float64 `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// The protocol used to perform health checks.
	// Experimental.
	HealthCheckProtocol HealthCheckProtocol `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy.
	// Experimental.
	HealthCheckThreshold *float64 `field:"optional" json:"healthCheckThreshold" yaml:"healthCheckThreshold"`
	// Override the destination ports used to route traffic to an endpoint.
	//
	// Unless overridden, the port used to hit the endpoint will be the same as the port
	// that traffic arrives on at the listener.
	// Experimental.
	PortOverrides *[]*PortOverride `field:"optional" json:"portOverrides" yaml:"portOverrides"`
	// The AWS Region where the endpoint group is located.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The percentage of traffic to send to this AWS Region.
	//
	// The percentage is applied to the traffic that would otherwise have been
	// routed to the Region based on optimal routing. Additional traffic is
	// distributed to other endpoint groups for this listener.
	// Experimental.
	TrafficDialPercentage *float64 `field:"optional" json:"trafficDialPercentage" yaml:"trafficDialPercentage"`
	// The Amazon Resource Name (ARN) of the listener.
	// Experimental.
	Listener IListener `field:"required" json:"listener" yaml:"listener"`
}

