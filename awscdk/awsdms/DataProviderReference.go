package awsdms


// A reference to a DataProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataProviderReference := &DataProviderReference{
//   	DataProviderArn: jsii.String("dataProviderArn"),
//   }
//
type DataProviderReference struct {
	// The DataProviderArn of the DataProvider resource.
	DataProviderArn *string `field:"required" json:"dataProviderArn" yaml:"dataProviderArn"`
}

