package awscodeguruprofiler

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProfilingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentPermissions interface{}
//
//   cfnProfilingGroupProps := &CfnProfilingGroupProps{
//   	ProfilingGroupName: jsii.String("profilingGroupName"),
//
//   	// the properties below are optional
//   	AgentPermissions: agentPermissions,
//   	AnomalyDetectionNotificationConfiguration: []interface{}{
//   		&ChannelProperty{
//   			ChannelUri: jsii.String("channelUri"),
//
//   			// the properties below are optional
//   			ChannelId: jsii.String("channelId"),
//   		},
//   	},
//   	ComputePlatform: jsii.String("computePlatform"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html
//
type CfnProfilingGroupProps struct {
	// The name of the profiling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-profilinggroupname
	//
	ProfilingGroupName *string `field:"required" json:"profilingGroupName" yaml:"profilingGroupName"`
	// The agent permissions attached to this profiling group.
	//
	// This action group grants `ConfigureAgent` and `PostAgentProfile` permissions to perform actions required by the profiling agent. The Json consists of key `Principals` .
	//
	// *Principals* : A list of string ARNs for the roles and users you want to grant access to the profiling group. Wildcards are not supported in the ARNs. You are allowed to provide up to 50 ARNs. An empty list is not permitted. This is a required key.
	//
	// For more information, see [Resource-based policies in CodeGuru Profiler](https://docs.aws.amazon.com/codeguru/latest/profiler-ug/resource-based-policies.html) in the *Amazon CodeGuru Profiler user guide* , [ConfigureAgent](https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ConfigureAgent.html) , and [PostAgentProfile](https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_PostAgentProfile.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-agentpermissions
	//
	AgentPermissions interface{} `field:"optional" json:"agentPermissions" yaml:"agentPermissions"`
	// Adds anomaly notifications for a profiling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-anomalydetectionnotificationconfiguration
	//
	AnomalyDetectionNotificationConfiguration interface{} `field:"optional" json:"anomalyDetectionNotificationConfiguration" yaml:"anomalyDetectionNotificationConfiguration"`
	// The compute platform of the profiling group.
	//
	// Use `AWSLambda` if your application runs on AWS Lambda. Use `Default` if your application runs on a compute platform that is not AWS Lambda , such an Amazon EC2 instance, an on-premises server, or a different platform. If not specified, `Default` is used. This property is immutable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-computeplatform
	//
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// A list of tags to add to the created profiling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html#cfn-codeguruprofiler-profilinggroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

