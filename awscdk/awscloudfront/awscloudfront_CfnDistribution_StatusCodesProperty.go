package awscloudfront


// A complex data type for the status codes that you specify that, when returned by a primary origin, trigger CloudFront to failover to a second origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statusCodesProperty := &statusCodesProperty{
//   	items: []interface{}{
//   		jsii.Number(123),
//   	},
//   	quantity: jsii.Number(123),
//   }
//
type CfnDistribution_StatusCodesProperty struct {
	// The items (status codes) for an origin group.
	Items interface{} `field:"required" json:"items" yaml:"items"`
	// The number of status codes.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

