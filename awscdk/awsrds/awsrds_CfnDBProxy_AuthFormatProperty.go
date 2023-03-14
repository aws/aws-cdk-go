package awsrds


// Specifies the details of authentication used by a proxy to log in as a specific database user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authFormatProperty := &authFormatProperty{
//   	authScheme: jsii.String("authScheme"),
//   	clientPasswordAuthType: jsii.String("clientPasswordAuthType"),
//   	description: jsii.String("description"),
//   	iamAuth: jsii.String("iamAuth"),
//   	secretArn: jsii.String("secretArn"),
//   }
//
type CfnDBProxy_AuthFormatProperty struct {
	// The type of authentication that the proxy uses for connections from the proxy to the underlying database.
	//
	// Valid Values: `SECRETS`.
	AuthScheme *string `field:"optional" json:"authScheme" yaml:"authScheme"`
	// Specifies the details of authentication used by a proxy to log in as a specific database user.
	ClientPasswordAuthType *string `field:"optional" json:"clientPasswordAuthType" yaml:"clientPasswordAuthType"`
	// A user-specified description about the authentication used by a proxy to log in as a specific database user.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
	//
	// The `ENABLED` value is valid only for proxies with RDS for Microsoft SQL Server.
	//
	// Valid Values: `ENABLED | DISABLED | REQUIRED`.
	IamAuth *string `field:"optional" json:"iamAuth" yaml:"iamAuth"`
	// The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

