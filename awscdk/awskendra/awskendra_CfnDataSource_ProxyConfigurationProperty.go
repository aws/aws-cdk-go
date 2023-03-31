package awskendra


// Provides the configuration information for a web proxy to connect to website hosts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proxyConfigurationProperty := &proxyConfigurationProperty{
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//
//   	// the properties below are optional
//   	credentials: jsii.String("credentials"),
//   }
//
type CfnDataSource_ProxyConfigurationProperty struct {
	// The name of the website host you want to connect to via a web proxy server.
	//
	// For example, the host name of https://a.example.com/page1.html is "a.example.com".
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port number of the website host you want to connect to via a web proxy server.
	//
	// For example, the port for https://a.example.com/page1.html is 443, the standard port for HTTPS.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Your secret ARN, which you can create in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html).
	//
	// The credentials are optional. You use a secret if web proxy credentials are required to connect to a website host. Amazon Kendra currently support basic authentication to connect to a web proxy server. The secret stores your credentials.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
}

