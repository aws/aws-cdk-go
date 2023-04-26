package awsquicksight


// The option that determines the text display weight, or boldness.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fontWeightProperty := &FontWeightProperty{
//   	Name: jsii.String("name"),
//   }
//
type CfnTemplate_FontWeightProperty struct {
	// The lexical name for the level of boldness of the text display.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

