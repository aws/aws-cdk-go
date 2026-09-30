package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   geospatialStaticFileSourceProperty := &GeospatialStaticFileSourceProperty{
//   	StaticFileId: jsii.String("staticFileId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialstaticfilesource.html
//
type CfnTemplatePropsMixin_GeospatialStaticFileSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialstaticfilesource.html#cfn-quicksight-template-geospatialstaticfilesource-staticfileid
	//
	StaticFileId *string `field:"optional" json:"staticFileId" yaml:"staticFileId"`
}

