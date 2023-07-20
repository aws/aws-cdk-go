package awssecurityhub


// A map filter for querying findings.
//
// Each map filter provides the field to check, the value to look for, and the comparison operator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mapFilterProperty := &MapFilterProperty{
//   	Comparison: jsii.String("comparison"),
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-mapfilter.html
//
type CfnAutomationRule_MapFilterProperty struct {
	// The condition to apply to the key value when querying for findings with a map filter.
	//
	// To search for values that exactly match the filter value, use `EQUALS` . For example, for the `ResourceTags` field, the filter `Department EQUALS Security` matches findings that have the value `Security` for the tag `Department` .
	//
	// To search for values other than the filter value, use `NOT_EQUALS` . For example, for the `ResourceTags` field, the filter `Department NOT_EQUALS Finance` matches findings that do not have the value `Finance` for the tag `Department` .
	//
	// `EQUALS` filters on the same field are joined by `OR` . A finding matches if it matches any one of those filters.
	//
	// `NOT_EQUALS` filters on the same field are joined by `AND` . A finding matches only if it matches all of those filters.
	//
	// You cannot have both an `EQUALS` filter and a `NOT_EQUALS` filter on the same field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-mapfilter.html#cfn-securityhub-automationrule-mapfilter-comparison
	//
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The key of the map filter.
	//
	// For example, for `ResourceTags` , `Key` identifies the name of the tag. For `UserDefinedFields` , `Key` is the name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-mapfilter.html#cfn-securityhub-automationrule-mapfilter-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the key in the map filter.
	//
	// Filter values are case sensitive. For example, one of the values for a tag called `Department` might be `Security` . If you provide `security` as the filter value, then there is no match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-mapfilter.html#cfn-securityhub-automationrule-mapfilter-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

