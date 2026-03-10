package previewawscodebuildevents


// Type definition for Badge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   badge := &Badge{
//   	BadgeEnabled: []*string{
//   		jsii.String("badgeEnabled"),
//   	},
//   	BadgeRequestUrl: []*string{
//   		jsii.String("badgeRequestUrl"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Badge struct {
	// badgeEnabled property.
	//
	// Specify an array of string values to match this event if the actual value of badgeEnabled is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BadgeEnabled *[]*string `field:"optional" json:"badgeEnabled" yaml:"badgeEnabled"`
	// badgeRequestUrl property.
	//
	// Specify an array of string values to match this event if the actual value of badgeRequestUrl is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BadgeRequestUrl *[]*string `field:"optional" json:"badgeRequestUrl" yaml:"badgeRequestUrl"`
}

