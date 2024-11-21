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
	// Ingest Protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-ingestprotocol
	//
	// Default: - "RTMPS".
	//
	IngestProtocol *string `field:"optional" json:"ingestProtocol" yaml:"ingestProtocol"`
	// Whether ingest configuration allows insecure ingest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-insecureingest
	//
	// Default: - false.
	//
	InsecureIngest interface{} `field:"optional" json:"insecureIngest" yaml:"insecureIngest"`
	// IngestConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-name
	//
	// Default: - "-".
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Stage ARN.
	//
	// A value other than an empty string indicates that stage is linked to IngestConfiguration. Default: "" (recording is disabled).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-stagearn
	//
	// Default: - "".
	//
	StageArn *string `field:"optional" json:"stageArn" yaml:"stageArn"`
	// A list of key-value pairs that contain metadata for the asset model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// User defined indentifier for participant associated with IngestConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html#cfn-ivs-ingestconfiguration-userid
	//
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

