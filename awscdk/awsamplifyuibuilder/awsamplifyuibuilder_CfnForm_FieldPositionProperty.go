package awsamplifyuibuilder


// Describes the field position.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldPositionProperty := &fieldPositionProperty{
//   	below: jsii.String("below"),
//   	fixed: jsii.String("fixed"),
//   	rightOf: jsii.String("rightOf"),
//   }
//
type CfnForm_FieldPositionProperty struct {
	// The field position is below the field specified by the string.
	Below *string `field:"optional" json:"below" yaml:"below"`
	// The field position is fixed and doesn't change in relation to other fields.
	Fixed *string `field:"optional" json:"fixed" yaml:"fixed"`
	// The field position is to the right of the field specified by the string.
	RightOf *string `field:"optional" json:"rightOf" yaml:"rightOf"`
}

