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
type CfnAnalysis_DataSetIdentifierDeclarationProperty struct {
	// The Amazon Resource Name (ARN) of the data set.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// The identifier of the data set, typically the data set's name.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

