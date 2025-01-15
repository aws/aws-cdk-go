package awscustomerprofiles


// The criteria that a specific object attribute must meet to trigger the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectAttributeProperty := &ObjectAttributeProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//
//   	// the properties below are optional
//   	FieldName: jsii.String("fieldName"),
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-objectattribute.html
//
type CfnEventTrigger_ObjectAttributeProperty struct {
	// The operator used to compare an attribute against a list of values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-objectattribute.html#cfn-customerprofiles-eventtrigger-objectattribute-comparisonoperator
	//
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// A list of attribute values used for comparison.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-objectattribute.html#cfn-customerprofiles-eventtrigger-objectattribute-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// A field defined within an object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-objectattribute.html#cfn-customerprofiles-eventtrigger-objectattribute-fieldname
	//
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// An attribute contained within a source object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-objectattribute.html#cfn-customerprofiles-eventtrigger-objectattribute-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
}

