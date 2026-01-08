package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnS3TableIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnS3TableIntegrationProps := &CfnS3TableIntegrationProps{
//   	Encryption: &EncryptionConfigProperty{
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	LogSources: []interface{}{
//   		&LogSourceProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Identifier: jsii.String("identifier"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html
//
type CfnS3TableIntegrationProps struct {
	// Defines the encryption configuration for S3 Table integrations, including the encryption algorithm and KMS key settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html#cfn-observabilityadmin-s3tableintegration-encryption
	//
	Encryption interface{} `field:"required" json:"encryption" yaml:"encryption"`
	// The Amazon Resource Name (ARN) of the IAM role that grants permissions for the S3 Table integration to access necessary resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html#cfn-observabilityadmin-s3tableintegration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A data source with an S3 Table integration for query access in the `logs` namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html#cfn-observabilityadmin-s3tableintegration-logsources
	//
	LogSources interface{} `field:"optional" json:"logSources" yaml:"logSources"`
	// The key-value pairs to associate with the S3 Table integration resource for categorization and management purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html#cfn-observabilityadmin-s3tableintegration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

