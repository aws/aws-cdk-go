package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnFilterProps := &cfnFilterProps{
//   	action: jsii.String("action"),
//   	description: jsii.String("description"),
//   	detectorId: jsii.String("detectorId"),
//   	findingCriteria: &findingCriteriaProperty{
//   		criterion: criterion,
//   		itemType: &conditionProperty{
//   			eq: []*string{
//   				jsii.String("eq"),
//   			},
//   			equalTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			greaterThan: jsii.Number(123),
//   			greaterThanOrEqual: jsii.Number(123),
//   			gt: jsii.Number(123),
//   			gte: jsii.Number(123),
//   			lessThan: jsii.Number(123),
//   			lessThanOrEqual: jsii.Number(123),
//   			lt: jsii.Number(123),
//   			lte: jsii.Number(123),
//   			neq: []*string{
//   				jsii.String("neq"),
//   			},
//   			notEquals: []*string{
//   				jsii.String("notEquals"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	rank: jsii.Number(123),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFilterProps struct {
	// Specifies the action that is to be applied to the findings that match the filter.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The description of the filter.
	//
	// Valid special characters include period (.), underscore (_), dash (-), and whitespace. The new line character is considered to be an invalid input for description.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The ID of the detector belonging to the GuardDuty account that you want to create a filter for.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// Represents the criteria to be used in the filter for querying findings.
	FindingCriteria interface{} `field:"required" json:"findingCriteria" yaml:"findingCriteria"`
	// The name of the filter.
	//
	// Valid characters include period (.), underscore (_), dash (-), and alphanumeric characters. A whitespace is considered to be an invalid character.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the position of the filter in the list of current filters.
	//
	// Also specifies the order in which this filter is applied to the findings. The minimum value for this property is 1 and the maximum is 100.
	//
	// By default, filters may not be created in the same order as they are ranked. To ensure that the filters are created in the expected order, you can use an optional attribute, [DependsOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) , with the following syntax: `"DependsOn":[ "ObjectName" ]` .
	Rank *float64 `field:"required" json:"rank" yaml:"rank"`
	// The tags to be added to a new filter resource.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

