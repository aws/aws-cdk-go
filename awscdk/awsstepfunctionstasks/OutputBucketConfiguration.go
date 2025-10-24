package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// S3 bucket configuration for the output data.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//
//   var outputBucket IBucket
//   var trainingBucket IBucket
//   var validationBucket IBucket
//   var kmsKey IKey
//   var vpc IVpc
//
//
//   model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())
//
//   task := tasks.NewBedrockCreateModelCustomizationJob(this, jsii.String("CreateModelCustomizationJob"), &BedrockCreateModelCustomizationJobProps{
//   	BaseModel: model,
//   	ClientRequestToken: jsii.String("MyToken"),
//   	CustomizationType: tasks.CustomizationType_FINE_TUNING,
//   	CustomModelKmsKey: kmsKey,
//   	CustomModelName: jsii.String("MyCustomModel"),
//   	 // required
//   	CustomModelTags: []CustomModelTag{
//   		&CustomModelTag{
//   			Key: jsii.String("key1"),
//   			Value: jsii.String("value1"),
//   		},
//   	},
//   	HyperParameters: map[string]*string{
//   		"batchSize": jsii.String("10"),
//   	},
//   	JobName: jsii.String("MyCustomizationJob"),
//   	 // required
//   	JobTags: []CustomModelTag{
//   		&CustomModelTag{
//   			Key: jsii.String("key2"),
//   			Value: jsii.String("value2"),
//   		},
//   	},
//   	OutputData: &OutputBucketConfiguration{
//   		Bucket: outputBucket,
//   		 // required
//   		Path: jsii.String("output-data/"),
//   	},
//   	TrainingData: &TrainingBucketConfiguration{
//   		Bucket: trainingBucket,
//   		Path: jsii.String("training-data/data.json"),
//   	},
//   	 // required
//   	// If you don't provide validation data, you have to specify `Evaluation percentage` hyperparameter.
//   	ValidationData: []ValidationBucketConfiguration{
//   		&ValidationBucketConfiguration{
//   			Bucket: validationBucket,
//   			Path: jsii.String("validation-data/data.json"),
//   		},
//   	},
//   	VpcConfig: map[string][]ISecurityGroup{
//   		"securityGroups": []ISecurityGroup{
//   			ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   				"vpc": vpc,
//   			}),
//   		},
//   		"subnets": vpc.privateSubnets,
//   	},
//   })
//
type OutputBucketConfiguration struct {
	// The S3 bucket.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Path to file or directory within the bucket.
	// Default: - root of the bucket.
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

