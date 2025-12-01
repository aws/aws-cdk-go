package awslambda


// The [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) for a Lambda function. To deploy a function defined as a container image, you specify the location of a container image in the Amazon ECR registry. For a .zip file deployment package, you can specify the location of an object in Amazon S3. For Node.js and Python functions, you can specify the function code inline in the template.
//
// > When you specify source code inline for a Node.js function, the `index` file that CloudFormation creates uses the extension `.js` . This means that Node.js treats the file as a CommonJS module.
//
// Changes to a deployment package in Amazon S3 or a container image in ECR are not detected automatically during stack updates. To update the function code, change the object key or version in the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeProperty := &CodeProperty{
//   	ImageUri: jsii.String("imageUri"),
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Key: jsii.String("s3Key"),
//   	S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	SourceKmsKeyArn: jsii.String("sourceKmsKeyArn"),
//   	ZipFile: jsii.String("zipFile"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html
//
type CfnFunction_CodeProperty struct {
	// URI of a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the Amazon ECR registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-imageuri
	//
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// An Amazon S3 bucket in the same AWS Region as your function.
	//
	// The bucket can be in a different AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 key of the deployment package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3key
	//
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// For versioned objects, the version of the deployment package object to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3objectversion
	//
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// The ARN of the AWS Key Management Service ( AWS  ) customer managed key that's used to encrypt your function's .zip deployment package. If you don't provide a customer managed key, Lambda uses an [AWS owned key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-sourcekmskeyarn
	//
	SourceKmsKeyArn *string `field:"optional" json:"sourceKmsKeyArn" yaml:"sourceKmsKeyArn"`
	// (Node.js and Python) The source code of your Lambda function. If you include your function source inline with this parameter, CloudFormation places it in a file named `index` and zips it to create a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) . This zip file cannot exceed 4MB. For the `Handler` property, the first part of the handler identifier must be `index` . For example, `index.handler` .
	//
	// > When you specify source code inline for a Node.js function, the `index` file that CloudFormation creates uses the extension `.js` . This means that Node.js treats the file as a CommonJS module.
	// >
	// > When using Node.js 24 or later, Node.js can automatically detect if a `.js` file should be treated as CommonJS or as an ES module. To enable auto-detection, add the `--experimental-detect-module` flag to the `NODE_OPTIONS` environment variable. For more information, see [Experimental Node.js features](https://docs.aws.amazon.com//lambda/latest/dg/lambda-nodejs.html#nodejs-experimental-features) .
	//
	// For JSON, you must escape quotes and special characters such as newline ( `\n` ) with a backslash.
	//
	// If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides a response module ( [cfn-response](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-lambda-function-code-cfnresponsemodule.html) ) that simplifies sending responses. See [Using AWS Lambda with AWS CloudFormation](https://docs.aws.amazon.com/lambda/latest/dg/services-cloudformation.html) for details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-zipfile
	//
	ZipFile *string `field:"optional" json:"zipFile" yaml:"zipFile"`
}

