package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCalculatedAttributeDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCalculatedAttributeDefinitionProps := &CfnCalculatedAttributeDefinitionProps{
//   	AttributeDetails: &AttributeDetailsProperty{
//   		Attributes: []interface{}{
//   			&AttributeItemProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Expression: jsii.String("expression"),
//   	},
//   	CalculatedAttributeName: jsii.String("calculatedAttributeName"),
//   	DomainName: jsii.String("domainName"),
//   	Statistic: jsii.String("statistic"),
//
//   	// the properties below are optional
//   	Conditions: &ConditionsProperty{
//   		ObjectCount: jsii.Number(123),
//   		Range: &RangeProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//
//   			// the properties below are optional
//   			TimestampFormat: jsii.String("timestampFormat"),
//   			TimestampSource: jsii.String("timestampSource"),
//   			ValueRange: &ValueRangeProperty{
//   				End: jsii.Number(123),
//   				Start: jsii.Number(123),
//   			},
//   		},
//   		Threshold: &ThresholdProperty{
//   			Operator: jsii.String("operator"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseHistoricalData: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html
//
type CfnCalculatedAttributeDefinitionProps struct {
	// Mathematical expression and a list of attribute items specified in that expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-attributedetails
	//
	AttributeDetails interface{} `field:"required" json:"attributeDetails" yaml:"attributeDetails"`
	// The name of an attribute defined in a profile object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-calculatedattributename
	//
	CalculatedAttributeName *string `field:"required" json:"calculatedAttributeName" yaml:"calculatedAttributeName"`
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The aggregation operation to perform for the calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-statistic
	//
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// The conditions including range, object count, and threshold for the calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The description of the calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Whether to use historical data for the calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-usehistoricaldata
	//
	UseHistoricalData interface{} `field:"optional" json:"useHistoricalData" yaml:"useHistoricalData"`
}

