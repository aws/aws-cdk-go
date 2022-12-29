package awslambda


// The [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) for a Lambda function. To deploy a function defined as a container image, you specify the location of a container image in the Amazon ECR registry. For a .zip file deployment package, you can specify the location of an object in Amazon S3. For Node.js and Python functions, you can specify the function code inline in the template.
//
// Changes to a deployment package in Amazon S3 are not detected automatically during stack updates. To update the function code, change the object key or version in the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeProperty := &codeProperty{
//   	imageUri: jsii.String("imageUri"),
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3Key: jsii.String("s3Key"),
//   	s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	zipFile: jsii.String("zipFile"),
//   }
//
type CfnFunction_CodeProperty struct {
	// URI of a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the Amazon ECR registry.
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// An Amazon S3 bucket in the same AWS Region as your function.
	//
	// The bucket can be in a different AWS account.
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 key of the deployment package.
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// For versioned objects, the version of the deployment package object to use.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// (Node.js and Python) The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named `index` and zips it to create a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) . This zip file cannot exceed 4MB. For the `Handler` property, the first part of the handler identifier must be `index` . For example, `index.handler` .
	//
	// For JSON, you must escape quotes and special characters such as newline ( `\n` ) with a backslash.
	//
	// If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides a response module ( [cfn-response](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-lambda-function-code-cfnresponsemodule.html) ) that simplifies sending responses. See [Using AWS Lambda with AWS CloudFormation](https://docs.aws.amazon.com/lambda/latest/dg/services-cloudformation.html) for details.
	ZipFile *string `field:"optional" json:"zipFile" yaml:"zipFile"`
}

