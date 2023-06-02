package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a PublicHostedZone.
//
// Example:
//   subZone := route53.NewPublicHostedZone(this, jsii.String("SubZone"), &PublicHostedZoneProps{
//   	ZoneName: jsii.String("sub.someexample.com"),
//   })
//
//   // import the delegation role by constructing the roleArn
//   delegationRoleArn := awscdk.stack_Of(this).FormatArn(&ArnComponents{
//   	Region: jsii.String(""),
//   	 // IAM is global in each partition
//   	Service: jsii.String("iam"),
//   	Account: jsii.String("parent-account-id"),
//   	Resource: jsii.String("role"),
//   	ResourceName: jsii.String("MyDelegationRole"),
//   })
//   delegationRole := iam.Role_FromRoleArn(this, jsii.String("DelegationRole"), delegationRoleArn)
//
//   // create the record
//   // create the record
//   route53.NewCrossAccountZoneDelegationRecord(this, jsii.String("delegate"), &CrossAccountZoneDelegationRecordProps{
//   	DelegatedZone: subZone,
//   	ParentHostedZoneName: jsii.String("someexample.com"),
//   	 // or you can use parentHostedZoneId
//   	DelegationRole: DelegationRole,
//   })
//
type PublicHostedZoneProps struct {
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
	// Whether to add a trailing dot to the zone name.
	AddTrailingDot *bool `field:"optional" json:"addTrailingDot" yaml:"addTrailingDot"`
	// Any comments that you want to include about the hosted zone.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	QueryLogsLogGroupArn *string `field:"optional" json:"queryLogsLogGroupArn" yaml:"queryLogsLogGroupArn"`
	// Whether to create a CAA record to restrict certificate authorities allowed to issue certificates for this domain to Amazon only.
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
	// Deprecated: Create the Role yourself and call `hostedZone.grantDelegation()`.
	CrossAccountZoneDelegationPrincipal awsiam.IPrincipal `field:"optional" json:"crossAccountZoneDelegationPrincipal" yaml:"crossAccountZoneDelegationPrincipal"`
	// The name of the role created for cross account delegation.
	// Deprecated: Create the Role yourself and call `hostedZone.grantDelegation()`.
	CrossAccountZoneDelegationRoleName *string `field:"optional" json:"crossAccountZoneDelegationRoleName" yaml:"crossAccountZoneDelegationRoleName"`
}

