package awscognito


// The minimum and maximum values of an attribute that is of the number data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberAttributeConstraintsProperty := &NumberAttributeConstraintsProperty{
//   	MaxValue: jsii.String("maxValue"),
//   	MinValue: jsii.String("minValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html
//
type CfnUserPool_NumberAttributeConstraintsProperty struct {
	// The maximum length of a number attribute value.
	//
	// Must be a number less than or equal to `2^1023` , represented as a string with a length of 131072 characters or fewer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html#cfn-cognito-userpool-numberattributeconstraints-maxvalue
	//
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum value of an attribute that is of the number data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html#cfn-cognito-userpool-numberattributeconstraints-minvalue
	//
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
}

