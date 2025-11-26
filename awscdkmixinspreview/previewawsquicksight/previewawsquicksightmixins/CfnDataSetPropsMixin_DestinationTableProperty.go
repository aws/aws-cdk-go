package previewawsquicksightmixins


// Defines a destination table in data preparation that receives the final transformed data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationTableProperty := &DestinationTableProperty{
//   	Alias: jsii.String("alias"),
//   	Source: &DestinationTableSourceProperty{
//   		TransformOperationId: jsii.String("transformOperationId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-destinationtable.html
//
type CfnDataSetPropsMixin_DestinationTableProperty struct {
	// Alias for the destination table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-destinationtable.html#cfn-quicksight-dataset-destinationtable-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The source configuration that specifies which transform operation provides data to this destination table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-destinationtable.html#cfn-quicksight-dataset-destinationtable-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

