package awsservicecatalog


// A reference to a CloudFormationProvisionedProduct resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudFormationProvisionedProductReference := &CloudFormationProvisionedProductReference{
//   	ProvisionedProductId: jsii.String("provisionedProductId"),
//   }
//
type CloudFormationProvisionedProductReference struct {
	// The ProvisionedProductId of the CloudFormationProvisionedProduct resource.
	ProvisionedProductId *string `field:"required" json:"provisionedProductId" yaml:"provisionedProductId"`
}

