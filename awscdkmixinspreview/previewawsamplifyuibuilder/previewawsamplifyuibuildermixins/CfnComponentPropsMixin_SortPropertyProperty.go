package previewawsamplifyuibuildermixins


// The `SortProperty` property specifies how to sort the data that you bind to a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sortPropertyProperty := &SortPropertyProperty{
//   	Direction: jsii.String("direction"),
//   	Field: jsii.String("field"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-sortproperty.html
//
type CfnComponentPropsMixin_SortPropertyProperty struct {
	// The direction of the sort, either ascending or descending.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-sortproperty.html#cfn-amplifyuibuilder-component-sortproperty-direction
	//
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// The field to perform the sort on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-sortproperty.html#cfn-amplifyuibuilder-component-sortproperty-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
}

