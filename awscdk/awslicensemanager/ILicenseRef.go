package awslicensemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a License.
// Experimental.
type ILicenseRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILicenseRef
type jsiiProxy_ILicenseRef struct {
	internal.Type__constructsIConstruct
}

