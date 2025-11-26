package previewawsguarddutyevents


// Type definition for AccessControlList.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessControlList := &AccessControlList{
//   	AllowsPublicReadAccess: []*string{
//   		jsii.String("allowsPublicReadAccess"),
//   	},
//   	AllowsPublicWriteAccess: []*string{
//   		jsii.String("allowsPublicWriteAccess"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_AccessControlList struct {
	// allowsPublicReadAccess property.
	//
	// Specify an array of string values to match this event if the actual value of allowsPublicReadAccess is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AllowsPublicReadAccess *[]*string `field:"optional" json:"allowsPublicReadAccess" yaml:"allowsPublicReadAccess"`
	// allowsPublicWriteAccess property.
	//
	// Specify an array of string values to match this event if the actual value of allowsPublicWriteAccess is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AllowsPublicWriteAccess *[]*string `field:"optional" json:"allowsPublicWriteAccess" yaml:"allowsPublicWriteAccess"`
}

