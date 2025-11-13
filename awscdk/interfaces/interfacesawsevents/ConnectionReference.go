package interfacesawsevents


// A reference to a Connection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionReference := &ConnectionReference{
//   	ConnectionArn: jsii.String("connectionArn"),
//   	ConnectionName: jsii.String("connectionName"),
//   }
//
type ConnectionReference struct {
	// The ARN of the Connection resource.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The Name of the Connection resource.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
}

