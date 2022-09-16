package awsappmesh


// An object that represents a target and its relative weight.
//
// Traffic is distributed across targets according to their relative weight. For example, a weighted target with a relative weight of 50 receives five times as much traffic as one with a relative weight of 10. The total weight for all targets combined must be less than or equal to 100.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   weightedTargetProperty := &weightedTargetProperty{
//   	virtualNode: jsii.String("virtualNode"),
//   	weight: jsii.Number(123),
//
//   	// the properties below are optional
//   	port: jsii.Number(123),
//   }
//
type CfnRoute_WeightedTargetProperty struct {
	// The virtual node to associate with the weighted target.
	VirtualNode *string `field:"required" json:"virtualNode" yaml:"virtualNode"`
	// The relative weight of the weighted target.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// `CfnRoute.WeightedTargetProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

