package awswellarchitected


// Properties for CfnReviewTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnReviewTemplateMixinProps := &CfnReviewTemplateMixinProps{
//   	Description: jsii.String("description"),
//   	Lenses: []*string{
//   		jsii.String("lenses"),
//   	},
//   	Notes: jsii.String("notes"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateName: jsii.String("templateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html
//
type CfnReviewTemplateMixinProps struct {
	// The review template description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The lenses applied to the review template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-lenses
	//
	Lenses *[]*string `field:"optional" json:"lenses" yaml:"lenses"`
	// The notes associated with the review template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-notes
	//
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The tags assigned to the review template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-tags
	//
	Tags *[]*CfnReviewTemplatePropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The name of the review template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

