package awsqbusiness


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributeFilterProperty_ attributeFilterProperty
//
//   attributeFilterProperty := &attributeFilterProperty{
//   	AndAllFilters: []interface{}{
//   		attributeFilterProperty_,
//   	},
//   	ContainsAll: &DocumentAttributeProperty{
//   		Name: jsii.String("name"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	ContainsAny: &DocumentAttributeProperty{
//   		Name: jsii.String("name"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	EqualsTo: &DocumentAttributeProperty{
//   		Name: jsii.String("name"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	GreaterThan: &DocumentAttributeProperty{
//   		Name: jsii.String("name"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	GreaterThanOrEquals: &DocumentAttributeProperty{
//   		Name: jsii.String("name"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	LessThan: &DocumentAttributeProperty{
//   		Name: jsii.String("name"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	LessThanOrEquals: &DocumentAttributeProperty{
//   		Name: jsii.String("name"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	NotFilter: attributeFilterProperty_,
//   	OrAllFilters: []interface{}{
//   		attributeFilterProperty_,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html
//
type CfnDataAccessor_AttributeFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-andallfilters
	//
	AndAllFilters interface{} `field:"optional" json:"andAllFilters" yaml:"andAllFilters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-containsall
	//
	ContainsAll interface{} `field:"optional" json:"containsAll" yaml:"containsAll"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-containsany
	//
	ContainsAny interface{} `field:"optional" json:"containsAny" yaml:"containsAny"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-equalsto
	//
	EqualsTo interface{} `field:"optional" json:"equalsTo" yaml:"equalsTo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-greaterthan
	//
	GreaterThan interface{} `field:"optional" json:"greaterThan" yaml:"greaterThan"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-greaterthanorequals
	//
	GreaterThanOrEquals interface{} `field:"optional" json:"greaterThanOrEquals" yaml:"greaterThanOrEquals"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-lessthan
	//
	LessThan interface{} `field:"optional" json:"lessThan" yaml:"lessThan"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-lessthanorequals
	//
	LessThanOrEquals interface{} `field:"optional" json:"lessThanOrEquals" yaml:"lessThanOrEquals"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-notfilter
	//
	NotFilter interface{} `field:"optional" json:"notFilter" yaml:"notFilter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-attributefilter.html#cfn-qbusiness-dataaccessor-attributefilter-orallfilters
	//
	OrAllFilters interface{} `field:"optional" json:"orAllFilters" yaml:"orAllFilters"`
}

