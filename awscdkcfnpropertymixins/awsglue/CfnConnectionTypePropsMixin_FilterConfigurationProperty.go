package awsglue


// Configuration that defines how filter predicates are applied to REST API requests, supporting both query parameter and filter string strategies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   filterConfigurationProperty := &FilterConfigurationProperty{
//   	BetweenConfiguration: &BetweenConfigurationProperty{
//   		HighBoundKey: jsii.String("highBoundKey"),
//   		LowBoundKey: jsii.String("lowBoundKey"),
//   		Template: jsii.String("template"),
//   	},
//   	DateTimeFormat: jsii.String("dateTimeFormat"),
//   	FilterMode: jsii.String("filterMode"),
//   	FilterStringConfiguration: &FilterStringConfigurationProperty{
//   		QueryParameterName: jsii.String("queryParameterName"),
//   		QuoteCharacter: jsii.String("quoteCharacter"),
//   		QuoteStringValues: jsii.Boolean(false),
//   	},
//   	OperatorMappings: map[string]*string{
//   		"operatorMappingsKey": jsii.String("operatorMappings"),
//   	},
//   	StripQuotes: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterconfiguration.html
//
type CfnConnectionTypePropsMixin_FilterConfigurationProperty struct {
	// Configuration that defines how BETWEEN range filter operations are translated into REST API request parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterconfiguration.html#cfn-glue-connectiontype-filterconfiguration-betweenconfiguration
	//
	BetweenConfiguration interface{} `field:"optional" json:"betweenConfiguration" yaml:"betweenConfiguration"`
	// The global date and time format for filter expressions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterconfiguration.html#cfn-glue-connectiontype-filterconfiguration-datetimeformat
	//
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// The strategy for applying filters to requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterconfiguration.html#cfn-glue-connectiontype-filterconfiguration-filtermode
	//
	FilterMode *string `field:"optional" json:"filterMode" yaml:"filterMode"`
	// Configuration for constructing filter expression strings when using the FILTER_STRING filter mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterconfiguration.html#cfn-glue-connectiontype-filterconfiguration-filterstringconfiguration
	//
	FilterStringConfiguration interface{} `field:"optional" json:"filterStringConfiguration" yaml:"filterStringConfiguration"`
	// A map of logical filter operators to their API-specific string representations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterconfiguration.html#cfn-glue-connectiontype-filterconfiguration-operatormappings
	//
	OperatorMappings interface{} `field:"optional" json:"operatorMappings" yaml:"operatorMappings"`
	// Indicates whether surrounding double quotes should be stripped from filter values before processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterconfiguration.html#cfn-glue-connectiontype-filterconfiguration-stripquotes
	//
	StripQuotes interface{} `field:"optional" json:"stripQuotes" yaml:"stripQuotes"`
}

