package mixinsawsquicksight


// The style configuration of the reference line.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceLineStyleConfigurationProperty := &ReferenceLineStyleConfigurationProperty{
//   	Color: jsii.String("color"),
//   	Pattern: jsii.String("pattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinestyleconfiguration.html
//
type CfnTemplatePropsMixin_ReferenceLineStyleConfigurationProperty struct {
	// The hex color of the reference line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinestyleconfiguration.html#cfn-quicksight-template-referencelinestyleconfiguration-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The pattern type of the line style. Choose one of the following options:.
	//
	// - `SOLID`
	// - `DASHED`
	// - `DOTTED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinestyleconfiguration.html#cfn-quicksight-template-referencelinestyleconfiguration-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

