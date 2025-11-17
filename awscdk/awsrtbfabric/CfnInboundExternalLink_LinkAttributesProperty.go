package awsrtbfabric


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linkAttributesProperty := &LinkAttributesProperty{
//   	CustomerProvidedId: jsii.String("customerProvidedId"),
//   	ResponderErrorMasking: []interface{}{
//   		&ResponderErrorMaskingForHttpCodeProperty{
//   			Action: jsii.String("action"),
//   			HttpCode: jsii.String("httpCode"),
//   			LoggingTypes: []*string{
//   				jsii.String("loggingTypes"),
//   			},
//
//   			// the properties below are optional
//   			ResponseLoggingPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-linkattributes.html
//
type CfnInboundExternalLink_LinkAttributesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-linkattributes.html#cfn-rtbfabric-inboundexternallink-linkattributes-customerprovidedid
	//
	CustomerProvidedId *string `field:"optional" json:"customerProvidedId" yaml:"customerProvidedId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-inboundexternallink-linkattributes.html#cfn-rtbfabric-inboundexternallink-linkattributes-respondererrormasking
	//
	ResponderErrorMasking interface{} `field:"optional" json:"responderErrorMasking" yaml:"responderErrorMasking"`
}

