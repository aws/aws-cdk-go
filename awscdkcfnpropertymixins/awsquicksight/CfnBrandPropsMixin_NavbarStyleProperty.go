package awsquicksight


// The navigation bar style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   navbarStyleProperty := &NavbarStyleProperty{
//   	ContextualNavbar: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   	GlobalNavbar: &PaletteProperty{
//   		Background: jsii.String("background"),
//   		Foreground: jsii.String("foreground"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-navbarstyle.html
//
type CfnBrandPropsMixin_NavbarStyleProperty struct {
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-navbarstyle.html#cfn-quicksight-brand-navbarstyle-contextualnavbar
	//
	ContextualNavbar interface{} `field:"optional" json:"contextualNavbar" yaml:"contextualNavbar"`
	// The palette colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-navbarstyle.html#cfn-quicksight-brand-navbarstyle-globalnavbar
	//
	GlobalNavbar interface{} `field:"optional" json:"globalNavbar" yaml:"globalNavbar"`
}

