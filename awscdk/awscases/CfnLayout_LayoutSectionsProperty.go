package awscases


// Ordered list containing different kinds of sections that can be added.
//
// A LayoutSections object can only contain one section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   layoutSectionsProperty := &LayoutSectionsProperty{
//   	Sections: []interface{}{
//   		&SectionProperty{
//   			FieldGroup: &FieldGroupProperty{
//   				Fields: []interface{}{
//   					&FieldItemProperty{
//   						Id: jsii.String("id"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-layoutsections.html
//
type CfnLayout_LayoutSectionsProperty struct {
	// Ordered list containing different kinds of sections that can be added.
	//
	// A LayoutSections object can only contain one section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-layoutsections.html#cfn-cases-layout-layoutsections-sections
	//
	Sections interface{} `field:"optional" json:"sections" yaml:"sections"`
}

