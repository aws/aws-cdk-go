package awslambda


// Properties for CfnWebFunctionRevisionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWebFunctionRevisionMixinProps := &CfnWebFunctionRevisionMixinProps{
//   	BuildConfig: &BuildConfigProperty{
//   		CodeConfig: &CodeConfigProperty{
//   			S3Object: &S3ObjectProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   				VersionId: jsii.String("versionId"),
//   			},
//   		},
//   		RuntimeConfig: &RuntimeConfigProperty{
//   			Runtime: jsii.String("runtime"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FunctionName: jsii.String("functionName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	ServiceConfig: &ServiceConfigProperty{
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		MaxConcurrencyPerEnvironment: jsii.Number(123),
//   		TelemetryConfig: &TelemetryConfigProperty{
//   			LoggingConfig: &LoggingConfigProperty{
//   				ApplicationLogLevel: jsii.String("applicationLogLevel"),
//   				LogGroup: jsii.String("logGroup"),
//   				SystemLogLevel: jsii.String("systemLogLevel"),
//   			},
//   		},
//   		TimeoutSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html
//
type CfnWebFunctionRevisionMixinProps struct {
	// The build configuration for the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-buildconfig
	//
	BuildConfig interface{} `field:"optional" json:"buildConfig" yaml:"buildConfig"`
	// A description of the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the web function this revision belongs to.
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-functionname
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The ARN of the KMS key used to encrypt the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The service configuration for the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-serviceconfig
	//
	ServiceConfig interface{} `field:"optional" json:"serviceConfig" yaml:"serviceConfig"`
}

