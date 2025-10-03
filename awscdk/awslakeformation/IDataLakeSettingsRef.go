package awslakeformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataLakeSettings.
// Experimental.
type IDataLakeSettingsRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataLakeSettingsRef
type jsiiProxy_IDataLakeSettingsRef struct {
	internal.Type__constructsIConstruct
}

