package awsamplifyuibuilder


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
type CfnForm_FieldPositionProperty struct {
	// `CfnForm.FieldPositionProperty.Below`.
	Below *string `field:"optional" json:"below" yaml:"below"`
	// `CfnForm.FieldPositionProperty.Fixed`.
	Fixed *string `field:"optional" json:"fixed" yaml:"fixed"`
	// `CfnForm.FieldPositionProperty.RightOf`.
	RightOf *string `field:"optional" json:"rightOf" yaml:"rightOf"`
}

