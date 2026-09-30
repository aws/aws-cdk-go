package interfacesawsec2


// A reference to a InstanceEventWindow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceEventWindowReference := &InstanceEventWindowReference{
//   	InstanceEventWindowArn: jsii.String("instanceEventWindowArn"),
//   }
//
type InstanceEventWindowReference struct {
	// The Arn of the InstanceEventWindow resource.
	InstanceEventWindowArn *string `field:"required" json:"instanceEventWindowArn" yaml:"instanceEventWindowArn"`
}

