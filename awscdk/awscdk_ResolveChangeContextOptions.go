// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options that can be changed while doing a recursive resolve.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   resolveChangeContextOptions := &resolveChangeContextOptions{
//   	allowIntrinsicKeys: jsii.Boolean(false),
//   }
//
type ResolveChangeContextOptions struct {
	// Change the 'allowIntrinsicKeys' option.
	AllowIntrinsicKeys *bool `field:"optional" json:"allowIntrinsicKeys" yaml:"allowIntrinsicKeys"`
}

