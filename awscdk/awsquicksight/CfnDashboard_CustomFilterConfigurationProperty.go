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
type CfnDashboard_CustomFilterConfigurationProperty struct {
	// `CfnDashboard.CustomFilterConfigurationProperty.MatchOperator`.
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// `CfnDashboard.CustomFilterConfigurationProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnDashboard.CustomFilterConfigurationProperty.CategoryValue`.
	CategoryValue *string `field:"optional" json:"categoryValue" yaml:"categoryValue"`
	// `CfnDashboard.CustomFilterConfigurationProperty.ParameterName`.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// `CfnDashboard.CustomFilterConfigurationProperty.SelectAllOptions`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}

