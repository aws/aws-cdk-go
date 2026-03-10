package previewawsdevopsguruevents


// Type definition for Tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagType := &TagType{
//   	AppBoundaryKey: []*string{
//   		jsii.String("appBoundaryKey"),
//   	},
//   	TagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewAnomalyAssociation_TagType struct {
	// appBoundaryKey property.
	//
	// Specify an array of string values to match this event if the actual value of appBoundaryKey is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AppBoundaryKey *[]*string `field:"optional" json:"appBoundaryKey" yaml:"appBoundaryKey"`
	// tagValues property.
	//
	// Specify an array of string values to match this event if the actual value of tagValues is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}

