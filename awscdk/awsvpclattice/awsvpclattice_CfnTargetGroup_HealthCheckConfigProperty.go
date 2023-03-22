package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigProperty := &healthCheckConfigProperty{
//   	enabled: jsii.Boolean(false),
//   	healthCheckIntervalSeconds: jsii.Number(123),
//   	healthCheckTimeoutSeconds: jsii.Number(123),
//   	healthyThresholdCount: jsii.Number(123),
//   	matcher: &matcherProperty{
//   		httpCode: jsii.String("httpCode"),
//   	},
//   	path: jsii.String("path"),
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	unhealthyThresholdCount: jsii.Number(123),
//   }
//
type CfnTargetGroup_HealthCheckConfigProperty struct {
	// `CfnTargetGroup.HealthCheckConfigProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `CfnTargetGroup.HealthCheckConfigProperty.HealthCheckIntervalSeconds`.
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// `CfnTargetGroup.HealthCheckConfigProperty.HealthCheckTimeoutSeconds`.
	HealthCheckTimeoutSeconds *float64 `field:"optional" json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// `CfnTargetGroup.HealthCheckConfigProperty.HealthyThresholdCount`.
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// `CfnTargetGroup.HealthCheckConfigProperty.Matcher`.
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// `CfnTargetGroup.HealthCheckConfigProperty.Path`.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// `CfnTargetGroup.HealthCheckConfigProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// `CfnTargetGroup.HealthCheckConfigProperty.Protocol`.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// `CfnTargetGroup.HealthCheckConfigProperty.UnhealthyThresholdCount`.
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
}

