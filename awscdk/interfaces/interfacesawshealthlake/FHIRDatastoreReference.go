package interfacesawshealthlake


// A reference to a FHIRDatastore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fHIRDatastoreReference := &FHIRDatastoreReference{
//   	DatastoreArn: jsii.String("datastoreArn"),
//   	DatastoreId: jsii.String("datastoreId"),
//   }
//
type FHIRDatastoreReference struct {
	// The ARN of the FHIRDatastore resource.
	DatastoreArn *string `field:"required" json:"datastoreArn" yaml:"datastoreArn"`
	// The DatastoreId of the FHIRDatastore resource.
	DatastoreId *string `field:"required" json:"datastoreId" yaml:"datastoreId"`
}

