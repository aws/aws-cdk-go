package awscomprehend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEntityRecognizerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnEntityRecognizerMixinProps := &CfnEntityRecognizerMixinProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	InputDataConfig: &EntityRecognizerInputDataConfigProperty{
//   		Annotations: &EntityRecognizerAnnotationsProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   			TestS3Uri: jsii.String("testS3Uri"),
//   		},
//   		AugmentedManifests: []interface{}{
//   			&AugmentedManifestsListItemProperty{
//   				AnnotationDataS3Uri: jsii.String("annotationDataS3Uri"),
//   				AttributeNames: []*string{
//   					jsii.String("attributeNames"),
//   				},
//   				DocumentType: jsii.String("documentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   				SourceDocumentsS3Uri: jsii.String("sourceDocumentsS3Uri"),
//   				Split: jsii.String("split"),
//   			},
//   		},
//   		DataFormat: jsii.String("dataFormat"),
//   		Documents: &EntityRecognizerDocumentsProperty{
//   			InputFormat: jsii.String("inputFormat"),
//   			S3Uri: jsii.String("s3Uri"),
//   			TestS3Uri: jsii.String("testS3Uri"),
//   		},
//   		EntityList: &EntityRecognizerEntityListProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		EntityTypes: []interface{}{
//   			&EntityTypesListItemProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	LanguageCode: jsii.String("languageCode"),
//   	ModelKmsKeyId: jsii.String("modelKmsKeyId"),
//   	ModelPolicy: jsii.String("modelPolicy"),
//   	RecognizerName: jsii.String("recognizerName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VersionName: jsii.String("versionName"),
//   	VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html
//
type CfnEntityRecognizerMixinProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Specifies the format and location of the input data for an entity recognizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-inputdataconfig
	//
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// The language of the input documents.
	//
	// All documents must be in the same language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// ID for the AWS KMS key that Amazon Comprehend uses to encrypt trained custom models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-modelkmskeyid
	//
	ModelKmsKeyId *string `field:"optional" json:"modelKmsKeyId" yaml:"modelKmsKeyId"`
	// The JSON resource-based policy to attach to your custom entity recognizer model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-modelpolicy
	//
	ModelPolicy *string `field:"optional" json:"modelPolicy" yaml:"modelPolicy"`
	// The name given to the entity recognizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-recognizername
	//
	RecognizerName *string `field:"optional" json:"recognizerName" yaml:"recognizerName"`
	// Tags to associate with the entity recognizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The version name given to the entity recognizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-versionname
	//
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
	// ID for the AWS KMS key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-volumekmskeyid
	//
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-entityrecognizer.html#cfn-comprehend-entityrecognizer-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

