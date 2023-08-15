package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Options to add the multi user rotation.
//
// Example:
//   import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
//
//   var myImportedSecret secret
//   var cluster databaseCluster
//
//
//   cluster.AddRotationMultiUser(jsii.String("MyUser"), &RotationMultiUserOptions{
//   	Secret: myImportedSecret,
//   })
//
type RotationMultiUserOptions struct {
	// The secret to rotate.
	//
	// It must be a JSON string with the following format:
	// ```
	// {
	//   "engine": <required: must be set to 'mongo'>,
	//   "host": <required: instance host name>,
	//   "username": <required: username>,
	//   "password": <required: password>,
	//   "dbname": <optional: database name>,
	//   "port": <optional: if not specified, default port 27017 will be used>,
	//   "masterarn": <required: the arn of the master secret which will be used to create users/change passwords>
	//   "ssl": <optional: if not specified, defaults to false. This must be true if being used for DocumentDB rotations
	//          where the cluster has TLS enabled>
	// }
	// ```.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	// Default: Duration.days(30)
	//
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
}

