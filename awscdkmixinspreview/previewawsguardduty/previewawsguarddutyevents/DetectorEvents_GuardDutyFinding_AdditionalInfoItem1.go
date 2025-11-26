package previewawsguarddutyevents


// Type definition for AdditionalInfoItem_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   additionalInfoItem1 := &AdditionalInfoItem1{
//   	AccessKeyId: []*string{
//   		jsii.String("accessKeyId"),
//   	},
//   	IpAddressV4: []*string{
//   		jsii.String("ipAddressV4"),
//   	},
//   	PrincipalId: []*string{
//   		jsii.String("principalId"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_AdditionalInfoItem1 struct {
	// accessKeyId property.
	//
	// Specify an array of string values to match this event if the actual value of accessKeyId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccessKeyId *[]*string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// ipAddressV4 property.
	//
	// Specify an array of string values to match this event if the actual value of ipAddressV4 is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IpAddressV4 *[]*string `field:"optional" json:"ipAddressV4" yaml:"ipAddressV4"`
	// principalId property.
	//
	// Specify an array of string values to match this event if the actual value of principalId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrincipalId *[]*string `field:"optional" json:"principalId" yaml:"principalId"`
}

