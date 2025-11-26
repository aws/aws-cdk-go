package previewawsguarddutyevents


// Type definition for InstanceDetailsItem_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceDetailsItem1 := &InstanceDetailsItem1{
//   	ProductCodeId: []*string{
//   		jsii.String("productCodeId"),
//   	},
//   	ProductCodeType: []*string{
//   		jsii.String("productCodeType"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_InstanceDetailsItem1 struct {
	// productCodeId property.
	//
	// Specify an array of string values to match this event if the actual value of productCodeId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProductCodeId *[]*string `field:"optional" json:"productCodeId" yaml:"productCodeId"`
	// productCodeType property.
	//
	// Specify an array of string values to match this event if the actual value of productCodeType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProductCodeType *[]*string `field:"optional" json:"productCodeType" yaml:"productCodeType"`
}

