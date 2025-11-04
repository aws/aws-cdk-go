package awscdkeksv2alpha


// Taint interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   taintSpec := &TaintSpec{
//   	Effect: eks_v2_alpha.TaintEffect_NO_SCHEDULE,
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// Experimental.
type TaintSpec struct {
	// Effect type.
	// Default: - None.
	//
	// Experimental.
	Effect TaintEffect `field:"optional" json:"effect" yaml:"effect"`
	// Taint key.
	// Default: - None.
	//
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Taint value.
	// Default: - None.
	//
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

