package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Where a WebFunction endpoint serves traffic.
//
// A regional location carries its replica Regions, and a home-Region one has none, so a
// combination the service rejects cannot be expressed.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var codeBucket IBucket
//
//
//   lambda.NewWebFunction(this, jsii.String("GlobalApi"), &WebFunctionProps{
//   	Code: lambda.WebFunctionCode_FromBucket(codeBucket, jsii.String("app.zip"), jsii.String("version-id")),
//   	Runtime: lambda.WebFunctionRuntime_NODEJS_24_X(),
//   	DefaultEndpointAuth: lambda.WebFunctionAuthType_IAM_AUTH,
//   	DefaultEndpoint: &WebFunctionDefaultEndpointOptions{
//   		EndpointLocation: lambda.WebFunctionEndpointLocation_MultiRegion([]*string{
//   			jsii.String("us-east-1"),
//   			jsii.String("eu-west-1"),
//   			jsii.String("ap-northeast-1"),
//   		}),
//   	},
//   })
//
type WebFunctionEndpointLocation interface {
	// The endpoint type this location maps to.
	EndpointType() WebFunctionEndpointType
	// The replica Regions, for a regional location, as given.
	ReplicaRegions() *[]*string
}

// The jsii proxy struct for WebFunctionEndpointLocation
type jsiiProxy_WebFunctionEndpointLocation struct {
	_ byte // padding
}

func (j *jsiiProxy_WebFunctionEndpointLocation) EndpointType() WebFunctionEndpointType {
	var returns WebFunctionEndpointType
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebFunctionEndpointLocation) ReplicaRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaRegions",
		&returns,
	)
	return returns
}


// Serve from the stack's Region only, under one regional domain name.
func WebFunctionEndpointLocation_HomeRegion() WebFunctionEndpointLocation {
	_init_.Initialize()

	var returns WebFunctionEndpointLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionEndpointLocation",
		"homeRegion",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Serve from the stack's Region and the replica Regions under one global domain name, with requests routed to the nearest Region.
//
// The stack's own Region always serves traffic, so listing it has no effect, and
// `replicaRegions` must hold at least one other Region. The list is treated as a set: order
// and duplicates do not matter. Adding or removing a Region replaces the endpoint; a
// deploy-time list (for example from a `CfnParameter`) cannot feed the default endpoint
// name, so changing its value at deployment also needs a new `endpointName`.
func WebFunctionEndpointLocation_MultiRegion(replicaRegions *[]*string) WebFunctionEndpointLocation {
	_init_.Initialize()

	if err := validateWebFunctionEndpointLocation_MultiRegionParameters(replicaRegions); err != nil {
		panic(err)
	}
	var returns WebFunctionEndpointLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionEndpointLocation",
		"multiRegion",
		[]interface{}{replicaRegions},
		&returns,
	)

	return returns
}

// Serve from the stack's Region and the replica Regions, with a separate domain name in each Region and no managed routing between them.
//
// The same rules for `replicaRegions` apply as for `multiRegion()`.
func WebFunctionEndpointLocation_PerRegion(replicaRegions *[]*string) WebFunctionEndpointLocation {
	_init_.Initialize()

	if err := validateWebFunctionEndpointLocation_PerRegionParameters(replicaRegions); err != nil {
		panic(err)
	}
	var returns WebFunctionEndpointLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionEndpointLocation",
		"perRegion",
		[]interface{}{replicaRegions},
		&returns,
	)

	return returns
}

