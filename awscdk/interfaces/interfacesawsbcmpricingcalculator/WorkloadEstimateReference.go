package interfacesawsbcmpricingcalculator


// A reference to a WorkloadEstimate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workloadEstimateReference := &WorkloadEstimateReference{
//   	WorkloadEstimateArn: jsii.String("workloadEstimateArn"),
//   }
//
type WorkloadEstimateReference struct {
	// The Arn of the WorkloadEstimate resource.
	WorkloadEstimateArn *string `field:"required" json:"workloadEstimateArn" yaml:"workloadEstimateArn"`
}

