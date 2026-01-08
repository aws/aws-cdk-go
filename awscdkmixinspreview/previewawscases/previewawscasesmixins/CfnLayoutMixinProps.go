package previewawscasesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLayoutPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLayoutMixinProps := &CfnLayoutMixinProps{
//   	Content: &LayoutContentProperty{
//   		Basic: &BasicLayoutProperty{
//   			MoreInfo: &LayoutSectionsProperty{
//   				Sections: []interface{}{
//   					&SectionProperty{
//   						FieldGroup: &FieldGroupProperty{
//   							Fields: []interface{}{
//   								&FieldItemProperty{
//   									Id: jsii.String("id"),
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   			TopPanel: &LayoutSectionsProperty{
//   				Sections: []interface{}{
//   					&SectionProperty{
//   						FieldGroup: &FieldGroupProperty{
//   							Fields: []interface{}{
//   								&FieldItemProperty{
//   									Id: jsii.String("id"),
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html
//
type CfnLayoutMixinProps struct {
	// Object to store union of different versions of layout content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html#cfn-cases-layout-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// The unique identifier of the Cases domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html#cfn-cases-layout-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The name of the layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html#cfn-cases-layout-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html#cfn-cases-layout-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

