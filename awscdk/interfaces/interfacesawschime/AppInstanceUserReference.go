package interfacesawschime


// A reference to a AppInstanceUser resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appInstanceUserReference := &AppInstanceUserReference{
//   	AppInstanceUserArn: jsii.String("appInstanceUserArn"),
//   }
//
type AppInstanceUserReference struct {
	// The AppInstanceUserArn of the AppInstanceUser resource.
	AppInstanceUserArn *string `field:"required" json:"appInstanceUserArn" yaml:"appInstanceUserArn"`
}

