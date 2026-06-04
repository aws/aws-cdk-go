package awsbedrockagentcorealpha


// Attributes for importing an existing workload identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   workloadIdentityAttributes := &WorkloadIdentityAttributes{
//   	WorkloadIdentityArn: jsii.String("workloadIdentityArn"),
//   	WorkloadIdentityName: jsii.String("workloadIdentityName"),
//
//   	// the properties below are optional
//   	CreatedTime: jsii.String("createdTime"),
//   	LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type WorkloadIdentityAttributes struct {
	// ARN of the workload identity.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	WorkloadIdentityArn *string `field:"required" json:"workloadIdentityArn" yaml:"workloadIdentityArn"`
	// Name of the workload identity.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	WorkloadIdentityName *string `field:"required" json:"workloadIdentityName" yaml:"workloadIdentityName"`
	// Resource creation time.
	// Default: - not set.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Resource last-updated time.
	// Default: - not set.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
}

