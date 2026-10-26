package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnBrandPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnBrandMixinProps := &CfnBrandMixinProps{
//   	BrandDefinition: &BrandDefinitionProperty{
//   		ApplicationTheme: &ApplicationThemeProperty{
//   			BrandColorPalette: &BrandColorPaletteProperty{
//   				Accent: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Danger: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Dimension: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Info: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Measure: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Primary: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Secondary: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Success: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Warning: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   			},
//   			BrandElementStyle: &BrandElementStyleProperty{
//   				NavbarStyle: &NavbarStyleProperty{
//   					ContextualNavbar: &PaletteProperty{
//   						Background: jsii.String("background"),
//   						Foreground: jsii.String("foreground"),
//   					},
//   					GlobalNavbar: &PaletteProperty{
//   						Background: jsii.String("background"),
//   						Foreground: jsii.String("foreground"),
//   					},
//   				},
//   			},
//   		},
//   		BrandName: jsii.String("brandName"),
//   		Description: jsii.String("description"),
//   	},
//   	BrandId: jsii.String("brandId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-brand.html
//
type CfnBrandMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-brand.html#cfn-quicksight-brand-branddefinition
	//
	BrandDefinition interface{} `field:"optional" json:"brandDefinition" yaml:"brandDefinition"`
	// The ID of the QuickSight brand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-brand.html#cfn-quicksight-brand-brandid
	//
	BrandId *string `field:"optional" json:"brandId" yaml:"brandId"`
	// The tags for the brand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-brand.html#cfn-quicksight-brand-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

