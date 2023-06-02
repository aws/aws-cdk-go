package awsiottwinmaker


// The property group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyGroupProperty := &PropertyGroupProperty{
//   	GroupType: jsii.String("groupType"),
//   	PropertyNames: []*string{
//   		jsii.String("propertyNames"),
//   	},
//   }
//
type CfnEntity_PropertyGroupProperty struct {
	// The group type.
	GroupType *string `field:"optional" json:"groupType" yaml:"groupType"`
	// The property names.
	PropertyNames *[]*string `field:"optional" json:"propertyNames" yaml:"propertyNames"`
}

