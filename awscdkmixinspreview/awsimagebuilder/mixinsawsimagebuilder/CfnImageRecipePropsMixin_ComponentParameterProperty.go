package mixinsawsimagebuilder


// Contains a key/value pair that sets the named component parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   componentParameterProperty := &ComponentParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentparameter.html
//
type CfnImageRecipePropsMixin_ComponentParameterProperty struct {
	// The name of the component parameter to set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentparameter.html#cfn-imagebuilder-imagerecipe-componentparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sets the value for the named component parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentparameter.html#cfn-imagebuilder-imagerecipe-componentparameter-value
	//
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

