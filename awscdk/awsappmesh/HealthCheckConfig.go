package awsappmesh


// All Properties for Health Checks for mesh endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfig := &HealthCheckConfig{
//   	VirtualGatewayHealthCheck: &VirtualGatewayHealthCheckPolicyProperty{
//   		HealthyThreshold: jsii.Number(123),
//   		IntervalMillis: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		TimeoutMillis: jsii.Number(123),
//   		UnhealthyThreshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		Path: jsii.String("path"),
//   		Port: jsii.Number(123),
//   	},
//   	VirtualNodeHealthCheck: &HealthCheckProperty{
//   		HealthyThreshold: jsii.Number(123),
//   		IntervalMillis: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		TimeoutMillis: jsii.Number(123),
//   		UnhealthyThreshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		Path: jsii.String("path"),
//   		Port: jsii.Number(123),
//   	},
//   }
//
type HealthCheckConfig struct {
	// VirtualGateway CFN configuration for Health Checks.
	VirtualGatewayHealthCheck *CfnVirtualGateway_VirtualGatewayHealthCheckPolicyProperty `field:"optional" json:"virtualGatewayHealthCheck" yaml:"virtualGatewayHealthCheck"`
	// VirtualNode CFN configuration for Health Checks.
	VirtualNodeHealthCheck *CfnVirtualNode_HealthCheckProperty `field:"optional" json:"virtualNodeHealthCheck" yaml:"virtualNodeHealthCheck"`
}

