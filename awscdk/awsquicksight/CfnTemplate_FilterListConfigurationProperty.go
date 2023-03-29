package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterListConfigurationProperty := &FilterListConfigurationProperty{
//   	MatchOperator: jsii.String("matchOperator"),
//
//   	// the properties below are optional
//   	CategoryValues: []*string{
//   		jsii.String("categoryValues"),
//   	},
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
type CfnTemplate_FilterListConfigurationProperty struct {
	// `CfnTemplate.FilterListConfigurationProperty.MatchOperator`.
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// `CfnTemplate.FilterListConfigurationProperty.CategoryValues`.
	CategoryValues *[]*string `field:"optional" json:"categoryValues" yaml:"categoryValues"`
	// `CfnTemplate.FilterListConfigurationProperty.SelectAllOptions`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}

