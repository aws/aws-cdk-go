package interfacesawsdataexchange


// A reference to a DataSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetReference := &DataSetReference{
//   	DataSetArn: jsii.String("dataSetArn"),
//   }
//
type DataSetReference struct {
	// The Arn of the DataSet resource.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
}

