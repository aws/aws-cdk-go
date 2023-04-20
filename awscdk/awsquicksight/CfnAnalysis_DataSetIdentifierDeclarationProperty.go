package awsquicksight


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
	// `CfnAnalysis.DataSetIdentifierDeclarationProperty.DataSetArn`.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// `CfnAnalysis.DataSetIdentifierDeclarationProperty.Identifier`.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

