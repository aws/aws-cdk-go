package awseks


// Taint interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taintSpec := &TaintSpec{
//   	Effect: awscdk.Aws_eks.TaintEffect_NO_SCHEDULE,
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
type TaintSpec struct {
	// Effect type.
	// Default: - None.
	//
	Effect TaintEffect `field:"optional" json:"effect" yaml:"effect"`
	// Taint key.
	// Default: - None.
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Taint value.
	// Default: - None.
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

