package awsamplifyuibuilder


// The `FieldPosition` property specifies the field position.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldPositionProperty := &FieldPositionProperty{
//   	Below: jsii.String("below"),
//   	Fixed: jsii.String("fixed"),
//   	RightOf: jsii.String("rightOf"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldposition.html
//
type CfnForm_FieldPositionProperty struct {
	// The field position is below the field specified by the string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldposition.html#cfn-amplifyuibuilder-form-fieldposition-below
	//
	Below *string `field:"optional" json:"below" yaml:"below"`
	// The field position is fixed and doesn't change in relation to other fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldposition.html#cfn-amplifyuibuilder-form-fieldposition-fixed
	//
	Fixed *string `field:"optional" json:"fixed" yaml:"fixed"`
	// The field position is to the right of the field specified by the string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldposition.html#cfn-amplifyuibuilder-form-fieldposition-rightof
	//
	RightOf *string `field:"optional" json:"rightOf" yaml:"rightOf"`
}

