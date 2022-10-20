package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a CrossAccountZoneDelegationRecord.
//
// Example:
//   subZone := route53.NewPublicHostedZone(this, jsii.String("SubZone"), &publicHostedZoneProps{
//   	zoneName: jsii.String("sub.someexample.com"),
//   })
//
//   // import the delegation role by constructing the roleArn
//   delegationRoleArn := awscdk.stack.of(this).formatArn(&arnComponents{
//   	region: jsii.String(""),
//   	 // IAM is global in each partition
//   	service: jsii.String("iam"),
//   	account: jsii.String("parent-account-id"),
//   	resource: jsii.String("role"),
//   	resourceName: jsii.String("MyDelegationRole"),
//   })
//   delegationRole := iam.role.fromRoleArn(this, jsii.String("DelegationRole"), delegationRoleArn)
//
//   // create the record
//   // create the record
//   route53.NewCrossAccountZoneDelegationRecord(this, jsii.String("delegate"), &crossAccountZoneDelegationRecordProps{
//   	delegatedZone: subZone,
//   	parentHostedZoneName: jsii.String("someexample.com"),
//   	 // or you can use parentHostedZoneId
//   	delegationRole: delegationRole,
//   })
//
type CrossAccountZoneDelegationRecordProps struct {
	// The zone to be delegated.
	DelegatedZone IHostedZone `field:"required" json:"delegatedZone" yaml:"delegatedZone"`
	// The delegation role in the parent account.
	DelegationRole awsiam.IRole `field:"required" json:"delegationRole" yaml:"delegationRole"`
	// The hosted zone id in the parent account.
	ParentHostedZoneId *string `field:"optional" json:"parentHostedZoneId" yaml:"parentHostedZoneId"`
	// The hosted zone name in the parent account.
	ParentHostedZoneName *string `field:"optional" json:"parentHostedZoneName" yaml:"parentHostedZoneName"`
	// The removal policy to apply to the record set.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
}

