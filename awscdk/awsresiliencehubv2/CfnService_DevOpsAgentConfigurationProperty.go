package awsresiliencehubv2


// Links the service to an AWS DevOps Agent Agent Space.
//
// Provide AgentSpaceArn to connect the service; omit this property to leave an existing link unchanged. The DevOps Agent integration is governed by the service's resilience policy and operates in analysis-only mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   devOpsAgentConfigurationProperty := &DevOpsAgentConfigurationProperty{
//   	AgentSpaceArn: jsii.String("agentSpaceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-devopsagentconfiguration.html
//
type CfnService_DevOpsAgentConfigurationProperty struct {
	// The ARN of the DevOps Agent Agent Space to link to this service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-devopsagentconfiguration.html#cfn-resiliencehubv2-service-devopsagentconfiguration-agentspacearn
	//
	AgentSpaceArn *string `field:"required" json:"agentSpaceArn" yaml:"agentSpaceArn"`
}

