package previewawsquicksightmixins


// An element in the hierarchical path to a table within a data source, containing both name and identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tablePathElementProperty := &TablePathElementProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tablepathelement.html
//
type CfnDataSetPropsMixin_TablePathElementProperty struct {
	// The unique identifier of the path element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tablepathelement.html#cfn-quicksight-dataset-tablepathelement-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the path element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tablepathelement.html#cfn-quicksight-dataset-tablepathelement-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

