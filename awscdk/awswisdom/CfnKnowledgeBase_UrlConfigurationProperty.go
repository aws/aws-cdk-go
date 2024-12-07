package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   urlConfigurationProperty := &UrlConfigurationProperty{
//   	SeedUrls: []interface{}{
//   		&SeedUrlProperty{
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-urlconfiguration.html
//
type CfnKnowledgeBase_UrlConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-urlconfiguration.html#cfn-wisdom-knowledgebase-urlconfiguration-seedurls
	//
	SeedUrls interface{} `field:"optional" json:"seedUrls" yaml:"seedUrls"`
}

