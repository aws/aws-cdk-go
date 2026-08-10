package interfacesawssso


// A reference to a ApplicationProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationProviderReference := &ApplicationProviderReference{
//   	ApplicationProviderArn: jsii.String("applicationProviderArn"),
//   }
//
type ApplicationProviderReference struct {
	// The ApplicationProviderArn of the ApplicationProvider resource.
	ApplicationProviderArn *string `field:"required" json:"applicationProviderArn" yaml:"applicationProviderArn"`
}

