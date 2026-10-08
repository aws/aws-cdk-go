package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFunctionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFunctionMixinProps := &CfnFunctionMixinProps{
//   	AwsServiceRequestConfiguration: &AwsServiceRequestConfigurationProperty{
//   		Body: jsii.String("body"),
//   		Headers: map[string]*string{
//   			"headersKey": jsii.String("headers"),
//   		},
//   		MethodType: jsii.String("methodType"),
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		RequestTimeoutMilliseconds: jsii.Number(123),
//   		Runtime: jsii.String("runtime"),
//   		TargetRegion: jsii.String("targetRegion"),
//   		TargetService: jsii.String("targetService"),
//   		Url: jsii.String("url"),
//   	},
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
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		Runtime: jsii.String("runtime"),
//   	},
//   	Description: jsii.String("description"),
//   	FunctionId: jsii.String("functionId"),
//   	FunctionType: jsii.String("functionType"),
//   	HttpRequestConfiguration: &HttpRequestConfigurationProperty{
//   		Body: jsii.String("body"),
//   		Headers: map[string]*string{
//   			"headersKey": jsii.String("headers"),
//   		},
//   		MethodType: jsii.String("methodType"),
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		RequestTimeoutMilliseconds: jsii.Number(123),
//   		Runtime: jsii.String("runtime"),
//   		Url: jsii.String("url"),
//   	},
//   	SequentialExecutorConfiguration: &SequentialExecutorConfigurationProperty{
//   		FunctionList: []interface{}{
//   			&FunctionRefProperty{
//   				Alias: jsii.String("alias"),
//   				FunctionId: jsii.String("functionId"),
//   				RunCondition: jsii.String("runCondition"),
//   			},
//   		},
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		Runtime: jsii.String("runtime"),
//   		TimeoutMilliseconds: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VastRequestConfiguration: &VastRequestConfigurationProperty{
//   		Body: jsii.String("body"),
//   		Headers: map[string]*string{
//   			"headersKey": jsii.String("headers"),
//   		},
//   		MethodType: jsii.String("methodType"),
//   		Output: map[string]*string{
//   			"outputKey": jsii.String("output"),
//   		},
//   		RequestTimeoutMilliseconds: jsii.Number(123),
//   		Runtime: jsii.String("runtime"),
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html
//
type CfnFunctionMixinProps struct {
	// The configuration for an AWS_SERVICE_REQUEST function.
	//
	// Contains the target service, target Region, and request parameters that the function uses to call an AWS service API. For more information, see AWS_SERVICE_REQUEST (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-types-aws-service-request.html) in the MediaTailor User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-awsservicerequestconfiguration
	//
	AwsServiceRequestConfiguration interface{} `field:"optional" json:"awsServiceRequestConfiguration" yaml:"awsServiceRequestConfiguration"`
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
	// The unique identifier for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-functionid
	//
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-functiontype
	//
	FunctionType *string `field:"optional" json:"functionType" yaml:"functionType"`
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
	// The configuration for a VAST_REQUEST function.
	//
	// Specifies the HTTP method, URL, headers, body, timeout, and output expressions for a request to a VAST endpoint. MediaTailor parses the response as VAST and resolves wrapper redirects, then makes the parsed ads available to the function's output expressions. For more information, see Function types and composition (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-types.html) in the MediaTailor User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-function.html#cfn-mediatailor-function-vastrequestconfiguration
	//
	VastRequestConfiguration interface{} `field:"optional" json:"vastRequestConfiguration" yaml:"vastRequestConfiguration"`
}

