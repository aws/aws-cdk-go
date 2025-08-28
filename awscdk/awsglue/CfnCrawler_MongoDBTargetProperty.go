package awsglue


// Specifies an Amazon DocumentDB or MongoDB data store to crawl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mongoDBTargetProperty := &MongoDBTargetProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-mongodbtarget.html
//
type CfnCrawler_MongoDBTargetProperty struct {
	// The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-mongodbtarget.html#cfn-glue-crawler-mongodbtarget-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The path of the Amazon DocumentDB or MongoDB target (database/collection).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-mongodbtarget.html#cfn-glue-crawler-mongodbtarget-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

