package interfacesawsopensearchserverless


// A reference to a CollectionIndex resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collectionIndexReference := &CollectionIndexReference{
//   	CollectionIndexId: jsii.String("collectionIndexId"),
//   	IndexName: jsii.String("indexName"),
//   }
//
type CollectionIndexReference struct {
	// The Id of the CollectionIndex resource.
	CollectionIndexId *string `field:"required" json:"collectionIndexId" yaml:"collectionIndexId"`
	// The IndexName of the CollectionIndex resource.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
}

