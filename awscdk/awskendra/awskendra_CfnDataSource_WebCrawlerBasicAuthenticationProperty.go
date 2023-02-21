package awskendra


// Provides the configuration information to connect to websites that require basic user authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerBasicAuthenticationProperty := &webCrawlerBasicAuthenticationProperty{
//   	credentials: jsii.String("credentials"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_WebCrawlerBasicAuthenticationProperty struct {
	// Your secret ARN, which you can create in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html).
	//
	// You use a secret if basic authentication credentials are required to connect to a website. The secret stores your credentials of user name and password.
	Credentials *string `field:"required" json:"credentials" yaml:"credentials"`
	// The name of the website host you want to connect to using authentication credentials.
	//
	// For example, the host name of https://a.example.com/page1.html is "a.example.com".
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port number of the website host you want to connect to using authentication credentials.
	//
	// For example, the port for https://a.example.com/page1.html is 443, the standard port for HTTPS.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

