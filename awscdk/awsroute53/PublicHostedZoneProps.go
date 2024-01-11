package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a PublicHostedZone.
//
// Example:
//   parentZone := route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
//   	ZoneName: jsii.String("someexample.com"),
//   })
//   crossAccountRole := iam.NewRole(this, jsii.String("CrossAccountRole"), &RoleProps{
//   	// The role name must be predictable
//   	RoleName: jsii.String("MyDelegationRole"),
//   	// The other account
//   	AssumedBy: iam.NewAccountPrincipal(jsii.String("12345678901")),
//   	// You can scope down this role policy to be least privileged.
//   	// If you want the other account to be able to manage specific records,
//   	// you can scope down by resource and/or normalized record names
//   	InlinePolicies: map[string]policyDocument{
//   		"crossAccountPolicy": iam.NewPolicyDocument(&PolicyDocumentProps{
//   			"statements": []PolicyStatement{
//   				iam.NewPolicyStatement(&PolicyStatementProps{
//   					"sid": jsii.String("ListHostedZonesByName"),
//   					"effect": iam.Effect_ALLOW,
//   					"actions": []*string{
//   						jsii.String("route53:ListHostedZonesByName"),
//   					},
//   					"resources": []*string{
//   						jsii.String("*"),
//   					},
//   				}),
//   				iam.NewPolicyStatement(&PolicyStatementProps{
//   					"sid": jsii.String("GetHostedZoneAndChangeResourceRecordSet"),
//   					"effect": iam.Effect_ALLOW,
//   					"actions": []*string{
//   						jsii.String("route53:GetHostedZone"),
//   						jsii.String("route53:ChangeResourceRecordSet"),
//   					},
//   					// This example assumes the RecordSet subdomain.somexample.com
//   					// is contained in the HostedZone
//   					"resources": []*string{
//   						jsii.String("arn:aws:route53:::hostedzone/HZID00000000000000000"),
//   					},
//   					"conditions": map[string]interface{}{
//   						"ForAllValues:StringLike": map[string][]*string{
//   							"route53:ChangeResourceRecordSetsNormalizedRecordNames": []*string{
//   								jsii.String("subdomain.someexample.com"),
//   							},
//   						},
//   					},
//   				}),
//   			},
//   		}),
//   	},
//   })
//   parentZone.GrantDelegation(crossAccountRole)
//
type PublicHostedZoneProps struct {
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
	// Whether to add a trailing dot to the zone name.
	// Default: true.
	//
	AddTrailingDot *bool `field:"optional" json:"addTrailingDot" yaml:"addTrailingDot"`
	// Any comments that you want to include about the hosted zone.
	// Default: none.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	// Default: disabled.
	//
	QueryLogsLogGroupArn *string `field:"optional" json:"queryLogsLogGroupArn" yaml:"queryLogsLogGroupArn"`
	// Whether to create a CAA record to restrict certificate authorities allowed to issue certificates for this domain to Amazon only.
	// Default: false.
	//
	CaaAmazon *bool `field:"optional" json:"caaAmazon" yaml:"caaAmazon"`
	// A principal which is trusted to assume a role for zone delegation.
	//
	// If supplied, this will create a Role in the same account as the Hosted
	// Zone, which can be assumed by the `CrossAccountZoneDelegationRecord` to
	// create a delegation record to a zone in a different account.
	//
	// Be sure to indicate the account(s) that you trust to create delegation
	// records, using either `iam.AccountPrincipal` or `iam.OrganizationPrincipal`.
	//
	// If you are planning to use `iam.ServicePrincipal`s here, be sure to include
	// region-specific service principals for every opt-in region you are going to
	// be delegating to; or don't use this feature and create separate roles
	// with appropriate permissions for every opt-in region instead.
	// Default: - No delegation configuration.
	//
	// Deprecated: Create the Role yourself and call `hostedZone.grantDelegation()`.
	CrossAccountZoneDelegationPrincipal awsiam.IPrincipal `field:"optional" json:"crossAccountZoneDelegationPrincipal" yaml:"crossAccountZoneDelegationPrincipal"`
	// The name of the role created for cross account delegation.
	// Default: - A role name is generated automatically.
	//
	// Deprecated: Create the Role yourself and call `hostedZone.grantDelegation()`.
	CrossAccountZoneDelegationRoleName *string `field:"optional" json:"crossAccountZoneDelegationRoleName" yaml:"crossAccountZoneDelegationRoleName"`
}

