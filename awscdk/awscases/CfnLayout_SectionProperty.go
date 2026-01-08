package awscases


// This represents a sections within a panel or tab of the page layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionProperty := &SectionProperty{
//   	FieldGroup: &FieldGroupProperty{
//   		Fields: []interface{}{
//   			&FieldItemProperty{
//   				Id: jsii.String("id"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-section.html
//
type CfnLayout_SectionProperty struct {
	// Consists of a group of fields and associated properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-layout-section.html#cfn-cases-layout-section-fieldgroup
	//
	FieldGroup interface{} `field:"required" json:"fieldGroup" yaml:"fieldGroup"`
}

