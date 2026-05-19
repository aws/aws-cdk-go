package awsbedrockagentcore


// Attributes for importing an existing workload identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type WorkloadIdentityAttributes struct {
	// ARN of the workload identity.
	WorkloadIdentityArn *string `field:"required" json:"workloadIdentityArn" yaml:"workloadIdentityArn"`
	// Name of the workload identity.
	WorkloadIdentityName *string `field:"required" json:"workloadIdentityName" yaml:"workloadIdentityName"`
	// Resource creation time.
	// Default: - not set.
	//
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Resource last-updated time.
	// Default: - not set.
	//
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
}

