package previewawslogsevents


// Type definition for SessionContext.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sessionContext := &SessionContext{
//   	Attributes: &Attributes{
//   		CreationDate: []*string{
//   			jsii.String("creationDate"),
//   		},
//   		MfaAuthenticated: []*string{
//   			jsii.String("mfaAuthenticated"),
//   		},
//   	},
//   	SessionIssuer: &SessionIssuer{
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   		UserName: []*string{
//   			jsii.String("userName"),
//   		},
//   	},
//   }
//
// Experimental.
type LogGroupEvents_AWSAPICallViaCloudTrail_SessionContext struct {
	// attributes property.
	//
	// Specify an array of string values to match this event if the actual value of attributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attributes *LogGroupEvents_AWSAPICallViaCloudTrail_Attributes `field:"optional" json:"attributes" yaml:"attributes"`
	// sessionIssuer property.
	//
	// Specify an array of string values to match this event if the actual value of sessionIssuer is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SessionIssuer *LogGroupEvents_AWSAPICallViaCloudTrail_SessionIssuer `field:"optional" json:"sessionIssuer" yaml:"sessionIssuer"`
}

