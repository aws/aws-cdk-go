package interfacesawssagemaker


// A reference to a ComputeQuota resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeQuotaReference := &ComputeQuotaReference{
//   	ComputeQuotaArn: jsii.String("computeQuotaArn"),
//   }
//
type ComputeQuotaReference struct {
	// The ComputeQuotaArn of the ComputeQuota resource.
	ComputeQuotaArn *string `field:"required" json:"computeQuotaArn" yaml:"computeQuotaArn"`
}

