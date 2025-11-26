package previewawsappmeshmixins


// An object that represents an HTTP or HTTP/2 route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpRouteProperty := &HttpRouteProperty{
//   	Action: &HttpRouteActionProperty{
//   		WeightedTargets: []interface{}{
//   			&WeightedTargetProperty{
//   				Port: jsii.Number(123),
//   				VirtualNode: jsii.String("virtualNode"),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Match: &HttpRouteMatchProperty{
//   		Headers: []interface{}{
//   			&HttpRouteHeaderProperty{
//   				Invert: jsii.Boolean(false),
//   				Match: &HeaderMatchMethodProperty{
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   					Range: &MatchRangeProperty{
//   						End: jsii.Number(123),
//   						Start: jsii.Number(123),
//   					},
//   					Regex: jsii.String("regex"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Method: jsii.String("method"),
//   		Path: &HttpPathMatchProperty{
//   			Exact: jsii.String("exact"),
//   			Regex: jsii.String("regex"),
//   		},
//   		Port: jsii.Number(123),
//   		Prefix: jsii.String("prefix"),
//   		QueryParameters: []interface{}{
//   			&QueryParameterProperty{
//   				Match: &HttpQueryParameterMatchProperty{
//   					Exact: jsii.String("exact"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Scheme: jsii.String("scheme"),
//   	},
//   	RetryPolicy: &HttpRetryPolicyProperty{
//   		HttpRetryEvents: []*string{
//   			jsii.String("httpRetryEvents"),
//   		},
//   		MaxRetries: jsii.Number(123),
//   		PerRetryTimeout: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		TcpRetryEvents: []*string{
//   			jsii.String("tcpRetryEvents"),
//   		},
//   	},
//   	Timeout: &HttpTimeoutProperty{
//   		Idle: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		PerRequest: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproute.html
//
type CfnRoutePropsMixin_HttpRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproute.html#cfn-appmesh-route-httproute-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproute.html#cfn-appmesh-route-httproute-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// An object that represents a retry policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproute.html#cfn-appmesh-route-httproute-retrypolicy
	//
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents types of timeouts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httproute.html#cfn-appmesh-route-httproute-timeout
	//
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

