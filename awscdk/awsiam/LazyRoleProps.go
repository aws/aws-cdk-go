package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a LazyRole.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var managedPolicy managedPolicy
//   var policyDocument policyDocument
//   var principal iPrincipal
//
//   lazyRoleProps := &LazyRoleProps{
//   	AssumedBy: principal,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExternalIds: []*string{
//   		jsii.String("externalIds"),
//   	},
//   	InlinePolicies: map[string]*policyDocument{
//   		"inlinePoliciesKey": policyDocument,
//   	},
//   	ManagedPolicies: []iManagedPolicy{
//   		managedPolicy,
//   	},
//   	MaxSessionDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   	Path: jsii.String("path"),
//   	PermissionsBoundary: managedPolicy,
//   	RoleName: jsii.String("roleName"),
//   }
//
type LazyRoleProps struct {
	// The IAM principal (i.e. `new ServicePrincipal('sns.amazonaws.com')`) which can assume this role.
	//
	// You can later modify the assume role policy document by accessing it via
	// the `assumeRolePolicy` property.
	AssumedBy IPrincipal `field:"required" json:"assumedBy" yaml:"assumedBy"`
	// A description of the role.
	//
	// It can be up to 1000 characters long.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of IDs that the role assumer needs to provide one of when assuming this role.
	//
	// If the configured and provided external IDs do not match, the
	// AssumeRole operation will fail.
	// Default: No external ID required.
	//
	ExternalIds *[]*string `field:"optional" json:"externalIds" yaml:"externalIds"`
	// A list of named policies to inline into this role.
	//
	// These policies will be
	// created with the role, whereas those added by ``addToPolicy`` are added
	// using a separate CloudFormation resource (allowing a way around circular
	// dependencies that could otherwise be introduced).
	// Default: - No policy is inlined in the Role resource.
	//
	InlinePolicies *map[string]PolicyDocument `field:"optional" json:"inlinePolicies" yaml:"inlinePolicies"`
	// A list of managed policies associated with this role.
	//
	// You can add managed policies later using
	// `addManagedPolicy(ManagedPolicy.fromAwsManagedPolicyName(policyName))`.
	// Default: - No managed policies.
	//
	ManagedPolicies *[]IManagedPolicy `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// The maximum session duration that you want to set for the specified role.
	//
	// This setting can have a value from 1 hour (3600sec) to 12 (43200sec) hours.
	//
	// Anyone who assumes the role from the AWS CLI or API can use the
	// DurationSeconds API parameter or the duration-seconds CLI parameter to
	// request a longer session. The MaxSessionDuration setting determines the
	// maximum duration that can be requested using the DurationSeconds
	// parameter.
	//
	// If users don't specify a value for the DurationSeconds parameter, their
	// security credentials are valid for one hour by default. This applies when
	// you use the AssumeRole* API operations or the assume-role* CLI operations
	// but does not apply when you use those operations to create a console URL.
	// Default: Duration.hours(1)
	//
	MaxSessionDuration awscdk.Duration `field:"optional" json:"maxSessionDuration" yaml:"maxSessionDuration"`
	// The path associated with this role.
	//
	// For information about IAM paths, see
	// Friendly Names and Paths in IAM User Guide.
	// Default: /.
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// AWS supports permissions boundaries for IAM entities (users or roles).
	//
	// A permissions boundary is an advanced feature for using a managed policy
	// to set the maximum permissions that an identity-based policy can grant to
	// an IAM entity. An entity's permissions boundary allows it to perform only
	// the actions that are allowed by both its identity-based policies and its
	// permissions boundaries.
	// Default: - No permissions boundary.
	//
	PermissionsBoundary IManagedPolicy `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// A name for the IAM role.
	//
	// For valid values, see the RoleName parameter for
	// the CreateRole action in the IAM API Reference.
	//
	// IMPORTANT: If you specify a name, you cannot perform updates that require
	// replacement of this resource. You can perform updates that require no or
	// some interruption. If you must replace the resource, specify a new name.
	//
	// If you specify a name, you must specify the CAPABILITY_NAMED_IAM value to
	// acknowledge your template's capabilities. For more information, see
	// Acknowledging IAM Resources in AWS CloudFormation Templates.
	// Default: - AWS CloudFormation generates a unique physical ID and uses that ID
	// for the role name.
	//
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

