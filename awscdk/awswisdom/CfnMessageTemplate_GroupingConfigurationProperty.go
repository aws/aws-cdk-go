package awswisdom


// The configuration information of the grouping of Amazon Q in Connect users.
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
	//
	// The following is the list of supported criteria values.
	//
	// - `RoutingProfileArn` : Grouping the users by their [Amazon Connect routing profile ARN](https://docs.aws.amazon.com/connect/latest/APIReference/API_RoutingProfile.html) . User should have [SearchRoutingProfile](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchRoutingProfiles.html) and [DescribeRoutingProfile](https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeRoutingProfile.html) permissions when setting criteria to this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-groupingconfiguration.html#cfn-wisdom-messagetemplate-groupingconfiguration-criteria
	//
	Criteria *string `field:"required" json:"criteria" yaml:"criteria"`
	// The list of values that define different groups of Amazon Q in Connect users.
	//
	// - When setting `criteria` to `RoutingProfileArn` , you need to provide a list of ARNs of [Amazon Connect routing profiles](https://docs.aws.amazon.com/connect/latest/APIReference/API_RoutingProfile.html) as values of this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-groupingconfiguration.html#cfn-wisdom-messagetemplate-groupingconfiguration-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

