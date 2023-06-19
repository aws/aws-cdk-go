package awscdk


// Options for the `stack.exportValue()` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   exportValueOptions := &ExportValueOptions{
//   	Name: jsii.String("name"),
//   }
//
// Experimental.
type ExportValueOptions struct {
	// The name of the export to create.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

