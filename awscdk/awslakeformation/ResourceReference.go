package awslakeformation


// A reference to a Resource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceReference := &ResourceReference{
//   	ResourceId: jsii.String("resourceId"),
//   }
//
type ResourceReference struct {
	// The Id of the Resource resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

