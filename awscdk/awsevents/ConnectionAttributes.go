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
type ConnectionAttributes struct {
	// The ARN of the connection created.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The Name for the connection.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The ARN for the secret created for the connection.
	ConnectionSecretArn *string `field:"required" json:"connectionSecretArn" yaml:"connectionSecretArn"`
}

