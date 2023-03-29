package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableInlineVisualizationProperty := &TableInlineVisualizationProperty{
//   	DataBars: &DataBarsOptionsProperty{
//   		FieldId: jsii.String("fieldId"),
//
//   		// the properties below are optional
//   		NegativeColor: jsii.String("negativeColor"),
//   		PositiveColor: jsii.String("positiveColor"),
//   	},
//   }
//
type CfnAnalysis_TableInlineVisualizationProperty struct {
	// `CfnAnalysis.TableInlineVisualizationProperty.DataBars`.
	DataBars interface{} `field:"optional" json:"dataBars" yaml:"dataBars"`
}

