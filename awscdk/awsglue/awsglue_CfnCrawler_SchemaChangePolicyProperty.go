package awsglue


// The policy that specifies update and delete behaviors for the crawler.
//
// The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The `SchemaChangePolicy` does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the `SchemaChangePolicy` on a crawler.
//
// The SchemaChangePolicy consists of two components, `UpdateBehavior` and `DeleteBehavior` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaChangePolicyProperty := &schemaChangePolicyProperty{
//   	deleteBehavior: jsii.String("deleteBehavior"),
//   	updateBehavior: jsii.String("updateBehavior"),
//   }
//
type CfnCrawler_SchemaChangePolicyProperty struct {
	// The deletion behavior when the crawler finds a deleted object.
	//
	// A value of `LOG` specifies that if a table or partition is found to no longer exist, do not delete it, only log that it was found to no longer exist.
	//
	// A value of `DELETE_FROM_DATABASE` specifies that if a table or partition is found to have been removed, delete it from the database.
	//
	// A value of `DEPRECATE_IN_DATABASE` specifies that if a table has been found to no longer exist, to add a property to the table that says "DEPRECATED" and includes a timestamp with the time of deprecation.
	DeleteBehavior *string `field:"optional" json:"deleteBehavior" yaml:"deleteBehavior"`
	// The update behavior when the crawler finds a changed schema.
	//
	// A value of `LOG` specifies that if a table or a partition already exists, and a change is detected, do not update it, only log that a change was detected. Add new tables and new partitions (including on existing tables).
	//
	// A value of `UPDATE_IN_DATABASE` specifies that if a table or partition already exists, and a change is detected, update it. Add new tables and partitions.
	UpdateBehavior *string `field:"optional" json:"updateBehavior" yaml:"updateBehavior"`
}

