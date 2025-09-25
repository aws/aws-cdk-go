package awsquicksight


// A reference to a DataSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetReference := &DataSetReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DataSetArn: jsii.String("dataSetArn"),
//   	DataSetId: jsii.String("dataSetId"),
//   }
//
type DataSetReference struct {
	// The AwsAccountId of the DataSet resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the DataSet resource.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// The DataSetId of the DataSet resource.
	DataSetId *string `field:"required" json:"dataSetId" yaml:"dataSetId"`
}

