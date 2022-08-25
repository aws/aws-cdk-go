package awsstepfunctionstasks


// Configuration for a shuffle option for input data in a channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shuffleConfig := &shuffleConfig{
//   	seed: jsii.Number(123),
//   }
//
type ShuffleConfig struct {
	// Determines the shuffling order.
	Seed *float64 `field:"required" json:"seed" yaml:"seed"`
}

