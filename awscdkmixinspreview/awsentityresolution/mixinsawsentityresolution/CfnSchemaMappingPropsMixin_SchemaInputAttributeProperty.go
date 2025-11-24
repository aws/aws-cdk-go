package mixinsawsentityresolution


// A configuration object for defining input data fields in AWS Entity Resolution .
//
// The `SchemaInputAttribute` specifies how individual fields in your input data should be processed and matched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaInputAttributeProperty := &SchemaInputAttributeProperty{
//   	FieldName: jsii.String("fieldName"),
//   	GroupName: jsii.String("groupName"),
//   	Hashed: jsii.Boolean(false),
//   	MatchKey: jsii.String("matchKey"),
//   	SubType: jsii.String("subType"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html
//
type CfnSchemaMappingPropsMixin_SchemaInputAttributeProperty struct {
	// A string containing the field name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-fieldname
	//
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// A string that instructs AWS Entity Resolution to combine several columns into a unified column with the identical attribute type.
	//
	// For example, when working with columns such as `NAME_FIRST` , `NAME_MIDDLE` , and `NAME_LAST` , assigning them a common `groupName` will prompt AWS Entity Resolution to concatenate them into a single value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Indicates if the column values are hashed in the schema input.
	//
	// If the value is set to `TRUE` , the column values are hashed.
	//
	// If the value is set to `FALSE` , the column values are cleartext.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-hashed
	//
	Hashed interface{} `field:"optional" json:"hashed" yaml:"hashed"`
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
	// The type of the attribute, selected from a list of values.
	//
	// LiveRamp supports: `NAME` | `NAME_FIRST` | `NAME_MIDDLE` | `NAME_LAST` | `ADDRESS` | `ADDRESS_STREET1` | `ADDRESS_STREET2` | `ADDRESS_STREET3` | `ADDRESS_CITY` | `ADDRESS_STATE` | `ADDRESS_COUNTRY` | `ADDRESS_POSTALCODE` | `PHONE` | `PHONE_NUMBER` | `EMAIL_ADDRESS` | `UNIQUE_ID` | `PROVIDER_ID`
	//
	// TransUnion supports: `NAME` | `NAME_FIRST` | `NAME_LAST` | `ADDRESS` | `ADDRESS_CITY` | `ADDRESS_STATE` | `ADDRESS_COUNTRY` | `ADDRESS_POSTALCODE` | `PHONE_NUMBER` | `EMAIL_ADDRESS` | `UNIQUE_ID` | `IPV4` | `IPV6` | `MAID`
	//
	// Unified ID 2.0 supports: `PHONE_NUMBER` | `EMAIL_ADDRESS` | `UNIQUE_ID`
	//
	// > Normalization is only supported for `NAME` , `ADDRESS` , `PHONE` , and `EMAIL_ADDRESS` .
	// >
	// > If you want to normalize `NAME_FIRST` , `NAME_MIDDLE` , and `NAME_LAST` , you must group them by assigning them to the `NAME` `groupName` .
	// >
	// > If you want to normalize `ADDRESS_STREET1` , `ADDRESS_STREET2` , `ADDRESS_STREET3` , `ADDRESS_CITY` , `ADDRESS_STATE` , `ADDRESS_COUNTRY` , and `ADDRESS_POSTALCODE` , you must group them by assigning them to the `ADDRESS` `groupName` .
	// >
	// > If you want to normalize `PHONE_NUMBER` and `PHONE_COUNTRYCODE` , you must group them by assigning them to the `PHONE` `groupName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-schemamapping-schemainputattribute.html#cfn-entityresolution-schemamapping-schemainputattribute-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

