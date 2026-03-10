package previewawsmanagedblockchainevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.managedblockchain@ManagedBlockchainProposalStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedBlockchainProposalStatusChangeProps := &ManagedBlockchainProposalStatusChangeProps{
//   	Description: []*string{
//   		jsii.String("description"),
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
//   	ExpirationDate: []*string{
//   		jsii.String("expirationDate"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	NetworkId: []*string{
//   		jsii.String("networkId"),
//   	},
//   	ProposalId: []*string{
//   		jsii.String("proposalId"),
//   	},
//   	ProposedByMemberId: []*string{
//   		jsii.String("proposedByMemberId"),
//   	},
//   	ProposedByMemberName: []*string{
//   		jsii.String("proposedByMemberName"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type ManagedBlockchainProposalStatusChange_ManagedBlockchainProposalStatusChangeProps struct {
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// expirationDate property.
	//
	// Specify an array of string values to match this event if the actual value of expirationDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExpirationDate *[]*string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
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
	// proposalId property.
	//
	// Specify an array of string values to match this event if the actual value of proposalId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProposalId *[]*string `field:"optional" json:"proposalId" yaml:"proposalId"`
	// proposedByMemberId property.
	//
	// Specify an array of string values to match this event if the actual value of proposedByMemberId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProposedByMemberId *[]*string `field:"optional" json:"proposedByMemberId" yaml:"proposedByMemberId"`
	// proposedByMemberName property.
	//
	// Specify an array of string values to match this event if the actual value of proposedByMemberName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProposedByMemberName *[]*string `field:"optional" json:"proposedByMemberName" yaml:"proposedByMemberName"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

