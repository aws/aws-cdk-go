package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResponseProperty := &customResponseProperty{
//   	responseCode: jsii.Number(123),
//
//   	// the properties below are optional
//   	customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   	responseHeaders: []interface{}{
//   		&customHTTPHeaderProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_CustomResponseProperty struct {
	// `CfnRuleGroup.CustomResponseProperty.ResponseCode`.
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
	// `CfnRuleGroup.CustomResponseProperty.CustomResponseBodyKey`.
	CustomResponseBodyKey *string `field:"optional" json:"customResponseBodyKey" yaml:"customResponseBodyKey"`
	// `CfnRuleGroup.CustomResponseProperty.ResponseHeaders`.
	ResponseHeaders interface{} `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
}

