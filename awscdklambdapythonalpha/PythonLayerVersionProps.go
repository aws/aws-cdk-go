package awscdklambdapythonalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for PythonLayerVersion.
//
// Example:
//   python.NewPythonLayerVersion(this, jsii.String("MyLayer"), &PythonLayerVersionProps{
//   	Entry: jsii.String("/path/to/my/layer"),
//   })
//
// Experimental.
type PythonLayerVersionProps struct {
	// The description the this Lambda Layer.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the layer.
	// Default: - A name will be generated.
	//
	// Experimental.
	LayerVersionName *string `field:"optional" json:"layerVersionName" yaml:"layerVersionName"`
	// The SPDX licence identifier or URL to the license file for this layer.
	// Default: - No license information will be recorded.
	//
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Whether to retain this version of the layer when a new version is added or when the stack is deleted.
	// Default: RemovalPolicy.DESTROY
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The path to the root directory of the lambda layer.
	// Experimental.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Bundling options to use for this function.
	//
	// Use this to specify custom bundling options like
	// the bundling Docker image, asset hash type, custom hash, architecture, etc.
	// Default: - Use the default bundling Docker image, with x86_64 architecture.
	//
	// Experimental.
	Bundling *BundlingOptions `field:"optional" json:"bundling" yaml:"bundling"`
	// The system architectures compatible with this layer.
	// Default: [Architecture.X86_64]
	//
	// Experimental.
	CompatibleArchitectures *[]awslambda.Architecture `field:"optional" json:"compatibleArchitectures" yaml:"compatibleArchitectures"`
	// The runtimes compatible with the python layer.
	// Default: - Only Python 3.7 is supported.
	//
	// Experimental.
	CompatibleRuntimes *[]awslambda.Runtime `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
}

