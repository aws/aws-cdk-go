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
//   	ConcurrentExecutorConfiguration: &ConcurrentExecutorConfigurationProperty{
//   		FunctionList: []interface{}{
//   			&FunctionRefProperty{
//   				Alias: jsii.String("alias"),
//   				FunctionId: jsii.String("functionId"),
//   				RunCondition: jsii.String("runCondition"),
//   			},
//   		},
//   		MaxConcurrency: jsii.Number(123),
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		Runtime: jsii.String("runtime"),
//   		TimeoutMilliseconds: jsii.Number(123),
//   	},
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
//   				Alias: jsii.String("alias"),
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
	// The configuration for a CONCURRENT_EXECUTOR function.
	//
	// A CONCURRENT_EXECUTOR runs a set of child functions in parallel, up to a maximum concurrency, and combines their output when all functions complete. For more information about functions, see Working with functions (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions.html) in the MediaTailor User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-concurrentexecutorconfiguration
	//
	ConcurrentExecutorConfiguration interface{} `field:"optional" json:"concurrentExecutorConfiguration" yaml:"concurrentExecutorConfiguration"`
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
	// The configuration for a SEQUENTIAL_EXECUTOR function.
	//
	// A SEQUENTIAL_EXECUTOR runs an ordered list of child functions one at a time, passing data between them. For more information about functions, see Working with functions (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions.html) in the MediaTailor User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-sequentialexecutorconfiguration
	//
	SequentialExecutorConfiguration interface{} `field:"optional" json:"sequentialExecutorConfiguration" yaml:"sequentialExecutorConfiguration"`
	// The tags to assign to the function resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

