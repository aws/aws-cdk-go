package previewawscognitomixins


// The minimum and maximum length values of an attribute that is of the string type, for example `custom:department` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stringAttributeConstraintsProperty := &StringAttributeConstraintsProperty{
//   	MaxLength: jsii.String("maxLength"),
//   	MinLength: jsii.String("minLength"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html
//
type CfnUserPoolPropsMixin_StringAttributeConstraintsProperty struct {
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

