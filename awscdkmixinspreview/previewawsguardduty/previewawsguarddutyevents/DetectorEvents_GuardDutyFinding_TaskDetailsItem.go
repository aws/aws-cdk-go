package previewawsguarddutyevents


// Type definition for TaskDetailsItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskDetailsItem := &TaskDetailsItem{
//   	Image: []*string{
//   		jsii.String("image"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_TaskDetailsItem struct {
	// image property.
	//
	// Specify an array of string values to match this event if the actual value of image is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Image *[]*string `field:"optional" json:"image" yaml:"image"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}

