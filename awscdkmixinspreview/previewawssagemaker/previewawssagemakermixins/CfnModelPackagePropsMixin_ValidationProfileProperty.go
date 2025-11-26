package previewawssagemakermixins


// Contains data, such as the inputs and targeted instance types that are used in the process of validating the model package.
//
// The data provided in the validation profile is made available to your buyers on AWS Marketplace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   validationProfileProperty := &ValidationProfileProperty{
//   	ProfileName: jsii.String("profileName"),
//   	TransformJobDefinition: &TransformJobDefinitionProperty{
//   		BatchStrategy: jsii.String("batchStrategy"),
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		MaxConcurrentTransforms: jsii.Number(123),
//   		MaxPayloadInMb: jsii.Number(123),
//   		TransformInput: &TransformInputProperty{
//   			CompressionType: jsii.String("compressionType"),
//   			ContentType: jsii.String("contentType"),
//   			DataSource: &DataSourceProperty{
//   				S3DataSource: &S3DataSourceProperty{
//   					S3DataType: jsii.String("s3DataType"),
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   			SplitType: jsii.String("splitType"),
//   		},
//   		TransformOutput: &TransformOutputProperty{
//   			Accept: jsii.String("accept"),
//   			AssembleWith: jsii.String("assembleWith"),
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			S3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   		TransformResources: &TransformResourcesProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationprofile.html
//
type CfnModelPackagePropsMixin_ValidationProfileProperty struct {
	// The name of the profile for the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationprofile.html#cfn-sagemaker-modelpackage-validationprofile-profilename
	//
	ProfileName *string `field:"optional" json:"profileName" yaml:"profileName"`
	// The `TransformJobDefinition` object that describes the transform job used for the validation of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationprofile.html#cfn-sagemaker-modelpackage-validationprofile-transformjobdefinition
	//
	TransformJobDefinition interface{} `field:"optional" json:"transformJobDefinition" yaml:"transformJobDefinition"`
}

