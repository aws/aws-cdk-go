package interfacesawshealthimaging


// A reference to a Datastore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datastoreReference := &DatastoreReference{
//   	DatastoreArn: jsii.String("datastoreArn"),
//   	DatastoreId: jsii.String("datastoreId"),
//   }
//
type DatastoreReference struct {
	// The ARN of the Datastore resource.
	DatastoreArn *string `field:"required" json:"datastoreArn" yaml:"datastoreArn"`
	// The DatastoreId of the Datastore resource.
	DatastoreId *string `field:"required" json:"datastoreId" yaml:"datastoreId"`
}

