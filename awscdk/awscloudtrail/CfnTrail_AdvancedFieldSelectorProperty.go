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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html
//
type CfnTrail_AdvancedFieldSelectorProperty struct {
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
	// - `AWS::AppConfig::Configuration`
	// - `AWS::B2BI::Transformer`
	// - `AWS::Bedrock::AgentAlias`
	// - `AWS::Bedrock::FlowAlias`
	// - `AWS::Bedrock::Guardrail`
	// - `AWS::Bedrock::KnowledgeBase`
	// - `AWS::Cassandra::Table`
	// - `AWS::CloudFront::KeyValueStore`
	// - `AWS::CloudTrail::Channel`
	// - `AWS::CloudWatch::Metric`
	// - `AWS::CodeWhisperer::Customization`
	// - `AWS::CodeWhisperer::Profile`
	// - `AWS::Cognito::IdentityPool`
	// - `AWS::DynamoDB::Stream`
	// - `AWS::DynamoDB::Table`
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
	// - `AWS::Lambda::Function`
	// - `AWS::MachineLearning::MlModel`
	// - `AWS::ManagedBlockchain::Network`
	// - `AWS::ManagedBlockchain::Node`
	// - `AWS::MedicalImaging::Datastore`
	// - `AWS::NeptuneGraph::Graph`
	// - `AWS::PaymentCryptography::Alias`
	// - `AWS::PaymentCryptography::Key`
	// - `AWS::PCAConnectorAD::Connector`
	// - `AWS::PCAConnectorSCEP::Connector`
	// - `AWS::QApps:QApp`
	// - `AWS::QBusiness::Application`
	// - `AWS::QBusiness::DataSource`
	// - `AWS::QBusiness::Index`
	// - `AWS::QBusiness::WebExperience`
	// - `AWS::RDS::DBCluster`
	// - `AWS::S3::AccessPoint`
	// - `AWS::S3::Object`
	// - `AWS::S3Express::Object`
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
	// - *`resources.ARN`* - You can use any operator with `resources.ARN` , but if you use `Equals` or `NotEquals` , the value must exactly match the ARN of a valid resource of the type you've speciﬁed in the template as the value of resources.type. To log all data events for all objects in a specific S3 bucket, use the `StartsWith` operator, and include only the bucket ARN as the matching value. For information about filtering on the `resources.ARN` field, see [Filtering data events by resources.ARN](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/filtering-data-events.html#filtering-data-events-resourcearn) in the *AWS CloudTrail User Guide* .
	//
	// > You can't use the `resources.ARN` field to filter resource types that do not have ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-field
	//
	Field *string `field:"required" json:"field" yaml:"field"`
	// An operator that includes events that match the last few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-endswith
	//
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// An operator that includes events that match the exact value of the event record field specified as the value of `Field` .
	//
	// This is the only valid operator that you can use with the `readOnly` , `eventCategory` , and `resources.type` fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-equals
	//
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// An operator that excludes events that match the last few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-notendswith
	//
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// An operator that excludes events that match the exact value of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-notequals
	//
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// An operator that excludes events that match the first few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-notstartswith
	//
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// An operator that includes events that match the first few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedfieldselector.html#cfn-cloudtrail-trail-advancedfieldselector-startswith
	//
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

