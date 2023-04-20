package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppProps := &CfnAppProps{
//   	AppName: jsii.String("appName"),
//   	AppType: jsii.String("appType"),
//   	DomainId: jsii.String("domainId"),
//   	UserProfileName: jsii.String("userProfileName"),
//
//   	// the properties below are optional
//   	ResourceSpec: &ResourceSpecProperty{
//   		InstanceType: jsii.String("instanceType"),
//   		SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppProps struct {
	// The name of the app.
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The type of app.
	//
	// *Allowed Values* : `JupyterServer | KernelGateway | RSessionGateway | RStudioServerPro | TensorBoard | Canvas`.
	AppType *string `field:"required" json:"appType" yaml:"appType"`
	// The domain ID.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The user profile name.
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
	// Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
	ResourceSpec interface{} `field:"optional" json:"resourceSpec" yaml:"resourceSpec"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

