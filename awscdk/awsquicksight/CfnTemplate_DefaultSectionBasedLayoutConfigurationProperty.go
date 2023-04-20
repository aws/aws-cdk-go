package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultSectionBasedLayoutConfigurationProperty := &DefaultSectionBasedLayoutConfigurationProperty{
//   	CanvasSizeOptions: &SectionBasedLayoutCanvasSizeOptionsProperty{
//   		PaperCanvasSizeOptions: &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   			PaperMargin: &SpacingProperty{
//   				Bottom: jsii.String("bottom"),
//   				Left: jsii.String("left"),
//   				Right: jsii.String("right"),
//   				Top: jsii.String("top"),
//   			},
//   			PaperOrientation: jsii.String("paperOrientation"),
//   			PaperSize: jsii.String("paperSize"),
//   		},
//   	},
//   }
//
type CfnTemplate_DefaultSectionBasedLayoutConfigurationProperty struct {
	// `CfnTemplate.DefaultSectionBasedLayoutConfigurationProperty.CanvasSizeOptions`.
	CanvasSizeOptions interface{} `field:"required" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
}

