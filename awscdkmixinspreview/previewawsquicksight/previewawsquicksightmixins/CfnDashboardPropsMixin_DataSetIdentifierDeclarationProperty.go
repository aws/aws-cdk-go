package previewawsquicksightmixins


// A data set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetIdentifierDeclarationProperty := &DataSetIdentifierDeclarationProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	Identifier: jsii.String("identifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datasetidentifierdeclaration.html
//
type CfnDashboardPropsMixin_DataSetIdentifierDeclarationProperty struct {
	// The Amazon Resource Name (ARN) of the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datasetidentifierdeclaration.html#cfn-quicksight-dashboard-datasetidentifierdeclaration-datasetarn
	//
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// The identifier of the data set, typically the data set's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datasetidentifierdeclaration.html#cfn-quicksight-dashboard-datasetidentifierdeclaration-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

