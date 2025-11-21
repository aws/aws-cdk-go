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
	// The name of the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-integrationname
	//
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
	// The Amazon Resource Name (ARN) of the database to use as the source for replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-sourcearn
	//
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The Amazon Resource Name (ARN) of the Glue data warehouse to use as the target for replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-targetarn
	//
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An optional set of non-secret key value pairs that contains additional contextual information about the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-additionalencryptioncontext
	//
	AdditionalEncryptionContext interface{} `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-datafilter
	//
	DataFilter *string `field:"optional" json:"dataFilter" yaml:"dataFilter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration settings for the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-integrationconfig
	//
	IntegrationConfig interface{} `field:"optional" json:"integrationConfig" yaml:"integrationConfig"`
	// An KMS key identifier for the key to use to encrypt the integration.
	//
	// If you don't specify an encryption key, the default AWS owned KMS key is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integration.html#cfn-glue-integration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

