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
type CfnDashboard_DataSetIdentifierDeclarationProperty struct {
	// `CfnDashboard.DataSetIdentifierDeclarationProperty.DataSetArn`.
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// `CfnDashboard.DataSetIdentifierDeclarationProperty.Identifier`.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

