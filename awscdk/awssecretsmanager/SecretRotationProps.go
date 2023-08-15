package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Construction properties for a SecretRotation.
//
// Example:
//   var myUserSecret secret
//   var myMasterSecret secret
//   var myDatabase iConnectable
//   var myVpc vpc
//
//
//   secretsmanager.NewSecretRotation(this, jsii.String("SecretRotation"), &SecretRotationProps{
//   	Application: secretsmanager.SecretRotationApplication_MYSQL_ROTATION_MULTI_USER(),
//   	Secret: myUserSecret,
//   	 // The secret that will be rotated
//   	MasterSecret: myMasterSecret,
//   	 // The secret used for the rotation
//   	Target: myDatabase,
//   	Vpc: myVpc,
//   })
//
type SecretRotationProps struct {
	// The serverless application for the rotation.
	Application SecretRotationApplication `field:"required" json:"application" yaml:"application"`
	// The secret to rotate. It must be a JSON string with the following format:.
	//
	// ```
	// {
	//   "engine": <required: database engine>,
	//   "host": <required: instance host name>,
	//   "username": <required: username>,
	//   "password": <required: password>,
	//   "dbname": <optional: database name>,
	//   "port": <optional: if not specified, default port will be used>,
	//   "masterarn": <required for multi user rotation: the arn of the master secret which will be used to create users/change passwords>
	// }
	// ```
	//
	// This is typically the case for a secret referenced from an `AWS::SecretsManager::SecretTargetAttachment`
	// or an `ISecret` returned by the `attach()` method of `Secret`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html
	//
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
	// The target service or database.
	Target awsec2.IConnectable `field:"required" json:"target" yaml:"target"`
	// The VPC where the Lambda rotation function will run.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	// Default: Duration.days(30)
	//
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// The VPC interface endpoint to use for the Secrets Manager API.
	//
	// If you enable private DNS hostnames for your VPC private endpoint (the default), you don't
	// need to specify an endpoint. The standard Secrets Manager DNS hostname the Secrets Manager
	// CLI and SDKs use by default (https://secretsmanager.<region>.amazonaws.com) automatically
	// resolves to your VPC endpoint.
	// Default: https://secretsmanager.<region>.amazonaws.com
	//
	Endpoint awsec2.IInterfaceVpcEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Characters which should not appear in the generated password.
	// Default: - no additional characters are explicitly excluded.
	//
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret for a multi user rotation scheme.
	// Default: - single user rotation scheme.
	//
	MasterSecret ISecret `field:"optional" json:"masterSecret" yaml:"masterSecret"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	// Default: true.
	//
	RotateImmediatelyOnUpdate *bool `field:"optional" json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// The security group for the Lambda rotation function.
	// Default: - a new security group is created.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The type of subnets in the VPC where the Lambda rotation function will run.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

