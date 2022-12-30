package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   removeHeadersConfigProperty := &removeHeadersConfigProperty{
//   	items: []interface{}{
//   		&removeHeaderProperty{
//   			header: jsii.String("header"),
//   		},
//   	},
//   }
//
type CfnResponseHeadersPolicy_RemoveHeadersConfigProperty struct {
	// `CfnResponseHeadersPolicy.RemoveHeadersConfigProperty.Items`.
	Items interface{} `field:"required" json:"items" yaml:"items"`
}

