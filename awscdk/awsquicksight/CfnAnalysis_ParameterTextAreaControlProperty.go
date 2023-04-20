package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterTextAreaControlProperty := &ParameterTextAreaControlProperty{
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Delimiter: jsii.String("delimiter"),
//   	DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   		PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TitleOptions: &LabelOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   }
//
type CfnAnalysis_ParameterTextAreaControlProperty struct {
	// `CfnAnalysis.ParameterTextAreaControlProperty.ParameterControlId`.
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// `CfnAnalysis.ParameterTextAreaControlProperty.SourceParameterName`.
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// `CfnAnalysis.ParameterTextAreaControlProperty.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `CfnAnalysis.ParameterTextAreaControlProperty.Delimiter`.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// `CfnAnalysis.ParameterTextAreaControlProperty.DisplayOptions`.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

