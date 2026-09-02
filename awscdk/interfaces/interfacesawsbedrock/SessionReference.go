package interfacesawsbedrock


// A reference to a Session resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionReference := &SessionReference{
//   	SessionArn: jsii.String("sessionArn"),
//   }
//
type SessionReference struct {
	// The SessionArn of the Session resource.
	SessionArn *string `field:"required" json:"sessionArn" yaml:"sessionArn"`
}

