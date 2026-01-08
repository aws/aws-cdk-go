package previewawscasesmixins


// Content specific to `BasicLayout` type.
//
// It configures fields in the top panel and More Info tab of agent application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   basicLayoutProperty := &BasicLayoutProperty{
//   	MoreInfo: &LayoutSectionsProperty{
//   		Sections: []interface{}{
//   			&SectionProperty{
//   				FieldGroup: &FieldGroupProperty{
//   					Fields: []interface{}{
//   						&FieldItemProperty{
//   							Id: jsii.String("id"),
//   						},
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	TopPanel: &LayoutSectionsProperty{
//   		Sections: []interface{}{
//   			&SectionProperty{
//   				FieldGroup: &FieldGroupProperty{
//   					Fields: []interface{}{
//   						&FieldItemProperty{
//   							Id: jsii.String("id"),
//   						},
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-basiclayout.html
//
type CfnLayoutPropsMixin_BasicLayoutProperty struct {
	// This represents sections in a tab of the page layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-basiclayout.html#cfn-cases-layout-basiclayout-moreinfo
	//
	MoreInfo interface{} `field:"optional" json:"moreInfo" yaml:"moreInfo"`
	// This represents sections in a panel of the page layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-basiclayout.html#cfn-cases-layout-basiclayout-toppanel
	//
	TopPanel interface{} `field:"optional" json:"topPanel" yaml:"topPanel"`
}

