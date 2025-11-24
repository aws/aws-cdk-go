package mixinsawskendra


// Provides the configuration information to connect to websites that require basic user authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webCrawlerBasicAuthenticationProperty := &WebCrawlerBasicAuthenticationProperty{
//   	Credentials: jsii.String("credentials"),
//   	Host: jsii.String("host"),
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerbasicauthentication.html
//
type CfnDataSourcePropsMixin_WebCrawlerBasicAuthenticationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret.
	//
	// You create a secret to store your credentials in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html)
	//
	// You use a secret if basic authentication credentials are required to connect to a website. The secret stores your credentials of user name and password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerbasicauthentication.html#cfn-kendra-datasource-webcrawlerbasicauthentication-credentials
	//
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// The name of the website host you want to connect to using authentication credentials.
	//
	// For example, the host name of https://a.example.com/page1.html is "a.example.com".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerbasicauthentication.html#cfn-kendra-datasource-webcrawlerbasicauthentication-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The port number of the website host you want to connect to using authentication credentials.
	//
	// For example, the port for https://a.example.com/page1.html is 443, the standard port for HTTPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerbasicauthentication.html#cfn-kendra-datasource-webcrawlerbasicauthentication-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

