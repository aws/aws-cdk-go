package awscdk


// Options when Applying an Aspect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   aspectOptions := &AspectOptions{
//   	Priority: jsii.Number(123),
//   }
//
type AspectOptions struct {
	// The priority value to apply on an Aspect.
	//
	// Priority must be a non-negative integer.
	// Default: - AspectPriority.DEFAULT
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

