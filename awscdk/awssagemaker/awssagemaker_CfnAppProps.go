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
//   cfnAppProps := &cfnAppProps{
//   	appName: jsii.String("appName"),
//   	appType: jsii.String("appType"),
//   	domainId: jsii.String("domainId"),
//   	userProfileName: jsii.String("userProfileName"),
//
//   	// the properties below are optional
//   	resourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

