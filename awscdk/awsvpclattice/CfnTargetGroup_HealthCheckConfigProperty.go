package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigProperty := &HealthCheckConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	HealthCheckIntervalSeconds: jsii.Number(123),
//   	HealthCheckTimeoutSeconds: jsii.Number(123),
//   	HealthyThresholdCount: jsii.Number(123),
//   	Matcher: &MatcherProperty{
//   		HttpCode: jsii.String("httpCode"),
//   	},
//   	Path: jsii.String("path"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	UnhealthyThresholdCount: jsii.Number(123),
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

