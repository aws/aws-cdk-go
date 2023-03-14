package awsquicksight


// Dataset reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetReferenceProperty := &dataSetReferenceProperty{
//   	dataSetArn: jsii.String("dataSetArn"),
//   	dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   }
//
type CfnTemplate_DataSetReferenceProperty struct {
	// Dataset Amazon Resource Name (ARN).
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Dataset placeholder.
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

