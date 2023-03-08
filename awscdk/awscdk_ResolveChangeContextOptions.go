// An experiment to bundle the entire CDK into a single module
package awscdk


// Options that can be changed while doing a recursive resolve.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   resolveChangeContextOptions := &resolveChangeContextOptions{
//   	allowIntrinsicKeys: jsii.Boolean(false),
//   }
//
// Experimental.
type ResolveChangeContextOptions struct {
	// Change the 'allowIntrinsicKeys' option.
	// Experimental.
	AllowIntrinsicKeys *bool `field:"optional" json:"allowIntrinsicKeys" yaml:"allowIntrinsicKeys"`
}

