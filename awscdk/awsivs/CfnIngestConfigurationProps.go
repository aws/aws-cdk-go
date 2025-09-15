package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIngestConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIngestConfigurationProps := &CfnIngestConfigurationProps{
//   	IngestProtocol: jsii.String("ingestProtocol"),
//   	InsecureIngest: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	StageArn: jsii.String("stageArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserId: jsii.String("userId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html
//
type CfnIngestConfigurationProps struct {
	// Type of ingest protocol that the user employs for broadcasting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-ingestprotocol
	//
	// Default: - "RTMPS".
	//
	IngestProtocol *string `field:"optional" json:"ingestProtocol" yaml:"ingestProtocol"`
	// Whether the channel allows insecure RTMP ingest.
	//
	// Default: `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-insecureingest
	//
	// Default: - false.
	//
	InsecureIngest interface{} `field:"optional" json:"insecureIngest" yaml:"insecureIngest"`
	// Ingest name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-name
	//
	// Default: - "-".
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// ARN of the stage with which the IngestConfiguration is associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-stagearn
	//
	// Default: - "".
	//
	StageArn *string `field:"optional" json:"stageArn" yaml:"stageArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Customer-assigned name to help identify the participant using the IngestConfiguration;
	//
	// this can be used to link a participant to a user in the customer’s own systems. This can be any UTF-8 encoded text. *This field is exposed to all stage participants and should not be used for personally identifying, confidential, or sensitive information.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-userid
	//
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

