package interfacesawsappfabric


// A reference to a IngestionDestination resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestionDestinationReference := &IngestionDestinationReference{
//   	IngestionDestinationArn: jsii.String("ingestionDestinationArn"),
//   }
//
type IngestionDestinationReference struct {
	// The Arn of the IngestionDestination resource.
	IngestionDestinationArn *string `field:"required" json:"ingestionDestinationArn" yaml:"ingestionDestinationArn"`
}

