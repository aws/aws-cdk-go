package interfacesawsiotanalytics


// A reference to a Datastore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datastoreReference := &DatastoreReference{
//   	DatastoreName: jsii.String("datastoreName"),
//   }
//
type DatastoreReference struct {
	// The DatastoreName of the Datastore resource.
	DatastoreName *string `field:"required" json:"datastoreName" yaml:"datastoreName"`
}

