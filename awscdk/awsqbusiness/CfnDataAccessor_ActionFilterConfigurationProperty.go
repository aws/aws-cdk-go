package awsqbusiness


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributeFilterProperty_ attributeFilterProperty
//
//   actionFilterConfigurationProperty := &ActionFilterConfigurationProperty{
//   	DocumentAttributeFilter: &attributeFilterProperty{
//   		AndAllFilters: []interface{}{
//   			attributeFilterProperty_,
//   		},
//   		ContainsAll: &DocumentAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		ContainsAny: &DocumentAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		EqualsTo: &DocumentAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		GreaterThan: &DocumentAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		GreaterThanOrEquals: &DocumentAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		LessThan: &DocumentAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		LessThanOrEquals: &DocumentAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		NotFilter: attributeFilterProperty_,
//   		OrAllFilters: []interface{}{
//   			attributeFilterProperty_,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-actionfilterconfiguration.html
//
type CfnDataAccessor_ActionFilterConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-actionfilterconfiguration.html#cfn-qbusiness-dataaccessor-actionfilterconfiguration-documentattributefilter
	//
	DocumentAttributeFilter interface{} `field:"required" json:"documentAttributeFilter" yaml:"documentAttributeFilter"`
}

