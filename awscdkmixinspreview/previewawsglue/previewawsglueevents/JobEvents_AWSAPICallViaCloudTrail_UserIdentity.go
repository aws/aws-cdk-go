package previewawsglueevents


// Type definition for UserIdentity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userIdentity := &UserIdentity{
//   	AccessKeyId: []*string{
//   		jsii.String("accessKeyId"),
//   	},
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	InvokedBy: []*string{
//   		jsii.String("invokedBy"),
//   	},
//   	PrincipalId: []*string{
//   		jsii.String("principalId"),
//   	},
//   	SessionContext: &SessionContext{
//   		Attributes: &Attributes{
//   			CreationDate: []*string{
//   				jsii.String("creationDate"),
//   			},
//   			MfaAuthenticated: []*string{
//   				jsii.String("mfaAuthenticated"),
//   			},
//   		},
//   		SessionIssuer: &SessionIssuer{
//   			AccountId: []*string{
//   				jsii.String("accountId"),
//   			},
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			PrincipalId: []*string{
//   				jsii.String("principalId"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   			UserName: []*string{
//   				jsii.String("userName"),
//   			},
//   		},
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type JobEvents_AWSAPICallViaCloudTrail_UserIdentity struct {
	// accessKeyId property.
	//
	// Specify an array of string values to match this event if the actual value of accessKeyId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccessKeyId *[]*string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// invokedBy property.
	//
	// Specify an array of string values to match this event if the actual value of invokedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InvokedBy *[]*string `field:"optional" json:"invokedBy" yaml:"invokedBy"`
	// principalId property.
	//
	// Specify an array of string values to match this event if the actual value of principalId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrincipalId *[]*string `field:"optional" json:"principalId" yaml:"principalId"`
	// sessionContext property.
	//
	// Specify an array of string values to match this event if the actual value of sessionContext is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SessionContext *JobEvents_AWSAPICallViaCloudTrail_SessionContext `field:"optional" json:"sessionContext" yaml:"sessionContext"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

