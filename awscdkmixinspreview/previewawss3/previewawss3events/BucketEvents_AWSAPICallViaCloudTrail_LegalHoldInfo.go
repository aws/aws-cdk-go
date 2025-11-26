package previewawss3events


// Type definition for LegalHoldInfo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   legalHoldInfo := &LegalHoldInfo{
//   	IsUnderLegalHold: []*string{
//   		jsii.String("isUnderLegalHold"),
//   	},
//   	LastModifiedTime: []*string{
//   		jsii.String("lastModifiedTime"),
//   	},
//   }
//
// Experimental.
type BucketEvents_AWSAPICallViaCloudTrail_LegalHoldInfo struct {
	// isUnderLegalHold property.
	//
	// Specify an array of string values to match this event if the actual value of isUnderLegalHold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IsUnderLegalHold *[]*string `field:"optional" json:"isUnderLegalHold" yaml:"isUnderLegalHold"`
	// lastModifiedTime property.
	//
	// Specify an array of string values to match this event if the actual value of lastModifiedTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedTime *[]*string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
}

