package awswellarchitected


// Properties for defining a `CfnLens`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLensProps := &CfnLensProps{
//   	JsonString: jsii.String("jsonString"),
//   	LensVersion: jsii.String("lensVersion"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-lens.html
//
type CfnLensProps struct {
	// The JSON representation of a lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-lens.html#cfn-wellarchitected-lens-jsonstring
	//
	JsonString *string `field:"optional" json:"jsonString" yaml:"jsonString"`
	// The version of the lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-lens.html#cfn-wellarchitected-lens-lensversion
	//
	LensVersion *string `field:"optional" json:"lensVersion" yaml:"lensVersion"`
	// The tags assigned to the lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-lens.html#cfn-wellarchitected-lens-tags
	//
	Tags *[]*CfnLens_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

