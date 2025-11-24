package mixinsawsquicksight


// The source of the static file that contains the image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetImageStaticFileSourceProperty := &SheetImageStaticFileSourceProperty{
//   	StaticFileId: jsii.String("staticFileId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagestaticfilesource.html
//
type CfnDashboardPropsMixin_SheetImageStaticFileSourceProperty struct {
	// The ID of the static file that contains the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetimagestaticfilesource.html#cfn-quicksight-dashboard-sheetimagestaticfilesource-staticfileid
	//
	StaticFileId *string `field:"optional" json:"staticFileId" yaml:"staticFileId"`
}

