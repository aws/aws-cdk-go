package awscdk


// Options for the `stack.exportValue()` method.
//
// Example:
//   var stack Stack
//
//
//   stack.ExportValue(myBucket.BucketName, &ExportValueOptions{
//   	Name: jsii.String("TheAwesomeBucket"),
//   	Description: jsii.String("The name of an S3 bucket"),
//   })
//
type ExportValueOptions struct {
	// The description of the outputs.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the export to create.
	// Default: - A name is automatically chosen.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

