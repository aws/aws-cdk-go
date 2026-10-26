package awsquicksight


// The element style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brandElementStyleProperty := &BrandElementStyleProperty{
//   	NavbarStyle: &NavbarStyleProperty{
//   		ContextualNavbar: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		GlobalNavbar: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandelementstyle.html
//
type CfnBrand_BrandElementStyleProperty struct {
	// The navigation bar style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-brandelementstyle.html#cfn-quicksight-brand-brandelementstyle-navbarstyle
	//
	NavbarStyle interface{} `field:"optional" json:"navbarStyle" yaml:"navbarStyle"`
}

