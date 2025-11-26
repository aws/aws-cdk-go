package previewawsappmeshmixins


// An object that represents the health check policy for a virtual gateway's listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayHealthCheckPolicyProperty := &VirtualGatewayHealthCheckPolicyProperty{
//   	HealthyThreshold: jsii.Number(123),
//   	IntervalMillis: jsii.Number(123),
//   	Path: jsii.String("path"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	TimeoutMillis: jsii.Number(123),
//   	UnhealthyThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayHealthCheckPolicyProperty struct {
	// The number of consecutive successful health checks that must occur before declaring the listener healthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.html#cfn-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy-healthythreshold
	//
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The time period in milliseconds between each health check execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.html#cfn-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy-intervalmillis
	//
	IntervalMillis *float64 `field:"optional" json:"intervalMillis" yaml:"intervalMillis"`
	// The destination path for the health check request.
	//
	// This value is only used if the specified protocol is HTTP or HTTP/2. For any other protocol, this value is ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.html#cfn-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The destination port for the health check request.
	//
	// This port must match the port defined in the `PortMapping` for the listener.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.html#cfn-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol for the health check request.
	//
	// If you specify `grpc` , then your service must conform to the [GRPC Health Checking Protocol](https://docs.aws.amazon.com/https://github.com/grpc/grpc/blob/master/doc/health-checking.md) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.html#cfn-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The amount of time to wait when receiving a response from the health check, in milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.html#cfn-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy-timeoutmillis
	//
	TimeoutMillis *float64 `field:"optional" json:"timeoutMillis" yaml:"timeoutMillis"`
	// The number of consecutive failed health checks that must occur before declaring a virtual gateway unhealthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy.html#cfn-appmesh-virtualgateway-virtualgatewayhealthcheckpolicy-unhealthythreshold
	//
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

