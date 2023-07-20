package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var criterion interface{}
//
//   cfnFilterProps := &CfnFilterProps{
//   	Action: jsii.String("action"),
//   	Description: jsii.String("description"),
//   	DetectorId: jsii.String("detectorId"),
//   	FindingCriteria: &FindingCriteriaProperty{
//   		Criterion: criterion,
//   		ItemType: &ConditionProperty{
//   			Eq: []*string{
//   				jsii.String("eq"),
//   			},
//   			EqualTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			GreaterThan: jsii.Number(123),
//   			GreaterThanOrEqual: jsii.Number(123),
//   			Gt: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			LessThan: jsii.Number(123),
//   			LessThanOrEqual: jsii.Number(123),
//   			Lt: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   			Neq: []*string{
//   				jsii.String("neq"),
//   			},
//   			NotEquals: []*string{
//   				jsii.String("notEquals"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Rank: jsii.Number(123),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html
//
type CfnFilterProps struct {
	// Specifies the action that is to be applied to the findings that match the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The description of the filter.
	//
	// Valid characters include alphanumeric characters, and special characters such as hyphen, period, colon, underscore, parentheses ( `{ }` , `[ ]` , and `( )` ), forward slash, horizontal tab, vertical tab, newline, form feed, return, and whitespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The ID of the detector belonging to the GuardDuty account that you want to create a filter for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-detectorid
	//
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// Represents the criteria to be used in the filter for querying findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-findingcriteria
	//
	FindingCriteria interface{} `field:"required" json:"findingCriteria" yaml:"findingCriteria"`
	// The name of the filter.
	//
	// Valid characters include period (.), underscore (_), dash (-), and alphanumeric characters. A whitespace is considered to be an invalid character.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the position of the filter in the list of current filters.
	//
	// Also specifies the order in which this filter is applied to the findings. The minimum value for this property is 1 and the maximum is 100.
	//
	// By default, filters may not be created in the same order as they are ranked. To ensure that the filters are created in the expected order, you can use an optional attribute, [DependsOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) , with the following syntax: `"DependsOn":[ "ObjectName" ]` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-rank
	//
	Rank *float64 `field:"required" json:"rank" yaml:"rank"`
	// The tags to be added to a new filter resource.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

