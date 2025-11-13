package awsbedrockagentcore


// Representation of a code configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfigurationProperty := &CodeConfigurationProperty{
//   	Code: &CodeProperty{
//   		S3: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Prefix: jsii.String("prefix"),
//
//   			// the properties below are optional
//   			VersionId: jsii.String("versionId"),
//   		},
//   	},
//   	EntryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	Runtime: jsii.String("runtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-codeconfiguration.html
//
type CfnRuntime_CodeConfigurationProperty struct {
	// Object represents source code from zip file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-codeconfiguration.html#cfn-bedrockagentcore-runtime-codeconfiguration-code
	//
	Code interface{} `field:"required" json:"code" yaml:"code"`
	// List of entry points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-codeconfiguration.html#cfn-bedrockagentcore-runtime-codeconfiguration-entrypoint
	//
	EntryPoint *[]*string `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// Managed runtime types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-codeconfiguration.html#cfn-bedrockagentcore-runtime-codeconfiguration-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
}

