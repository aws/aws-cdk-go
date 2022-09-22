package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Options to add the multi user rotation.
//
// Example:
//   user := awscdk.NewUser(this, jsii.String("User"), &userProps{
//   	cluster: cluster,
//   	databaseName: jsii.String("databaseName"),
//   })
//   cluster.addRotationMultiUser(jsii.String("MultiUserRotation"), &rotationMultiUserOptions{
//   	secret: user.secret,
//   })
//
// Experimental.
type RotationMultiUserOptions struct {
	// The secret to rotate.
	//
	// It must be a JSON string with the following format:
	// ```
	// {
	//    "engine": <required: database engine>,
	//    "host": <required: instance host name>,
	//    "username": <required: username>,
	//    "password": <required: password>,
	//    "dbname": <optional: database name>,
	//    "port": <optional: if not specified, default port will be used>,
	//    "masterarn": <required: the arn of the master secret which will be used to create users/change passwords>
	// }
	// ```.
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	// Experimental.
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
}

