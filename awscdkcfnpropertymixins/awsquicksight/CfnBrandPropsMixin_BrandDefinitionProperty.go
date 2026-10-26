package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   brandDefinitionProperty := &BrandDefinitionProperty{
//   	ApplicationTheme: &ApplicationThemeProperty{
//   		BrandColorPalette: &BrandColorPaletteProperty{
//   			Accent: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			Danger: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			Dimension: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			Info: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			Measure: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			Primary: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			Secondary: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			Success: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   			Warning: &PaletteProperty{
//   				Background: jsii.String("background"),
//   				Foreground: jsii.String("foreground"),
//   			},
//   		},
//   		BrandElementStyle: &BrandElementStyleProperty{
//   			NavbarStyle: &NavbarStyleProperty{
//   				ContextualNavbar: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				GlobalNavbar: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   			},
//   		},
//   	},
//   	BrandName: jsii.String("brandName"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-branddefinition.html
//
type CfnBrandPropsMixin_BrandDefinitionProperty struct {
	// The application theme.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-branddefinition.html#cfn-quicksight-brand-branddefinition-applicationtheme
	//
	ApplicationTheme interface{} `field:"optional" json:"applicationTheme" yaml:"applicationTheme"`
	// The name of the brand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-branddefinition.html#cfn-quicksight-brand-branddefinition-brandname
	//
	BrandName *string `field:"optional" json:"brandName" yaml:"brandName"`
	// The description of the brand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-brand-branddefinition.html#cfn-quicksight-brand-branddefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

