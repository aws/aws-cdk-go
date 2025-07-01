package awscloudwatch


// Props of the spacer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spacerProps := &SpacerProps{
//   	Height: jsii.Number(123),
//   	Width: jsii.Number(123),
//   }
//
type SpacerProps struct {
	// Height of the spacer.
	// Default: : 1.
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Width of the spacer.
	// Default: 1.
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

