package awsappmesh


// An object that represents the criteria for determining a request match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcRouteMatchProperty := &grpcRouteMatchProperty{
//   	metadata: []interface{}{
//   		&grpcRouteMetadataProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			invert: jsii.Boolean(false),
//   			match: &grpcRouteMetadataMatchMethodProperty{
//   				exact: jsii.String("exact"),
//   				prefix: jsii.String("prefix"),
//   				range: &matchRangeProperty{
//   					end: jsii.Number(123),
//   					start: jsii.Number(123),
//   				},
//   				regex: jsii.String("regex"),
//   				suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	methodName: jsii.String("methodName"),
//   	port: jsii.Number(123),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type CfnRoute_GrpcRouteMatchProperty struct {
	// An object that represents the data to match from the request.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The method name to match from the request.
	//
	// If you specify a name, you must also specify a `serviceName` .
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
	// `CfnRoute.GrpcRouteMatchProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The fully qualified domain name for the service to match from the request.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

