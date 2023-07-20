package awsappmesh


// An object that represents the health check policy for a virtual node's listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckProperty := &HealthCheckProperty{
//   	HealthyThreshold: jsii.Number(123),
//   	IntervalMillis: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	TimeoutMillis: jsii.Number(123),
//   	UnhealthyThreshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	Path: jsii.String("path"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-healthcheck.html
//
type CfnVirtualNode_HealthCheckProperty struct {
	// The number of consecutive successful health checks that must occur before declaring listener healthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-healthcheck.html#cfn-appmesh-virtualnode-healthcheck-healthythreshold
	//
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period in milliseconds between each health check execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-healthcheck.html#cfn-appmesh-virtualnode-healthcheck-intervalmillis
	//
	IntervalMillis *float64 `field:"required" json:"intervalMillis" yaml:"intervalMillis"`
	// The protocol for the health check request.
	//
	// If you specify `grpc` , then your service must conform to the [GRPC Health Checking Protocol](https://docs.aws.amazon.com/https://github.com/grpc/grpc/blob/master/doc/health-checking.md) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-healthcheck.html#cfn-appmesh-virtualnode-healthcheck-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The amount of time to wait when receiving a response from the health check, in milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-healthcheck.html#cfn-appmesh-virtualnode-healthcheck-timeoutmillis
	//
	TimeoutMillis *float64 `field:"required" json:"timeoutMillis" yaml:"timeoutMillis"`
	// The number of consecutive failed health checks that must occur before declaring a virtual node unhealthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-healthcheck.html#cfn-appmesh-virtualnode-healthcheck-unhealthythreshold
	//
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// The destination path for the health check request.
	//
	// This value is only used if the specified protocol is HTTP or HTTP/2. For any other protocol, this value is ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-healthcheck.html#cfn-appmesh-virtualnode-healthcheck-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The destination port for the health check request.
	//
	// This port must match the port defined in the `PortMapping` for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-healthcheck.html#cfn-appmesh-virtualnode-healthcheck-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

