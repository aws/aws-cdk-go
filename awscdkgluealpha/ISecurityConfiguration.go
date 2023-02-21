// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
)

// Interface representing a created or an imported `SecurityConfiguration`.
// Experimental.
type ISecurityConfiguration interface {
	awscdk.IResource
	// The name of the security configuration.
	// Experimental.
	SecurityConfigurationName() *string
}

// The jsii proxy for ISecurityConfiguration
type jsiiProxy_ISecurityConfiguration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISecurityConfiguration) SecurityConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationName",
		&returns,
	)
	return returns
}

