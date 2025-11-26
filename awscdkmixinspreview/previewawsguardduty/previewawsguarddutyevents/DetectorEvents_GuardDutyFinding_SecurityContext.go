package previewawsguarddutyevents


// Type definition for SecurityContext.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityContext := &SecurityContext{
//   	Privileged: []*string{
//   		jsii.String("privileged"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_SecurityContext struct {
	// privileged property.
	//
	// Specify an array of string values to match this event if the actual value of privileged is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Privileged *[]*string `field:"optional" json:"privileged" yaml:"privileged"`
}

