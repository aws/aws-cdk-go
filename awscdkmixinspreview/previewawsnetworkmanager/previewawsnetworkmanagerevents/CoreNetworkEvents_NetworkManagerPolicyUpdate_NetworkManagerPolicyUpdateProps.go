package previewawsnetworkmanagerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for CoreNetwork aws.networkmanager@NetworkManagerPolicyUpdate event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkManagerPolicyUpdateProps := &NetworkManagerPolicyUpdateProps{
//   	ChangeDescription: []*string{
//   		jsii.String("changeDescription"),
//   	},
//   	ChangeType: []*string{
//   		jsii.String("changeType"),
//   	},
//   	CoreNetworkArn: []*string{
//   		jsii.String("coreNetworkArn"),
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
//   	PolicyVersionId: []*string{
//   		jsii.String("policyVersionId"),
//   	},
//   }
//
// Experimental.
type CoreNetworkEvents_NetworkManagerPolicyUpdate_NetworkManagerPolicyUpdateProps struct {
	// changeDescription property.
	//
	// Specify an array of string values to match this event if the actual value of changeDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChangeDescription *[]*string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// changeType property.
	//
	// Specify an array of string values to match this event if the actual value of changeType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ChangeType *[]*string `field:"optional" json:"changeType" yaml:"changeType"`
	// coreNetworkArn property.
	//
	// Specify an array of string values to match this event if the actual value of coreNetworkArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the CoreNetwork reference.
	//
	// Experimental.
	CoreNetworkArn *[]*string `field:"optional" json:"coreNetworkArn" yaml:"coreNetworkArn"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// policyVersionId property.
	//
	// Specify an array of string values to match this event if the actual value of policyVersionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyVersionId *[]*string `field:"optional" json:"policyVersionId" yaml:"policyVersionId"`
}

