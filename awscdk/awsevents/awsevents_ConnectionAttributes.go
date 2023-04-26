package awsevents


// Interface with properties necessary to import a reusable Connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionAttributes := &ConnectionAttributes{
//   	ConnectionArn: jsii.String("connectionArn"),
//   	ConnectionName: jsii.String("connectionName"),
//   	ConnectionSecretArn: jsii.String("connectionSecretArn"),
//   }
//
// Experimental.
type ConnectionAttributes struct {
	// The ARN of the connection created.
	// Experimental.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The Name for the connection.
	// Experimental.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The ARN for the secret created for the connection.
	// Experimental.
	ConnectionSecretArn *string `field:"required" json:"connectionSecretArn" yaml:"connectionSecretArn"`
}

