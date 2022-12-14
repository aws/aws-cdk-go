package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockProperty := &blockProperty{
//   	customResponse: &customResponseProperty{
//   		responseCode: jsii.Number(123),
//
//   		// the properties below are optional
//   		customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   		responseHeaders: []interface{}{
//   			&customHTTPHeaderProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_BlockProperty struct {
	// `CfnRuleGroup.BlockProperty.CustomResponse`.
	CustomResponse interface{} `field:"optional" json:"customResponse" yaml:"customResponse"`
}

