package previewawslakeformationmixins


// The LF-tag key and values attached to a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lFTagProperty := &LFTagProperty{
//   	TagKey: jsii.String("tagKey"),
//   	TagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftag.html
//
type CfnPrincipalPermissionsPropsMixin_LFTagProperty struct {
	// The key-name for the LF-tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftag.html#cfn-lakeformation-principalpermissions-lftag-tagkey
	//
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftag.html#cfn-lakeformation-principalpermissions-lftag-tagvalues
	//
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}

