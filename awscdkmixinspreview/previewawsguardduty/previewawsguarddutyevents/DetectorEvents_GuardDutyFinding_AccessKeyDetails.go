package previewawsguarddutyevents


// Type definition for AccessKeyDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessKeyDetails := &AccessKeyDetails{
//   	AccessKeyId: []*string{
//   		jsii.String("accessKeyId"),
//   	},
//   	PrincipalId: []*string{
//   		jsii.String("principalId"),
//   	},
//   	UserName: []*string{
//   		jsii.String("userName"),
//   	},
//   	UserType: []*string{
//   		jsii.String("userType"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_AccessKeyDetails struct {
	// accessKeyId property.
	//
	// Specify an array of string values to match this event if the actual value of accessKeyId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccessKeyId *[]*string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// principalId property.
	//
	// Specify an array of string values to match this event if the actual value of principalId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrincipalId *[]*string `field:"optional" json:"principalId" yaml:"principalId"`
	// userName property.
	//
	// Specify an array of string values to match this event if the actual value of userName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserName *[]*string `field:"optional" json:"userName" yaml:"userName"`
	// userType property.
	//
	// Specify an array of string values to match this event if the actual value of userType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserType *[]*string `field:"optional" json:"userType" yaml:"userType"`
}

