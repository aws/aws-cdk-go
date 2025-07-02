package awsglue


// Specifies an Amazon DynamoDB table to crawl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBTargetProperty := &DynamoDBTargetProperty{
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-dynamodbtarget.html
//
type CfnCrawler_DynamoDBTargetProperty struct {
	// The name of the DynamoDB table to crawl.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-dynamodbtarget.html#cfn-glue-crawler-dynamodbtarget-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

