package previewawsguarddutyevents


// Type definition for Owner.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   owner := &Owner{
//   	Id: []*string{
//   		jsii.String("id"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_Owner struct {
	// id property.
	//
	// Specify an array of string values to match this event if the actual value of id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
}

