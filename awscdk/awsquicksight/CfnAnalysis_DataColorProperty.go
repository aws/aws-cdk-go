package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataColorProperty := &DataColorProperty{
//   	Color: jsii.String("color"),
//   	DataValue: jsii.Number(123),
//   }
//
type CfnAnalysis_DataColorProperty struct {
	// `CfnAnalysis.DataColorProperty.Color`.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// `CfnAnalysis.DataColorProperty.DataValue`.
	DataValue *float64 `field:"optional" json:"dataValue" yaml:"dataValue"`
}

