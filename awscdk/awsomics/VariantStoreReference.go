package awsomics


// A reference to a VariantStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variantStoreReference := &VariantStoreReference{
//   	VariantStoreName: jsii.String("variantStoreName"),
//   }
//
type VariantStoreReference struct {
	// The Name of the VariantStore resource.
	VariantStoreName *string `field:"required" json:"variantStoreName" yaml:"variantStoreName"`
}

