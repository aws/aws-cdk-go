package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   lambda.NewLayerVersion(this, jsii.String("MyLayer"), &LayerVersionProps{
//   	RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	CompatibleArchitectures: []architecture{
//   		lambda.*architecture_X86_64(),
//   		lambda.*architecture_ARM_64(),
//   	},
//   })
//
type LayerVersionProps struct {
	// The description the this Lambda Layer.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the layer.
	// Default: - A name will be generated.
	//
	LayerVersionName *string `field:"optional" json:"layerVersionName" yaml:"layerVersionName"`
	// The SPDX licence identifier or URL to the license file for this layer.
	// Default: - No license information will be recorded.
	//
	License *string `field:"optional" json:"license" yaml:"license"`
	// Whether to retain this version of the layer when a new version is added or when the stack is deleted.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The content of this Layer.
	//
	// Using `Code.fromInline` is not supported.
	Code Code `field:"required" json:"code" yaml:"code"`
	// The system architectures compatible with this layer.
	// Default: [Architecture.X86_64]
	//
	CompatibleArchitectures *[]Architecture `field:"optional" json:"compatibleArchitectures" yaml:"compatibleArchitectures"`
	// The runtimes compatible with this Layer.
	// Default: - All runtimes are supported.
	//
	CompatibleRuntimes *[]Runtime `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
}

