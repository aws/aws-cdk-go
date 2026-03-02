package interfacesawsopensearchserverless


// A reference to a CollectionGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collectionGroupReference := &CollectionGroupReference{
//   	CollectionGroupArn: jsii.String("collectionGroupArn"),
//   	CollectionGroupId: jsii.String("collectionGroupId"),
//   }
//
type CollectionGroupReference struct {
	// The ARN of the CollectionGroup resource.
	CollectionGroupArn *string `field:"required" json:"collectionGroupArn" yaml:"collectionGroupArn"`
	// The Id of the CollectionGroup resource.
	CollectionGroupId *string `field:"required" json:"collectionGroupId" yaml:"collectionGroupId"`
}

