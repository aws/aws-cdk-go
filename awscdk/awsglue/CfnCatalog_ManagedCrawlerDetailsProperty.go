package awsglue


// The state of the crawler Glue manages on your behalf for a synchronization mode.
//
// Returned by the service; Glue rejects it as input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedCrawlerDetailsProperty := &ManagedCrawlerDetailsProperty{
//   	CrawlerArn: jsii.String("crawlerArn"),
//   	Error: &ErrorDetailProperty{
//   		ErrorCode: jsii.String("errorCode"),
//   		ErrorMessage: jsii.String("errorMessage"),
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-managedcrawlerdetails.html
//
type CfnCatalog_ManagedCrawlerDetailsProperty struct {
	// The Amazon Resource Name (ARN) of the crawler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-managedcrawlerdetails.html#cfn-glue-catalog-managedcrawlerdetails-crawlerarn
	//
	CrawlerArn *string `field:"optional" json:"crawlerArn" yaml:"crawlerArn"`
	// Details about an error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-managedcrawlerdetails.html#cfn-glue-catalog-managedcrawlerdetails-error
	//
	Error interface{} `field:"optional" json:"error" yaml:"error"`
	// The provisioning state of the crawler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-managedcrawlerdetails.html#cfn-glue-catalog-managedcrawlerdetails-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

