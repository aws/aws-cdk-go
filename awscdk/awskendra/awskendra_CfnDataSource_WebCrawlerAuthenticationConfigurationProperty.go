package awskendra


// Provides the configuration information to connect to websites that require user authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerAuthenticationConfigurationProperty := &webCrawlerAuthenticationConfigurationProperty{
//   	basicAuthentication: []interface{}{
//   		&webCrawlerBasicAuthenticationProperty{
//   			credentials: jsii.String("credentials"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnDataSource_WebCrawlerAuthenticationConfigurationProperty struct {
	// The list of configuration information that's required to connect to and crawl a website host using basic authentication credentials.
	//
	// The list includes the name and port number of the website host.
	BasicAuthentication interface{} `field:"optional" json:"basicAuthentication" yaml:"basicAuthentication"`
}

