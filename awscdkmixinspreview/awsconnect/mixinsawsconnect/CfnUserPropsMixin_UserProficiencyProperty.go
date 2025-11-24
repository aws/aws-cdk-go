package mixinsawsconnect


// > A predefined attribute must be created before using `UserProficiencies` in the Cloudformation *User* template.
//
// For more information, see [Predefined attributes](https://docs.aws.amazon.com/connect/latest/adminguide/predefined-attributes.html) .
//
// Proficiency of a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userProficiencyProperty := &UserProficiencyProperty{
//   	AttributeName: jsii.String("attributeName"),
//   	AttributeValue: jsii.String("attributeValue"),
//   	Level: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userproficiency.html
//
type CfnUserPropsMixin_UserProficiencyProperty struct {
	// The name of user’s proficiency.
	//
	// You must use a predefined attribute name that is present in the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userproficiency.html#cfn-connect-user-userproficiency-attributename
	//
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// The value of user’s proficiency.
	//
	// You must use a predefined attribute value that is present in the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userproficiency.html#cfn-connect-user-userproficiency-attributevalue
	//
	AttributeValue *string `field:"optional" json:"attributeValue" yaml:"attributeValue"`
	// The level of the proficiency.
	//
	// The valid values are 1, 2, 3, 4 and 5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-userproficiency.html#cfn-connect-user-userproficiency-level
	//
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

