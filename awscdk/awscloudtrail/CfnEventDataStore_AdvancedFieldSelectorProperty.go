package awscloudtrail


// A single selector statement in an advanced event selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedFieldSelectorProperty := &AdvancedFieldSelectorProperty{
//   	Field: jsii.String("field"),
//
//   	// the properties below are optional
//   	EndsWith: []*string{
//   		jsii.String("endsWith"),
//   	},
//   	EqualTo: []*string{
//   		jsii.String("equalTo"),
//   	},
//   	NotEndsWith: []*string{
//   		jsii.String("notEndsWith"),
//   	},
//   	NotEquals: []*string{
//   		jsii.String("notEquals"),
//   	},
//   	NotStartsWith: []*string{
//   		jsii.String("notStartsWith"),
//   	},
//   	StartsWith: []*string{
//   		jsii.String("startsWith"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html
//
type CfnEventDataStore_AdvancedFieldSelectorProperty struct {
	// A field in a CloudTrail event record on which to filter events to be logged.
	//
	// For event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS , the field is used only for selecting events as filtering is not supported.
	//
	// For CloudTrail management events, supported fields include `readOnly` , `eventCategory` , and `eventSource` .
	//
	// For CloudTrail data events, supported fields include `readOnly` , `eventCategory` , `eventName` , `resources.type` , and `resources.ARN` .
	//
	// For event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS , the only supported field is `eventCategory` .
	//
	// - *`readOnly`* - Optional. Can be set to `Equals` a value of `true` or `false` . If you do not add this field, CloudTrail logs both `read` and `write` events. A value of `true` logs only `read` events. A value of `false` logs only `write` events.
	// - *`eventSource`* - For filtering management events only. This can be set to `NotEquals` `kms.amazonaws.com` or `NotEquals` `rdsdata.amazonaws.com` .
	// - *`eventName`* - Can use any operator. You can use it to ﬁlter in or ﬁlter out any data event logged to CloudTrail, such as `PutBucket` or `GetSnapshotBlock` . You can have multiple values for this ﬁeld, separated by commas.
	// - *`eventCategory`* - This is required and must be set to `Equals` .
	//
	// - For CloudTrail management events, the value must be `Management` .
	// - For CloudTrail data events, the value must be `Data` .
	//
	// The following are used only for event data stores:
	//
	// - For CloudTrail Insights events, the value must be `Insight` .
	// - For AWS Config configuration items, the value must be `ConfigurationItem` .
	// - For Audit Manager evidence, the value must be `Evidence` .
	// - For non- AWS events, the value must be `ActivityAuditLog` .
	// - *`resources.type`* - This ﬁeld is required for CloudTrail data events. `resources.type` can only use the `Equals` operator, and the value can be one of the following:
	//
	// - `AWS::DynamoDB::Table`
	// - `AWS::Lambda::Function`
	// - `AWS::S3::Object`
	// - `AWS::AppConfig::Configuration`
	// - `AWS::B2BI::Transformer`
	// - `AWS::Bedrock::AgentAlias`
	// - `AWS::Bedrock::KnowledgeBase`
	// - `AWS::Cassandra::Table`
	// - `AWS::CloudFront::KeyValueStore`
	// - `AWS::CloudTrail::Channel`
	// - `AWS::CloudWatch::Metric`
	// - `AWS::CodeWhisperer::Customization`
	// - `AWS::CodeWhisperer::Profile`
	// - `AWS::Cognito::IdentityPool`
	// - `AWS::DynamoDB::Stream`
	// - `AWS::EC2::Snapshot`
	// - `AWS::EMRWAL::Workspace`
	// - `AWS::FinSpace::Environment`
	// - `AWS::Glue::Table`
	// - `AWS::GreengrassV2::ComponentVersion`
	// - `AWS::GreengrassV2::Deployment`
	// - `AWS::GuardDuty::Detector`
	// - `AWS::IoT::Certificate`
	// - `AWS::IoT::Thing`
	// - `AWS::IoTSiteWise::Asset`
	// - `AWS::IoTSiteWise::TimeSeries`
	// - `AWS::IoTTwinMaker::Entity`
	// - `AWS::IoTTwinMaker::Workspace`
	// - `AWS::KendraRanking::ExecutionPlan`
	// - `AWS::Kinesis::Stream`
	// - `AWS::Kinesis::StreamConsumer`
	// - `AWS::KinesisVideo::Stream`
	// - `AWS::MachineLearning::MlModel`
	// - `AWS::ManagedBlockchain::Network`
	// - `AWS::ManagedBlockchain::Node`
	// - `AWS::MedicalImaging::Datastore`
	// - `AWS::NeptuneGraph::Graph`
	// - `AWS::PCAConnectorAD::Connector`
	// - `AWS::PCAConnectorSCEP::Connector`
	// - `AWS::QApps:QApp`
	// - `AWS::QBusiness::Application`
	// - `AWS::QBusiness::DataSource`
	// - `AWS::QBusiness::Index`
	// - `AWS::QBusiness::WebExperience`
	// - `AWS::RDS::DBCluster`
	// - `AWS::S3::AccessPoint`
	// - `AWS::S3ObjectLambda::AccessPoint`
	// - `AWS::S3Outposts::Object`
	// - `AWS::SageMaker::Endpoint`
	// - `AWS::SageMaker::ExperimentTrialComponent`
	// - `AWS::SageMaker::FeatureGroup`
	// - `AWS::ServiceDiscovery::Namespace`
	// - `AWS::ServiceDiscovery::Service`
	// - `AWS::SCN::Instance`
	// - `AWS::SNS::PlatformEndpoint`
	// - `AWS::SNS::Topic`
	// - `AWS::SQS::Queue`
	// - `AWS::SSM::ManagedNode`
	// - `AWS::SSMMessages::ControlChannel`
	// - `AWS::StepFunctions::StateMachine`
	// - `AWS::SWF::Domain`
	// - `AWS::ThinClient::Device`
	// - `AWS::ThinClient::Environment`
	// - `AWS::Timestream::Database`
	// - `AWS::Timestream::Table`
	// - `AWS::VerifiedPermissions::PolicyStore`
	// - `AWS::XRay::Trace`
	//
	// You can have only one `resources.type` ﬁeld per selector. To log data events on more than one resource type, add another selector.
	// - *`resources.ARN`* - You can use any operator with `resources.ARN` , but if you use `Equals` or `NotEquals` , the value must exactly match the ARN of a valid resource of the type you've speciﬁed in the template as the value of resources.type.
	//
	// > You can't use the `resources.ARN` field to filter resource types that do not have ARNs.
	//
	// The `resources.ARN` field can be set one of the following.
	//
	// If resources.type equals `AWS::S3::Object` , the ARN must be in one of the following formats. To log all data events for all objects in a specific S3 bucket, use the `StartsWith` operator, and include only the bucket ARN as the matching value.
	//
	// The trailing slash is intentional; do not exclude it. Replace the text between less than and greater than symbols (<>) with resource-specific information.
	//
	// - `arn:<partition>:s3:::<bucket_name>/`
	// - `arn:<partition>:s3:::<bucket_name>/<object_path>/`
	//
	// When resources.type equals `AWS::DynamoDB::Table` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:dynamodb:<region>:<account_ID>:table/<table_name>`
	//
	// When resources.type equals `AWS::Lambda::Function` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:lambda:<region>:<account_ID>:function:<function_name>`
	//
	// When resources.type equals `AWS::AppConfig::Configuration` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:appconfig:<region>:<account_ID>:application/<application_ID>/environment/<environment_ID>/configuration/<configuration_profile_ID>`
	//
	// When resources.type equals `AWS::B2BI::Transformer` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:b2bi:<region>:<account_ID>:transformer/<transformer_ID>`
	//
	// When resources.type equals `AWS::Bedrock::AgentAlias` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:bedrock:<region>:<account_ID>:agent-alias/<agent_ID>/<alias_ID>`
	//
	// When resources.type equals `AWS::Bedrock::KnowledgeBase` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:bedrock:<region>:<account_ID>:knowledge-base/<knowledge_base_ID>`
	//
	// When resources.type equals `AWS::Cassandra::Table` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:cassandra:<region>:<account_ID>:/keyspace/<keyspace_name>/table/<table_name>`
	//
	// When resources.type equals `AWS::CloudFront::KeyValueStore` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:cloudfront:<region>:<account_ID>:key-value-store/<KVS_name>`
	//
	// When resources.type equals `AWS::CloudTrail::Channel` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:cloudtrail:<region>:<account_ID>:channel/<channel_UUID>`
	//
	// When resources.type equals `AWS::CodeWhisperer::Customization` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:codewhisperer:<region>:<account_ID>:customization/<customization_ID>`
	//
	// When resources.type equals `AWS::CodeWhisperer::Profile` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:codewhisperer:<region>:<account_ID>:profile/<profile_ID>`
	//
	// When resources.type equals `AWS::Cognito::IdentityPool` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:cognito-identity:<region>:<account_ID>:identitypool/<identity_pool_ID>`
	//
	// When `resources.type` equals `AWS::DynamoDB::Stream` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:dynamodb:<region>:<account_ID>:table/<table_name>/stream/<date_time>`
	//
	// When `resources.type` equals `AWS::EC2::Snapshot` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:ec2:<region>::snapshot/<snapshot_ID>`
	//
	// When `resources.type` equals `AWS::EMRWAL::Workspace` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:emrwal:<region>:<account_ID>:workspace/<workspace_name>`
	//
	// When `resources.type` equals `AWS::FinSpace::Environment` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:finspace:<region>:<account_ID>:environment/<environment_ID>`
	//
	// When `resources.type` equals `AWS::Glue::Table` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:glue:<region>:<account_ID>:table/<database_name>/<table_name>`
	//
	// When `resources.type` equals `AWS::GreengrassV2::ComponentVersion` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:greengrass:<region>:<account_ID>:components/<component_name>`
	//
	// When `resources.type` equals `AWS::GreengrassV2::Deployment` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:greengrass:<region>:<account_ID>:deployments/<deployment_ID`
	//
	// When `resources.type` equals `AWS::GuardDuty::Detector` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:guardduty:<region>:<account_ID>:detector/<detector_ID>`
	//
	// When `resources.type` equals `AWS::IoT::Certificate` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:iot:<region>:<account_ID>:cert/<certificate_ID>`
	//
	// When `resources.type` equals `AWS::IoT::Thing` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:iot:<region>:<account_ID>:thing/<thing_ID>`
	//
	// When `resources.type` equals `AWS::IoTSiteWise::Asset` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:iotsitewise:<region>:<account_ID>:asset/<asset_ID>`
	//
	// When `resources.type` equals `AWS::IoTSiteWise::TimeSeries` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:iotsitewise:<region>:<account_ID>:timeseries/<timeseries_ID>`
	//
	// When `resources.type` equals `AWS::IoTTwinMaker::Entity` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:iottwinmaker:<region>:<account_ID>:workspace/<workspace_ID>/entity/<entity_ID>`
	//
	// When `resources.type` equals `AWS::IoTTwinMaker::Workspace` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:iottwinmaker:<region>:<account_ID>:workspace/<workspace_ID>`
	//
	// When `resources.type` equals `AWS::KendraRanking::ExecutionPlan` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:kendra-ranking:<region>:<account_ID>:rescore-execution-plan/<rescore_execution_plan_ID>`
	//
	// When `resources.type` equals `AWS::Kinesis::Stream` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:kinesis:<region>:<account_ID>:stream/<stream_name>`
	//
	// When `resources.type` equals `AWS::Kinesis::Stream` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:kinesis:<region>:<account_ID>:<stream_type>/<stream_name>/consumer/<consumer_name>:<consumer_creation_timestamp>`
	//
	// When `resources.type` equals `AWS::KinesisVideo::Stream` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:kinesisvideo:<region>:<account_ID>:stream/<stream_name>/<creation_time>`
	//
	// When `resources.type` equals `AWS::MachineLearning::MlModel` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:machinelearning:<region>:<account_ID>:mlmodel/<model_ID>`
	//
	// When `resources.type` equals `AWS::ManagedBlockchain::Network` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:managedblockchain:::networks/<network_name>`
	//
	// When `resources.type` equals `AWS::ManagedBlockchain::Node` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:managedblockchain:<region>:<account_ID>:nodes/<node_ID>`
	//
	// When `resources.type` equals `AWS::MedicalImaging::Datastore` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:medical-imaging:<region>:<account_ID>:datastore/<data_store_ID>`
	//
	// When `resources.type` equals `AWS::NeptuneGraph::Graph` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:neptune-graph:<region>:<account_ID>:graph/<graph_ID>`
	//
	// When `resources.type` equals `AWS::PCAConnectorAD::Connector` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:pca-connector-ad:<region>:<account_ID>:connector/<connector_ID>`
	//
	// When `resources.type` equals `AWS::PCAConnectorSCEP::Connector` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:pca-connector-scep:<region>:<account_ID>:connector/<connector_ID>`
	//
	// When `resources.type` equals `AWS::QApps:QApp` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:qapps:<region>:<account_ID>:application/<application_UUID>/qapp/<qapp_UUID>`
	//
	// When `resources.type` equals `AWS::QBusiness::Application` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:qbusiness:<region>:<account_ID>:application/<application_ID>`
	//
	// When `resources.type` equals `AWS::QBusiness::DataSource` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:qbusiness:<region>:<account_ID>:application/<application_ID>/index/<index_ID>/data-source/<datasource_ID>`
	//
	// When `resources.type` equals `AWS::QBusiness::Index` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:qbusiness:<region>:<account_ID>:application/<application_ID>/index/<index_ID>`
	//
	// When `resources.type` equals `AWS::QBusiness::WebExperience` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:qbusiness:<region>:<account_ID>:application/<application_ID>/web-experience/<web_experience_ID>`
	//
	// When `resources.type` equals `AWS::RDS::DBCluster` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:rds:<region>:<account_ID>:cluster/<cluster_name>`
	//
	// When `resources.type` equals `AWS::S3::AccessPoint` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in one of the following formats. To log events on all objects in an S3 access point, we recommend that you use only the access point ARN, don’t include the object path, and use the `StartsWith` or `NotStartsWith` operators.
	//
	// - `arn:<partition>:s3:<region>:<account_ID>:accesspoint/<access_point_name>`
	// - `arn:<partition>:s3:<region>:<account_ID>:accesspoint/<access_point_name>/object/<object_path>`
	//
	// When `resources.type` equals `AWS::S3ObjectLambda::AccessPoint` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:s3-object-lambda:<region>:<account_ID>:accesspoint/<access_point_name>`
	//
	// When `resources.type` equals `AWS::S3Outposts::Object` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:s3-outposts:<region>:<account_ID>:<object_path>`
	//
	// When `resources.type` equals `AWS::SageMaker::Endpoint` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:sagemaker:<region>:<account_ID>:endpoint/<endpoint_name>`
	//
	// When `resources.type` equals `AWS::SageMaker::ExperimentTrialComponent` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:sagemaker:<region>:<account_ID>:experiment-trial-component/<experiment_trial_component_name>`
	//
	// When `resources.type` equals `AWS::SageMaker::FeatureGroup` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:sagemaker:<region>:<account_ID>:feature-group/<feature_group_name>`
	//
	// When `resources.type` equals `AWS::SCN::Instance` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:scn:<region>:<account_ID>:instance/<instance_ID>`
	//
	// When `resources.type` equals `AWS::ServiceDiscovery::Namespace` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:servicediscovery:<region>:<account_ID>:namespace/<namespace_ID>`
	//
	// When `resources.type` equals `AWS::ServiceDiscovery::Service` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:servicediscovery:<region>:<account_ID>:service/<service_ID>`
	//
	// When `resources.type` equals `AWS::SNS::PlatformEndpoint` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:sns:<region>:<account_ID>:endpoint/<endpoint_type>/<endpoint_name>/<endpoint_ID>`
	//
	// When `resources.type` equals `AWS::SNS::Topic` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:sns:<region>:<account_ID>:<topic_name>`
	//
	// When `resources.type` equals `AWS::SQS::Queue` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:sqs:<region>:<account_ID>:<queue_name>`
	//
	// When `resources.type` equals `AWS::SSM::ManagedNode` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in one of the following formats:
	//
	// - `arn:<partition>:ssm:<region>:<account_ID>:managed-instance/<instance_ID>`
	// - `arn:<partition>:ec2:<region>:<account_ID>:instance/<instance_ID>`
	//
	// When `resources.type` equals `AWS::SSMMessages::ControlChannel` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:ssmmessages:<region>:<account_ID>:control-channel/<channel_ID>`
	//
	// When `resources.type` equals `AWS::StepFunctions::StateMachine` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in one of the following formats:
	//
	// - `arn:<partition>:states:<region>:<account_ID>:stateMachine:<stateMachine_name>`
	// - `arn:<partition>:states:<region>:<account_ID>:stateMachine:<stateMachine_name>/<label_name>`
	//
	// When `resources.type` equals `AWS::SWF::Domain` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:swf:<region>:<account_ID>:domain/<domain_name>`
	//
	// When `resources.type` equals `AWS::ThinClient::Device` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:thinclient:<region>:<account_ID>:device/<device_ID>`
	//
	// When `resources.type` equals `AWS::ThinClient::Environment` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:thinclient:<region>:<account_ID>:environment/<environment_ID>`
	//
	// When `resources.type` equals `AWS::Timestream::Database` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:timestream:<region>:<account_ID>:database/<database_name>`
	//
	// When `resources.type` equals `AWS::Timestream::Table` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:timestream:<region>:<account_ID>:database/<database_name>/table/<table_name>`
	//
	// When resources.type equals `AWS::VerifiedPermissions::PolicyStore` , and the operator is set to `Equals` or `NotEquals` , the ARN must be in the following format:
	//
	// - `arn:<partition>:verifiedpermissions:<region>:<account_ID>:policy-store/<policy_store_UUID>`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-field
	//
	Field *string `field:"required" json:"field" yaml:"field"`
	// An operator that includes events that match the last few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-endswith
	//
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// An operator that includes events that match the exact value of the event record field specified as the value of `Field` .
	//
	// This is the only valid operator that you can use with the `readOnly` , `eventCategory` , and `resources.type` fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-equals
	//
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// An operator that excludes events that match the last few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-notendswith
	//
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// An operator that excludes events that match the exact value of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-notequals
	//
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// An operator that excludes events that match the first few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-notstartswith
	//
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// An operator that includes events that match the first few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-startswith
	//
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

