package interfacesawsopsworks


// A reference to a App resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appReference := &AppReference{
//   	AppId: jsii.String("appId"),
//   }
//
type AppReference struct {
	// The Id of the App resource.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
}

