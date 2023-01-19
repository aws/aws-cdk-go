package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAppImageConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppImageConfigProps := &cfnAppImageConfigProps{
//   	appImageConfigName: jsii.String("appImageConfigName"),
//
//   	// the properties below are optional
//   	kernelGatewayImageConfig: &kernelGatewayImageConfigProperty{
//   		kernelSpecs: []interface{}{
//   			&kernelSpecProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				displayName: jsii.String("displayName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		fileSystemConfig: &fileSystemConfigProperty{
//   			defaultGid: jsii.Number(123),
//   			defaultUid: jsii.Number(123),
//   			mountPath: jsii.String("mountPath"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppImageConfigProps struct {
	// The name of the AppImageConfig.
	//
	// Must be unique to your account.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The configuration for the file system and kernels in the SageMaker image.
	KernelGatewayImageConfig interface{} `field:"optional" json:"kernelGatewayImageConfig" yaml:"kernelGatewayImageConfig"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

