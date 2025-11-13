package interfacesawsrekognition


// A reference to a Collection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collectionReference := &CollectionReference{
//   	CollectionArn: jsii.String("collectionArn"),
//   	CollectionId: jsii.String("collectionId"),
//   }
//
type CollectionReference struct {
	// The ARN of the Collection resource.
	CollectionArn *string `field:"required" json:"collectionArn" yaml:"collectionArn"`
	// The CollectionId of the Collection resource.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
}

