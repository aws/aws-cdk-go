package awsglue


// Defines a field in an entity schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldDefinitionProperty := &FieldDefinitionProperty{
//   	FieldDataType: jsii.String("fieldDataType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	FilterOverrides: &FilterOverridesProperty{
//   		BetweenConfiguration: &BetweenConfigurationProperty{
//   			HighBoundKey: jsii.String("highBoundKey"),
//   			LowBoundKey: jsii.String("lowBoundKey"),
//   			Template: jsii.String("template"),
//   		},
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
//   		FieldName: jsii.String("fieldName"),
//   		OperatorMappings: map[string]*string{
//   			"operatorMappingsKey": jsii.String("operatorMappings"),
//   		},
//   	},
//   	IsNullable: jsii.Boolean(false),
//   	IsOrderable: jsii.Boolean(false),
//   	IsPartitionable: jsii.Boolean(false),
//   	IsQueryable: jsii.Boolean(false),
//   	ResponseDateFormat: jsii.String("responseDateFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html
//
type CfnConnectionType_FieldDefinitionProperty struct {
	// The data type of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html#cfn-glue-connectiontype-fielddefinition-fielddatatype
	//
	FieldDataType *string `field:"required" json:"fieldDataType" yaml:"fieldDataType"`
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html#cfn-glue-connectiontype-fielddefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration that defines per-field overrides for filter behavior, allowing individual fields to customize how filter operations are applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html#cfn-glue-connectiontype-fielddefinition-filteroverrides
	//
	FilterOverrides interface{} `field:"optional" json:"filterOverrides" yaml:"filterOverrides"`
	// Indicates whether this field can contain null values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html#cfn-glue-connectiontype-fielddefinition-isnullable
	//
	IsNullable interface{} `field:"optional" json:"isNullable" yaml:"isNullable"`
	// Indicates whether this field can be used for ordering results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html#cfn-glue-connectiontype-fielddefinition-isorderable
	//
	IsOrderable interface{} `field:"optional" json:"isOrderable" yaml:"isOrderable"`
	// Indicates whether this field can be used for partitioning queries to the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html#cfn-glue-connectiontype-fielddefinition-ispartitionable
	//
	IsPartitionable interface{} `field:"optional" json:"isPartitionable" yaml:"isPartitionable"`
	// Indicates whether this field can be used in filter predicates when querying data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html#cfn-glue-connectiontype-fielddefinition-isqueryable
	//
	IsQueryable interface{} `field:"optional" json:"isQueryable" yaml:"isQueryable"`
	// The format pattern for parsing date values from API responses.
	//
	// Accepts Java DateTimeFormatter patterns, EPOCH_SECONDS, or EPOCH_MILLIS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-fielddefinition.html#cfn-glue-connectiontype-fielddefinition-responsedateformat
	//
	ResponseDateFormat *string `field:"optional" json:"responseDateFormat" yaml:"responseDateFormat"`
}

