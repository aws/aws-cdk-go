// Package jsii contains the functionaility needed for jsii packages to
// initialize their dependencies and themselves. Users should never need to use this package
// directly. If you find you need to - please report a bug at
// https://github.com/aws/jsii/issues/new/choose
package jsii

import (
	_                                      "embed"

	_jsii_                                 "github.com/aws/jsii-runtime-go/runtime"

	awscdk                                 "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	awscdkiotalpha                         "github.com/aws/aws-cdk-go/awscdkiotalpha/v2/jsii"
	awscdkioteventsalpha                   "github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2/jsii"
	awscdkkinesisfirehosealpha             "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha/v2/jsii"
	awscdkkinesisfirehosedestinationsalpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha/v2/jsii"
	constructs                             "github.com/aws/constructs-go/constructs/v10/jsii"
)

//go:embed aws-cdk-aws-iot-actions-alpha-2.181.1-alpha.0.tgz
var tarball []byte

// Initialize loads the necessary packages in the @jsii/kernel to support the enclosing module.
// The implementation is idempotent (and hence safe to be called over and over).
func Initialize() {
	// Ensure all dependencies are initialized
	awscdkiotalpha.Initialize()
	awscdkioteventsalpha.Initialize()
	awscdkkinesisfirehosealpha.Initialize()
	awscdkkinesisfirehosedestinationsalpha.Initialize()
	awscdk.Initialize()
	constructs.Initialize()

	// Load this library into the kernel
	_jsii_.Load("@aws-cdk/aws-iot-actions-alpha", "2.181.1-alpha.0", tarball)
}
