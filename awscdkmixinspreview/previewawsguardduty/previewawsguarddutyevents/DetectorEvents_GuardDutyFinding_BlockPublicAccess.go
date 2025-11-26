package previewawsguarddutyevents


// Type definition for BlockPublicAccess.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   blockPublicAccess := &BlockPublicAccess{
//   	BlockPublicAcls: []*string{
//   		jsii.String("blockPublicAcls"),
//   	},
//   	BlockPublicPolicy: []*string{
//   		jsii.String("blockPublicPolicy"),
//   	},
//   	IgnorePublicAcls: []*string{
//   		jsii.String("ignorePublicAcls"),
//   	},
//   	RestrictPublicBuckets: []*string{
//   		jsii.String("restrictPublicBuckets"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_BlockPublicAccess struct {
	// blockPublicAcls property.
	//
	// Specify an array of string values to match this event if the actual value of blockPublicAcls is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BlockPublicAcls *[]*string `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// blockPublicPolicy property.
	//
	// Specify an array of string values to match this event if the actual value of blockPublicPolicy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BlockPublicPolicy *[]*string `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// ignorePublicAcls property.
	//
	// Specify an array of string values to match this event if the actual value of ignorePublicAcls is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IgnorePublicAcls *[]*string `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// restrictPublicBuckets property.
	//
	// Specify an array of string values to match this event if the actual value of restrictPublicBuckets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RestrictPublicBuckets *[]*string `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

