package previewawsbatchevents


// Type definition for MountPoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mountPoint := &MountPoint{
//   	ContainerPath: []*string{
//   		jsii.String("containerPath"),
//   	},
//   	ReadOnly: []*string{
//   		jsii.String("readOnly"),
//   	},
//   	SourceVolume: []*string{
//   		jsii.String("sourceVolume"),
//   	},
//   }
//
// Experimental.
type BatchJobStateChange_MountPoint struct {
	// containerPath property.
	//
	// Specify an array of string values to match this event if the actual value of containerPath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerPath *[]*string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// readOnly property.
	//
	// Specify an array of string values to match this event if the actual value of readOnly is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReadOnly *[]*string `field:"optional" json:"readOnly" yaml:"readOnly"`
	// sourceVolume property.
	//
	// Specify an array of string values to match this event if the actual value of sourceVolume is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceVolume *[]*string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

