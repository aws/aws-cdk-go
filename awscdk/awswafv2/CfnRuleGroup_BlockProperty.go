package awswafv2


// Block traffic towards application.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-block.html
//
type CfnRuleGroup_BlockProperty struct {
	// Custom response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-block.html#cfn-wafv2-rulegroup-block-customresponse
	//
	CustomResponse interface{} `field:"optional" json:"customResponse" yaml:"customResponse"`
}

