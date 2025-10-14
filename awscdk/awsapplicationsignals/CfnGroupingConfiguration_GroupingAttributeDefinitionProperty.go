package awsapplicationsignals


// A structure that defines how services should be grouped based on specific attributes.
//
// This includes the friendly name for the grouping, the source keys to derive values from, and an optional default value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupingAttributeDefinitionProperty := &GroupingAttributeDefinitionProperty{
//   	GroupingName: jsii.String("groupingName"),
//   	GroupingSourceKeys: []*string{
//   		jsii.String("groupingSourceKeys"),
//   	},
//
//   	// the properties below are optional
//   	DefaultGroupingValue: jsii.String("defaultGroupingValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-groupingconfiguration-groupingattributedefinition.html
//
type CfnGroupingConfiguration_GroupingAttributeDefinitionProperty struct {
	// The friendly name for this grouping attribute, such as `BusinessUnit` or `Environment` .
	//
	// This name is used to identify the grouping in the console and APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-groupingconfiguration-groupingattributedefinition.html#cfn-applicationsignals-groupingconfiguration-groupingattributedefinition-groupingname
	//
	GroupingName *string `field:"required" json:"groupingName" yaml:"groupingName"`
	// An array of source keys used to derive the grouping attribute value from telemetry data, AWS tags, or other sources.
	//
	// For example, ["business_unit", "team"] would look for values in those fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-groupingconfiguration-groupingattributedefinition.html#cfn-applicationsignals-groupingconfiguration-groupingattributedefinition-groupingsourcekeys
	//
	GroupingSourceKeys *[]*string `field:"required" json:"groupingSourceKeys" yaml:"groupingSourceKeys"`
	// The default value to use for this grouping attribute when no value can be derived from the source keys.
	//
	// This ensures all services have a grouping value even if the source data is missing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-groupingconfiguration-groupingattributedefinition.html#cfn-applicationsignals-groupingconfiguration-groupingattributedefinition-defaultgroupingvalue
	//
	DefaultGroupingValue *string `field:"optional" json:"defaultGroupingValue" yaml:"defaultGroupingValue"`
}

