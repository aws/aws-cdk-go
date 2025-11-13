package interfacesawsamplifyuibuilder


// A reference to a Component resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentReference := &ComponentReference{
//   	AppId: jsii.String("appId"),
//   	ComponentId: jsii.String("componentId"),
//   	EnvironmentName: jsii.String("environmentName"),
//   }
//
type ComponentReference struct {
	// The AppId of the Component resource.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The Id of the Component resource.
	ComponentId *string `field:"required" json:"componentId" yaml:"componentId"`
	// The EnvironmentName of the Component resource.
	EnvironmentName *string `field:"required" json:"environmentName" yaml:"environmentName"`
}

