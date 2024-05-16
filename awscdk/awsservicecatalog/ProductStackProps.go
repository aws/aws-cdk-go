package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
)

// Product stack props.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   type lambdaProduct struct {
//   	productStack
//   }
//
//   func newLambdaProduct(scope construct, id *string, props productStackProps) *lambdaProduct {
//   	this := &lambdaProduct{}
//   	servicecatalog.NewProductStack_Override(this, scope, id, props)
//
//   	lambda.NewFunction(this, jsii.String("LambdaProduct"), &FunctionProps{
//   		Runtime: lambda.Runtime_PYTHON_3_9(),
//   		Code: lambda.Code_FromAsset(jsii.String("./assets")),
//   		Handler: jsii.String("index.handler"),
//   	})
//   	return this
//   }
//
//   userDefinedBucket := awscdk.NewBucket(this, jsii.String("UserDefinedBucket"), &BucketProps{
//   	BucketName: jsii.String("user-defined-bucket-for-product-stack-assets"),
//   })
//
//   product := servicecatalog.NewCloudFormationProduct(this, jsii.String("Product"), &CloudFormationProductProps{
//   	ProductName: jsii.String("My Product"),
//   	Owner: jsii.String("Product Owner"),
//   	ProductVersions: []cloudFormationProductVersion{
//   		&cloudFormationProductVersion{
//   			ProductVersionName: jsii.String("v1"),
//   			CloudFormationTemplate: servicecatalog.CloudFormationTemplate_FromProductStack(NewLambdaProduct(this, jsii.String("LambdaFunctionProduct"), &productStackProps{
//   				AssetBucket: userDefinedBucket,
//   			})),
//   		},
//   	},
//   })
//
type ProductStackProps struct {
	// A Bucket can be passed to store assets, enabling ProductStack Asset support.
	// Default: - No Bucket provided and Assets will not be supported.
	//
	AssetBucket awss3.IBucket `field:"optional" json:"assetBucket" yaml:"assetBucket"`
	// The amount of memory (in MiB) to allocate to the AWS Lambda function which replicates the files from the CDK bucket to the destination bucket.
	//
	// If you are deploying large files, you will need to increase this number
	// accordingly.
	// Default: 128.
	//
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// A ServerSideEncryption can be enabled to encrypt assets that are put into assetBucket.
	// Default: - No encryption is used.
	//
	ServerSideEncryption awss3deployment.ServerSideEncryption `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// For AWS_KMS ServerSideEncryption a KMS KeyId must be provided which will be used to encrypt assets.
	// Default: - No KMS KeyId and SSE_KMS encryption cannot be used.
	//
	ServerSideEncryptionAwsKmsKeyId *string `field:"optional" json:"serverSideEncryptionAwsKmsKeyId" yaml:"serverSideEncryptionAwsKmsKeyId"`
}

