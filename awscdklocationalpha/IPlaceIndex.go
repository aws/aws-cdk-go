// The CDK Construct Library for AWS::Location
package awscdklocationalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdklocationalpha/v2/internal"
)

// A Place Index.
// Experimental.
type IPlaceIndex interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the place index resource.
	// Experimental.
	PlaceIndexArn() *string
	// The name of the place index.
	// Experimental.
	PlaceIndexName() *string
}

// The jsii proxy for IPlaceIndex
type jsiiProxy_IPlaceIndex struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPlaceIndex) PlaceIndexArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placeIndexArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlaceIndex) PlaceIndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placeIndexName",
		&returns,
	)
	return returns
}

