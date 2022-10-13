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
	Description *string `field:"required" json:"description" yaml:"description"`
	// The ID of the detector belonging to the GuardDuty account that you want to create a filter for.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// Represents the criteria to be used in the filter for querying findings.
	FindingCriteria interface{} `field:"required" json:"findingCriteria" yaml:"findingCriteria"`
	// The name of the filter.
	//
	// Minimum length of 3. Maximum length of 64. Valid characters include alphanumeric characters, dot (.), underscore (_), and dash (-). Spaces are not allowed.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::GuardDuty::Filter.Rank`.
	Rank *float64 `field:"required" json:"rank" yaml:"rank"`
	// `AWS::GuardDuty::Filter.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

