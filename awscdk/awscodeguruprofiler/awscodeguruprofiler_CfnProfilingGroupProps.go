package awscodeguruprofiler

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnProfilingGroupProps := &cfnProfilingGroupProps{
//   	profilingGroupName: jsii.String("profilingGroupName"),
//
//   	// the properties below are optional
//   	agentPermissions: agentPermissions,
//   	anomalyDetectionNotificationConfiguration: []interface{}{
//   		&channelProperty{
//   			channelUri: jsii.String("channelUri"),
//
//   			// the properties below are optional
//   			channelId: jsii.String("channelId"),
//   		},
//   	},
//   	computePlatform: jsii.String("computePlatform"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProfilingGroupProps struct {
	// The name of the profiling group.
	ProfilingGroupName *string `field:"required" json:"profilingGroupName" yaml:"profilingGroupName"`
	// The agent permissions attached to this profiling group.
	//
	// This action group grants `ConfigureAgent` and `PostAgentProfile` permissions to perform actions required by the profiling agent. The Json consists of key `Principals` .
	//
	// *Principals* : A list of string ARNs for the roles and users you want to grant access to the profiling group. Wildcards are not supported in the ARNs. You are allowed to provide up to 50 ARNs. An empty list is not permitted. This is a required key.
	//
	// For more information, see [Resource-based policies in CodeGuru Profiler](https://docs.aws.amazon.com/codeguru/latest/profiler-ug/resource-based-policies.html) in the *Amazon CodeGuru Profiler user guide* , [ConfigureAgent](https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ConfigureAgent.html) , and [PostAgentProfile](https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_PostAgentProfile.html) .
	AgentPermissions interface{} `field:"optional" json:"agentPermissions" yaml:"agentPermissions"`
	// Adds anomaly notifications for a profiling group.
	AnomalyDetectionNotificationConfiguration interface{} `field:"optional" json:"anomalyDetectionNotificationConfiguration" yaml:"anomalyDetectionNotificationConfiguration"`
	// The compute platform of the profiling group.
	//
	// Use `AWSLambda` if your application runs on AWS Lambda. Use `Default` if your application runs on a compute platform that is not AWS Lambda , such an Amazon EC2 instance, an on-premises server, or a different platform. If not specified, `Default` is used. This property is immutable.
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// A list of tags to add to the created profiling group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

