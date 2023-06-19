package awslambda


// A version/weight pair for routing traffic to Lambda functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var version version
//
//   versionWeight := &VersionWeight{
//   	Version: version,
//   	Weight: jsii.Number(123),
//   }
//
// Experimental.
type VersionWeight struct {
	// The version to route traffic to.
	// Experimental.
	Version IVersion `field:"required" json:"version" yaml:"version"`
	// How much weight to assign to this version (0..1).
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

