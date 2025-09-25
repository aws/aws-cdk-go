package awsiotfleetwise


// A reference to a SignalCatalog resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signalCatalogReference := &SignalCatalogReference{
//   	SignalCatalogArn: jsii.String("signalCatalogArn"),
//   	SignalCatalogName: jsii.String("signalCatalogName"),
//   }
//
type SignalCatalogReference struct {
	// The ARN of the SignalCatalog resource.
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// The Name of the SignalCatalog resource.
	SignalCatalogName *string `field:"required" json:"signalCatalogName" yaml:"signalCatalogName"`
}

