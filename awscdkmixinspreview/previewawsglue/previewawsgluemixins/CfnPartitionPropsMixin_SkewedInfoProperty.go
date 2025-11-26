package previewawsgluemixins


// Specifies skewed values in a table.
//
// Skewed values are those that occur with very high frequency.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var skewedColumnValueLocationMaps interface{}
//
//   skewedInfoProperty := &SkewedInfoProperty{
//   	SkewedColumnNames: []*string{
//   		jsii.String("skewedColumnNames"),
//   	},
//   	SkewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   	SkewedColumnValues: []*string{
//   		jsii.String("skewedColumnValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html
//
type CfnPartitionPropsMixin_SkewedInfoProperty struct {
	// A list of names of columns that contain skewed values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html#cfn-glue-partition-skewedinfo-skewedcolumnnames
	//
	SkewedColumnNames *[]*string `field:"optional" json:"skewedColumnNames" yaml:"skewedColumnNames"`
	// A mapping of skewed values to the columns that contain them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html#cfn-glue-partition-skewedinfo-skewedcolumnvaluelocationmaps
	//
	SkewedColumnValueLocationMaps interface{} `field:"optional" json:"skewedColumnValueLocationMaps" yaml:"skewedColumnValueLocationMaps"`
	// A list of values that appear so frequently as to be considered skewed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html#cfn-glue-partition-skewedinfo-skewedcolumnvalues
	//
	SkewedColumnValues *[]*string `field:"optional" json:"skewedColumnValues" yaml:"skewedColumnValues"`
}

