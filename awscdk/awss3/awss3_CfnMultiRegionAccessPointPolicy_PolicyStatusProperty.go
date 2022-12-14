package awss3


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyStatusProperty := &policyStatusProperty{
//   	isPublic: jsii.String("isPublic"),
//   }
//
type CfnMultiRegionAccessPointPolicy_PolicyStatusProperty struct {
	// `CfnMultiRegionAccessPointPolicy.PolicyStatusProperty.IsPublic`.
	IsPublic *string `field:"required" json:"isPublic" yaml:"isPublic"`
}

