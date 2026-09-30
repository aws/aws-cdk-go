package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Registers a custom connection type in Glue based on the configuration provided.
//
// This enables customers to configure custom connectors for any data source with REST-based APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionType := awscdk.Aws_glue.NewCfnConnectionType(this, jsii.String("MyCfnConnectionType"), &CfnConnectionTypeProps{
//   	ConnectionType: jsii.String("connectionType"),
//   	RestConfiguration: &RestConfigurationProperty{
//   		EntityConfigurations: map[string]interface{}{
//   			"entityConfigurationsKey": &EntityConfigurationProperty{
//   				"schema": map[string]interface{}{
//   					"schemaKey": &FieldDefinitionProperty{
//   						"fieldDataType": jsii.String("fieldDataType"),
//   						"name": jsii.String("name"),
//
//   						// the properties below are optional
//   						"filterOverrides": &FilterOverridesProperty{
//   							"betweenConfiguration": &BetweenConfigurationProperty{
//   								"highBoundKey": jsii.String("highBoundKey"),
//   								"lowBoundKey": jsii.String("lowBoundKey"),
//   								"template": jsii.String("template"),
//   							},
//   							"dateTimeFormat": jsii.String("dateTimeFormat"),
//   							"fieldName": jsii.String("fieldName"),
//   							"operatorMappings": map[string]*string{
//   								"operatorMappingsKey": jsii.String("operatorMappings"),
//   							},
//   						},
//   						"isNullable": jsii.Boolean(false),
//   						"isOrderable": jsii.Boolean(false),
//   						"isPartitionable": jsii.Boolean(false),
//   						"isQueryable": jsii.Boolean(false),
//   						"responseDateFormat": jsii.String("responseDateFormat"),
//   					},
//   				},
//   				"sourceConfiguration": &SourceConfigurationProperty{
//   					"filterConfiguration": &FilterConfigurationProperty{
//   						"filterMode": jsii.String("filterMode"),
//
//   						// the properties below are optional
//   						"betweenConfiguration": &BetweenConfigurationProperty{
//   							"highBoundKey": jsii.String("highBoundKey"),
//   							"lowBoundKey": jsii.String("lowBoundKey"),
//   							"template": jsii.String("template"),
//   						},
//   						"dateTimeFormat": jsii.String("dateTimeFormat"),
//   						"filterStringConfiguration": &FilterStringConfigurationProperty{
//   							"queryParameterName": jsii.String("queryParameterName"),
//
//   							// the properties below are optional
//   							"quoteCharacter": jsii.String("quoteCharacter"),
//   							"quoteStringValues": jsii.Boolean(false),
//   						},
//   						"operatorMappings": map[string]*string{
//   							"operatorMappingsKey": jsii.String("operatorMappings"),
//   						},
//   						"stripQuotes": jsii.Boolean(false),
//   					},
//   					"paginationConfiguration": &PaginationConfigurationProperty{
//   						"cursorConfiguration": &CursorConfigurationProperty{
//   							"nextPage": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							"limitParameter": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//   						},
//   						"offsetConfiguration": &OffsetConfigurationProperty{
//   							"limitParameter": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//   							"offsetParameter": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//   						},
//   					},
//   					"requestMethod": jsii.String("requestMethod"),
//   					"requestParameters": []interface{}{
//   						&ConnectorPropertyProperty{
//   							"name": jsii.String("name"),
//   							"propertyType": jsii.String("propertyType"),
//   							"required": jsii.Boolean(false),
//
//   							// the properties below are optional
//   							"allowedValues": []*string{
//   								jsii.String("allowedValues"),
//   							},
//   							"defaultValue": jsii.String("defaultValue"),
//   							"keyOverride": jsii.String("keyOverride"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   						},
//   					},
//   					"requestPath": jsii.String("requestPath"),
//   					"responseConfiguration": &ResponseConfigurationProperty{
//   						"resultPath": jsii.String("resultPath"),
//
//   						// the properties below are optional
//   						"errorPath": jsii.String("errorPath"),
//   					},
//   				},
//   			},
//   		},
//   		GlobalSourceConfiguration: &SourceConfigurationProperty{
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				FilterMode: jsii.String("filterMode"),
//
//   				// the properties below are optional
//   				BetweenConfiguration: &BetweenConfigurationProperty{
//   					HighBoundKey: jsii.String("highBoundKey"),
//   					LowBoundKey: jsii.String("lowBoundKey"),
//   					Template: jsii.String("template"),
//   				},
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				FilterStringConfiguration: &FilterStringConfigurationProperty{
//   					QueryParameterName: jsii.String("queryParameterName"),
//
//   					// the properties below are optional
//   					QuoteCharacter: jsii.String("quoteCharacter"),
//   					QuoteStringValues: jsii.Boolean(false),
//   				},
//   				OperatorMappings: map[string]*string{
//   					"operatorMappingsKey": jsii.String("operatorMappings"),
//   				},
//   				StripQuotes: jsii.Boolean(false),
//   			},
//   			PaginationConfiguration: &PaginationConfigurationProperty{
//   				CursorConfiguration: &CursorConfigurationProperty{
//   					NextPage: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					LimitParameter: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//   				},
//   				OffsetConfiguration: &OffsetConfigurationProperty{
//   					LimitParameter: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//   					OffsetParameter: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//   				},
//   			},
//   			RequestMethod: jsii.String("requestMethod"),
//   			RequestParameters: []interface{}{
//   				&ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   			},
//   			RequestPath: jsii.String("requestPath"),
//   			ResponseConfiguration: &ResponseConfigurationProperty{
//   				ResultPath: jsii.String("resultPath"),
//
//   				// the properties below are optional
//   				ErrorPath: jsii.String("errorPath"),
//   			},
//   		},
//   		ValidationEndpointConfiguration: &ValidationEndpointConfigurationProperty{
//   			RequestMethod: jsii.String("requestMethod"),
//   			RequestPath: jsii.String("requestPath"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ConnectionProperties: &ConnectionPropertiesConfigurationProperty{
//   		AdditionalRequestParameters: []interface{}{
//   			&ConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   		},
//   		Url: &ConnectorPropertyProperty{
//   			Name: jsii.String("name"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   		},
//   	},
//   	ConnectorAuthenticationConfiguration: &ConnectorAuthenticationConfigurationProperty{
//   		AuthenticationTypes: []*string{
//   			jsii.String("authenticationTypes"),
//   		},
//
//   		// the properties below are optional
//   		BasicAuthenticationProperties: &BasicAuthenticationPropertiesProperty{
//   			Password: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   			Username: &SecretConnectorPropertyProperty{
//   				Name: jsii.String("name"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				KeyOverride: jsii.String("keyOverride"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   			},
//   		},
//   		CustomAuthenticationProperties: &CustomAuthenticationPropertiesProperty{
//   			AuthenticationParameters: []interface{}{
//   				&SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   			},
//   		},
//   		OAuth2Properties: &ConnectorOAuth2PropertiesProperty{
//   			OAuth2GrantType: jsii.String("oAuth2GrantType"),
//
//   			// the properties below are optional
//   			AuthorizationCodeProperties: &ConnectorAuthorizationCodePropertiesProperty{
//   				AuthorizationCode: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ClientId: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ClientSecret: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ContentType: jsii.String("contentType"),
//   				Prompt: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				RedirectUri: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				RequestMethod: jsii.String("requestMethod"),
//   				Scope: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrl: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						Name: jsii.String("name"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   					},
//   				},
//   			},
//   			ClientCredentialsProperties: &ClientCredentialsPropertiesProperty{
//   				ClientId: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ClientSecret: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				ContentType: jsii.String("contentType"),
//   				RequestMethod: jsii.String("requestMethod"),
//   				Scope: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrl: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						Name: jsii.String("name"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   					},
//   				},
//   			},
//   			JwtBearerProperties: &JWTBearerPropertiesProperty{
//   				ContentType: jsii.String("contentType"),
//   				JwtToken: &SecretConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				RequestMethod: jsii.String("requestMethod"),
//   				TokenUrl: &ConnectorPropertyProperty{
//   					Name: jsii.String("name"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						Name: jsii.String("name"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IntegrationType: jsii.String("integrationType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html
//
type CfnConnectionType interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsglue.IConnectionTypeRef
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN) of the registered connection type.
	AttrConnectionTypeArn() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Configuration that defines the base URL and additional request parameters needed during connection creation.
	ConnectionProperties() interface{}
	SetConnectionProperties(val interface{})
	// The name of the connection type.
	ConnectionType() *string
	SetConnectionType(val *string)
	// A reference to a ConnectionType resource.
	ConnectionTypeRef() *interfacesawsglue.ConnectionTypeReference
	// Configuration that defines supported authentication types and required properties.
	ConnectorAuthenticationConfiguration() interface{}
	SetConnectorAuthenticationConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the connection type.
	Description() *string
	SetDescription(val *string)
	Env() *interfaces.ResourceEnvironment
	// The integration type for the connection.
	IntegrationType() *string
	SetIntegrationType(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Configuration for HTTP request and response handling.
	RestConfiguration() interface{}
	SetRestConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags to assign to the connection type.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This method has been renamed to `addResourceDependency` to more clearly
	// set it apart from `construct.node.addDependency`. See the documentation
	// of that function for more details.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method has been renamed to `addResourceDependency`, which makes it
	// more clear that this method operates at a different level from the
	// construct-level `construct.node.addDependency()` mechanism.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	CfnPropertyName(cdkPropertyName *string) *string
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources and stacks this resource depends on.
	//
	// For resources depended on directly, returns the `CfnResource` object. For
	// dependencies on other stacks, returns the `Stack` object. The order of the
	// array is not guaranteed.
	ObtainDependencies() *[]interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Deprecated: Use `removeResourceDependency` instead.
	RemoveDependency(target awscdk.CfnResource)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveResourceDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnConnectionType
type jsiiProxy_CfnConnectionType struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsglueIConnectionTypeRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnConnectionType) AttrConnectionTypeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConnectionTypeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) ConnectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) ConnectionTypeRef() *interfacesawsglue.ConnectionTypeReference {
	var returns *interfacesawsglue.ConnectionTypeReference
	_jsii_.Get(
		j,
		"connectionTypeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) ConnectorAuthenticationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectorAuthenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) IntegrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) RestConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionType) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::ConnectionType`.
func NewCfnConnectionType(scope constructs.Construct, id *string, props *CfnConnectionTypeProps) CfnConnectionType {
	_init_.Initialize()

	if err := validateNewCfnConnectionTypeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectionType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnConnectionType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::ConnectionType`.
func NewCfnConnectionType_Override(c CfnConnectionType, scope constructs.Construct, id *string, props *CfnConnectionTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnConnectionType",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnectionType)SetConnectionProperties(val interface{}) {
	if err := j.validateSetConnectionPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProperties",
		val,
	)
}

func (j *jsiiProxy_CfnConnectionType)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_CfnConnectionType)SetConnectorAuthenticationConfiguration(val interface{}) {
	if err := j.validateSetConnectorAuthenticationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorAuthenticationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnConnectionType)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnConnectionType)SetIntegrationType(val *string) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_CfnConnectionType)SetRestConfiguration(val interface{}) {
	if err := j.validateSetRestConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnConnectionType)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnConnectionType_ArnForConnectionType(resource interfacesawsglue.IConnectionTypeRef) *string {
	_init_.Initialize()

	if err := validateCfnConnectionType_ArnForConnectionTypeParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnConnectionType",
		"arnForConnectionType",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnConnectionType.
func CfnConnectionType_IsCfnConnectionType(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectionType_IsCfnConnectionTypeParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnConnectionType",
		"isCfnConnectionType",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnectionType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectionType_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnConnectionType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnConnectionType_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectionType_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnConnectionType",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnConnectionType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectionType_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnConnectionType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectionType_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnConnectionType",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectionType) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnectionType) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectionType) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectionType) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnectionType) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnectionType) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnectionType) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnectionType) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnConnectionType) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnConnectionType) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConnectionType) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectionType) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectionType) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnConnectionType) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnectionType) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectionType) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnectionType) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectionType) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectionType) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnConnectionType) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnConnectionType) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectionType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectionType) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnConnectionType) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

