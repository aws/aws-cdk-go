package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customFilterConfigurationProperty := &CustomFilterConfigurationProperty{
//   	MatchOperator: jsii.String("matchOperator"),
//   	NullOption: jsii.String("nullOption"),
//
//   	// the properties below are optional
//   	CategoryValue: jsii.String("categoryValue"),
//   	ParameterName: jsii.String("parameterName"),
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
type CfnAnalysis_CustomFilterConfigurationProperty struct {
	// `CfnAnalysis.CustomFilterConfigurationProperty.MatchOperator`.
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// `CfnAnalysis.CustomFilterConfigurationProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnAnalysis.CustomFilterConfigurationProperty.CategoryValue`.
	CategoryValue *string `field:"optional" json:"categoryValue" yaml:"categoryValue"`
	// `CfnAnalysis.CustomFilterConfigurationProperty.ParameterName`.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// `CfnAnalysis.CustomFilterConfigurationProperty.SelectAllOptions`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}

