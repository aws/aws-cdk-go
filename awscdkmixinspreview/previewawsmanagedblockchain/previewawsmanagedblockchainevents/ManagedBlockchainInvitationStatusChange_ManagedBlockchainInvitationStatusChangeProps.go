package previewawsmanagedblockchainevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.managedblockchain@ManagedBlockchainInvitationStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedBlockchainInvitationStatusChangeProps := &ManagedBlockchainInvitationStatusChangeProps{
//   	AcceptedAt: []*string{
//   		jsii.String("acceptedAt"),
//   	},
//   	CreatedAt: []*string{
//   		jsii.String("createdAt"),
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
//   	ExpiresAt: []*string{
//   		jsii.String("expiresAt"),
//   	},
//   	InvitationId: []*string{
//   		jsii.String("invitationId"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	NetworkId: []*string{
//   		jsii.String("networkId"),
//   	},
//   	RejectedAt: []*string{
//   		jsii.String("rejectedAt"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type ManagedBlockchainInvitationStatusChange_ManagedBlockchainInvitationStatusChangeProps struct {
	// acceptedAt property.
	//
	// Specify an array of string values to match this event if the actual value of acceptedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AcceptedAt *[]*string `field:"optional" json:"acceptedAt" yaml:"acceptedAt"`
	// createdAt property.
	//
	// Specify an array of string values to match this event if the actual value of createdAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// expiresAt property.
	//
	// Specify an array of string values to match this event if the actual value of expiresAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExpiresAt *[]*string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// invitationId property.
	//
	// Specify an array of string values to match this event if the actual value of invitationId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InvitationId *[]*string `field:"optional" json:"invitationId" yaml:"invitationId"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// networkId property.
	//
	// Specify an array of string values to match this event if the actual value of networkId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkId *[]*string `field:"optional" json:"networkId" yaml:"networkId"`
	// rejectedAt property.
	//
	// Specify an array of string values to match this event if the actual value of rejectedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RejectedAt *[]*string `field:"optional" json:"rejectedAt" yaml:"rejectedAt"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

