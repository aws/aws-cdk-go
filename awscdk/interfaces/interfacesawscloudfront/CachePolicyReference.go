package interfacesawscloudfront


// A reference to a CachePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cachePolicyReference := &CachePolicyReference{
//   	CachePolicyId: jsii.String("cachePolicyId"),
//   }
//
type CachePolicyReference struct {
	// The Id of the CachePolicy resource.
	CachePolicyId *string `field:"required" json:"cachePolicyId" yaml:"cachePolicyId"`
}

