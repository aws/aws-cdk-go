package mixinsawsbedrock


// The configuration of web URLs that you want to crawl.
//
// You should be authorized to crawl the URLs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   urlConfigurationProperty := &UrlConfigurationProperty{
//   	SeedUrls: []interface{}{
//   		&SeedUrlProperty{
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-urlconfiguration.html
//
type CfnDataSourcePropsMixin_UrlConfigurationProperty struct {
	// One or more seed or starting point URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-urlconfiguration.html#cfn-bedrock-datasource-urlconfiguration-seedurls
	//
	SeedUrls interface{} `field:"optional" json:"seedUrls" yaml:"seedUrls"`
}

