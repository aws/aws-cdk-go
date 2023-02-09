package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStoredQuery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStoredQueryProps := &cfnStoredQueryProps{
//   	queryExpression: jsii.String("queryExpression"),
//   	queryName: jsii.String("queryName"),
//
//   	// the properties below are optional
//   	queryDescription: jsii.String("queryDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStoredQueryProps struct {
	// The expression of the query.
	//
	// For example, `SELECT resourceId, resourceType, supplementaryConfiguration.BucketVersioningConfiguration.status WHERE resourceType = 'AWS::S3::Bucket' AND supplementaryConfiguration.BucketVersioningConfiguration.status = 'Off'.`
	QueryExpression *string `field:"required" json:"queryExpression" yaml:"queryExpression"`
	// The name of the query.
	QueryName *string `field:"required" json:"queryName" yaml:"queryName"`
	// A unique description for the query.
	QueryDescription *string `field:"optional" json:"queryDescription" yaml:"queryDescription"`
	// An array of key-value pairs to apply to this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

