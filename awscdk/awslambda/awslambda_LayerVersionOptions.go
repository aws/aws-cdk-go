package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Non runtime options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   layerVersionOptions := &layerVersionOptions{
//   	description: jsii.String("description"),
//   	layerVersionName: jsii.String("layerVersionName"),
//   	license: jsii.String("license"),
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   }
//
type LayerVersionOptions struct {
	// The description the this Lambda Layer.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the layer.
	LayerVersionName *string `field:"optional" json:"layerVersionName" yaml:"layerVersionName"`
	// The SPDX licence identifier or URL to the license file for this layer.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Whether to retain this version of the layer when a new version is added or when the stack is deleted.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

