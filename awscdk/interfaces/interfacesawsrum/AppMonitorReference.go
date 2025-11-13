package interfacesawsrum


// A reference to a AppMonitor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMonitorReference := &AppMonitorReference{
//   	AppMonitorName: jsii.String("appMonitorName"),
//   }
//
type AppMonitorReference struct {
	// The Name of the AppMonitor resource.
	AppMonitorName *string `field:"required" json:"appMonitorName" yaml:"appMonitorName"`
}

