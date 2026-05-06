package awsrtbfabric


// Describes the configuration of an auto scaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingGroupsConfigurationProperty := &AutoScalingGroupsConfigurationProperty{
//   	AutoScalingGroupNameList: []*string{
//   		jsii.String("autoScalingGroupNameList"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	HealthCheckConfig: &HealthCheckConfigProperty{
//   		Path: jsii.String("path"),
//   		Port: jsii.Number(123),
//
//   		// the properties below are optional
//   		HealthyThresholdCount: jsii.Number(123),
//   		IntervalSeconds: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		StatusCodeMatcher: jsii.String("statusCodeMatcher"),
//   		TimeoutMs: jsii.Number(123),
//   		UnhealthyThresholdCount: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.html
//
type CfnResponderGateway_AutoScalingGroupsConfigurationProperty struct {
	// The names of the auto scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.html#cfn-rtbfabric-respondergateway-autoscalinggroupsconfiguration-autoscalinggroupnamelist
	//
	AutoScalingGroupNameList *[]*string `field:"required" json:"autoScalingGroupNameList" yaml:"autoScalingGroupNameList"`
	// The role ARN of the auto scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.html#cfn-rtbfabric-respondergateway-autoscalinggroupsconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.html#cfn-rtbfabric-respondergateway-autoscalinggroupsconfiguration-healthcheckconfig
	//
	HealthCheckConfig interface{} `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
}

