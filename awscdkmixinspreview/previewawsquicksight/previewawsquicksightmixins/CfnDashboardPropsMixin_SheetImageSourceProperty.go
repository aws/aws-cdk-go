package previewawsquicksightmixins


// The source of the image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetImageSourceProperty := &SheetImageSourceProperty{
//   	SheetImageStaticFileSource: &SheetImageStaticFileSourceProperty{
//   		StaticFileId: jsii.String("staticFileId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagesource.html
//
type CfnDashboardPropsMixin_SheetImageSourceProperty struct {
	// The source of the static file that contains the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagesource.html#cfn-quicksight-dashboard-sheetimagesource-sheetimagestaticfilesource
	//
	SheetImageStaticFileSource interface{} `field:"optional" json:"sheetImageStaticFileSource" yaml:"sheetImageStaticFileSource"`
}

