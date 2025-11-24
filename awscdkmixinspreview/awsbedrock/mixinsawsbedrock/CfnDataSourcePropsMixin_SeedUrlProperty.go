package mixinsawsbedrock


// The seed or starting point URL.
//
// You should be authorized to crawl the URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   seedUrlProperty := &SeedUrlProperty{
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-seedurl.html
//
type CfnDataSourcePropsMixin_SeedUrlProperty struct {
	// A seed or starting point URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-seedurl.html#cfn-bedrock-datasource-seedurl-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

