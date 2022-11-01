package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Construction properties for a SecretRotation.
//
// Example:
//   var mySecret secret
//   var myDatabase iConnectable
//   var myVpc vpc
//
//
//   secretsmanager.NewSecretRotation(this, jsii.String("SecretRotation"), &secretRotationProps{
//   	application: secretsmanager.secretRotationApplication_MYSQL_ROTATION_SINGLE_USER(),
//   	 // MySQL single user scheme
//   	secret: mySecret,
//   	target: myDatabase,
//   	 // a Connectable
//   	vpc: myVpc,
//   	 // The VPC where the secret rotation application will be deployed
//   	excludeCharacters: jsii.String(" %+:;{}"),
//   })
//
// Experimental.
type SecretRotationProps struct {
	// The serverless application for the rotation.
	// Experimental.
	Application SecretRotationApplication `field:"required" json:"application" yaml:"application"`
	// The secret to rotate. It must be a JSON string with the following format:.
	//
	// ```
	// {
	//    "engine": <required: database engine>,
	//    "host": <required: instance host name>,
	//    "username": <required: username>,
	//    "password": <required: password>,
	//    "dbname": <optional: database name>,
	//    "port": <optional: if not specified, default port will be used>,
	//    "masterarn": <required for multi user rotation: the arn of the master secret which will be used to create users/change passwords>
	// }
	// ```
	//
	// This is typically the case for a secret referenced from an `AWS::SecretsManager::SecretTargetAttachment`
	// or an `ISecret` returned by the `attach()` method of `Secret`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html
	//
	// Experimental.
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
	// The target service or database.
	// Experimental.
	Target awsec2.IConnectable `field:"required" json:"target" yaml:"target"`
	// The VPC where the Lambda rotation function will run.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	// Experimental.
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// The VPC interface endpoint to use for the Secrets Manager API.
	//
	// If you enable private DNS hostnames for your VPC private endpoint (the default), you don't
	// need to specify an endpoint. The standard Secrets Manager DNS hostname the Secrets Manager
	// CLI and SDKs use by default (https://secretsmanager.<region>.amazonaws.com) automatically
	// resolves to your VPC endpoint.
	// Experimental.
	Endpoint awsec2.IInterfaceVpcEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Characters which should not appear in the generated password.
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret for a multi user rotation scheme.
	// Experimental.
	MasterSecret ISecret `field:"optional" json:"masterSecret" yaml:"masterSecret"`
	// The security group for the Lambda rotation function.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The type of subnets in the VPC where the Lambda rotation function will run.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

