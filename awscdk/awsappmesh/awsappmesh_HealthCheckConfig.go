package awsappmesh


// All Properties for Health Checks for mesh endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfig := &healthCheckConfig{
//   	virtualGatewayHealthCheck: &virtualGatewayHealthCheckPolicyProperty{
//   		healthyThreshold: jsii.Number(123),
//   		intervalMillis: jsii.Number(123),
//   		protocol: jsii.String("protocol"),
//   		timeoutMillis: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		path: jsii.String("path"),
//   		port: jsii.Number(123),
//   	},
//   	virtualNodeHealthCheck: &healthCheckProperty{
//   		healthyThreshold: jsii.Number(123),
//   		intervalMillis: jsii.Number(123),
//   		protocol: jsii.String("protocol"),
//   		timeoutMillis: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		path: jsii.String("path"),
//   		port: jsii.Number(123),
//   	},
//   }
//
type HealthCheckConfig struct {
	// VirtualGateway CFN configuration for Health Checks.
	VirtualGatewayHealthCheck *CfnVirtualGateway_VirtualGatewayHealthCheckPolicyProperty `field:"optional" json:"virtualGatewayHealthCheck" yaml:"virtualGatewayHealthCheck"`
	// VirtualNode CFN configuration for Health Checks.
	VirtualNodeHealthCheck *CfnVirtualNode_HealthCheckProperty `field:"optional" json:"virtualNodeHealthCheck" yaml:"virtualNodeHealthCheck"`
}

