package mixinsawsrtbfabric


// Describes the attributes of a link.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   			ResponseLoggingPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-linkattributes.html
//
type CfnLinkPropsMixin_LinkAttributesProperty struct {
	// The customer-provided unique identifier of the link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-linkattributes.html#cfn-rtbfabric-link-linkattributes-customerprovidedid
	//
	CustomerProvidedId *string `field:"optional" json:"customerProvidedId" yaml:"customerProvidedId"`
	// Describes the masking for HTTP error codes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-link-linkattributes.html#cfn-rtbfabric-link-linkattributes-respondererrormasking
	//
	ResponderErrorMasking interface{} `field:"optional" json:"responderErrorMasking" yaml:"responderErrorMasking"`
}

