package mixinsawssynthetics


// Use this structure to input your script code for the canary.
//
// This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script is passed into the canary directly, the script code is contained in the value of `Script` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeProperty := &CodeProperty{
//   	BlueprintTypes: []*string{
//   		jsii.String("blueprintTypes"),
//   	},
//   	Dependencies: []interface{}{
//   		&DependencyProperty{
//   			Reference: jsii.String("reference"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Handler: jsii.String("handler"),
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Key: jsii.String("s3Key"),
//   	S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	Script: jsii.String("script"),
//   	SourceLocationArn: jsii.String("sourceLocationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html
//
type CfnCanaryPropsMixin_CodeProperty struct {
	// `BlueprintTypes` are a list of templates that enable simplified canary creation.
	//
	// You can create canaries for common monitoring scenarios by providing only a JSON configuration file instead of writing custom scripts. `multi-checks` is the only supported value.
	//
	// When you specify `BlueprintTypes` , the `Handler` field cannot be specified since the blueprint provides a pre-defined entry point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-blueprinttypes
	//
	BlueprintTypes *[]*string `field:"optional" json:"blueprintTypes" yaml:"blueprintTypes"`
	// List of Lambda layers to attach to the canary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-dependencies
	//
	Dependencies interface{} `field:"optional" json:"dependencies" yaml:"dependencies"`
	// The entry point to use for the source code when running the canary.
	//
	// For canaries that use the `syn-python-selenium-1.0` runtime or a `syn-nodejs.puppeteer` runtime earlier than `syn-nodejs.puppeteer-3.4` , the handler must be specified as `*fileName* .handler` . For `syn-python-selenium-1.1` , `syn-nodejs.puppeteer-3.4` , and later runtimes, the handler can be specified as `*fileName* . *functionName*` , or you can specify a folder where canary scripts reside as `*folder* / *fileName* . *functionName*` .
	//
	// This field is required when you don't specify `BlueprintTypes` and is not allowed when you specify `BlueprintTypes` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-handler
	//
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// If your canary script is located in S3, specify the bucket name here.
	//
	// The bucket must already exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 key of your script.
	//
	// For more information, see [Working with Amazon S3 Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingObjects.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3key
	//
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// The Amazon S3 version ID of your script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3objectversion
	//
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// If you input your canary script directly into the canary instead of referring to an S3 location, the value of this parameter is the script in plain text.
	//
	// It can be up to 5 MB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-script
	//
	Script *string `field:"optional" json:"script" yaml:"script"`
	// The ARN of the Lambda layer where Synthetics stores the canary script code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-sourcelocationarn
	//
	SourceLocationArn *string `field:"optional" json:"sourceLocationArn" yaml:"sourceLocationArn"`
}

