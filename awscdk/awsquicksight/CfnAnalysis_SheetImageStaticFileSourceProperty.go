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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimagestaticfilesource.html
//
type CfnAnalysis_SheetImageStaticFileSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetimagestaticfilesource.html#cfn-quicksight-analysis-sheetimagestaticfilesource-staticfileid
	//
	StaticFileId *string `field:"required" json:"staticFileId" yaml:"staticFileId"`
}

