package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFunction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunctionProps := &CfnFunctionProps{
//   	FunctionId: jsii.String("functionId"),
//   	FunctionType: jsii.String("functionType"),
//
//   	// the properties below are optional
//   	CustomOutputConfiguration: &CustomOutputConfigurationProperty{
//   		Runtime: jsii.String("runtime"),
//
//   		// the properties below are optional
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	HttpRequestConfiguration: &HttpRequestConfigurationProperty{
//   		MethodType: jsii.String("methodType"),
//   		RequestTimeoutMilliseconds: jsii.Number(123),
//   		Runtime: jsii.String("runtime"),
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		Body: jsii.String("body"),
//   		Headers: map[string]*string{
//   			"headersKey": jsii.String("headers"),
//   		},
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   	},
//   	SequentialExecutorConfiguration: &SequentialExecutorConfigurationProperty{
//   		FunctionList: []interface{}{
//   			&FunctionRefProperty{
//   				FunctionId: jsii.String("functionId"),
//   				RunCondition: jsii.String("runCondition"),
//   			},
//   		},
//   		Runtime: jsii.String("runtime"),
//   		TimeoutMilliseconds: jsii.Number(123),
//
//   		// the properties below are optional
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html
//
type CfnFunctionProps struct {
	// The unique identifier for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-functionid
	//
	FunctionId *string `field:"required" json:"functionId" yaml:"functionId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-functiontype
	//
	FunctionType *string `field:"required" json:"functionType" yaml:"functionType"`
	// Configuration for custom output functions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-customoutputconfiguration
	//
	CustomOutputConfiguration interface{} `field:"optional" json:"customOutputConfiguration" yaml:"customOutputConfiguration"`
	// A description of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration for HTTP request functions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-httprequestconfiguration
	//
	HttpRequestConfiguration interface{} `field:"optional" json:"httpRequestConfiguration" yaml:"httpRequestConfiguration"`
	// Configuration for sequential executor functions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-sequentialexecutorconfiguration
	//
	SequentialExecutorConfiguration interface{} `field:"optional" json:"sequentialExecutorConfiguration" yaml:"sequentialExecutorConfiguration"`
	// The tags to assign to the function resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

