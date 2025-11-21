package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb"
)

// Construction properties for TableGrants.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryptedResource IEncryptedResource
//   var resourceWithPolicyV2 IResourceWithPolicyV2
//   var tableRef ITableRef
//
//   tableGrantsProps := &TableGrantsProps{
//   	Table: tableRef,
//
//   	// the properties below are optional
//   	EncryptedResource: encryptedResource,
//   	HasIndex: jsii.Boolean(false),
//   	PolicyResource: resourceWithPolicyV2,
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
type TableGrantsProps struct {
	// The table to grant permissions on.
	Table interfacesawsdynamodb.ITableRef `field:"required" json:"table" yaml:"table"`
	// The encrypted resource on which actions will be allowed.
	// Default: - No permission is added to the KMS key, even if it exists.
	//
	EncryptedResource awsiam.IEncryptedResource `field:"optional" json:"encryptedResource" yaml:"encryptedResource"`
	// Whether this table has indexes.
	//
	// If so, permissions are granted on all table indexes as well.
	// Default: false.
	//
	HasIndex *bool `field:"optional" json:"hasIndex" yaml:"hasIndex"`
	// The resource with policy on which actions will be allowed.
	// Default: - No resource policy is created.
	//
	PolicyResource awsiam.IResourceWithPolicyV2 `field:"optional" json:"policyResource" yaml:"policyResource"`
	// Additional regions other than the main one that this table is replicated to.
	// Default: - No regions.
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

