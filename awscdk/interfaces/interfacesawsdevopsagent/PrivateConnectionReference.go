package interfacesawsdevopsagent


// A reference to a PrivateConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateConnectionReference := &PrivateConnectionReference{
//   	PrivateConnectionArn: jsii.String("privateConnectionArn"),
//   	PrivateConnectionName: jsii.String("privateConnectionName"),
//   }
//
type PrivateConnectionReference struct {
	// The ARN of the PrivateConnection resource.
	PrivateConnectionArn *string `field:"required" json:"privateConnectionArn" yaml:"privateConnectionArn"`
	// The Name of the PrivateConnection resource.
	PrivateConnectionName *string `field:"required" json:"privateConnectionName" yaml:"privateConnectionName"`
}

