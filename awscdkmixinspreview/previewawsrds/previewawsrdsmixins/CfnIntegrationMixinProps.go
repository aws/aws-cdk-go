package previewawsrdsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIntegrationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIntegrationMixinProps := &CfnIntegrationMixinProps{
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	DataFilter: jsii.String("dataFilter"),
//   	Description: jsii.String("description"),
//   	IntegrationName: jsii.String("integrationName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SourceArn: jsii.String("sourceArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetArn: jsii.String("targetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html
//
type CfnIntegrationMixinProps struct {
	// An optional set of non-secret key–value pairs that contains additional contextual information about the data.
	//
	// For more information, see [Encryption context](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the *AWS Key Management Service Developer Guide* .
	//
	// You can only include this parameter if you specify the `KMSKeyId` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html#cfn-rds-integration-additionalencryptioncontext
	//
	AdditionalEncryptionContext interface{} `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// Data filters for the integration.
	//
	// These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html#cfn-rds-integration-datafilter
	//
	DataFilter *string `field:"optional" json:"dataFilter" yaml:"dataFilter"`
	// A description of the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html#cfn-rds-integration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html#cfn-rds-integration-integrationname
	//
	IntegrationName *string `field:"optional" json:"integrationName" yaml:"integrationName"`
	// The AWS Key Management System ( AWS KMS) key identifier for the key to use to encrypt the integration.
	//
	// If you don't specify an encryption key, RDS uses a default AWS owned key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html#cfn-rds-integration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The Amazon Resource Name (ARN) of the database to use as the source for replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html#cfn-rds-integration-sourcearn
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// An optional array of key-value pairs to apply to this integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html#cfn-rds-integration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the Redshift data warehouse to use as the target for replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-integration.html#cfn-rds-integration-targetarn
	//
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

