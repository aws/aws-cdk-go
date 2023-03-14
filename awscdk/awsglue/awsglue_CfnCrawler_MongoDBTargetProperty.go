package awsglue


// Specifies an Amazon DocumentDB or MongoDB data store to crawl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mongoDBTargetProperty := &mongoDBTargetProperty{
//   	connectionName: jsii.String("connectionName"),
//   	path: jsii.String("path"),
//   }
//
type CfnCrawler_MongoDBTargetProperty struct {
	// The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The path of the Amazon DocumentDB or MongoDB target (database/collection).
	Path *string `field:"optional" json:"path" yaml:"path"`
}

