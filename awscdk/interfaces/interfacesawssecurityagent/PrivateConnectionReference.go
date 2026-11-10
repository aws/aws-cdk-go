package interfacesawssecurityagent


// A reference to a PrivateConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateConnectionReference := &PrivateConnectionReference{
//   	PrivateConnectionArn: jsii.String("privateConnectionArn"),
//   }
//
type PrivateConnectionReference struct {
	// The Arn of the PrivateConnection resource.
	PrivateConnectionArn *string `field:"required" json:"privateConnectionArn" yaml:"privateConnectionArn"`
}

