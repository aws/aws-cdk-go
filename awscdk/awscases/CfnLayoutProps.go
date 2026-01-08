package awscases

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLayout`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLayoutProps := &CfnLayoutProps{
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
//
//   							// the properties below are optional
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
//
//   							// the properties below are optional
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DomainId: jsii.String("domainId"),
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
type CfnLayoutProps struct {
	// Object to store union of different versions of layout content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html#cfn-cases-layout-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// The name of the layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html#cfn-cases-layout-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier of the Cases domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html#cfn-cases-layout-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-layout.html#cfn-cases-layout-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

