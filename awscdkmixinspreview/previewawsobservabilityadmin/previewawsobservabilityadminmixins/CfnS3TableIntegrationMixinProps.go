package previewawsobservabilityadminmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnS3TableIntegrationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnS3TableIntegrationMixinProps := &CfnS3TableIntegrationMixinProps{
//   	Encryption: &EncryptionConfigProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//   	},
//   	LogSources: []interface{}{
//   		&LogSourceProperty{
//   			Identifier: jsii.String("identifier"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
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
type CfnS3TableIntegrationMixinProps struct {
	// Defines the encryption configuration for S3 Table integrations, including the encryption algorithm and KMS key settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html#cfn-observabilityadmin-s3tableintegration-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// The CloudWatch Logs data sources to associate with the S3 Table Integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html#cfn-observabilityadmin-s3tableintegration-logsources
	//
	LogSources interface{} `field:"optional" json:"logSources" yaml:"logSources"`
	// The ARN of the role used to access the S3 Table Integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html#cfn-observabilityadmin-s3tableintegration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-s3tableintegration.html#cfn-observabilityadmin-s3tableintegration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

