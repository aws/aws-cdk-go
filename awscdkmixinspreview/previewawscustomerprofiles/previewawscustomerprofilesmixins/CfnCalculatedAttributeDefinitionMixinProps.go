package previewawscustomerprofilesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCalculatedAttributeDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCalculatedAttributeDefinitionMixinProps := &CfnCalculatedAttributeDefinitionMixinProps{
//   	AttributeDetails: &AttributeDetailsProperty{
//   		Attributes: []interface{}{
//   			&AttributeItemProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Expression: jsii.String("expression"),
//   	},
//   	CalculatedAttributeName: jsii.String("calculatedAttributeName"),
//   	Conditions: &ConditionsProperty{
//   		ObjectCount: jsii.Number(123),
//   		Range: &RangeProperty{
//   			TimestampFormat: jsii.String("timestampFormat"),
//   			TimestampSource: jsii.String("timestampSource"),
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
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
//   	DomainName: jsii.String("domainName"),
//   	Statistic: jsii.String("statistic"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseHistoricalData: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html
//
type CfnCalculatedAttributeDefinitionMixinProps struct {
	// Mathematical expression and a list of attribute items specified in that expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-attributedetails
	//
	AttributeDetails interface{} `field:"optional" json:"attributeDetails" yaml:"attributeDetails"`
	// The name of an attribute defined in a profile object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-calculatedattributename
	//
	CalculatedAttributeName *string `field:"optional" json:"calculatedAttributeName" yaml:"calculatedAttributeName"`
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
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The aggregation operation to perform for the calculated attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-statistic
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Whether historical data ingested before the Calculated Attribute was created should be included in calculations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-calculatedattributedefinition.html#cfn-customerprofiles-calculatedattributedefinition-usehistoricaldata
	//
	UseHistoricalData interface{} `field:"optional" json:"useHistoricalData" yaml:"useHistoricalData"`
}

