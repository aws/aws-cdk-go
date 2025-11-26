package previewawss3events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Bucket aws.s3@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AdditionalEventData: &AdditionalEventData{
//   		AuthenticationMethod: []*string{
//   			jsii.String("authenticationMethod"),
//   		},
//   		BytesTransferredIn: []*string{
//   			jsii.String("bytesTransferredIn"),
//   		},
//   		BytesTransferredOut: []*string{
//   			jsii.String("bytesTransferredOut"),
//   		},
//   		CipherSuite: []*string{
//   			jsii.String("cipherSuite"),
//   		},
//   		ObjectRetentionInfo: &ObjectRetentionInfo{
//   			LegalHoldInfo: &LegalHoldInfo{
//   				IsUnderLegalHold: []*string{
//   					jsii.String("isUnderLegalHold"),
//   				},
//   				LastModifiedTime: []*string{
//   					jsii.String("lastModifiedTime"),
//   				},
//   			},
//   			RetentionInfo: &RetentionInfo{
//   				LastModifiedTime: []*string{
//   					jsii.String("lastModifiedTime"),
//   				},
//   				RetainUntilMode: []*string{
//   					jsii.String("retainUntilMode"),
//   				},
//   				RetainUntilTime: []*string{
//   					jsii.String("retainUntilTime"),
//   				},
//   			},
//   		},
//   		SignatureVersion: []*string{
//   			jsii.String("signatureVersion"),
//   		},
//   		XAmzId2: []*string{
//   			jsii.String("xAmzId2"),
//   		},
//   	},
//   	AwsRegion: []*string{
//   		jsii.String("awsRegion"),
//   	},
//   	ErrorCode: []*string{
//   		jsii.String("errorCode"),
//   	},
//   	ErrorMessage: []*string{
//   		jsii.String("errorMessage"),
//   	},
//   	EventCategory: []*string{
//   		jsii.String("eventCategory"),
//   	},
//   	EventId: []*string{
//   		jsii.String("eventId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventSource: []*string{
//   		jsii.String("eventSource"),
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	EventVersion: []*string{
//   		jsii.String("eventVersion"),
//   	},
//   	ManagementEvent: []*string{
//   		jsii.String("managementEvent"),
//   	},
//   	ReadOnly: []*string{
//   		jsii.String("readOnly"),
//   	},
//   	RecipientAccountId: []*string{
//   		jsii.String("recipientAccountId"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	RequestParameters: &RequestParameters{
//   		BucketName: []*string{
//   			jsii.String("bucketName"),
//   		},
//   		Host: []*string{
//   			jsii.String("host"),
//   		},
//   		Key: []*string{
//   			jsii.String("key"),
//   		},
//   		LegalHold: []*string{
//   			jsii.String("legalHold"),
//   		},
//   		Retention: []*string{
//   			jsii.String("retention"),
//   		},
//   	},
//   	Resources: []AwsapiCallViaCloudTrailItem{
//   		&AwsapiCallViaCloudTrailItem{
//   			AccountId: []*string{
//   				jsii.String("accountId"),
//   			},
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	ResponseElements: []*string{
//   		jsii.String("responseElements"),
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	TlsDetails: &TlsDetails{
//   		CipherSuite: []*string{
//   			jsii.String("cipherSuite"),
//   		},
//   		ClientProvidedHostHeader: []*string{
//   			jsii.String("clientProvidedHostHeader"),
//   		},
//   		TlsVersion: []*string{
//   			jsii.String("tlsVersion"),
//   		},
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	UserIdentity: &UserIdentity{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		SessionContext: &SessionContext{
//   			Attributes: &Attributes{
//   				CreationDate: []*string{
//   					jsii.String("creationDate"),
//   				},
//   				MfaAuthenticated: []*string{
//   					jsii.String("mfaAuthenticated"),
//   				},
//   			},
//   			SessionIssuer: &SessionIssuer{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Arn: []*string{
//   					jsii.String("arn"),
//   				},
//   				PrincipalId: []*string{
//   					jsii.String("principalId"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				UserName: []*string{
//   					jsii.String("userName"),
//   				},
//   			},
//   			WebIdFederationData: []*string{
//   				jsii.String("webIdFederationData"),
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	VpcEndpointId: []*string{
//   		jsii.String("vpcEndpointId"),
//   	},
//   }
//
// Experimental.
type BucketEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// additionalEventData property.
	//
	// Specify an array of string values to match this event if the actual value of additionalEventData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AdditionalEventData *BucketEvents_AWSAPICallViaCloudTrail_AdditionalEventData `field:"optional" json:"additionalEventData" yaml:"additionalEventData"`
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// errorMessage property.
	//
	// Specify an array of string values to match this event if the actual value of errorMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorMessage *[]*string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// eventCategory property.
	//
	// Specify an array of string values to match this event if the actual value of eventCategory is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventCategory *[]*string `field:"optional" json:"eventCategory" yaml:"eventCategory"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// managementEvent property.
	//
	// Specify an array of string values to match this event if the actual value of managementEvent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ManagementEvent *[]*string `field:"optional" json:"managementEvent" yaml:"managementEvent"`
	// readOnly property.
	//
	// Specify an array of string values to match this event if the actual value of readOnly is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReadOnly *[]*string `field:"optional" json:"readOnly" yaml:"readOnly"`
	// recipientAccountId property.
	//
	// Specify an array of string values to match this event if the actual value of recipientAccountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RecipientAccountId *[]*string `field:"optional" json:"recipientAccountId" yaml:"recipientAccountId"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *BucketEvents_AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// resources property.
	//
	// Specify an array of string values to match this event if the actual value of resources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Resources *[]*BucketEvents_AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem `field:"optional" json:"resources" yaml:"resources"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *[]*string `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// tlsDetails property.
	//
	// Specify an array of string values to match this event if the actual value of tlsDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TlsDetails *BucketEvents_AWSAPICallViaCloudTrail_TlsDetails `field:"optional" json:"tlsDetails" yaml:"tlsDetails"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *BucketEvents_AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
	// vpcEndpointId property.
	//
	// Specify an array of string values to match this event if the actual value of vpcEndpointId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcEndpointId *[]*string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

