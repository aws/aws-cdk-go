package awsiotanalytics


// The datastore activity that specifies where to store the processed data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datastoreProperty := &datastoreProperty{
//   	datastoreName: jsii.String("datastoreName"),
//   	name: jsii.String("name"),
//   }
//
type CfnPipeline_DatastoreProperty struct {
	// The name of the data store where processed messages are stored.
	DatastoreName *string `field:"required" json:"datastoreName" yaml:"datastoreName"`
	// The name of the datastore activity.
	Name *string `field:"required" json:"name" yaml:"name"`
}

