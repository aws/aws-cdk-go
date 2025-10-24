package awsqbusiness


// Specifies an allowed action and its associated filter configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributeFilterProperty_ AttributeFilterProperty
//
//   actionConfigurationProperty := &ActionConfigurationProperty{
//   	Action: jsii.String("action"),
//
//   	// the properties below are optional
//   	FilterConfiguration: &ActionFilterConfigurationProperty{
//   		DocumentAttributeFilter: &AttributeFilterProperty{
//   			AndAllFilters: []interface{}{
//   				attributeFilterProperty_,
//   			},
//   			ContainsAll: &DocumentAttributeProperty{
//   				Name: jsii.String("name"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			ContainsAny: &DocumentAttributeProperty{
//   				Name: jsii.String("name"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			EqualsTo: &DocumentAttributeProperty{
//   				Name: jsii.String("name"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			GreaterThan: &DocumentAttributeProperty{
//   				Name: jsii.String("name"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			GreaterThanOrEquals: &DocumentAttributeProperty{
//   				Name: jsii.String("name"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			LessThan: &DocumentAttributeProperty{
//   				Name: jsii.String("name"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			LessThanOrEquals: &DocumentAttributeProperty{
//   				Name: jsii.String("name"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			NotFilter: attributeFilterProperty_,
//   			OrAllFilters: []interface{}{
//   				attributeFilterProperty_,
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-actionconfiguration.html
//
type CfnDataAccessor_ActionConfigurationProperty struct {
	// The Amazon Q Business action that is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-actionconfiguration.html#cfn-qbusiness-dataaccessor-actionconfiguration-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The filter configuration for the action, if any.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-actionconfiguration.html#cfn-qbusiness-dataaccessor-actionconfiguration-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
}

