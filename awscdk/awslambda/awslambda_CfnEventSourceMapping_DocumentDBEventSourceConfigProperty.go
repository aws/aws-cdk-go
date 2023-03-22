package awslambda


// Specific configuration settings for a DocumentDB event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentDBEventSourceConfigProperty := &documentDBEventSourceConfigProperty{
//   	collectionName: jsii.String("collectionName"),
//   	databaseName: jsii.String("databaseName"),
//   	fullDocument: jsii.String("fullDocument"),
//   }
//
type CfnEventSourceMapping_DocumentDBEventSourceConfigProperty struct {
	// The name of the collection to consume within the database.
	//
	// If you do not specify a collection, Lambda consumes all collections.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// The name of the database to consume within the DocumentDB cluster.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Determines what DocumentDB sends to your event stream during document update operations.
	//
	// If set to UpdateLookup, DocumentDB sends a delta describing the changes, along with a copy of the entire document. Otherwise, DocumentDB sends only a partial document that contains the changes.
	FullDocument *string `field:"optional" json:"fullDocument" yaml:"fullDocument"`
}

