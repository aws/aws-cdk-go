package awsglue


// Specifies a data store in Amazon Simple Storage Service (Amazon S3).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3TargetProperty := &s3TargetProperty{
//   	connectionName: jsii.String("connectionName"),
//   	dlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   	eventQueueArn: jsii.String("eventQueueArn"),
//   	exclusions: []*string{
//   		jsii.String("exclusions"),
//   	},
//   	path: jsii.String("path"),
//   	sampleSize: jsii.Number(123),
//   }
//
type CfnCrawler_S3TargetProperty struct {
	// The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC).
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// `CfnCrawler.S3TargetProperty.DlqEventQueueArn`.
	DlqEventQueueArn *string `field:"optional" json:"dlqEventQueueArn" yaml:"dlqEventQueueArn"`
	// `CfnCrawler.S3TargetProperty.EventQueueArn`.
	EventQueueArn *string `field:"optional" json:"eventQueueArn" yaml:"eventQueueArn"`
	// A list of glob patterns used to exclude from the crawl.
	//
	// For more information, see [Catalog Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) .
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The path to the Amazon S3 target.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset.
	//
	// If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
}

