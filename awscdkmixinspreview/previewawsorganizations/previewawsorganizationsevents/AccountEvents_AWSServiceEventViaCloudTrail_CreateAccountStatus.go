package previewawsorganizationsevents


// Type definition for CreateAccountStatus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createAccountStatus := &CreateAccountStatus{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	AccountName: []*string{
//   		jsii.String("accountName"),
//   	},
//   	CompletedTimestamp: []*string{
//   		jsii.String("completedTimestamp"),
//   	},
//   	Id: []*string{
//   		jsii.String("id"),
//   	},
//   	RequestedTimestamp: []*string{
//   		jsii.String("requestedTimestamp"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   }
//
// Experimental.
type AccountEvents_AWSServiceEventViaCloudTrail_CreateAccountStatus struct {
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Account reference.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// accountName property.
	//
	// Specify an array of string values to match this event if the actual value of accountName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountName *[]*string `field:"optional" json:"accountName" yaml:"accountName"`
	// completedTimestamp property.
	//
	// Specify an array of string values to match this event if the actual value of completedTimestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CompletedTimestamp *[]*string `field:"optional" json:"completedTimestamp" yaml:"completedTimestamp"`
	// id property.
	//
	// Specify an array of string values to match this event if the actual value of id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// requestedTimestamp property.
	//
	// Specify an array of string values to match this event if the actual value of requestedTimestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestedTimestamp *[]*string `field:"optional" json:"requestedTimestamp" yaml:"requestedTimestamp"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
}

