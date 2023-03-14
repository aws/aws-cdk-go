package awsglue


// Specifies an Amazon DynamoDB table to crawl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBTargetProperty := &dynamoDBTargetProperty{
//   	path: jsii.String("path"),
//   }
//
type CfnCrawler_DynamoDBTargetProperty struct {
	// The name of the DynamoDB table to crawl.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

