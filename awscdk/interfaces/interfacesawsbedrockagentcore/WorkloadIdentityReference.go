package interfacesawsbedrockagentcore


// A reference to a WorkloadIdentity resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workloadIdentityReference := &WorkloadIdentityReference{
//   	WorkloadIdentityArn: jsii.String("workloadIdentityArn"),
//   	WorkloadIdentityName: jsii.String("workloadIdentityName"),
//   }
//
type WorkloadIdentityReference struct {
	// The ARN of the WorkloadIdentity resource.
	WorkloadIdentityArn *string `field:"required" json:"workloadIdentityArn" yaml:"workloadIdentityArn"`
	// The Name of the WorkloadIdentity resource.
	WorkloadIdentityName *string `field:"required" json:"workloadIdentityName" yaml:"workloadIdentityName"`
}

