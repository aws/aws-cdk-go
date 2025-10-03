package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Asset.
// Experimental.
type IAssetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssetRef
type jsiiProxy_IAssetRef struct {
	internal.Type__constructsIConstruct
}

