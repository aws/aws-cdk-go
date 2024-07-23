package regioninfo

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Provides default values for certain regional information points.
//
// This class is no longer needed because service principals are no longer needed except in very specific cases
// that are handled in the IAM ServicePrincipal class.
// Deprecated: - Service principals are now globally `<SERVICE>.amazonaws.com`, use iam.ServicePrincipal instead.
type Default interface {
}

// The jsii proxy struct for Default
type jsiiProxy_Default struct {
	_ byte // padding
}

// Computes a "standard" AWS Service principal for a given service, region and suffix.
//
// This is useful for example when
// you need to compute a service principal name, but you do not have a synthesize-time region literal available (so
// all you have is `{ "Ref": "AWS::Region" }`). This way you get the same defaulting behavior that is normally used
// for built-in data.
// Deprecated: - Service principals are now globally `<SERVICE>.amazonaws.com`, use iam.ServicePrincipal instead.
func Default_ServicePrincipal(serviceFqn *string, region *string, urlSuffix *string) *string {
	_init_.Initialize()

	if err := validateDefault_ServicePrincipalParameters(serviceFqn, region, urlSuffix); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.region_info.Default",
		"servicePrincipal",
		[]interface{}{serviceFqn, region, urlSuffix},
		&returns,
	)

	return returns
}

func Default_VPC_ENDPOINT_SERVICE_NAME_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.region_info.Default",
		"VPC_ENDPOINT_SERVICE_NAME_PREFIX",
		&returns,
	)
	return returns
}

