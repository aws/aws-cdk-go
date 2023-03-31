package awsamplifyuibuilder


// The `SortProperty` property specifies how to sort the data that you bind to a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sortPropertyProperty := &sortPropertyProperty{
//   	direction: jsii.String("direction"),
//   	field: jsii.String("field"),
//   }
//
type CfnComponent_SortPropertyProperty struct {
	// The direction of the sort, either ascending or descending.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The field to perform the sort on.
	Field *string `field:"required" json:"field" yaml:"field"`
}

