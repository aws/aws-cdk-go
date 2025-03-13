package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAnalysisTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnalysisTemplateProps := &CfnAnalysisTemplateProps{
//   	Format: jsii.String("format"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	Source: &AnalysisSourceProperty{
//   		Text: jsii.String("text"),
//   	},
//
//   	// the properties below are optional
//   	AnalysisParameters: []interface{}{
//   		&AnalysisParameterProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			DefaultValue: jsii.String("defaultValue"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html
//
type CfnAnalysisTemplateProps struct {
	// The format of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// The identifier for a membership resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-membershipidentifier
	//
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// The name of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// The parameters of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-analysisparameters
	//
	AnalysisParameters interface{} `field:"optional" json:"analysisParameters" yaml:"analysisParameters"`
	// The description of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

