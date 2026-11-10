package interfacesawsappfabric


// A reference to a AppBundle resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appBundleReference := &AppBundleReference{
//   	AppBundleArn: jsii.String("appBundleArn"),
//   }
//
type AppBundleReference struct {
	// The Arn of the AppBundle resource.
	AppBundleArn *string `field:"required" json:"appBundleArn" yaml:"appBundleArn"`
}

