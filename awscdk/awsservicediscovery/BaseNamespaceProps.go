package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseNamespaceProps := &BaseNamespaceProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type BaseNamespaceProps struct {
	// A name for the Namespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the Namespace.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

