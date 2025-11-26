package previewawsorganizationsevents


// Type definition for UserIdentity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userIdentity := &UserIdentity{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	InvokedBy: []*string{
//   		jsii.String("invokedBy"),
//   	},
//   }
//
// Experimental.
type AccountEvents_AWSServiceEventViaCloudTrail_UserIdentity struct {
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// invokedBy property.
	//
	// Specify an array of string values to match this event if the actual value of invokedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InvokedBy *[]*string `field:"optional" json:"invokedBy" yaml:"invokedBy"`
}

