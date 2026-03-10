package previewawsdevopsguruevents


// Type definition for AnomalySourceMetadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   anomalySourceMetadata := &AnomalySourceMetadata{
//   	Source: []*string{
//   		jsii.String("source"),
//   	},
//   	SourceResourceArn: []*string{
//   		jsii.String("sourceResourceArn"),
//   	},
//   	SourceResourceType: []*string{
//   		jsii.String("sourceResourceType"),
//   	},
//   }
//
// Experimental.
type DevOpsGuruInsightClosed_AnomalySourceMetadata struct {
	// source property.
	//
	// Specify an array of string values to match this event if the actual value of source is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Source *[]*string `field:"optional" json:"source" yaml:"source"`
	// sourceResourceArn property.
	//
	// Specify an array of string values to match this event if the actual value of sourceResourceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceResourceArn *[]*string `field:"optional" json:"sourceResourceArn" yaml:"sourceResourceArn"`
	// sourceResourceType property.
	//
	// Specify an array of string values to match this event if the actual value of sourceResourceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceResourceType *[]*string `field:"optional" json:"sourceResourceType" yaml:"sourceResourceType"`
}

