package awss3


// Properties for defining a `CfnMultiRegionAccessPointPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnMultiRegionAccessPointPolicyProps := &cfnMultiRegionAccessPointPolicyProps{
//   	mrapName: jsii.String("mrapName"),
//   	policy: policy,
//   }
//
type CfnMultiRegionAccessPointPolicyProps struct {
	// The name of the Multi-Region Access Point.
	MrapName *string `field:"required" json:"mrapName" yaml:"mrapName"`
	// The access policy associated with the Multi-Region Access Point.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
}

