package awsrds


// Properties for defining a `CfnDBProxy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBProxyProps := &cfnDBProxyProps{
//   	auth: []interface{}{
//   		&authFormatProperty{
//   			authScheme: jsii.String("authScheme"),
//   			clientPasswordAuthType: jsii.String("clientPasswordAuthType"),
//   			description: jsii.String("description"),
//   			iamAuth: jsii.String("iamAuth"),
//   			secretArn: jsii.String("secretArn"),
//   			userName: jsii.String("userName"),
//   		},
//   	},
//   	dbProxyName: jsii.String("dbProxyName"),
//   	engineFamily: jsii.String("engineFamily"),
//   	roleArn: jsii.String("roleArn"),
//   	vpcSubnetIds: []*string{
//   		jsii.String("vpcSubnetIds"),
//   	},
//
//   	// the properties below are optional
//   	debugLogging: jsii.Boolean(false),
//   	idleClientTimeout: jsii.Number(123),
//   	requireTls: jsii.Boolean(false),
//   	tags: []tagFormatProperty{
//   		&tagFormatProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnDBProxyProps struct {
	// The authorization mechanism that the proxy uses.
	Auth interface{} `field:"required" json:"auth" yaml:"auth"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region . An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The kinds of databases that the proxy can connect to.
	//
	// This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora.
	//
	// *Valid values* : `MYSQL` | `POSTGRESQL`.
	EngineFamily *string `field:"required" json:"engineFamily" yaml:"engineFamily"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging interface{} `field:"optional" json:"debugLogging" yaml:"debugLogging"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout *float64 `field:"optional" json:"idleClientTimeout" yaml:"idleClientTimeout"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls interface{} `field:"optional" json:"requireTls" yaml:"requireTls"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
	Tags *[]*CfnDBProxy_TagFormatProperty `field:"optional" json:"tags" yaml:"tags"`
	// One or more VPC security group IDs to associate with the new proxy.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

