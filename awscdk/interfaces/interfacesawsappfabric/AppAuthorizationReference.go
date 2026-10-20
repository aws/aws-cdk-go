package interfacesawsappfabric


// A reference to a AppAuthorization resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appAuthorizationReference := &AppAuthorizationReference{
//   	AppAuthorizationArn: jsii.String("appAuthorizationArn"),
//   }
//
type AppAuthorizationReference struct {
	// The AppAuthorizationArn of the AppAuthorization resource.
	AppAuthorizationArn *string `field:"required" json:"appAuthorizationArn" yaml:"appAuthorizationArn"`
}

