package interfacesawsopensearchserverless


// A reference to a Index resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   indexReference := &IndexReference{
//   	CollectionEndpoint: jsii.String("collectionEndpoint"),
//   	IndexName: jsii.String("indexName"),
//   }
//
type IndexReference struct {
	// The CollectionEndpoint of the Index resource.
	CollectionEndpoint *string `field:"required" json:"collectionEndpoint" yaml:"collectionEndpoint"`
	// The IndexName of the Index resource.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
}

