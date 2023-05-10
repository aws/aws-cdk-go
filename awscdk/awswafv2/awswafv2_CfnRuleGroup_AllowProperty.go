package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowProperty := &AllowProperty{
//   	CustomRequestHandling: &CustomRequestHandlingProperty{
//   		InsertHeaders: []interface{}{
//   			&CustomHTTPHeaderProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_AllowProperty struct {
	// `CfnRuleGroup.AllowProperty.CustomRequestHandling`.
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

