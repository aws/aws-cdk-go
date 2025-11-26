package previewawscloudfrontmixins


// A complex data type for the status codes that you specify that, when returned by a primary origin, trigger CloudFront to failover to a second origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   statusCodesProperty := &StatusCodesProperty{
//   	Items: []interface{}{
//   		jsii.Number(123),
//   	},
//   	Quantity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-statuscodes.html
//
type CfnDistributionPropsMixin_StatusCodesProperty struct {
	// The items (status codes) for an origin group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-statuscodes.html#cfn-cloudfront-distribution-statuscodes-items
	//
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// The number of status codes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-statuscodes.html#cfn-cloudfront-distribution-statuscodes-quantity
	//
	Quantity *float64 `field:"optional" json:"quantity" yaml:"quantity"`
}

