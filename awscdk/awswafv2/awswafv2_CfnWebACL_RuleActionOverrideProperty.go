package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleActionOverrideProperty := &ruleActionOverrideProperty{
//   	actionToUse: &ruleActionProperty{
//   		allow: &allowActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		block: &blockActionProperty{
//   			customResponse: &customResponseProperty{
//   				responseCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   				responseHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		captcha: &captchaActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		challenge: &challengeActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		count: &countActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnWebACL_RuleActionOverrideProperty struct {
	// `CfnWebACL.RuleActionOverrideProperty.ActionToUse`.
	ActionToUse interface{} `field:"required" json:"actionToUse" yaml:"actionToUse"`
	// `CfnWebACL.RuleActionOverrideProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
}

