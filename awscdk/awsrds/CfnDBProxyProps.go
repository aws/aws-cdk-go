package awsrds


// Properties for defining a `CfnDBProxy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBProxyProps := &CfnDBProxyProps{
//   	DbProxyName: jsii.String("dbProxyName"),
//   	EngineFamily: jsii.String("engineFamily"),
//   	RoleArn: jsii.String("roleArn"),
//   	VpcSubnetIds: []*string{
//   		jsii.String("vpcSubnetIds"),
//   	},
//
//   	// the properties below are optional
//   	Auth: []interface{}{
//   		&AuthFormatProperty{
//   			AuthScheme: jsii.String("authScheme"),
//   			ClientPasswordAuthType: jsii.String("clientPasswordAuthType"),
//   			Description: jsii.String("description"),
//   			IamAuth: jsii.String("iamAuth"),
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   	},
//   	DebugLogging: jsii.Boolean(false),
//   	DefaultAuthScheme: jsii.String("defaultAuthScheme"),
//   	EndpointNetworkType: jsii.String("endpointNetworkType"),
//   	IdleClientTimeout: jsii.Number(123),
//   	RequireTls: jsii.Boolean(false),
//   	Tags: []tagFormatProperty{
//   		&tagFormatProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetConnectionNetworkType: jsii.String("targetConnectionNetworkType"),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html
//
type CfnDBProxyProps struct {
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region . An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-dbproxyname
	//
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The kinds of databases that the proxy can connect to.
	//
	// This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for MySQL databases, specify `MYSQL` . For Aurora PostgreSQL and RDS for PostgreSQL databases, specify `POSTGRESQL` . For RDS for Microsoft SQL Server, specify `SQLSERVER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-enginefamily
	//
	EngineFamily *string `field:"required" json:"engineFamily" yaml:"engineFamily"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// One or more VPC subnet IDs to associate with the new proxy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsubnetids
	//
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// The authorization mechanism that the proxy uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// Specifies whether the proxy logs detailed connection and query information.
	//
	// When you enable `DebugLogging` , the proxy captures connection details and connection pool behavior from your queries. Debug logging increases CloudWatch costs and can impact proxy performance. Enable this option only when you need to troubleshoot connection or performance issues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-debuglogging
	//
	DebugLogging interface{} `field:"optional" json:"debugLogging" yaml:"debugLogging"`
	// The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-defaultauthscheme
	//
	DefaultAuthScheme *string `field:"optional" json:"defaultAuthScheme" yaml:"defaultAuthScheme"`
	// The network type of the DB proxy endpoint.
	//
	// The network type determines the IP version that the proxy endpoint supports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-endpointnetworktype
	//
	EndpointNetworkType *string `field:"optional" json:"endpointNetworkType" yaml:"endpointNetworkType"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// You can set this value higher or lower than the connection timeout limit for the associated database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-idleclienttimeout
	//
	IdleClientTimeout *float64 `field:"optional" json:"idleClientTimeout" yaml:"idleClientTimeout"`
	// Specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-requiretls
	//
	RequireTls interface{} `field:"optional" json:"requireTls" yaml:"requireTls"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-tags
	//
	Tags *[]*CfnDBProxy_TagFormatProperty `field:"optional" json:"tags" yaml:"tags"`
	// The network type that the proxy uses to connect to the target database.
	//
	// The network type determines the IP version that the proxy uses for connections to the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-targetconnectionnetworktype
	//
	TargetConnectionNetworkType *string `field:"optional" json:"targetConnectionNetworkType" yaml:"targetConnectionNetworkType"`
	// One or more VPC security group IDs to associate with the new proxy.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

