package previewawsecsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs"
)

// EventBridge event patterns for Cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clusterRef IClusterRef
//
//   clusterEvents := awscdkmixinspreview.Events.ClusterEvents_FromCluster(clusterRef)
//
// Experimental.
type ClusterEvents interface {
	// EventBridge event pattern for Cluster AWS API Call via CloudTrail.
	// Experimental.
	AwsAPICallViaCloudTrailPattern(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
	// EventBridge event pattern for Cluster ECS Container Instance State Change.
	// Experimental.
	EcsContainerInstanceStateChangePattern(options *ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for Cluster ECS Service Action.
	// Experimental.
	EcsServiceActionPattern(options *ECSServiceAction_ECSServiceActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Cluster ECS Task State Change.
	// Experimental.
	EcsTaskStateChangePattern(options *ECSTaskStateChange_ECSTaskStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for ClusterEvents
type jsiiProxy_ClusterEvents struct {
	_ byte // padding
}

// Create ClusterEvents from a Cluster reference.
// Experimental.
func ClusterEvents_FromCluster(clusterRef interfacesawsecs.IClusterRef) ClusterEvents {
	_init_.Initialize()

	if err := validateClusterEvents_FromClusterParameters(clusterRef); err != nil {
		panic(err)
	}
	var returns ClusterEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents",
		"fromCluster",
		[]interface{}{clusterRef},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterEvents) AwsAPICallViaCloudTrailPattern(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	if err := c.validateAwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterEvents) EcsContainerInstanceStateChangePattern(options *ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps) *awsevents.EventPattern {
	if err := c.validateEcsContainerInstanceStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"ecsContainerInstanceStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterEvents) EcsServiceActionPattern(options *ECSServiceAction_ECSServiceActionProps) *awsevents.EventPattern {
	if err := c.validateEcsServiceActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"ecsServiceActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterEvents) EcsTaskStateChangePattern(options *ECSTaskStateChange_ECSTaskStateChangeProps) *awsevents.EventPattern {
	if err := c.validateEcsTaskStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"ecsTaskStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

