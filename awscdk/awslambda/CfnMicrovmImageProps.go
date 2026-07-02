package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMicrovmImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMicrovmImageProps := &CfnMicrovmImageProps{
//   	AdditionalOsCapabilities: []*string{
//   		jsii.String("additionalOsCapabilities"),
//   	},
//   	BaseImageArn: jsii.String("baseImageArn"),
//   	BaseImageVersion: jsii.String("baseImageVersion"),
//   	BuildRoleArn: jsii.String("buildRoleArn"),
//   	CodeArtifact: &CodeArtifactProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   	CpuConfigurations: []interface{}{
//   		&CpuConfigurationProperty{
//   			Architecture: jsii.String("architecture"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EgressNetworkConnectors: []*string{
//   		jsii.String("egressNetworkConnectors"),
//   	},
//   	EnvironmentVariables: []interface{}{
//   		&EnvironmentVariableProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Hooks: &HooksProperty{
//   		MicrovmHooks: &MicrovmHooksProperty{
//   			Resume: jsii.String("resume"),
//   			ResumeTimeoutInSeconds: jsii.Number(123),
//   			Run: jsii.String("run"),
//   			RunTimeoutInSeconds: jsii.Number(123),
//   			Suspend: jsii.String("suspend"),
//   			SuspendTimeoutInSeconds: jsii.Number(123),
//   			Terminate: jsii.String("terminate"),
//   			TerminateTimeoutInSeconds: jsii.Number(123),
//   		},
//   		MicrovmImageHooks: &MicrovmImageHooksProperty{
//   			Ready: jsii.String("ready"),
//   			ReadyTimeoutInSeconds: jsii.Number(123),
//   			Validate: jsii.String("validate"),
//   			ValidateTimeoutInSeconds: jsii.Number(123),
//   		},
//   		Port: jsii.Number(123),
//   	},
//   	Logging: &LoggingProperty{
//   		CloudWatch: &CloudWatchLoggingProperty{
//   			LogGroup: jsii.String("logGroup"),
//   			LogStream: jsii.String("logStream"),
//   		},
//   		Disabled: jsii.Boolean(false),
//   	},
//   	Name: jsii.String("name"),
//   	Resources: []interface{}{
//   		&ResourcesProperty{
//   			MinimumMemoryInMiB: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html
//
type CfnMicrovmImageProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-additionaloscapabilities
	//
	AdditionalOsCapabilities *[]*string `field:"required" json:"additionalOsCapabilities" yaml:"additionalOsCapabilities"`
	// ARN of the base MicroVM image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-baseimagearn
	//
	BaseImageArn *string `field:"required" json:"baseImageArn" yaml:"baseImageArn"`
	// Specific version of the base MicroVM image to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-baseimageversion
	//
	BaseImageVersion *string `field:"required" json:"baseImageVersion" yaml:"baseImageVersion"`
	// ARN of the IAM build role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-buildrolearn
	//
	BuildRoleArn *string `field:"required" json:"buildRoleArn" yaml:"buildRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-codeartifact
	//
	CodeArtifact interface{} `field:"required" json:"codeArtifact" yaml:"codeArtifact"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-cpuconfigurations
	//
	CpuConfigurations interface{} `field:"required" json:"cpuConfigurations" yaml:"cpuConfigurations"`
	// Human-readable description of the MicroVM image and its purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-egressnetworkconnectors
	//
	EgressNetworkConnectors *[]*string `field:"required" json:"egressNetworkConnectors" yaml:"egressNetworkConnectors"`
	// Environment variables to set in the container during the snapshot build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-environmentvariables
	//
	EnvironmentVariables interface{} `field:"required" json:"environmentVariables" yaml:"environmentVariables"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-hooks
	//
	Hooks interface{} `field:"required" json:"hooks" yaml:"hooks"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-logging
	//
	Logging interface{} `field:"required" json:"logging" yaml:"logging"`
	// Unique name for the MicroVM image within the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-resources
	//
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// Key-value pairs to associate with the MicroVM image for organization and management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-microvmimage.html#cfn-lambda-microvmimage-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

