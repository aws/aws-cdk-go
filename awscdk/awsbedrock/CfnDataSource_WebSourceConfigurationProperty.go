package awsbedrock


// The configuration of the URL/URLs for the web content that you want to crawl.
//
// You should be authorized to crawl the URLs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSourceConfigurationProperty := &WebSourceConfigurationProperty{
//   	UrlConfiguration: &UrlConfigurationProperty{
//   		SeedUrls: []interface{}{
//   			&SeedUrlProperty{
//   				Url: jsii.String("url"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-websourceconfiguration.html
//
type CfnDataSource_WebSourceConfigurationProperty struct {
	// The configuration of the URL/URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-websourceconfiguration.html#cfn-bedrock-datasource-websourceconfiguration-urlconfiguration
	//
	UrlConfiguration interface{} `field:"required" json:"urlConfiguration" yaml:"urlConfiguration"`
}

