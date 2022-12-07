package awsiottwinmaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyGroupProperty := &propertyGroupProperty{
//   	groupType: jsii.String("groupType"),
//   	propertyNames: []*string{
//   		jsii.String("propertyNames"),
//   	},
//   }
//
type CfnComponentType_PropertyGroupProperty struct {
	// `CfnComponentType.PropertyGroupProperty.GroupType`.
	GroupType *string `field:"optional" json:"groupType" yaml:"groupType"`
	// `CfnComponentType.PropertyGroupProperty.PropertyNames`.
	PropertyNames *[]*string `field:"optional" json:"propertyNames" yaml:"propertyNames"`
}

