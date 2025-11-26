package previewawsbcmdataexportsmixins


// The tag structure that contains a tag key and value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceTagProperty := &ResourceTagProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-resourcetag.html
//
type CfnExportPropsMixin_ResourceTagProperty struct {
	// The key that's associated with the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-resourcetag.html#cfn-bcmdataexports-export-resourcetag-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value that's associated with the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-resourcetag.html#cfn-bcmdataexports-export-resourcetag-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

