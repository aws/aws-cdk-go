package interfacesawscloudtrail


// A reference to a EventDataStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventDataStoreReference := &EventDataStoreReference{
//   	EventDataStoreArn: jsii.String("eventDataStoreArn"),
//   }
//
type EventDataStoreReference struct {
	// The EventDataStoreArn of the EventDataStore resource.
	EventDataStoreArn *string `field:"required" json:"eventDataStoreArn" yaml:"eventDataStoreArn"`
}

