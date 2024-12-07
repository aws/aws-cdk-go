package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetImageStaticFileSourceProperty := &SheetImageStaticFileSourceProperty{
//   	StaticFileId: jsii.String("staticFileId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagestaticfilesource.html
//
type CfnTemplate_SheetImageStaticFileSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetimagestaticfilesource.html#cfn-quicksight-template-sheetimagestaticfilesource-staticfileid
	//
	StaticFileId *string `field:"required" json:"staticFileId" yaml:"staticFileId"`
}

