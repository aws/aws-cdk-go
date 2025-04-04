package awsappmesh


// An object that represents the criteria for determining a request match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grpcRouteMatchProperty := &GrpcRouteMatchProperty{
//   	Metadata: []interface{}{
//   		&GrpcRouteMetadataProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Invert: jsii.Boolean(false),
//   			Match: &GrpcRouteMetadataMatchMethodProperty{
//   				Exact: jsii.String("exact"),
//   				Prefix: jsii.String("prefix"),
//   				Range: &MatchRangeProperty{
//   					End: jsii.Number(123),
//   					Start: jsii.Number(123),
//   				},
//   				Regex: jsii.String("regex"),
//   				Suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	MethodName: jsii.String("methodName"),
//   	Port: jsii.Number(123),
//   	ServiceName: jsii.String("serviceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutematch.html
//
type CfnRoute_GrpcRouteMatchProperty struct {
	// An object that represents the data to match from the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutematch.html#cfn-appmesh-route-grpcroutematch-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The method name to match from the request.
	//
	// If you specify a name, you must also specify a `serviceName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutematch.html#cfn-appmesh-route-grpcroutematch-methodname
	//
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
	// The port number to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutematch.html#cfn-appmesh-route-grpcroutematch-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The fully qualified domain name for the service to match from the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcroutematch.html#cfn-appmesh-route-grpcroutematch-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

