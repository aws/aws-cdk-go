package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Note: This class currently stands as a namespaced container for importing an ECS Deployment Group defined outside the CDK app until CloudFormation supports provisioning ECS Deployment Groups.
//
// Until then it is closed (private constructor) and does not
// extend {@link Construct}.
type EcsDeploymentGroup interface {
}

// The jsii proxy struct for EcsDeploymentGroup
type jsiiProxy_EcsDeploymentGroup struct {
	_ byte // padding
}

// Import an ECS Deployment Group defined outside the CDK app.
//
// Returns: a Construct representing a reference to an existing Deployment Group.
func EcsDeploymentGroup_FromEcsDeploymentGroupAttributes(scope constructs.Construct, id *string, attrs *EcsDeploymentGroupAttributes) IEcsDeploymentGroup {
	_init_.Initialize()

	if err := validateEcsDeploymentGroup_FromEcsDeploymentGroupAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IEcsDeploymentGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentGroup",
		"fromEcsDeploymentGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

