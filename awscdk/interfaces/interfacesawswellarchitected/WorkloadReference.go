package interfacesawswellarchitected


// A reference to a Workload resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workloadReference := &WorkloadReference{
//   	WorkloadArn: jsii.String("workloadArn"),
//   }
//
type WorkloadReference struct {
	// The WorkloadArn of the Workload resource.
	WorkloadArn *string `field:"required" json:"workloadArn" yaml:"workloadArn"`
}

