package awsinspector


// A reference to a ResourceGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceGroupReference := &ResourceGroupReference{
//   	ResourceGroupArn: jsii.String("resourceGroupArn"),
//   }
//
type ResourceGroupReference struct {
	// The Arn of the ResourceGroup resource.
	ResourceGroupArn *string `field:"required" json:"resourceGroupArn" yaml:"resourceGroupArn"`
}

