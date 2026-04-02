package awsmediapackagev2alpha


// Options for configuring a StartTag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   startTagOptions := &StartTagOptions{
//   	Precise: jsii.Boolean(false),
//   }
//
// Experimental.
type StartTagOptions struct {
	// Specify the value for PRECISE within your EXT-X-START tag.
	//
	// Leave blank, or choose false, to use the default value NO. Choose true to use the value YES.
	// Default: false.
	//
	// Experimental.
	Precise *bool `field:"optional" json:"precise" yaml:"precise"`
}

