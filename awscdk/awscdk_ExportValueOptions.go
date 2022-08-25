// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for the `stack.exportValue()` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   exportValueOptions := &exportValueOptions{
//   	name: jsii.String("name"),
//   }
//
type ExportValueOptions struct {
	// The name of the export to create.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

