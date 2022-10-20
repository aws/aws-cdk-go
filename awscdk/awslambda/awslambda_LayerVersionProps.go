package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   lambda.NewLayerVersion(this, jsii.String("MyLayer"), &layerVersionProps{
//   	removalPolicy: awscdk.RemovalPolicy_RETAIN,
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	compatibleArchitectures: []architecture{
//   		lambda.*architecture_X86_64(),
//   		lambda.*architecture_ARM_64(),
//   	},
//   })
//
type LayerVersionProps struct {
	// The description the this Lambda Layer.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the layer.
	LayerVersionName *string `field:"optional" json:"layerVersionName" yaml:"layerVersionName"`
	// The SPDX licence identifier or URL to the license file for this layer.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Whether to retain this version of the layer when a new version is added or when the stack is deleted.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The content of this Layer.
	//
	// Using `Code.fromInline` is not supported.
	Code Code `field:"required" json:"code" yaml:"code"`
	// The system architectures compatible with this layer.
	CompatibleArchitectures *[]Architecture `field:"optional" json:"compatibleArchitectures" yaml:"compatibleArchitectures"`
	// The runtimes compatible with this Layer.
	CompatibleRuntimes *[]Runtime `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
}

