package awslogs


// This processor takes a list of objects that contain key fields, and converts them into a map of target keys.
//
// For more information about this processor including examples, see listToMap in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listToMapProperty := &ListToMapProperty{
//   	Key: jsii.String("key"),
//   	Source: jsii.String("source"),
//
//   	// the properties below are optional
//   	Flatten: jsii.Boolean(false),
//   	FlattenedElement: jsii.String("flattenedElement"),
//   	Target: jsii.String("target"),
//   	ValueKey: jsii.String("valueKey"),
//   }
//
type ListToMapProperty struct {
	// The key of the field to be extracted as keys in the generated map.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The key in the log event that has a list of objects that will be converted to a map.
	Source *string `field:"required" json:"source" yaml:"source"`
	// A Boolean value to indicate whether the list will be flattened into single items.
	// Default: false.
	//
	Flatten *bool `field:"optional" json:"flatten" yaml:"flatten"`
	// If you set flatten to true, use flattenedElement to specify which element, first or last, to keep.
	//
	// You must specify this parameter if flatten is true.
	// Default: - Must be specified if flatten is true and if flatten is false, has no effect.
	//
	FlattenedElement *string `field:"optional" json:"flattenedElement" yaml:"flattenedElement"`
	// The key of the field that will hold the generated map.
	// Default: - Stored at the root of the log event.
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// If this is specified, the values that you specify in this parameter will be extracted from the source objects and put into the values of the generated map.
	// Default: - Original objects in the source list will be put into the values of the generated map.
	//
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

