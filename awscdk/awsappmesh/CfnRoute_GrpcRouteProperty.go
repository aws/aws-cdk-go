package awsappmesh


// An object that represents a gRPC route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcRouteProperty := &GrpcRouteProperty{
//   	Action: &GrpcRouteActionProperty{
//   		WeightedTargets: []interface{}{
//   			&WeightedTargetProperty{
//   				VirtualNode: jsii.String("virtualNode"),
//   				Weight: jsii.Number(123),
//
//   				// the properties below are optional
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Match: &GrpcRouteMatchProperty{
//   		Metadata: []interface{}{
//   			&GrpcRouteMetadataProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Invert: jsii.Boolean(false),
//   				Match: &GrpcRouteMetadataMatchMethodProperty{
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   					Range: &MatchRangeProperty{
//   						End: jsii.Number(123),
//   						Start: jsii.Number(123),
//   					},
//   					Regex: jsii.String("regex"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   		MethodName: jsii.String("methodName"),
//   		Port: jsii.Number(123),
//   		ServiceName: jsii.String("serviceName"),
//   	},
//
//   	// the properties below are optional
//   	RetryPolicy: &GrpcRetryPolicyProperty{
//   		MaxRetries: jsii.Number(123),
//   		PerRetryTimeout: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		GrpcRetryEvents: []*string{
//   			jsii.String("grpcRetryEvents"),
//   		},
//   		HttpRetryEvents: []*string{
//   			jsii.String("httpRetryEvents"),
//   		},
//   		TcpRetryEvents: []*string{
//   			jsii.String("tcpRetryEvents"),
//   		},
//   	},
//   	Timeout: &GrpcTimeoutProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroute.html
//
type CfnRoute_GrpcRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroute.html#cfn-appmesh-route-grpcroute-action
	//
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroute.html#cfn-appmesh-route-grpcroute-match
	//
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// An object that represents a retry policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroute.html#cfn-appmesh-route-grpcroute-retrypolicy
	//
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents types of timeouts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroute.html#cfn-appmesh-route-grpcroute-timeout
	//
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

