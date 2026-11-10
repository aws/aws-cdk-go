package interfacesawsappfabric


// A reference to a Ingestion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestionReference := &IngestionReference{
//   	IngestionArn: jsii.String("ingestionArn"),
//   }
//
type IngestionReference struct {
	// The Arn of the Ingestion resource.
	IngestionArn *string `field:"required" json:"ingestionArn" yaml:"ingestionArn"`
}

