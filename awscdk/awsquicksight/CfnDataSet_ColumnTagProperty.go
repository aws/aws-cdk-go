package awsquicksight


// A tag for a column in a `[TagColumnOperation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_TagColumnOperation.html)` structure. This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnTagProperty := &ColumnTagProperty{
//   	ColumnDescription: &ColumnDescriptionProperty{
//   		Text: jsii.String("text"),
//   	},
//   	ColumnGeographicRole: jsii.String("columnGeographicRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columntag.html
//
type CfnDataSet_ColumnTagProperty struct {
	// A description for a column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columntag.html#cfn-quicksight-dataset-columntag-columndescription
	//
	ColumnDescription interface{} `field:"optional" json:"columnDescription" yaml:"columnDescription"`
	// A geospatial role for a column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columntag.html#cfn-quicksight-dataset-columntag-columngeographicrole
	//
	ColumnGeographicRole *string `field:"optional" json:"columnGeographicRole" yaml:"columnGeographicRole"`
}

