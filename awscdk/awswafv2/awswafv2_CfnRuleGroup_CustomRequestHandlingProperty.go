package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customRequestHandlingProperty := &customRequestHandlingProperty{
//   	insertHeaders: []interface{}{
//   		&customHTTPHeaderProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_CustomRequestHandlingProperty struct {
	// `CfnRuleGroup.CustomRequestHandlingProperty.InsertHeaders`.
	InsertHeaders interface{} `field:"required" json:"insertHeaders" yaml:"insertHeaders"`
}

