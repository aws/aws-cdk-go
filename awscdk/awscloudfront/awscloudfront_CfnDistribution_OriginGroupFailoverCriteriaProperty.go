package awscloudfront


// A complex data type that includes information about the failover criteria for an origin group, including the status codes for which CloudFront will failover from the primary origin to the second origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originGroupFailoverCriteriaProperty := &originGroupFailoverCriteriaProperty{
//   	statusCodes: &statusCodesProperty{
//   		items: []interface{}{
//   			jsii.Number(123),
//   		},
//   		quantity: jsii.Number(123),
//   	},
//   }
//
type CfnDistribution_OriginGroupFailoverCriteriaProperty struct {
	// The status codes that, when returned from the primary origin, will trigger CloudFront to failover to the second origin.
	StatusCodes interface{} `field:"required" json:"statusCodes" yaml:"statusCodes"`
}

