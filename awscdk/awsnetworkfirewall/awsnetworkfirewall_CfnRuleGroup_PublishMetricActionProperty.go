package awsnetworkfirewall


// Stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet.
//
// This setting defines a CloudWatch dimension value to be published.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publishMetricActionProperty := &publishMetricActionProperty{
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_PublishMetricActionProperty struct {
	// `CfnRuleGroup.PublishMetricActionProperty.Dimensions`.
	Dimensions interface{} `field:"required" json:"dimensions" yaml:"dimensions"`
}

