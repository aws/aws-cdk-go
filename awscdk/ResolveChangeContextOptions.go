package awscdk


// Options that can be changed while doing a recursive resolve.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   resolveChangeContextOptions := &ResolveChangeContextOptions{
//   	AllowIntrinsicKeys: jsii.Boolean(false),
//   	RemoveEmpty: jsii.Boolean(false),
//   }
//
type ResolveChangeContextOptions struct {
	// Change the 'allowIntrinsicKeys' option.
	// Default: - Unchanged.
	//
	AllowIntrinsicKeys *bool `field:"optional" json:"allowIntrinsicKeys" yaml:"allowIntrinsicKeys"`
	// Whether to remove undefined elements from arrays and objects when resolving.
	// Default: - Unchanged.
	//
	RemoveEmpty *bool `field:"optional" json:"removeEmpty" yaml:"removeEmpty"`
}

