package interfacesawss3tables


// A reference to a Namespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceReference := &NamespaceReference{
//   	Namespace: jsii.String("namespace"),
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   }
//
type NamespaceReference struct {
	// The Namespace of the Namespace resource.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The TableBucketARN of the Namespace resource.
	TableBucketArn *string `field:"required" json:"tableBucketArn" yaml:"tableBucketArn"`
}

