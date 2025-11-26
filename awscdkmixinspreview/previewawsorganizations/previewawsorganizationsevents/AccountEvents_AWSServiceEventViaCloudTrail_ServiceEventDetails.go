package previewawsorganizationsevents


// Type definition for ServiceEventDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceEventDetails := &ServiceEventDetails{
//   	CreateAccountStatus: &CreateAccountStatus{
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		AccountName: []*string{
//   			jsii.String("accountName"),
//   		},
//   		CompletedTimestamp: []*string{
//   			jsii.String("completedTimestamp"),
//   		},
//   		Id: []*string{
//   			jsii.String("id"),
//   		},
//   		RequestedTimestamp: []*string{
//   			jsii.String("requestedTimestamp"),
//   		},
//   		State: []*string{
//   			jsii.String("state"),
//   		},
//   	},
//   }
//
// Experimental.
type AccountEvents_AWSServiceEventViaCloudTrail_ServiceEventDetails struct {
	// createAccountStatus property.
	//
	// Specify an array of string values to match this event if the actual value of createAccountStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreateAccountStatus *AccountEvents_AWSServiceEventViaCloudTrail_CreateAccountStatus `field:"optional" json:"createAccountStatus" yaml:"createAccountStatus"`
}

