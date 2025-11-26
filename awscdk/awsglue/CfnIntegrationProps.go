package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationProps := &CfnIntegrationProps{
//   	IntegrationName: jsii.String("integrationName"),
//   	SourceArn: jsii.String("sourceArn"),
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	DataFilter: jsii.String("dataFilter"),
//   	Description: jsii.String("description"),
//   	IntegrationConfig: &IntegrationConfigProperty{
//   		ContinuousSync: jsii.Boolean(false),
//   		RefreshInterval: jsii.String("refreshInterval"),
//   		SourceProperties: map[string]*string{
//   			"sourcePropertiesKey": jsii.String("sourceProperties"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html
//
type CfnIntegrationProps struct {
	// A unique name for the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-integrationname
	//
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
	// The ARN for the source of the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-sourcearn
	//
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The ARN for the target of the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-targetarn
	//
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An optional set of non-secret key–value pairs that contains additional contextual information for encryption.
	//
	// This can only be provided if `KMSKeyId` is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-additionalencryptioncontext
	//
	AdditionalEncryptionContext interface{} `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// Selects source tables for the integration using Maxwell filter syntax.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-datafilter
	//
	DataFilter *string `field:"optional" json:"dataFilter" yaml:"dataFilter"`
	// A description for the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The structure used to define properties associated with the zero-ETL integration.
	//
	// For more information, see [IntegrationConfig structure.](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-integrations.html#aws-glue-api-integrations-IntegrationConfig)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-integrationconfig
	//
	IntegrationConfig interface{} `field:"optional" json:"integrationConfig" yaml:"integrationConfig"`
	// The ARN of a KMS key used for encrypting the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Metadata assigned to the resource consisting of a list of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

