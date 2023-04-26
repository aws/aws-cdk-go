package awsappmesh


// Properties for the Weighted Targets in the route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var virtualNode virtualNode
//
//   weightedTarget := &WeightedTarget{
//   	VirtualNode: virtualNode,
//
//   	// the properties below are optional
//   	Weight: jsii.Number(123),
//   }
//
// Experimental.
type WeightedTarget struct {
	// The VirtualNode the route points to.
	// Experimental.
	VirtualNode IVirtualNode `field:"required" json:"virtualNode" yaml:"virtualNode"`
	// The weight for the target.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

