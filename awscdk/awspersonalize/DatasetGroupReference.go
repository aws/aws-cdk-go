package awspersonalize


// A reference to a DatasetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetGroupReference := &DatasetGroupReference{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   }
//
type DatasetGroupReference struct {
	// The DatasetGroupArn of the DatasetGroup resource.
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
}

