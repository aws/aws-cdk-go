package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Physical ID of the custom resource.
//
// Example:
//   awsCustom := cr.NewAwsCustomResource(this, jsii.String("aws-custom"), &awsCustomResourceProps{
//   	onCreate: &awsSdkCall{
//   		service: jsii.String("..."),
//   		action: jsii.String("..."),
//   		parameters: map[string]*string{
//   			"text": jsii.String("..."),
//   		},
//   		physicalResourceId: cr.physicalResourceId.of(jsii.String("...")),
//   	},
//   	onUpdate: &awsSdkCall{
//   		service: jsii.String("..."),
//   		action: jsii.String("..."),
//   		parameters: map[string]interface{}{
//   			"text": jsii.String("..."),
//   			"resourceId": cr.NewPhysicalResourceIdReference(),
//   		},
//   	},
//   	policy: cr.awsCustomResourcePolicy.fromSdkCalls(&sdkCallsPolicyOptions{
//   		resources: cr.*awsCustomResourcePolicy_ANY_RESOURCE(),
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

	var returns PhysicalResourceId

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.PhysicalResourceId",
		"of",
		[]interface{}{id},
		&returns,
	)

	return returns
}

