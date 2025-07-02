package awsquicksight


// A data set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetIdentifierDeclarationProperty := &DataSetIdentifierDeclarationProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	Identifier: jsii.String("identifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datasetidentifierdeclaration.html
//
type CfnAnalysis_DataSetIdentifierDeclarationProperty struct {
	// The Amazon Resource Name (ARN) of the data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datasetidentifierdeclaration.html#cfn-quicksight-analysis-datasetidentifierdeclaration-datasetarn
	//
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// The identifier of the data set, typically the data set's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datasetidentifierdeclaration.html#cfn-quicksight-analysis-datasetidentifierdeclaration-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

