package previewawsgluemixins


// Specifies AWS Lake Formation configuration settings for the crawler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lakeFormationConfigurationProperty := &LakeFormationConfigurationProperty{
//   	AccountId: jsii.String("accountId"),
//   	UseLakeFormationCredentials: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-lakeformationconfiguration.html
//
type CfnCrawlerPropsMixin_LakeFormationConfigurationProperty struct {
	// Required for cross account crawls.
	//
	// For same account crawls as the target data, this can be left as null.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-lakeformationconfiguration.html#cfn-glue-crawler-lakeformationconfiguration-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Specifies whether to use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-lakeformationconfiguration.html#cfn-glue-crawler-lakeformationconfiguration-uselakeformationcredentials
	//
	UseLakeFormationCredentials interface{} `field:"optional" json:"useLakeFormationCredentials" yaml:"useLakeFormationCredentials"`
}

