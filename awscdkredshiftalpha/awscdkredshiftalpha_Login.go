// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Username and password combination.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc ec2.Vpc
//
//
//   awscdkredshiftalpha.NewCluster(stack, jsii.String("Redshift"), &clusterProps{
//   	masterUser: &login{
//   		masterUsername: jsii.String("admin"),
//   		masterPassword: cdk.secretValue_UnsafePlainText(jsii.String("tooshort")),
//   	},
//   	vpc: vpc,
//   	publiclyAccessible: jsii.Boolean(true),
//   	elasticIp: jsii.String("10.123.123.255"),
//   })
//
// Experimental.
type Login struct {
	// Username.
	// Experimental.
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Experimental.
	MasterPassword awscdk.SecretValue `field:"optional" json:"masterPassword" yaml:"masterPassword"`
}

