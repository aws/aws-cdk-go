package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Physical ID of the custom resource.
//
// Example:
//   crossAccountRoleArn := "arn:aws:iam::OTHERACCOUNT:role/CrossAccountRoleName" // arn of role deployed in separate account
//
//   callRegion := "us-west-1" // sdk call to be made in specified region (optional)
//
//    // sdk call to be made in specified region (optional)
//   cr.NewAwsCustomResource(this, jsii.String("CrossAccount"), &AwsCustomResourceProps{
//   	OnCreate: &AwsSdkCall{
//   		AssumedRoleArn: crossAccountRoleArn,
//   		Region: callRegion,
//   		 // optional
//   		Service: jsii.String("sts"),
//   		Action: jsii.String("GetCallerIdentity"),
//   		PhysicalResourceId: cr.PhysicalResourceId_Of(jsii.String("id")),
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromStatements([]PolicyStatement{
//   		iam.PolicyStatement_FromJson(map[string]*string{
//   			"Effect": jsii.String("Allow"),
//   			"Action": jsii.String("sts:AssumeRole"),
//   			"Resource": crossAccountRoleArn,
//   		}),
//   	}),
//   })
//
type PhysicalResourceId interface {
	// Literal string to be used as the physical id.
	Id() *string
	// Path to a response data element to be used as the physical id.
	ResponsePath() *string
}

// The jsii proxy struct for PhysicalResourceId
type jsiiProxy_PhysicalResourceId struct {
	_ byte // padding
}

func (j *jsiiProxy_PhysicalResourceId) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalResourceId) ResponsePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responsePath",
		&returns,
	)
	return returns
}


// Extract the physical resource id from the path (dot notation) to the data in the API call response.
func PhysicalResourceId_FromResponse(responsePath *string) PhysicalResourceId {
	_init_.Initialize()

	if err := validatePhysicalResourceId_FromResponseParameters(responsePath); err != nil {
		panic(err)
	}
	var returns PhysicalResourceId

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.PhysicalResourceId",
		"fromResponse",
		[]interface{}{responsePath},
		&returns,
	)

	return returns
}

// Explicit physical resource id.
func PhysicalResourceId_Of(id *string) PhysicalResourceId {
	_init_.Initialize()

	if err := validatePhysicalResourceId_OfParameters(id); err != nil {
		panic(err)
	}
	var returns PhysicalResourceId

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.PhysicalResourceId",
		"of",
		[]interface{}{id},
		&returns,
	)

	return returns
}

