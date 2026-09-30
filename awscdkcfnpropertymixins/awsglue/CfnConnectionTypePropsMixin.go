package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Registers a custom connection type in Glue based on the configuration provided.
//
// This enables customers to configure custom connectors for any data source with REST-based APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnConnectionTypePropsMixin := awscdkcfnpropertymixins.Aws_glue.NewCfnConnectionTypePropsMixin(&CfnConnectionTypeMixinProps{
//   	ConnectionProperties: &ConnectionPropertiesConfigurationProperty{
//   		AdditionalRequestParameters: []interface{}{
//   			&ConnectorPropertyProperty{
//   				AllowedValues: []*string{
//   					jsii.String("allowedValues"),
//   				},
//   				DefaultValue: jsii.String("defaultValue"),
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   		},
//   		Url: &ConnectorPropertyProperty{
//   			AllowedValues: []*string{
//   				jsii.String("allowedValues"),
//   			},
//   			DefaultValue: jsii.String("defaultValue"),
//   			KeyOverride: jsii.String("keyOverride"),
//   			Name: jsii.String("name"),
//   			PropertyLocation: jsii.String("propertyLocation"),
//   			PropertyType: jsii.String("propertyType"),
//   			Required: jsii.Boolean(false),
//   		},
//   	},
//   	ConnectionType: jsii.String("connectionType"),
//   	ConnectorAuthenticationConfiguration: &ConnectorAuthenticationConfigurationProperty{
//   		AuthenticationTypes: []*string{
//   			jsii.String("authenticationTypes"),
//   		},
//   		BasicAuthenticationProperties: &BasicAuthenticationPropertiesProperty{
//   			Password: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   			Username: &SecretConnectorPropertyProperty{
//   				KeyOverride: jsii.String("keyOverride"),
//   				Name: jsii.String("name"),
//   				PropertyLocation: jsii.String("propertyLocation"),
//   				PropertyType: jsii.String("propertyType"),
//   				Required: jsii.Boolean(false),
//   			},
//   		},
//   		CustomAuthenticationProperties: &CustomAuthenticationPropertiesProperty{
//   			AuthenticationParameters: []interface{}{
//   				&SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		OAuth2Properties: &ConnectorOAuth2PropertiesProperty{
//   			AuthorizationCodeProperties: &ConnectorAuthorizationCodePropertiesProperty{
//   				AuthorizationCode: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				AuthorizationCodeUrl: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ClientId: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ClientSecret: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ContentType: jsii.String("contentType"),
//   				Prompt: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				RedirectUri: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				RequestMethod: jsii.String("requestMethod"),
//   				Scope: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrl: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						Name: jsii.String("name"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			ClientCredentialsProperties: &ClientCredentialsPropertiesProperty{
//   				ClientId: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ClientSecret: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				ContentType: jsii.String("contentType"),
//   				RequestMethod: jsii.String("requestMethod"),
//   				Scope: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrl: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						Name: jsii.String("name"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			JwtBearerProperties: &JWTBearerPropertiesProperty{
//   				ContentType: jsii.String("contentType"),
//   				JwtToken: &SecretConnectorPropertyProperty{
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				RequestMethod: jsii.String("requestMethod"),
//   				TokenUrl: &ConnectorPropertyProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   				TokenUrlParameters: []interface{}{
//   					&ConnectorPropertyProperty{
//   						AllowedValues: []*string{
//   							jsii.String("allowedValues"),
//   						},
//   						DefaultValue: jsii.String("defaultValue"),
//   						KeyOverride: jsii.String("keyOverride"),
//   						Name: jsii.String("name"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						PropertyType: jsii.String("propertyType"),
//   						Required: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IntegrationType: jsii.String("integrationType"),
//   	RestConfiguration: &RestConfigurationProperty{
//   		EntityConfigurations: map[string]interface{}{
//   			"entityConfigurationsKey": &EntityConfigurationProperty{
//   				"schema": map[string]interface{}{
//   					"schemaKey": &FieldDefinitionProperty{
//   						"fieldDataType": jsii.String("fieldDataType"),
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
//   						"name": jsii.String("name"),
//   						"responseDateFormat": jsii.String("responseDateFormat"),
//   					},
//   				},
//   				"sourceConfiguration": &SourceConfigurationProperty{
//   					"filterConfiguration": &FilterConfigurationProperty{
//   						"betweenConfiguration": &BetweenConfigurationProperty{
//   							"highBoundKey": jsii.String("highBoundKey"),
//   							"lowBoundKey": jsii.String("lowBoundKey"),
//   							"template": jsii.String("template"),
//   						},
//   						"dateTimeFormat": jsii.String("dateTimeFormat"),
//   						"filterMode": jsii.String("filterMode"),
//   						"filterStringConfiguration": &FilterStringConfigurationProperty{
//   							"queryParameterName": jsii.String("queryParameterName"),
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
//   							"limitParameter": &ExtractedParameterProperty{
//   								"defaultValue": jsii.String("defaultValue"),
//   								"key": jsii.String("key"),
//   								"propertyLocation": jsii.String("propertyLocation"),
//   								"value": &ResponseExtractionMappingProperty{
//   									"contentPath": jsii.String("contentPath"),
//   									"headerKey": jsii.String("headerKey"),
//   								},
//   							},
//   							"nextPage": &ExtractedParameterProperty{
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
//   							"allowedValues": []*string{
//   								jsii.String("allowedValues"),
//   							},
//   							"defaultValue": jsii.String("defaultValue"),
//   							"keyOverride": jsii.String("keyOverride"),
//   							"name": jsii.String("name"),
//   							"propertyLocation": jsii.String("propertyLocation"),
//   							"propertyType": jsii.String("propertyType"),
//   							"required": jsii.Boolean(false),
//   						},
//   					},
//   					"requestPath": jsii.String("requestPath"),
//   					"responseConfiguration": &ResponseConfigurationProperty{
//   						"errorPath": jsii.String("errorPath"),
//   						"resultPath": jsii.String("resultPath"),
//   					},
//   				},
//   			},
//   		},
//   		GlobalSourceConfiguration: &SourceConfigurationProperty{
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				BetweenConfiguration: &BetweenConfigurationProperty{
//   					HighBoundKey: jsii.String("highBoundKey"),
//   					LowBoundKey: jsii.String("lowBoundKey"),
//   					Template: jsii.String("template"),
//   				},
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				FilterMode: jsii.String("filterMode"),
//   				FilterStringConfiguration: &FilterStringConfigurationProperty{
//   					QueryParameterName: jsii.String("queryParameterName"),
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
//   					LimitParameter: &ExtractedParameterProperty{
//   						DefaultValue: jsii.String("defaultValue"),
//   						Key: jsii.String("key"),
//   						PropertyLocation: jsii.String("propertyLocation"),
//   						Value: &ResponseExtractionMappingProperty{
//   							ContentPath: jsii.String("contentPath"),
//   							HeaderKey: jsii.String("headerKey"),
//   						},
//   					},
//   					NextPage: &ExtractedParameterProperty{
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
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					DefaultValue: jsii.String("defaultValue"),
//   					KeyOverride: jsii.String("keyOverride"),
//   					Name: jsii.String("name"),
//   					PropertyLocation: jsii.String("propertyLocation"),
//   					PropertyType: jsii.String("propertyType"),
//   					Required: jsii.Boolean(false),
//   				},
//   			},
//   			RequestPath: jsii.String("requestPath"),
//   			ResponseConfiguration: &ResponseConfigurationProperty{
//   				ErrorPath: jsii.String("errorPath"),
//   				ResultPath: jsii.String("resultPath"),
//   			},
//   		},
//   		ValidationEndpointConfiguration: &ValidationEndpointConfigurationProperty{
//   			RequestMethod: jsii.String("requestMethod"),
//   			RequestPath: jsii.String("requestPath"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connectiontype.html
//
type CfnConnectionTypePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnConnectionTypeMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConnectionTypePropsMixin
type jsiiProxy_CfnConnectionTypePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnConnectionTypePropsMixin) Props() *CfnConnectionTypeMixinProps {
	var returns *CfnConnectionTypeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionTypePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::ConnectionType`.
func NewCfnConnectionTypePropsMixin(props *CfnConnectionTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnConnectionTypePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConnectionTypePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectionTypePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionTypePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::ConnectionType`.
func NewCfnConnectionTypePropsMixin_Override(c CfnConnectionTypePropsMixin, props *CfnConnectionTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionTypePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnConnectionTypePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectionTypePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionTypePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectionTypePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_glue.CfnConnectionTypePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectionTypePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConnectionTypePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

