package awssagemaker


// Contains data, such as the inputs and targeted instance types that are used in the process of validating the model package.
//
// The data provided in the validation profile is made available to your buyers on AWS Marketplace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validationProfileProperty := &ValidationProfileProperty{
//   	ProfileName: jsii.String("profileName"),
//   	TransformJobDefinition: &TransformJobDefinitionProperty{
//   		TransformInput: &TransformInputProperty{
//   			DataSource: &DataSourceProperty{
//   				S3DataSource: &S3DataSourceProperty{
//   					S3DataType: jsii.String("s3DataType"),
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			CompressionType: jsii.String("compressionType"),
//   			ContentType: jsii.String("contentType"),
//   			SplitType: jsii.String("splitType"),
//   		},
//   		TransformOutput: &TransformOutputProperty{
//   			S3OutputPath: jsii.String("s3OutputPath"),
//
//   			// the properties below are optional
//   			Accept: jsii.String("accept"),
//   			AssembleWith: jsii.String("assembleWith"),
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   		TransformResources: &TransformResourcesProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//
//   		// the properties below are optional
//   		BatchStrategy: jsii.String("batchStrategy"),
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		MaxConcurrentTransforms: jsii.Number(123),
//   		MaxPayloadInMb: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationprofile.html
//
type CfnModelPackage_ValidationProfileProperty struct {
	// The name of the profile for the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationprofile.html#cfn-sagemaker-modelpackage-validationprofile-profilename
	//
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// The `TransformJobDefinition` object that describes the transform job used for the validation of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationprofile.html#cfn-sagemaker-modelpackage-validationprofile-transformjobdefinition
	//
	TransformJobDefinition interface{} `field:"required" json:"transformJobDefinition" yaml:"transformJobDefinition"`
}

