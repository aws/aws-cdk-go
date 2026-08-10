package awsscn


// The primary key field details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakeDatasetPrimaryKeyFieldProperty := &DataLakeDatasetPrimaryKeyFieldProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetprimarykeyfield.html
//
type CfnDataset_DataLakeDatasetPrimaryKeyFieldProperty struct {
	// The name of the primary key field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-datalakedatasetprimarykeyfield.html#cfn-scn-dataset-datalakedatasetprimarykeyfield-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

