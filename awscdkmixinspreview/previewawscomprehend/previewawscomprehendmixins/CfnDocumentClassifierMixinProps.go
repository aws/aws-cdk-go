package previewawscomprehendmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDocumentClassifierPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDocumentClassifierMixinProps := &CfnDocumentClassifierMixinProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	DocumentClassifierName: jsii.String("documentClassifierName"),
//   	InputDataConfig: &DocumentClassifierInputDataConfigProperty{
//   		AugmentedManifests: []interface{}{
//   			&AugmentedManifestsListItemProperty{
//   				AttributeNames: []*string{
//   					jsii.String("attributeNames"),
//   				},
//   				S3Uri: jsii.String("s3Uri"),
//   				Split: jsii.String("split"),
//   			},
//   		},
//   		DataFormat: jsii.String("dataFormat"),
//   		DocumentReaderConfig: &DocumentReaderConfigProperty{
//   			DocumentReadAction: jsii.String("documentReadAction"),
//   			DocumentReadMode: jsii.String("documentReadMode"),
//   			FeatureTypes: []*string{
//   				jsii.String("featureTypes"),
//   			},
//   		},
//   		Documents: &DocumentClassifierDocumentsProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   			TestS3Uri: jsii.String("testS3Uri"),
//   		},
//   		DocumentType: jsii.String("documentType"),
//   		LabelDelimiter: jsii.String("labelDelimiter"),
//   		S3Uri: jsii.String("s3Uri"),
//   		TestS3Uri: jsii.String("testS3Uri"),
//   	},
//   	LanguageCode: jsii.String("languageCode"),
//   	Mode: jsii.String("mode"),
//   	ModelKmsKeyId: jsii.String("modelKmsKeyId"),
//   	ModelPolicy: jsii.String("modelPolicy"),
//   	OutputDataConfig: &DocumentClassifierOutputDataConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html
//
type CfnDocumentClassifierMixinProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The name of the document classifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-documentclassifiername
	//
	DocumentClassifierName *string `field:"optional" json:"documentClassifierName" yaml:"documentClassifierName"`
	// Specifies the format and location of the input data for the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-inputdataconfig
	//
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// The language of the input documents.
	//
	// You can specify any of the languages supported by Amazon Comprehend. All documents must be in the same language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Indicates the mode in which the classifier will be trained.
	//
	// The classifier can be trained in multi-class (single-label) mode or multi-label mode. Multi-class mode identifies a single class label for each document and multi-label mode identifies one or more class labels for each document. Multiple labels for an individual document are separated by a delimiter. The default delimiter between labels is a pipe (|).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// ID for the AWS  key that Amazon Comprehend uses to encrypt trained custom models.
	//
	// The ModelKmsKeyId can be either of the following formats:
	//
	// - KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`
	// - Amazon Resource Name (ARN) of a KMS Key: `"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-modelkmskeyid
	//
	ModelKmsKeyId *string `field:"optional" json:"modelKmsKeyId" yaml:"modelKmsKeyId"`
	// The resource-based policy to attach to your custom document classifier model.
	//
	// You can use this policy to allow another AWS account to import your custom model.
	//
	// Provide your policy as a JSON body that you enter as a UTF-8 encoded string without line breaks. To provide valid JSON, enclose the attribute names and values in double quotes. If the JSON body is also enclosed in double quotes, then you must escape the double quotes that are inside the policy:
	//
	// `"{\"attribute\": \"value\", \"attribute\": [\"value\"]}"`
	//
	// To avoid escaping quotes, you can use single quotes to enclose the policy and double quotes to enclose the JSON names and values:
	//
	// `'{"attribute": "value", "attribute": ["value"]}'`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-modelpolicy
	//
	ModelPolicy *string `field:"optional" json:"modelPolicy" yaml:"modelPolicy"`
	// Provides output results configuration parameters for custom classifier jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-outputdataconfig
	//
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// Tags to associate with the document classifier.
	//
	// A tag is a key-value pair that adds as a metadata to a resource used by Amazon Comprehend. For example, a tag with "Sales" as the key might be added to a resource to indicate its use by the sales department.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The version name given to the newly created classifier.
	//
	// Version names can have a maximum of 256 characters. Alphanumeric characters, hyphens (-) and underscores (_) are allowed. The version name must be unique among all models with the same classifier name in the AWS account / AWS Region .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-versionname
	//
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.
	//
	// The VolumeKmsKeyId can be either of the following formats:
	//
	// - KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`
	// - Amazon Resource Name (ARN) of a KMS Key: `"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-volumekmskeyid
	//
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// Configuration parameters for a private Virtual Private Cloud (VPC) containing the resources you are using for your custom classifier.
	//
	// For more information, see [Amazon VPC](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-documentclassifier.html#cfn-comprehend-documentclassifier-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

