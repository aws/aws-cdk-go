package awsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataAccessor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributeFilterProperty_ attributeFilterProperty
//
//   cfnDataAccessorProps := &CfnDataAccessorProps{
//   	ActionConfigurations: []interface{}{
//   		&ActionConfigurationProperty{
//   			Action: jsii.String("action"),
//
//   			// the properties below are optional
//   			FilterConfiguration: &ActionFilterConfigurationProperty{
//   				DocumentAttributeFilter: &attributeFilterProperty{
//   					AndAllFilters: []interface{}{
//   						attributeFilterProperty_,
//   					},
//   					ContainsAll: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					ContainsAny: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					EqualsTo: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					GreaterThan: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					GreaterThanOrEquals: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					LessThan: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					LessThanOrEquals: &DocumentAttributeProperty{
//   						Name: jsii.String("name"),
//   						Value: &DocumentAttributeValueProperty{
//   							DateValue: jsii.String("dateValue"),
//   							LongValue: jsii.Number(123),
//   							StringListValue: []*string{
//   								jsii.String("stringListValue"),
//   							},
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					NotFilter: attributeFilterProperty_,
//   					OrAllFilters: []interface{}{
//   						attributeFilterProperty_,
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	ApplicationId: jsii.String("applicationId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html
//
type CfnDataAccessorProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-actionconfigurations
	//
	ActionConfigurations interface{} `field:"required" json:"actionConfigurations" yaml:"actionConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-principal
	//
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-dataaccessor.html#cfn-qbusiness-dataaccessor-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

