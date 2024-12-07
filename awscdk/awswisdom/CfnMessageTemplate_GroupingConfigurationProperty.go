package awswisdom


// The configuration information of the user groups that the message template is accessible to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupingConfigurationProperty := &GroupingConfigurationProperty{
//   	Criteria: jsii.String("criteria"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-groupingconfiguration.html
//
type CfnMessageTemplate_GroupingConfigurationProperty struct {
	// The criteria used for grouping Amazon Q in Connect users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-groupingconfiguration.html#cfn-wisdom-messagetemplate-groupingconfiguration-criteria
	//
	Criteria *string `field:"required" json:"criteria" yaml:"criteria"`
	// The list of values that define different groups of Amazon Q in Connect users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-groupingconfiguration.html#cfn-wisdom-messagetemplate-groupingconfiguration-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

