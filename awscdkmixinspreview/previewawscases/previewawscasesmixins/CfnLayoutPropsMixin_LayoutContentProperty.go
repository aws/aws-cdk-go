package previewawscasesmixins


// Object to store union of different versions of layout content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   layoutContentProperty := &LayoutContentProperty{
//   	Basic: &BasicLayoutProperty{
//   		MoreInfo: &LayoutSectionsProperty{
//   			Sections: []interface{}{
//   				&SectionProperty{
//   					FieldGroup: &FieldGroupProperty{
//   						Fields: []interface{}{
//   							&FieldItemProperty{
//   								Id: jsii.String("id"),
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		TopPanel: &LayoutSectionsProperty{
//   			Sections: []interface{}{
//   				&SectionProperty{
//   					FieldGroup: &FieldGroupProperty{
//   						Fields: []interface{}{
//   							&FieldItemProperty{
//   								Id: jsii.String("id"),
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-layoutcontent.html
//
type CfnLayoutPropsMixin_LayoutContentProperty struct {
	// Content specific to `BasicLayout` type.
	//
	// It configures fields in the top panel and More Info tab of agent application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-layoutcontent.html#cfn-cases-layout-layoutcontent-basic
	//
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
}

