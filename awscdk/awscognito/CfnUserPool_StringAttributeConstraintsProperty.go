package awscognito


// The minimum and maximum length values of an attribute that is of the string type, for example `custom:department` .
//
// This data type is part of [SchemaAttributeType](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SchemaAttributeType.html) . It defines the length constraints on string-type attributes that you configure in [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and displays the length constraints of all string-type attributes in the response to [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringAttributeConstraintsProperty := &StringAttributeConstraintsProperty{
//   	MaxLength: jsii.String("maxLength"),
//   	MinLength: jsii.String("minLength"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html
//
type CfnUserPool_StringAttributeConstraintsProperty struct {
	// The maximum length of a string attribute value.
	//
	// Must be a number less than or equal to `2^1023` , represented as a string with a length of 131072 characters or fewer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-maxlength
	//
	MaxLength *string `field:"optional" json:"maxLength" yaml:"maxLength"`
	// The minimum length of a string attribute value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-minlength
	//
	MinLength *string `field:"optional" json:"minLength" yaml:"minLength"`
}

