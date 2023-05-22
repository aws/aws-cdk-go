// An experiment to bundle the entire CDK into a single module
package awscdk


// A single dependency.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//
//   dependency := &Dependency{
//   	Source: construct,
//   	Target: construct,
//   }
//
// Experimental.
type Dependency struct {
	// Source the dependency.
	// Experimental.
	Source IConstruct `field:"required" json:"source" yaml:"source"`
	// Target of the dependency.
	// Experimental.
	Target IConstruct `field:"required" json:"target" yaml:"target"`
}

