package previewawsguarddutyevents


// Type definition for UserAgent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userAgent := &UserAgent{
//   	FullUserAgent: []*string{
//   		jsii.String("fullUserAgent"),
//   	},
//   	UserAgentCategory: []*string{
//   		jsii.String("userAgentCategory"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_UserAgent struct {
	// fullUserAgent property.
	//
	// Specify an array of string values to match this event if the actual value of fullUserAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FullUserAgent *[]*string `field:"optional" json:"fullUserAgent" yaml:"fullUserAgent"`
	// userAgentCategory property.
	//
	// Specify an array of string values to match this event if the actual value of userAgentCategory is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgentCategory *[]*string `field:"optional" json:"userAgentCategory" yaml:"userAgentCategory"`
}

