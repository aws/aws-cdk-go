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
//   	Port: jsii.Number(123),
//   	Weight: jsii.Number(123),
//   }
//
type WeightedTarget struct {
	// The VirtualNode the route points to.
	VirtualNode IVirtualNode `field:"required" json:"virtualNode" yaml:"virtualNode"`
	// The port to match from the request.
	// Default: - do not match on port.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The weight for the target.
	// Default: 1.
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

