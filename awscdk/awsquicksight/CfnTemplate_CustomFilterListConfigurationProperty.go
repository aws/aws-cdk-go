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
type CfnTemplate_CustomFilterListConfigurationProperty struct {
	// `CfnTemplate.CustomFilterListConfigurationProperty.MatchOperator`.
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// `CfnTemplate.CustomFilterListConfigurationProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnTemplate.CustomFilterListConfigurationProperty.CategoryValues`.
	CategoryValues *[]*string `field:"optional" json:"categoryValues" yaml:"categoryValues"`
	// `CfnTemplate.CustomFilterListConfigurationProperty.SelectAllOptions`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}

