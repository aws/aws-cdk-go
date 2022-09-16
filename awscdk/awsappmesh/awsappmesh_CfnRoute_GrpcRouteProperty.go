package awsappmesh


// An object that represents a gRPC route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcRouteProperty := &grpcRouteProperty{
//   	action: &grpcRouteActionProperty{
//   		weightedTargets: []interface{}{
//   			&weightedTargetProperty{
//   				virtualNode: jsii.String("virtualNode"),
//   				weight: jsii.Number(123),
//
//   				// the properties below are optional
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	match: &grpcRouteMatchProperty{
//   		metadata: []interface{}{
//   			&grpcRouteMetadataProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				invert: jsii.Boolean(false),
//   				match: &grpcRouteMetadataMatchMethodProperty{
//   					exact: jsii.String("exact"),
//   					prefix: jsii.String("prefix"),
//   					range: &matchRangeProperty{
//   						end: jsii.Number(123),
//   						start: jsii.Number(123),
//   					},
//   					regex: jsii.String("regex"),
//   					suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   		methodName: jsii.String("methodName"),
//   		port: jsii.Number(123),
//   		serviceName: jsii.String("serviceName"),
//   	},
//
//   	// the properties below are optional
//   	retryPolicy: &grpcRetryPolicyProperty{
//   		maxRetries: jsii.Number(123),
//   		perRetryTimeout: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		grpcRetryEvents: []*string{
//   			jsii.String("grpcRetryEvents"),
//   		},
//   		httpRetryEvents: []*string{
//   			jsii.String("httpRetryEvents"),
//   		},
//   		tcpRetryEvents: []*string{
//   			jsii.String("tcpRetryEvents"),
//   		},
//   	},
//   	timeout: &grpcTimeoutProperty{
//   		idle: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//   		perRequest: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnRoute_GrpcRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// An object that represents a retry policy.
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents types of timeouts.
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

