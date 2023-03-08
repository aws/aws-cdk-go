package awsdms


// Provides information that defines a Redis target endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Specifying endpoint settings for Redis as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redis.html#CHAP_Target.Redis.EndpointSettings) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redisSettingsProperty := &redisSettingsProperty{
//   	authPassword: jsii.String("authPassword"),
//   	authType: jsii.String("authType"),
//   	authUserName: jsii.String("authUserName"),
//   	port: jsii.Number(123),
//   	serverName: jsii.String("serverName"),
//   	sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   	sslSecurityProtocol: jsii.String("sslSecurityProtocol"),
//   }
//
type CfnEndpoint_RedisSettingsProperty struct {
	// The password provided with the `auth-role` and `auth-token` options of the `AuthType` setting for a Redis target endpoint.
	AuthPassword *string `field:"optional" json:"authPassword" yaml:"authPassword"`
	// The type of authentication to perform when connecting to a Redis target.
	//
	// Options include `none` , `auth-token` , and `auth-role` . The `auth-token` option requires an `AuthPassword` value to be provided. The `auth-role` option requires `AuthUserName` and `AuthPassword` values to be provided.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The user name provided with the `auth-role` option of the `AuthType` setting for a Redis target endpoint.
	AuthUserName *string `field:"optional" json:"authUserName" yaml:"authUserName"`
	// Transmission Control Protocol (TCP) port for the endpoint.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Fully qualified domain name of the endpoint.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The Amazon Resource Name (ARN) for the certificate authority (CA) that DMS uses to connect to your Redis target endpoint.
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// The connection to a Redis target endpoint using Transport Layer Security (TLS).
	//
	// Valid values include `plaintext` and `ssl-encryption` . The default is `ssl-encryption` . The `ssl-encryption` option makes an encrypted connection. Optionally, you can identify an Amazon Resource Name (ARN) for an SSL certificate authority (CA) using the `SslCaCertificateArn` setting. If an ARN isn't given for a CA, DMS uses the Amazon root CA.
	//
	// The `plaintext` option doesn't provide Transport Layer Security (TLS) encryption for traffic between endpoint and database.
	SslSecurityProtocol *string `field:"optional" json:"sslSecurityProtocol" yaml:"sslSecurityProtocol"`
}

