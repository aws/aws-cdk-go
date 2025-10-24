package awsbedrockagentcorealpha


// Attributes for importing an existing Runtime Endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   runtimeEndpointAttributes := &RuntimeEndpointAttributes{
//   	AgentRuntimeArn: jsii.String("agentRuntimeArn"),
//   	AgentRuntimeEndpointArn: jsii.String("agentRuntimeEndpointArn"),
//   	EndpointName: jsii.String("endpointName"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	Description: jsii.String("description"),
//   	EndpointId: jsii.String("endpointId"),
//   	LastUpdatedAt: jsii.String("lastUpdatedAt"),
//   	LiveVersion: jsii.String("liveVersion"),
//   	Status: jsii.String("status"),
//   	TargetVersion: jsii.String("targetVersion"),
//   }
//
// Experimental.
type RuntimeEndpointAttributes struct {
	// The ARN of the parent agent runtime.
	// Experimental.
	AgentRuntimeArn *string `field:"required" json:"agentRuntimeArn" yaml:"agentRuntimeArn"`
	// The ARN of the runtime endpoint.
	// Experimental.
	AgentRuntimeEndpointArn *string `field:"required" json:"agentRuntimeEndpointArn" yaml:"agentRuntimeEndpointArn"`
	// The name of the runtime endpoint.
	// Experimental.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// When the endpoint was created.
	// Default: - Creation time not available.
	//
	// Experimental.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The description of the runtime endpoint.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the runtime endpoint.
	// Default: - Endpoint ID not available.
	//
	// Experimental.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// When the endpoint was last updated.
	// Default: - Last update time not available.
	//
	// Experimental.
	LastUpdatedAt *string `field:"optional" json:"lastUpdatedAt" yaml:"lastUpdatedAt"`
	// The live version of the agent runtime that is currently serving requests.
	// Default: - Live version not available.
	//
	// Experimental.
	LiveVersion *string `field:"optional" json:"liveVersion" yaml:"liveVersion"`
	// The current status of the runtime endpoint.
	// Default: - Status not available.
	//
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The target version the endpoint is transitioning to (during updates).
	// Default: - Target version not available.
	//
	// Experimental.
	TargetVersion *string `field:"optional" json:"targetVersion" yaml:"targetVersion"`
}

