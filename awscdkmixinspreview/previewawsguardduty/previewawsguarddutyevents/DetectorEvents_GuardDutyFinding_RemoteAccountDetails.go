package previewawsguarddutyevents


// Type definition for RemoteAccountDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   remoteAccountDetails := &RemoteAccountDetails{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	Affiliated: []*string{
//   		jsii.String("affiliated"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_RemoteAccountDetails struct {
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// affiliated property.
	//
	// Specify an array of string values to match this event if the actual value of affiliated is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Affiliated *[]*string `field:"optional" json:"affiliated" yaml:"affiliated"`
}

