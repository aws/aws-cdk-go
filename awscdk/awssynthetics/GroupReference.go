package awssynthetics


// A reference to a Group resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupReference := &GroupReference{
//   	GroupName: jsii.String("groupName"),
//   }
//
type GroupReference struct {
	// The Name of the Group resource.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

