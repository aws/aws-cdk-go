package awskendra


// Provides the configuration information for a web proxy to connect to website hosts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proxyConfigurationProperty := &ProxyConfigurationProperty{
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//
//   	// the properties below are optional
//   	Credentials: jsii.String("credentials"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-proxyconfiguration.html
//
type CfnDataSource_ProxyConfigurationProperty struct {
	// The name of the website host you want to connect to via a web proxy server.
	//
	// For example, the host name of https://a.example.com/page1.html is "a.example.com".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-proxyconfiguration.html#cfn-kendra-datasource-proxyconfiguration-host
	//
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port number of the website host you want to connect to via a web proxy server.
	//
	// For example, the port for https://a.example.com/page1.html is 443, the standard port for HTTPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-proxyconfiguration.html#cfn-kendra-datasource-proxyconfiguration-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret.
	//
	// You create a secret to store your credentials in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html)
	//
	// The credentials are optional. You use a secret if web proxy credentials are required to connect to a website host. Amazon Kendra currently support basic authentication to connect to a web proxy server. The secret stores your credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-proxyconfiguration.html#cfn-kendra-datasource-proxyconfiguration-credentials
	//
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
}

