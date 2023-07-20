package awsrds


// Specifies the details of authentication used by a proxy to log in as a specific database user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authFormatProperty := &AuthFormatProperty{
//   	AuthScheme: jsii.String("authScheme"),
//   	ClientPasswordAuthType: jsii.String("clientPasswordAuthType"),
//   	Description: jsii.String("description"),
//   	IamAuth: jsii.String("iamAuth"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxy-authformat.html
//
type CfnDBProxy_AuthFormatProperty struct {
	// The type of authentication that the proxy uses for connections from the proxy to the underlying database.
	//
	// Valid Values: `SECRETS`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxy-authformat.html#cfn-rds-dbproxy-authformat-authscheme
	//
	AuthScheme *string `field:"optional" json:"authScheme" yaml:"authScheme"`
	// Specifies the details of authentication used by a proxy to log in as a specific database user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxy-authformat.html#cfn-rds-dbproxy-authformat-clientpasswordauthtype
	//
	ClientPasswordAuthType *string `field:"optional" json:"clientPasswordAuthType" yaml:"clientPasswordAuthType"`
	// A user-specified description about the authentication used by a proxy to log in as a specific database user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxy-authformat.html#cfn-rds-dbproxy-authformat-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
	//
	// The `ENABLED` value is valid only for proxies with RDS for Microsoft SQL Server.
	//
	// Valid Values: `ENABLED | DISABLED | REQUIRED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxy-authformat.html#cfn-rds-dbproxy-authformat-iamauth
	//
	IamAuth *string `field:"optional" json:"iamAuth" yaml:"iamAuth"`
	// The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxy-authformat.html#cfn-rds-dbproxy-authformat-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

