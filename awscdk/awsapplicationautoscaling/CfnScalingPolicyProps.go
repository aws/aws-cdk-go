package awsapplicationautoscaling


// Properties for defining a `CfnScalingPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScalingPolicyProps := &CfnScalingPolicyProps{
//   	PolicyName: jsii.String("policyName"),
//   	PolicyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	PredictiveScalingPolicyConfiguration: &PredictiveScalingPolicyConfigurationProperty{
//   		MetricSpecifications: []interface{}{
//   			&PredictiveScalingMetricSpecificationProperty{
//   				TargetValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				CustomizedCapacityMetricSpecification: &PredictiveScalingCustomizedCapacityMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&PredictiveScalingMetricDataQueryProperty{
//   							Expression: jsii.String("expression"),
//   							Id: jsii.String("id"),
//   							Label: jsii.String("label"),
//   							MetricStat: &PredictiveScalingMetricStatProperty{
//   								Metric: &PredictiveScalingMetricProperty{
//   									Dimensions: []interface{}{
//   										&PredictiveScalingMetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//   								},
//   								Stat: jsii.String("stat"),
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				CustomizedLoadMetricSpecification: &PredictiveScalingCustomizedLoadMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&PredictiveScalingMetricDataQueryProperty{
//   							Expression: jsii.String("expression"),
//   							Id: jsii.String("id"),
//   							Label: jsii.String("label"),
//   							MetricStat: &PredictiveScalingMetricStatProperty{
//   								Metric: &PredictiveScalingMetricProperty{
//   									Dimensions: []interface{}{
//   										&PredictiveScalingMetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//   								},
//   								Stat: jsii.String("stat"),
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				CustomizedScalingMetricSpecification: &PredictiveScalingCustomizedScalingMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&PredictiveScalingMetricDataQueryProperty{
//   							Expression: jsii.String("expression"),
//   							Id: jsii.String("id"),
//   							Label: jsii.String("label"),
//   							MetricStat: &PredictiveScalingMetricStatProperty{
//   								Metric: &PredictiveScalingMetricProperty{
//   									Dimensions: []interface{}{
//   										&PredictiveScalingMetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//   								},
//   								Stat: jsii.String("stat"),
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				PredefinedLoadMetricSpecification: &PredictiveScalingPredefinedLoadMetricProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				PredefinedMetricPairSpecification: &PredictiveScalingPredefinedMetricPairProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				PredefinedScalingMetricSpecification: &PredictiveScalingPredefinedScalingMetricProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		MaxCapacityBreachBehavior: jsii.String("maxCapacityBreachBehavior"),
//   		MaxCapacityBuffer: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   		SchedulingBufferTime: jsii.Number(123),
//   	},
//   	ResourceId: jsii.String("resourceId"),
//   	ScalableDimension: jsii.String("scalableDimension"),
//   	ScalingTargetId: jsii.String("scalingTargetId"),
//   	ServiceNamespace: jsii.String("serviceNamespace"),
//   	StepScalingPolicyConfiguration: &StepScalingPolicyConfigurationProperty{
//   		AdjustmentType: jsii.String("adjustmentType"),
//   		Cooldown: jsii.Number(123),
//   		MetricAggregationType: jsii.String("metricAggregationType"),
//   		MinAdjustmentMagnitude: jsii.Number(123),
//   		StepAdjustments: []interface{}{
//   			&StepAdjustmentProperty{
//   				ScalingAdjustment: jsii.Number(123),
//
//   				// the properties below are optional
//   				MetricIntervalLowerBound: jsii.Number(123),
//   				MetricIntervalUpperBound: jsii.Number(123),
//   			},
//   		},
//   	},
//   	TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   		TargetValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		CustomizedMetricSpecification: &CustomizedMetricSpecificationProperty{
//   			Dimensions: []interface{}{
//   				&MetricDimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MetricName: jsii.String("metricName"),
//   			Metrics: []interface{}{
//   				&TargetTrackingMetricDataQueryProperty{
//   					Expression: jsii.String("expression"),
//   					Id: jsii.String("id"),
//   					Label: jsii.String("label"),
//   					MetricStat: &TargetTrackingMetricStatProperty{
//   						Metric: &TargetTrackingMetricProperty{
//   							Dimensions: []interface{}{
//   								&TargetTrackingMetricDimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Stat: jsii.String("stat"),
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   			Namespace: jsii.String("namespace"),
//   			Statistic: jsii.String("statistic"),
//   			Unit: jsii.String("unit"),
//   		},
//   		DisableScaleIn: jsii.Boolean(false),
//   		PredefinedMetricSpecification: &PredefinedMetricSpecificationProperty{
//   			PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   			// the properties below are optional
//   			ResourceLabel: jsii.String("resourceLabel"),
//   		},
//   		ScaleInCooldown: jsii.Number(123),
//   		ScaleOutCooldown: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html
//
type CfnScalingPolicyProps struct {
	// The name of the scaling policy.
	//
	// Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing `AWS::ApplicationAutoScaling::ScalingPolicy` resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The scaling policy type.
	//
	// The following policy types are supported:
	//
	// `TargetTrackingScaling` —Not supported for Amazon EMR
	//
	// `StepScaling` —Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
	//
	// `PredictiveScaling` —Only supported for Amazon ECS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policytype
	//
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// The predictive scaling policy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration
	//
	PredictiveScalingPolicyConfiguration interface{} `field:"optional" json:"predictiveScalingPolicyConfiguration" yaml:"predictiveScalingPolicyConfiguration"`
	// The identifier of the resource associated with the scaling policy.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalabledimension
	//
	ScalableDimension *string `field:"optional" json:"scalableDimension" yaml:"scalableDimension"`
	// The CloudFormation-generated ID of an Application Auto Scaling scalable target.
	//
	// For more information about the ID, see the Return Value section of the `AWS::ApplicationAutoScaling::ScalableTarget` resource.
	//
	// > You must specify either the `ScalingTargetId` property, or the `ResourceId` , `ScalableDimension` , and `ServiceNamespace` properties, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalingtargetid
	//
	ScalingTargetId *string `field:"optional" json:"scalingTargetId" yaml:"scalingTargetId"`
	// The namespace of the AWS service that provides the resource, or a `custom-resource` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-servicenamespace
	//
	ServiceNamespace *string `field:"optional" json:"serviceNamespace" yaml:"serviceNamespace"`
	// A step scaling policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration
	//
	StepScalingPolicyConfiguration interface{} `field:"optional" json:"stepScalingPolicyConfiguration" yaml:"stepScalingPolicyConfiguration"`
	// A target tracking scaling policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration
	//
	TargetTrackingScalingPolicyConfiguration interface{} `field:"optional" json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
}

