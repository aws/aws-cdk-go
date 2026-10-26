package awsquicksight


// The application theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   applicationThemeProperty := &ApplicationThemeProperty{
//   	BrandColorPalette: &BrandColorPaletteProperty{
//   		Accent: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		Danger: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		Dimension: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		Info: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		Measure: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		Primary: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		Secondary: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		Success: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   		Warning: &PaletteProperty{
//   			Background: jsii.String("background"),
//   			Foreground: jsii.String("foreground"),
//   		},
//   	},
//   	BrandElementStyle: &BrandElementStyleProperty{
//   		NavbarStyle: &NavbarStyleProperty{
//   			ContextualNavbar: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			GlobalNavbar: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-applicationtheme.html
//
type CfnBrandPropsMixin_ApplicationThemeProperty struct {
	// The color palette.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-applicationtheme.html#cfn-quicksight-brand-applicationtheme-brandcolorpalette
	//
	BrandColorPalette interface{} `field:"optional" json:"brandColorPalette" yaml:"brandColorPalette"`
	// The element style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-applicationtheme.html#cfn-quicksight-brand-applicationtheme-brandelementstyle
	//
	BrandElementStyle interface{} `field:"optional" json:"brandElementStyle" yaml:"brandElementStyle"`
}

