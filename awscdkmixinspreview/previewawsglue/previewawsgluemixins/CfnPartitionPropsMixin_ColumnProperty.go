package previewawsgluemixins


// A column in a `Table` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnProperty := &ColumnProperty{
//   	Comment: jsii.String("comment"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-column.html
//
type CfnPartitionPropsMixin_ColumnProperty struct {
	// A free-form text comment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-column.html#cfn-glue-partition-column-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the `Column` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-column.html#cfn-glue-partition-column-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The data type of the `Column` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-column.html#cfn-glue-partition-column-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

