package previewawsguarddutyevents


// Type definition for DefaultServerSideEncryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultServerSideEncryption := &DefaultServerSideEncryption{
//   	EncryptionType: []*string{
//   		jsii.String("encryptionType"),
//   	},
//   	KmsMasterKeyArn: []*string{
//   		jsii.String("kmsMasterKeyArn"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_DefaultServerSideEncryption struct {
	// encryptionType property.
	//
	// Specify an array of string values to match this event if the actual value of encryptionType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EncryptionType *[]*string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// kmsMasterKeyArn property.
	//
	// Specify an array of string values to match this event if the actual value of kmsMasterKeyArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KmsMasterKeyArn *[]*string `field:"optional" json:"kmsMasterKeyArn" yaml:"kmsMasterKeyArn"`
}

