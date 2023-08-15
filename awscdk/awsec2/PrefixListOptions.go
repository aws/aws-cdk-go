package awsec2


// Options to add a prefix list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixListOptions := &PrefixListOptions{
//   	MaxEntries: jsii.Number(123),
//   }
//
type PrefixListOptions struct {
	// The maximum number of entries for the prefix list.
	// Default: Automatically-calculated.
	//
	MaxEntries *float64 `field:"optional" json:"maxEntries" yaml:"maxEntries"`
}

