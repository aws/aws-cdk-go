package previewawskmsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Key aws.kms@KMSCMKDeletion event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kMSCMKDeletionProps := &KMSCMKDeletionProps{
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
//   	KeyId: []*string{
//   		jsii.String("keyId"),
//   	},
//   }
//
// Experimental.
type KeyEvents_KMSCMKDeletion_KMSCMKDeletionProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// key-id property.
	//
	// Specify an array of string values to match this event if the actual value of key-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Key reference.
	//
	// Experimental.
	KeyId *[]*string `field:"optional" json:"keyId" yaml:"keyId"`
}

