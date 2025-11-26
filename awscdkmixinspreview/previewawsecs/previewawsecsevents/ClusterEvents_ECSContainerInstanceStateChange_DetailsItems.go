package previewawsecsevents


// Type definition for detailsItems.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   detailsItems := &DetailsItems{
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSContainerInstanceStateChange_DetailsItems struct {
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// value property.
	//
	// Specify an array of string values to match this event if the actual value of value is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

