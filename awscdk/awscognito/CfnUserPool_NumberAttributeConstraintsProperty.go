package awscognito


// The minimum and maximum values of an attribute that is of the number type, for example `custom:age` .
//
// This data type is part of [SchemaAttributeType](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SchemaAttributeType.html) . It defines the length constraints on number-type attributes that you configure in [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and displays the length constraints of all number-type attributes in the response to [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html)
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

