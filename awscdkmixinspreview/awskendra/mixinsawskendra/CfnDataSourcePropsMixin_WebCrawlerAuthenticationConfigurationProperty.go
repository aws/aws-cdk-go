package mixinsawskendra


// Provides the configuration information to connect to websites that require user authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webCrawlerAuthenticationConfigurationProperty := &WebCrawlerAuthenticationConfigurationProperty{
//   	BasicAuthentication: []interface{}{
//   		&WebCrawlerBasicAuthenticationProperty{
//   			Credentials: jsii.String("credentials"),
//   			Host: jsii.String("host"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerauthenticationconfiguration.html
//
type CfnDataSourcePropsMixin_WebCrawlerAuthenticationConfigurationProperty struct {
	// The list of configuration information that's required to connect to and crawl a website host using basic authentication credentials.
	//
	// The list includes the name and port number of the website host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerauthenticationconfiguration.html#cfn-kendra-datasource-webcrawlerauthenticationconfiguration-basicauthentication
	//
	BasicAuthentication interface{} `field:"optional" json:"basicAuthentication" yaml:"basicAuthentication"`
}

