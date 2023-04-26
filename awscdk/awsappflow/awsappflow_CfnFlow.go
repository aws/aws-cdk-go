package awsappflow

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AppFlow::Flow`.
//
// The `AWS::AppFlow::Flow` resource is an Amazon AppFlow resource type that specifies a new flow.
//
// > If you want to use AWS CloudFormation to create a connector profile for connectors that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth, or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow console to create the connector profile, and then use that connector profile in the flow creation CloudFormation template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlow := awscdk.Aws_appflow.NewCfnFlow(this, jsii.String("MyCfnFlow"), &CfnFlowProps{
//   	DestinationFlowConfigList: []interface{}{
//   		&DestinationFlowConfigProperty{
//   			ConnectorType: jsii.String("connectorType"),
//   			DestinationConnectorProperties: &DestinationConnectorPropertiesProperty{
//   				CustomConnector: &CustomConnectorDestinationPropertiesProperty{
//   					EntityName: jsii.String("entityName"),
//
//   					// the properties below are optional
//   					CustomProperties: map[string]*string{
//   						"customPropertiesKey": jsii.String("customProperties"),
//   					},
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   				EventBridge: &EventBridgeDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				LookoutMetrics: &LookoutMetricsDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//   				},
//   				Marketo: &MarketoDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				Redshift: &RedshiftDestinationPropertiesProperty{
//   					IntermediateBucketName: jsii.String("intermediateBucketName"),
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				S3: &S3DestinationPropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					S3OutputFormatConfig: &S3OutputFormatConfigProperty{
//   						AggregationConfig: &AggregationConfigProperty{
//   							AggregationType: jsii.String("aggregationType"),
//   							TargetFileSize: jsii.Number(123),
//   						},
//   						FileType: jsii.String("fileType"),
//   						PrefixConfig: &PrefixConfigProperty{
//   							PathPrefixHierarchy: []*string{
//   								jsii.String("pathPrefixHierarchy"),
//   							},
//   							PrefixFormat: jsii.String("prefixFormat"),
//   							PrefixType: jsii.String("prefixType"),
//   						},
//   						PreserveSourceDataTyping: jsii.Boolean(false),
//   					},
//   				},
//   				Salesforce: &SalesforceDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					DataTransferApi: jsii.String("dataTransferApi"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   				SapoData: &SAPODataDestinationPropertiesProperty{
//   					ObjectPath: jsii.String("objectPath"),
//
//   					// the properties below are optional
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					SuccessResponseHandlingConfig: &SuccessResponseHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   				Snowflake: &SnowflakeDestinationPropertiesProperty{
//   					IntermediateBucketName: jsii.String("intermediateBucketName"),
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				Upsolver: &UpsolverDestinationPropertiesProperty{
//   					BucketName: jsii.String("bucketName"),
//   					S3OutputFormatConfig: &UpsolverS3OutputFormatConfigProperty{
//   						PrefixConfig: &PrefixConfigProperty{
//   							PathPrefixHierarchy: []*string{
//   								jsii.String("pathPrefixHierarchy"),
//   							},
//   							PrefixFormat: jsii.String("prefixFormat"),
//   							PrefixType: jsii.String("prefixType"),
//   						},
//
//   						// the properties below are optional
//   						AggregationConfig: &AggregationConfigProperty{
//   							AggregationType: jsii.String("aggregationType"),
//   							TargetFileSize: jsii.Number(123),
//   						},
//   						FileType: jsii.String("fileType"),
//   					},
//
//   					// the properties below are optional
//   					BucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				Zendesk: &ZendeskDestinationPropertiesProperty{
//   					Object: jsii.String("object"),
//
//   					// the properties below are optional
//   					ErrorHandlingConfig: &ErrorHandlingConfigProperty{
//   						BucketName: jsii.String("bucketName"),
//   						BucketPrefix: jsii.String("bucketPrefix"),
//   						FailOnFirstError: jsii.Boolean(false),
//   					},
//   					IdFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					WriteOperationType: jsii.String("writeOperationType"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ApiVersion: jsii.String("apiVersion"),
//   			ConnectorProfileName: jsii.String("connectorProfileName"),
//   		},
//   	},
//   	FlowName: jsii.String("flowName"),
//   	SourceFlowConfig: &SourceFlowConfigProperty{
//   		ConnectorType: jsii.String("connectorType"),
//   		SourceConnectorProperties: &SourceConnectorPropertiesProperty{
//   			Amplitude: &AmplitudeSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			CustomConnector: &CustomConnectorSourcePropertiesProperty{
//   				EntityName: jsii.String("entityName"),
//
//   				// the properties below are optional
//   				CustomProperties: map[string]*string{
//   					"customPropertiesKey": jsii.String("customProperties"),
//   				},
//   			},
//   			Datadog: &DatadogSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Dynatrace: &DynatraceSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			GoogleAnalytics: &GoogleAnalyticsSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			InforNexus: &InforNexusSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Marketo: &MarketoSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Pardot: &PardotSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			S3: &S3SourcePropertiesProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketPrefix: jsii.String("bucketPrefix"),
//
//   				// the properties below are optional
//   				S3InputFormatConfig: &S3InputFormatConfigProperty{
//   					S3InputFileType: jsii.String("s3InputFileType"),
//   				},
//   			},
//   			Salesforce: &SalesforceSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//
//   				// the properties below are optional
//   				DataTransferApi: jsii.String("dataTransferApi"),
//   				EnableDynamicFieldUpdate: jsii.Boolean(false),
//   				IncludeDeletedRecords: jsii.Boolean(false),
//   			},
//   			SapoData: &SAPODataSourcePropertiesProperty{
//   				ObjectPath: jsii.String("objectPath"),
//   			},
//   			ServiceNow: &ServiceNowSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Singular: &SingularSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Slack: &SlackSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Trendmicro: &TrendmicroSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   			Veeva: &VeevaSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//
//   				// the properties below are optional
//   				DocumentType: jsii.String("documentType"),
//   				IncludeAllVersions: jsii.Boolean(false),
//   				IncludeRenditions: jsii.Boolean(false),
//   				IncludeSourceFiles: jsii.Boolean(false),
//   			},
//   			Zendesk: &ZendeskSourcePropertiesProperty{
//   				Object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		ApiVersion: jsii.String("apiVersion"),
//   		ConnectorProfileName: jsii.String("connectorProfileName"),
//   		IncrementalPullConfig: &IncrementalPullConfigProperty{
//   			DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	Tasks: []interface{}{
//   		&TaskProperty{
//   			SourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			TaskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			ConnectorOperator: &ConnectorOperatorProperty{
//   				Amplitude: jsii.String("amplitude"),
//   				CustomConnector: jsii.String("customConnector"),
//   				Datadog: jsii.String("datadog"),
//   				Dynatrace: jsii.String("dynatrace"),
//   				GoogleAnalytics: jsii.String("googleAnalytics"),
//   				InforNexus: jsii.String("inforNexus"),
//   				Marketo: jsii.String("marketo"),
//   				Pardot: jsii.String("pardot"),
//   				S3: jsii.String("s3"),
//   				Salesforce: jsii.String("salesforce"),
//   				SapoData: jsii.String("sapoData"),
//   				ServiceNow: jsii.String("serviceNow"),
//   				Singular: jsii.String("singular"),
//   				Slack: jsii.String("slack"),
//   				Trendmicro: jsii.String("trendmicro"),
//   				Veeva: jsii.String("veeva"),
//   				Zendesk: jsii.String("zendesk"),
//   			},
//   			DestinationField: jsii.String("destinationField"),
//   			TaskProperties: []interface{}{
//   				&TaskPropertiesObjectProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TriggerConfig: &TriggerConfigProperty{
//   		TriggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		ActivateFlowOnCreate: jsii.Boolean(false),
//   		TriggerProperties: &ScheduledTriggerPropertiesProperty{
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			DataPullMode: jsii.String("dataPullMode"),
//   			FirstExecutionFrom: jsii.Number(123),
//   			FlowErrorDeactivationThreshold: jsii.Number(123),
//   			ScheduleEndTime: jsii.Number(123),
//   			ScheduleOffset: jsii.Number(123),
//   			ScheduleStartTime: jsii.Number(123),
//   			TimeZone: jsii.String("timeZone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsArn: jsii.String("kmsArn"),
//   	MetadataCatalogConfig: &MetadataCatalogConfigProperty{
//   		GlueDataCatalog: &GlueDataCatalogProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			RoleArn: jsii.String("roleArn"),
//   			TablePrefix: jsii.String("tablePrefix"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFlow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The flow's Amazon Resource Name (ARN).
	AttrFlowArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A user-entered description of the flow.
	Description() *string
	SetDescription(val *string)
	// The configuration that controls how Amazon AppFlow places data in the destination connector.
	DestinationFlowConfigList() interface{}
	SetDestinationFlowConfigList(val interface{})
	// The specified name of the flow.
	//
	// Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	FlowName() *string
	SetFlowName(val *string)
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn() *string
	SetKmsArn(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// `AWS::AppFlow::Flow.MetadataCatalogConfig`.
	MetadataCatalogConfig() interface{}
	SetMetadataCatalogConfig(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Contains information about the configuration of the source connector used in the flow.
	SourceFlowConfig() interface{}
	SetSourceFlowConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tags used to organize, track, or control access for your flow.
	Tags() awscdk.TagManager
	// A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	Tasks() interface{}
	SetTasks(val interface{})
	// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	TriggerConfig() interface{}
	SetTriggerConfig(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFlow
type jsiiProxy_CfnFlow struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlow) AttrFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) DestinationFlowConfigList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationFlowConfigList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) FlowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) KmsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) MetadataCatalogConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataCatalogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) SourceFlowConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceFlowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Tasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) TriggerConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppFlow::Flow`.
func NewCfnFlow(scope awscdk.Construct, id *string, props *CfnFlowProps) CfnFlow {
	_init_.Initialize()

	if err := validateNewCfnFlowParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlow{}

	_jsii_.Create(
		"monocdk.aws_appflow.CfnFlow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppFlow::Flow`.
func NewCfnFlow_Override(c CfnFlow, scope awscdk.Construct, id *string, props *CfnFlowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appflow.CfnFlow",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlow)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetDestinationFlowConfigList(val interface{}) {
	if err := j.validateSetDestinationFlowConfigListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFlowConfigList",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetFlowName(val *string) {
	if err := j.validateSetFlowNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowName",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetKmsArn(val *string) {
	_jsii_.Set(
		j,
		"kmsArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetMetadataCatalogConfig(val interface{}) {
	if err := j.validateSetMetadataCatalogConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataCatalogConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetSourceFlowConfig(val interface{}) {
	if err := j.validateSetSourceFlowConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFlowConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetTasks(val interface{}) {
	if err := j.validateSetTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tasks",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetTriggerConfig(val interface{}) {
	if err := j.validateSetTriggerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnFlow_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnFlow",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFlow_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnFlow",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFlow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appflow.CfnFlow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlow_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appflow.CfnFlow",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlow) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFlow) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlow) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFlow) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFlow) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFlow) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFlow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFlow) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFlow) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFlow) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFlow) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlow) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFlow) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFlow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

