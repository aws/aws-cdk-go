package previewawss3events


// Type definition for AdditionalEventData.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   additionalEventData := &AdditionalEventData{
//   	AuthenticationMethod: []*string{
//   		jsii.String("authenticationMethod"),
//   	},
//   	BytesTransferredIn: []*string{
//   		jsii.String("bytesTransferredIn"),
//   	},
//   	BytesTransferredOut: []*string{
//   		jsii.String("bytesTransferredOut"),
//   	},
//   	CipherSuite: []*string{
//   		jsii.String("cipherSuite"),
//   	},
//   	ObjectRetentionInfo: &ObjectRetentionInfo{
//   		LegalHoldInfo: &LegalHoldInfo{
//   			IsUnderLegalHold: []*string{
//   				jsii.String("isUnderLegalHold"),
//   			},
//   			LastModifiedTime: []*string{
//   				jsii.String("lastModifiedTime"),
//   			},
//   		},
//   		RetentionInfo: &RetentionInfo{
//   			LastModifiedTime: []*string{
//   				jsii.String("lastModifiedTime"),
//   			},
//   			RetainUntilMode: []*string{
//   				jsii.String("retainUntilMode"),
//   			},
//   			RetainUntilTime: []*string{
//   				jsii.String("retainUntilTime"),
//   			},
//   		},
//   	},
//   	SignatureVersion: []*string{
//   		jsii.String("signatureVersion"),
//   	},
//   	XAmzId2: []*string{
//   		jsii.String("xAmzId2"),
//   	},
//   }
//
// Experimental.
type BucketEvents_AWSAPICallViaCloudTrail_AdditionalEventData struct {
	// AuthenticationMethod property.
	//
	// Specify an array of string values to match this event if the actual value of AuthenticationMethod is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthenticationMethod *[]*string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// bytesTransferredIn property.
	//
	// Specify an array of string values to match this event if the actual value of bytesTransferredIn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BytesTransferredIn *[]*string `field:"optional" json:"bytesTransferredIn" yaml:"bytesTransferredIn"`
	// bytesTransferredOut property.
	//
	// Specify an array of string values to match this event if the actual value of bytesTransferredOut is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BytesTransferredOut *[]*string `field:"optional" json:"bytesTransferredOut" yaml:"bytesTransferredOut"`
	// CipherSuite property.
	//
	// Specify an array of string values to match this event if the actual value of CipherSuite is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CipherSuite *[]*string `field:"optional" json:"cipherSuite" yaml:"cipherSuite"`
	// objectRetentionInfo property.
	//
	// Specify an array of string values to match this event if the actual value of objectRetentionInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ObjectRetentionInfo *BucketEvents_AWSAPICallViaCloudTrail_ObjectRetentionInfo `field:"optional" json:"objectRetentionInfo" yaml:"objectRetentionInfo"`
	// SignatureVersion property.
	//
	// Specify an array of string values to match this event if the actual value of SignatureVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SignatureVersion *[]*string `field:"optional" json:"signatureVersion" yaml:"signatureVersion"`
	// x-amz-id-2 property.
	//
	// Specify an array of string values to match this event if the actual value of x-amz-id-2 is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	XAmzId2 *[]*string `field:"optional" json:"xAmzId2" yaml:"xAmzId2"`
}

