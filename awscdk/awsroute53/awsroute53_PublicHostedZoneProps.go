package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties for a PublicHostedZone.
//
// Example:
//   parentZone := route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
//   	zoneName: jsii.String("someexample.com"),
//   	crossAccountZoneDelegationPrincipal: iam.NewAccountPrincipal(jsii.String("12345678901")),
//   	crossAccountZoneDelegationRoleName: jsii.String("MyDelegationRole"),
//   })
//
// Experimental.
type PublicHostedZoneProps struct {
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	// Experimental.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
	// Any comments that you want to include about the hosted zone.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	// Experimental.
	QueryLogsLogGroupArn *string `field:"optional" json:"queryLogsLogGroupArn" yaml:"queryLogsLogGroupArn"`
	// Whether to create a CAA record to restrict certificate authorities allowed to issue certificates for this domain to Amazon only.
	// Experimental.
	CaaAmazon *bool `field:"optional" json:"caaAmazon" yaml:"caaAmazon"`
	// A principal which is trusted to assume a role for zone delegation.
	// Experimental.
	CrossAccountZoneDelegationPrincipal awsiam.IPrincipal `field:"optional" json:"crossAccountZoneDelegationPrincipal" yaml:"crossAccountZoneDelegationPrincipal"`
	// The name of the role created for cross account delegation.
	// Experimental.
	CrossAccountZoneDelegationRoleName *string `field:"optional" json:"crossAccountZoneDelegationRoleName" yaml:"crossAccountZoneDelegationRoleName"`
}

