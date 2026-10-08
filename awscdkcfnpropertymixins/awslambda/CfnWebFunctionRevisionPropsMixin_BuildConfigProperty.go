package awslambda


// The build configuration for the revision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   buildConfigProperty := &BuildConfigProperty{
//   	CodeConfig: &CodeConfigProperty{
//   		S3Object: &S3ObjectProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			VersionId: jsii.String("versionId"),
//   		},
//   	},
//   	RuntimeConfig: &RuntimeConfigProperty{
//   		Runtime: jsii.String("runtime"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-buildconfig.html
//
type CfnWebFunctionRevisionPropsMixin_BuildConfigProperty struct {
	// The code configuration for the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-buildconfig.html#cfn-lambda-webfunctionrevision-buildconfig-codeconfig
	//
	CodeConfig interface{} `field:"optional" json:"codeConfig" yaml:"codeConfig"`
	// The runtime configuration for the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-buildconfig.html#cfn-lambda-webfunctionrevision-buildconfig-runtimeconfig
	//
	RuntimeConfig interface{} `field:"optional" json:"runtimeConfig" yaml:"runtimeConfig"`
}

