package awscloudwatch


// Props of the spacer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spacerProps := &spacerProps{
//   	height: jsii.Number(123),
//   	width: jsii.Number(123),
//   }
//
type SpacerProps struct {
	// Height of the spacer.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Width of the spacer.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

