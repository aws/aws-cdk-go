package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listControlDisplayOptionsProperty := &ListControlDisplayOptionsProperty{
//   	SearchOptions: &ListControlSearchOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	TitleOptions: &LabelOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnAnalysis_ListControlDisplayOptionsProperty struct {
	// `CfnAnalysis.ListControlDisplayOptionsProperty.SearchOptions`.
	SearchOptions interface{} `field:"optional" json:"searchOptions" yaml:"searchOptions"`
	// `CfnAnalysis.ListControlDisplayOptionsProperty.SelectAllOptions`.
	SelectAllOptions interface{} `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
	// `CfnAnalysis.ListControlDisplayOptionsProperty.TitleOptions`.
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

