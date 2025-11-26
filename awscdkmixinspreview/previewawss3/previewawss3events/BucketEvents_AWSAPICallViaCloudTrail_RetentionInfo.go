package previewawss3events


// Type definition for RetentionInfo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retentionInfo := &RetentionInfo{
//   	LastModifiedTime: []*string{
//   		jsii.String("lastModifiedTime"),
//   	},
//   	RetainUntilMode: []*string{
//   		jsii.String("retainUntilMode"),
//   	},
//   	RetainUntilTime: []*string{
//   		jsii.String("retainUntilTime"),
//   	},
//   }
//
// Experimental.
type BucketEvents_AWSAPICallViaCloudTrail_RetentionInfo struct {
	// lastModifiedTime property.
	//
	// Specify an array of string values to match this event if the actual value of lastModifiedTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedTime *[]*string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// retainUntilMode property.
	//
	// Specify an array of string values to match this event if the actual value of retainUntilMode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RetainUntilMode *[]*string `field:"optional" json:"retainUntilMode" yaml:"retainUntilMode"`
	// retainUntilTime property.
	//
	// Specify an array of string values to match this event if the actual value of retainUntilTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RetainUntilTime *[]*string `field:"optional" json:"retainUntilTime" yaml:"retainUntilTime"`
}

