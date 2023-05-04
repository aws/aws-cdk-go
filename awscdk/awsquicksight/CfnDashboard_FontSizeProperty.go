package awsquicksight


// The option that determines the text display size.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fontSizeProperty := &FontSizeProperty{
//   	Relative: jsii.String("relative"),
//   }
//
type CfnDashboard_FontSizeProperty struct {
	// The lexical name for the text size, proportional to its surrounding context.
	Relative *string `field:"optional" json:"relative" yaml:"relative"`
}

