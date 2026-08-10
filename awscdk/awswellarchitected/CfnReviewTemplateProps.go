package awswellarchitected


// Properties for defining a `CfnReviewTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReviewTemplateProps := &CfnReviewTemplateProps{
//   	Description: jsii.String("description"),
//   	Lenses: []*string{
//   		jsii.String("lenses"),
//   	},
//   	TemplateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	Notes: jsii.String("notes"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html
//
type CfnReviewTemplateProps struct {
	// The review template description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The lenses applied to the review template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-lenses
	//
	Lenses *[]*string `field:"required" json:"lenses" yaml:"lenses"`
	// The name of the review template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-templatename
	//
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// The notes associated with the review template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-notes
	//
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The tags assigned to the review template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-reviewtemplate.html#cfn-wellarchitected-reviewtemplate-tags
	//
	Tags *[]*CfnReviewTemplate_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

