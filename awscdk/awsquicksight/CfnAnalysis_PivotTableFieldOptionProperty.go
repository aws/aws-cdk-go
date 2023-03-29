package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableFieldOptionProperty := &PivotTableFieldOptionProperty{
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	CustomLabel: jsii.String("customLabel"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnAnalysis_PivotTableFieldOptionProperty struct {
	// `CfnAnalysis.PivotTableFieldOptionProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnAnalysis.PivotTableFieldOptionProperty.CustomLabel`.
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// `CfnAnalysis.PivotTableFieldOptionProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

