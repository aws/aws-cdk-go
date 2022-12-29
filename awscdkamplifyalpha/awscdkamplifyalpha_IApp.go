// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/internal"
)

// An Amplify Console application.
// Experimental.
type IApp interface {
	awscdk.IResource
	// The application id.
	// Experimental.
	AppId() *string
}

// The jsii proxy for IApp
type jsiiProxy_IApp struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApp) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

