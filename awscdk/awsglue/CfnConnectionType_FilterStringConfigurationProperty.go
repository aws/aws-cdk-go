package awsglue


// Configuration for constructing filter expression strings when using the FILTER_STRING filter mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterStringConfigurationProperty := &FilterStringConfigurationProperty{
//   	QueryParameterName: jsii.String("queryParameterName"),
//
//   	// the properties below are optional
//   	QuoteCharacter: jsii.String("quoteCharacter"),
//   	QuoteStringValues: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterstringconfiguration.html
//
type CfnConnectionType_FilterStringConfigurationProperty struct {
	// The query parameter name used to send the constructed filter expression string in API requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterstringconfiguration.html#cfn-glue-connectiontype-filterstringconfiguration-queryparametername
	//
	QueryParameterName *string `field:"required" json:"queryParameterName" yaml:"queryParameterName"`
	// The character used to quote values when QuoteStringValues is true.
	//
	// Defaults to double quotes if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterstringconfiguration.html#cfn-glue-connectiontype-filterstringconfiguration-quotecharacter
	//
	QuoteCharacter *string `field:"optional" json:"quoteCharacter" yaml:"quoteCharacter"`
	// Indicates whether string and date values should be wrapped with a quote character in the filter expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-filterstringconfiguration.html#cfn-glue-connectiontype-filterstringconfiguration-quotestringvalues
	//
	QuoteStringValues interface{} `field:"optional" json:"quoteStringValues" yaml:"quoteStringValues"`
}

