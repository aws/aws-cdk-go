package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customFilterListConfigurationProperty := &CustomFilterListConfigurationProperty{
//   	MatchOperator: jsii.String("matchOperator"),
//   	NullOption: jsii.String("nullOption"),
//
//   	// the properties below are optional
//   	CategoryValues: []*string{
//   		jsii.String("categoryValues"),
//   	},
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
type CfnAnalysis_CustomFilterListConfigurationProperty struct {
	// `CfnAnalysis.CustomFilterListConfigurationProperty.MatchOperator`.
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// `CfnAnalysis.CustomFilterListConfigurationProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnAnalysis.CustomFilterListConfigurationProperty.CategoryValues`.
	CategoryValues *[]*string `field:"optional" json:"categoryValues" yaml:"categoryValues"`
	// `CfnAnalysis.CustomFilterListConfigurationProperty.SelectAllOptions`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}

