package awsobservabilityadmin


// Specifies a field in the request to redact from WAF logs, such as headers, query parameters, or body content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldToMatchProperty := &FieldToMatchProperty{
//   	Method: jsii.String("method"),
//   	QueryString: jsii.String("queryString"),
//   	SingleHeader: &SingleHeaderProperty{
//   		Name: jsii.String("name"),
//   	},
//   	UriPath: jsii.String("uriPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-fieldtomatch.html
//
type CfnTelemetryRule_FieldToMatchProperty struct {
	// Redacts the HTTP method from WAF logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-fieldtomatch.html#cfn-observabilityadmin-telemetryrule-fieldtomatch-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Redacts the entire query string from WAF logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-fieldtomatch.html#cfn-observabilityadmin-telemetryrule-fieldtomatch-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Redacts a specific header field by name from WAF logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-fieldtomatch.html#cfn-observabilityadmin-telemetryrule-fieldtomatch-singleheader
	//
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Redacts the URI path from WAF logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-fieldtomatch.html#cfn-observabilityadmin-telemetryrule-fieldtomatch-uripath
	//
	UriPath *string `field:"optional" json:"uriPath" yaml:"uriPath"`
}

