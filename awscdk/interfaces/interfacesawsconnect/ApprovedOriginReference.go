package interfacesawsconnect


// A reference to a ApprovedOrigin resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   approvedOriginReference := &ApprovedOriginReference{
//   	InstanceId: jsii.String("instanceId"),
//   	Origin: jsii.String("origin"),
//   }
//
type ApprovedOriginReference struct {
	// The InstanceId of the ApprovedOrigin resource.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The Origin of the ApprovedOrigin resource.
	Origin *string `field:"required" json:"origin" yaml:"origin"`
}

