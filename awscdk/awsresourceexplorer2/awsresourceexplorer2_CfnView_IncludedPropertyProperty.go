package awsresourceexplorer2


// Information about an additional property that describes a resource, that you can optionally include in a view.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   includedPropertyProperty := &includedPropertyProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnView_IncludedPropertyProperty struct {
	// The name of the property that is included in this view.
	Name *string `field:"required" json:"name" yaml:"name"`
}

