package interfacesawscloudformation


// A reference to a CustomResource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResourceReference := &CustomResourceReference{
//   	CustomResourceId: jsii.String("customResourceId"),
//   }
//
type CustomResourceReference struct {
	// The Id of the CustomResource resource.
	CustomResourceId *string `field:"required" json:"customResourceId" yaml:"customResourceId"`
}

