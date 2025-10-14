package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FargateProfile.
// Experimental.
type IFargateProfileRef interface {
	constructs.IConstruct
	// A reference to a FargateProfile resource.
	// Experimental.
	FargateProfileRef() *FargateProfileReference
}

// The jsii proxy for IFargateProfileRef
type jsiiProxy_IFargateProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFargateProfileRef) FargateProfileRef() *FargateProfileReference {
	var returns *FargateProfileReference
	_jsii_.Get(
		j,
		"fargateProfileRef",
		&returns,
	)
	return returns
}

