package awsglue


// Configuration that defines per-field overrides for filter behavior, allowing individual fields to customize how filter operations are applied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterOverridesProperty := &FilterOverridesProperty{
//   	BetweenConfiguration: &BetweenConfigurationProperty{
//   		HighBoundKey: jsii.String("highBoundKey"),
//   		LowBoundKey: jsii.String("lowBoundKey"),
//   		Template: jsii.String("template"),
//   	},
//   	DateTimeFormat: jsii.String("dateTimeFormat"),
//   	FieldName: jsii.String("fieldName"),
//   	OperatorMappings: map[string]*string{
//   		"operatorMappingsKey": jsii.String("operatorMappings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filteroverrides.html
//
type CfnConnectionType_FilterOverridesProperty struct {
	// Configuration that defines how BETWEEN range filter operations are translated into REST API request parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filteroverrides.html#cfn-glue-connectiontype-filteroverrides-betweenconfiguration
	//
	BetweenConfiguration interface{} `field:"optional" json:"betweenConfiguration" yaml:"betweenConfiguration"`
	// The date and time format for filter expressions on this field, overriding the global DateTimeFormat.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filteroverrides.html#cfn-glue-connectiontype-filteroverrides-datetimeformat
	//
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// An override for the field name to use in filter expressions, if different from the schema field name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filteroverrides.html#cfn-glue-connectiontype-filteroverrides-fieldname
	//
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// A map of logical filter operators to their field-specific API representations, overriding the global operator mappings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filteroverrides.html#cfn-glue-connectiontype-filteroverrides-operatormappings
	//
	OperatorMappings interface{} `field:"optional" json:"operatorMappings" yaml:"operatorMappings"`
}

