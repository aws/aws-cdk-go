package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockProperty := &BlockProperty{
//   	CustomResponse: &CustomResponseProperty{
//   		ResponseCode: jsii.Number(123),
//
//   		// the properties below are optional
//   		CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   		ResponseHeaders: []interface{}{
//   			&CustomHTTPHeaderProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_BlockProperty struct {
	// `CfnRuleGroup.BlockProperty.CustomResponse`.
	CustomResponse interface{} `field:"optional" json:"customResponse" yaml:"customResponse"`
}

