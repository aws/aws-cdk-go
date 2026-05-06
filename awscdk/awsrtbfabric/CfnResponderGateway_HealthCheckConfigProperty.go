package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigProperty := &HealthCheckConfigProperty{
//   	Path: jsii.String("path"),
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	HealthyThresholdCount: jsii.Number(123),
//   	IntervalSeconds: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	StatusCodeMatcher: jsii.String("statusCodeMatcher"),
//   	TimeoutMs: jsii.Number(123),
//   	UnhealthyThresholdCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html
//
type CfnResponderGateway_HealthCheckConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html#cfn-rtbfabric-respondergateway-healthcheckconfig-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html#cfn-rtbfabric-respondergateway-healthcheckconfig-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html#cfn-rtbfabric-respondergateway-healthcheckconfig-healthythresholdcount
	//
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html#cfn-rtbfabric-respondergateway-healthcheckconfig-intervalseconds
	//
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html#cfn-rtbfabric-respondergateway-healthcheckconfig-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html#cfn-rtbfabric-respondergateway-healthcheckconfig-statuscodematcher
	//
	StatusCodeMatcher *string `field:"optional" json:"statusCodeMatcher" yaml:"statusCodeMatcher"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html#cfn-rtbfabric-respondergateway-healthcheckconfig-timeoutms
	//
	TimeoutMs *float64 `field:"optional" json:"timeoutMs" yaml:"timeoutMs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-healthcheckconfig.html#cfn-rtbfabric-respondergateway-healthcheckconfig-unhealthythresholdcount
	//
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
}

