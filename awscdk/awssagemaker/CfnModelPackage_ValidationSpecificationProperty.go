package awssagemaker


// Specifies batch transform jobs that SageMaker runs to validate your model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validationSpecificationProperty := &ValidationSpecificationProperty{
//   	ValidationProfiles: []interface{}{
//   		&ValidationProfileProperty{
//   			ProfileName: jsii.String("profileName"),
//   			TransformJobDefinition: &TransformJobDefinitionProperty{
//   				TransformInput: &TransformInputProperty{
//   					DataSource: &DataSourceProperty{
//   						S3DataSource: &S3DataSourceProperty{
//   							S3DataType: jsii.String("s3DataType"),
//   							S3Uri: jsii.String("s3Uri"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					CompressionType: jsii.String("compressionType"),
//   					ContentType: jsii.String("contentType"),
//   					SplitType: jsii.String("splitType"),
//   				},
//   				TransformOutput: &TransformOutputProperty{
//   					S3OutputPath: jsii.String("s3OutputPath"),
//
//   					// the properties below are optional
//   					Accept: jsii.String("accept"),
//   					AssembleWith: jsii.String("assembleWith"),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   				},
//   				TransformResources: &TransformResourcesProperty{
//   					InstanceCount: jsii.Number(123),
//   					InstanceType: jsii.String("instanceType"),
//
//   					// the properties below are optional
//   					VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   				},
//
//   				// the properties below are optional
//   				BatchStrategy: jsii.String("batchStrategy"),
//   				Environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				MaxConcurrentTransforms: jsii.Number(123),
//   				MaxPayloadInMb: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ValidationRole: jsii.String("validationRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationspecification.html
//
type CfnModelPackage_ValidationSpecificationProperty struct {
	// An array of `ModelPackageValidationProfile` objects, each of which specifies a batch transform job that SageMaker runs to validate your model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationspecification.html#cfn-sagemaker-modelpackage-validationspecification-validationprofiles
	//
	ValidationProfiles interface{} `field:"required" json:"validationProfiles" yaml:"validationProfiles"`
	// The IAM roles to be used for the validation of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-validationspecification.html#cfn-sagemaker-modelpackage-validationspecification-validationrole
	//
	ValidationRole *string `field:"required" json:"validationRole" yaml:"validationRole"`
}

