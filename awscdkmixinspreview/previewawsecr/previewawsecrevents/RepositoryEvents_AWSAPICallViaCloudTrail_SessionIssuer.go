package previewawsecrevents


// Type definition for SessionIssuer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sessionIssuer := &SessionIssuer{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	PrincipalId: []*string{
//   		jsii.String("principalId"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   	UserName: []*string{
//   		jsii.String("userName"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_SessionIssuer struct {
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
	// principalId property.
	//
	// Specify an array of string values to match this event if the actual value of principalId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrincipalId *[]*string `field:"optional" json:"principalId" yaml:"principalId"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
	// userName property.
	//
	// Specify an array of string values to match this event if the actual value of userName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserName *[]*string `field:"optional" json:"userName" yaml:"userName"`
}

