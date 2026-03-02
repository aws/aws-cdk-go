package interfacesawsdevopsagent


// A reference to a Service resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceReference := &ServiceReference{
//   	ServiceId: jsii.String("serviceId"),
//   }
//
type ServiceReference struct {
	// The ServiceId of the Service resource.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

