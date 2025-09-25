package awsgreengrass


// A reference to a GroupVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupVersionReference := &GroupVersionReference{
//   	GroupVersionId: jsii.String("groupVersionId"),
//   }
//
type GroupVersionReference struct {
	// The Id of the GroupVersion resource.
	GroupVersionId *string `field:"required" json:"groupVersionId" yaml:"groupVersionId"`
}

