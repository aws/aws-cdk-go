package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssetModel.
// Experimental.
type IAssetModelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssetModelRef
type jsiiProxy_IAssetModelRef struct {
	internal.Type__constructsIConstruct
}

