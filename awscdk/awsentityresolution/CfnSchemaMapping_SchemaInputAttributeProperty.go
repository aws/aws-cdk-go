package awsentityresolution


// An object containing `FieldName` , `Type` , `GroupName` , `MatchKey` , `Hashing` , and `SubType` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaInputAttributeProperty := &SchemaInputAttributeProperty{
//   	FieldName: jsii.String("fieldName"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	GroupName: jsii.String("groupName"),
//   	MatchKey: jsii.String("matchKey"),
//   	SubType: jsii.String("subType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html
//
type CfnSchemaMapping_SchemaInputAttributeProperty struct {
	// A string containing the field name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-fieldname
	//
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// The type of the attribute, selected from a list of values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A string that instructs AWS Entity Resolution to combine several columns into a unified column with the identical attribute type.
	//
	// For example, when working with columns such as `first_name` , `middle_name` , and `last_name` , assigning them a common `groupName` will prompt AWS Entity Resolution to concatenate them into a single value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// A key that allows grouping of multiple input attributes into a unified matching group.
	//
	// For example, consider a scenario where the source table contains various addresses, such as `business_address` and `shipping_address` . By assigning a `matchKey` called `address` to both attributes, AWS Entity Resolution will match records across these fields to create a consolidated matching group.
	//
	// If no `matchKey` is specified for a column, it won't be utilized for matching purposes but will still be included in the output table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-matchkey
	//
	MatchKey *string `field:"optional" json:"matchKey" yaml:"matchKey"`
	// The subtype of the attribute, selected from a list of values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-subtype
	//
	SubType *string `field:"optional" json:"subType" yaml:"subType"`
}

