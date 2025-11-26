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
	AwsAPICallViaCloudTrailPattern(options *ClusterEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
	// EventBridge event pattern for Cluster ECS Container Instance State Change.
	// Experimental.
	ECSContainerInstanceStateChangePattern(options *ClusterEvents_ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps) *awsevents.EventPattern
	// EventBridge event pattern for Cluster ECS Service Action.
	// Experimental.
	ECSServiceActionPattern(options *ClusterEvents_ECSServiceAction_ECSServiceActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Cluster ECS Task State Change.
	// Experimental.
	ECSTaskStateChangePattern(options *ClusterEvents_ECSTaskStateChange_ECSTaskStateChangeProps) *awsevents.EventPattern
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

func (c *jsiiProxy_ClusterEvents) AwsAPICallViaCloudTrailPattern(options *ClusterEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
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

func (c *jsiiProxy_ClusterEvents) ECSContainerInstanceStateChangePattern(options *ClusterEvents_ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps) *awsevents.EventPattern {
	if err := c.validateECSContainerInstanceStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"eCSContainerInstanceStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterEvents) ECSServiceActionPattern(options *ClusterEvents_ECSServiceAction_ECSServiceActionProps) *awsevents.EventPattern {
	if err := c.validateECSServiceActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"eCSServiceActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterEvents) ECSTaskStateChangePattern(options *ClusterEvents_ECSTaskStateChange_ECSTaskStateChangeProps) *awsevents.EventPattern {
	if err := c.validateECSTaskStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"eCSTaskStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

