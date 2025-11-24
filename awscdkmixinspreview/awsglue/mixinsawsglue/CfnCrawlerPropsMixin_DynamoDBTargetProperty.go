package mixinsawsglue


// Specifies an Amazon DynamoDB table to crawl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynamoDBTargetProperty := &DynamoDBTargetProperty{
//   	Path: jsii.String("path"),
//   	ScanAll: jsii.Boolean(false),
//   	ScanRate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-dynamodbtarget.html
//
type CfnCrawlerPropsMixin_DynamoDBTargetProperty struct {
	// The name of the DynamoDB table to crawl.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-dynamodbtarget.html#cfn-glue-crawler-dynamodbtarget-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Indicates whether to scan all the records, or to sample rows from the table.
	//
	// Scanning all the records can take a long time when the table is not a high throughput table.
	//
	// A value of `true` means to scan all records, while a value of `false` means to sample the records. If no value is specified, the value defaults to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-dynamodbtarget.html#cfn-glue-crawler-dynamodbtarget-scanall
	//
	ScanAll interface{} `field:"optional" json:"scanAll" yaml:"scanAll"`
	// The percentage of the configured read capacity units to use by the AWS Glue crawler.
	//
	// Read capacity units is a term defined by DynamoDB, and is a numeric value that acts as rate limiter for the number of reads that can be performed on that table per second.
	//
	// The valid values are null or a value between 0.1 to 1.5. A null value is used when user does not provide a value, and defaults to 0.5 of the configured Read Capacity Unit (for provisioned tables), or 0.25 of the max configured Read Capacity Unit (for tables using on-demand mode).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-dynamodbtarget.html#cfn-glue-crawler-dynamodbtarget-scanrate
	//
	ScanRate *float64 `field:"optional" json:"scanRate" yaml:"scanRate"`
}

