package awswafv2


// Inspect one of the headers in the web request, identified by name, for example, `User-Agent` or `Referer` .
//
// The name isn't case sensitive.
//
// You can filter and inspect all headers with the `FieldToMatch` setting `Headers` .
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example JSON: `"SingleHeader": { "Name": "haystack" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleHeaderProperty := &SingleHeaderProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-singleheader.html
//
type CfnLoggingConfiguration_SingleHeaderProperty struct {
	// The name of the query header to inspect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-singleheader.html#cfn-wafv2-loggingconfiguration-singleheader-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

