package interfacesawschime


// A reference to a AppInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appInstanceReference := &AppInstanceReference{
//   	AppInstanceArn: jsii.String("appInstanceArn"),
//   }
//
type AppInstanceReference struct {
	// The AppInstanceArn of the AppInstance resource.
	AppInstanceArn *string `field:"required" json:"appInstanceArn" yaml:"appInstanceArn"`
}

