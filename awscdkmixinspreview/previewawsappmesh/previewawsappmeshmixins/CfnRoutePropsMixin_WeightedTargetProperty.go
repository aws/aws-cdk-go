package previewawsappmeshmixins


// An object that represents a target and its relative weight.
//
// Traffic is distributed across targets according to their relative weight. For example, a weighted target with a relative weight of 50 receives five times as much traffic as one with a relative weight of 10. The total weight for all targets combined must be less than or equal to 100.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   weightedTargetProperty := &WeightedTargetProperty{
//   	Port: jsii.Number(123),
//   	VirtualNode: jsii.String("virtualNode"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-weightedtarget.html
//
type CfnRoutePropsMixin_WeightedTargetProperty struct {
	// The targeted port of the weighted object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-weightedtarget.html#cfn-appmesh-route-weightedtarget-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The virtual node to associate with the weighted target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-weightedtarget.html#cfn-appmesh-route-weightedtarget-virtualnode
	//
	VirtualNode *string `field:"optional" json:"virtualNode" yaml:"virtualNode"`
	// The relative weight of the weighted target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-weightedtarget.html#cfn-appmesh-route-weightedtarget-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

