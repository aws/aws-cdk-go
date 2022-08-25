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
//   weightedTarget := &weightedTarget{
//   	virtualNode: virtualNode,
//
//   	// the properties below are optional
//   	weight: jsii.Number(123),
//   }
//
type WeightedTarget struct {
	// The VirtualNode the route points to.
	VirtualNode IVirtualNode `field:"required" json:"virtualNode" yaml:"virtualNode"`
	// The weight for the target.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

