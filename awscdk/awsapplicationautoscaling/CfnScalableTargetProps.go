package awsapplicationautoscaling


// Properties for defining a `CfnScalableTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScalableTargetProps := &CfnScalableTargetProps{
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   	ResourceId: jsii.String("resourceId"),
//   	ScalableDimension: jsii.String("scalableDimension"),
//   	ServiceNamespace: jsii.String("serviceNamespace"),
//
//   	// the properties below are optional
//   	RoleArn: jsii.String("roleArn"),
//   	ScheduledActions: []interface{}{
//   		&ScheduledActionProperty{
//   			Schedule: jsii.String("schedule"),
//   			ScheduledActionName: jsii.String("scheduledActionName"),
//
//   			// the properties below are optional
//   			EndTime: NewDate(),
//   			ScalableTargetAction: &ScalableTargetActionProperty{
//   				MaxCapacity: jsii.Number(123),
//   				MinCapacity: jsii.Number(123),
//   			},
//   			StartTime: NewDate(),
//   			Timezone: jsii.String("timezone"),
//   		},
//   	},
//   	SuspendedState: &SuspendedStateProperty{
//   		DynamicScalingInSuspended: jsii.Boolean(false),
//   		DynamicScalingOutSuspended: jsii.Boolean(false),
//   		ScheduledScalingSuspended: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html
//
type CfnScalableTargetProps struct {
	// The maximum value that you plan to scale out to.
	//
	// When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-maxcapacity
	//
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum value that you plan to scale in to.
	//
	// When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-mincapacity
	//
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
	// The identifier of the resource associated with the scalable target.
	//
	// This string consists of the resource type and unique identifier.
	//
	// - ECS service - The resource type is `service` and the unique identifier is the cluster name and service name. Example: `service/my-cluster/my-service` .
	// - Spot Fleet - The resource type is `spot-fleet-request` and the unique identifier is the Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` .
	// - EMR cluster - The resource type is `instancegroup` and the unique identifier is the cluster ID and instance group ID. Example: `instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0` .
	// - AppStream 2.0 fleet - The resource type is `fleet` and the unique identifier is the fleet name. Example: `fleet/sample-fleet` .
	// - DynamoDB table - The resource type is `table` and the unique identifier is the table name. Example: `table/my-table` .
	// - DynamoDB global secondary index - The resource type is `index` and the unique identifier is the index name. Example: `table/my-table/index/my-table-index` .
	// - Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:my-db-cluster` .
	// - SageMaker endpoint variant - The resource type is `variant` and the unique identifier is the resource ID. Example: `endpoint/my-end-point/variant/KMeansClustering` .
	// - Custom resources are not supported with a resource type. This parameter must specify the `OutputValue` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource) .
	// - Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE` .
	// - Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE` .
	// - Lambda provisioned concurrency - The resource type is `function` and the unique identifier is the function name with a function version or alias name suffix that is not `$LATEST` . Example: `function:my-function:prod` or `function:my-function:1` .
	// - Amazon Keyspaces table - The resource type is `table` and the unique identifier is the table name. Example: `keyspace/mykeyspace/table/mytable` .
	// - Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: `arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5` .
	// - Amazon ElastiCache replication group - The resource type is `replication-group` and the unique identifier is the replication group name. Example: `replication-group/mycluster` .
	// - Amazon ElastiCache cache cluster - The resource type is `cache-cluster` and the unique identifier is the cache cluster name. Example: `cache-cluster/mycluster` .
	// - Neptune cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:mycluster` .
	// - SageMaker serverless endpoint - The resource type is `variant` and the unique identifier is the resource ID. Example: `endpoint/my-end-point/variant/KMeansClustering` .
	// - SageMaker inference component - The resource type is `inference-component` and the unique identifier is the resource ID. Example: `inference-component/my-inference-component` .
	// - Pool of WorkSpaces - The resource type is `workspacespool` and the unique identifier is the pool ID. Example: `workspacespool/wspool-123456` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-resourceid
	//
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The scalable dimension associated with the scalable target.
	//
	// This string consists of the service namespace, resource type, and scaling property.
	//
	// - `ecs:service:DesiredCount` - The task count of an ECS service.
	// - `elasticmapreduce:instancegroup:InstanceCount` - The instance count of an EMR Instance Group.
	// - `ec2:spot-fleet-request:TargetCapacity` - The target capacity of a Spot Fleet.
	// - `appstream:fleet:DesiredCapacity` - The capacity of an AppStream 2.0 fleet.
	// - `dynamodb:table:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB table.
	// - `dynamodb:table:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB table.
	// - `dynamodb:index:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB global secondary index.
	// - `dynamodb:index:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB global secondary index.
	// - `rds:cluster:ReadReplicaCount` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	// - `sagemaker:variant:DesiredInstanceCount` - The number of EC2 instances for a SageMaker model endpoint variant.
	// - `custom-resource:ResourceType:Property` - The scalable dimension for a custom resource provided by your own application or service.
	// - `comprehend:document-classifier-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend document classification endpoint.
	// - `comprehend:entity-recognizer-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.
	// - `lambda:function:ProvisionedConcurrency` - The provisioned concurrency for a Lambda function.
	// - `cassandra:table:ReadCapacityUnits` - The provisioned read capacity for an Amazon Keyspaces table.
	// - `cassandra:table:WriteCapacityUnits` - The provisioned write capacity for an Amazon Keyspaces table.
	// - `kafka:broker-storage:VolumeSize` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.
	// - `elasticache:cache-cluster:Nodes` - The number of nodes for an Amazon ElastiCache cache cluster.
	// - `elasticache:replication-group:NodeGroups` - The number of node groups for an Amazon ElastiCache replication group.
	// - `elasticache:replication-group:Replicas` - The number of replicas per node group for an Amazon ElastiCache replication group.
	// - `neptune:cluster:ReadReplicaCount` - The count of read replicas in an Amazon Neptune DB cluster.
	// - `sagemaker:variant:DesiredProvisionedConcurrency` - The provisioned concurrency for a SageMaker serverless endpoint.
	// - `sagemaker:inference-component:DesiredCopyCount` - The number of copies across an endpoint for a SageMaker inference component.
	// - `workspaces:workspacespool:DesiredUserSessions` - The number of user sessions for the WorkSpaces in the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-scalabledimension
	//
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// The namespace of the AWS service that provides the resource, or a `custom-resource` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-servicenamespace
	//
	ServiceNamespace *string `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
	// Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf.
	//
	// This can be either an IAM service role that Application Auto Scaling can assume to make calls to other AWS resources on your behalf, or a service-linked role for the specified service. For more information, see [How Application Auto Scaling works with IAM](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html) in the *Application Auto Scaling User Guide* .
	//
	// To automatically create a service-linked role (recommended), specify the full ARN of the service-linked role in your stack template. To find the exact ARN of the service-linked role for your AWS or custom resource, see the [Service-linked roles](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-service-linked-roles.html) topic in the *Application Auto Scaling User Guide* . Look for the ARN in the table at the bottom of the page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The scheduled actions for the scalable target.
	//
	// Duplicates aren't allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-scheduledactions
	//
	ScheduledActions interface{} `field:"optional" json:"scheduledActions" yaml:"scheduledActions"`
	// An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling.
	//
	// Setting the value of an attribute to `true` suspends the specified scaling activities. Setting it to `false` (default) resumes the specified scaling activities.
	//
	// *Suspension Outcomes*
	//
	// - For `DynamicScalingInSuspended` , while a suspension is in effect, all scale-in activities that are triggered by a scaling policy are suspended.
	// - For `DynamicScalingOutSuspended` , while a suspension is in effect, all scale-out activities that are triggered by a scaling policy are suspended.
	// - For `ScheduledScalingSuspended` , while a suspension is in effect, all scaling activities that involve scheduled actions are suspended.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-suspendedstate
	//
	SuspendedState interface{} `field:"optional" json:"suspendedState" yaml:"suspendedState"`
}

