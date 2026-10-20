package awslambda


// Properties for defining a `CfnWebFunctionRevision`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebFunctionRevisionProps := &CfnWebFunctionRevisionProps{
//   	BuildConfig: &BuildConfigProperty{
//   		CodeConfig: &CodeConfigProperty{
//   			S3Object: &S3ObjectProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				VersionId: jsii.String("versionId"),
//   			},
//   		},
//   		RuntimeConfig: &RuntimeConfigProperty{
//   			Runtime: jsii.String("runtime"),
//   		},
//   	},
//   	FunctionName: jsii.String("functionName"),
//   	ServiceConfig: &ServiceConfigProperty{
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//
//   		// the properties below are optional
//   		EgressNetworkConnectorArn: jsii.String("egressNetworkConnectorArn"),
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
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
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html
//
type CfnWebFunctionRevisionProps struct {
	// The build configuration for the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-buildconfig
	//
	BuildConfig interface{} `field:"required" json:"buildConfig" yaml:"buildConfig"`
	// The name of the web function this revision belongs to.
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-functionname
	//
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The service configuration for the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-serviceconfig
	//
	ServiceConfig interface{} `field:"required" json:"serviceConfig" yaml:"serviceConfig"`
	// A description of the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the KMS key used to encrypt the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-webfunctionrevision.html#cfn-lambda-webfunctionrevision-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

