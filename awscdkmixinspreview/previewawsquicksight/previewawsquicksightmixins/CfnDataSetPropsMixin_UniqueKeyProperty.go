package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   uniqueKeyProperty := &UniqueKeyProperty{
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uniquekey.html
//
type CfnDataSetPropsMixin_UniqueKeyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uniquekey.html#cfn-quicksight-dataset-uniquekey-columnnames
	//
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
}

