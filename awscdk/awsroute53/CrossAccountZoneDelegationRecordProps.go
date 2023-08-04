package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a CrossAccountZoneDelegationRecord.
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
type CrossAccountZoneDelegationRecordProps struct {
	// The zone to be delegated.
	DelegatedZone IHostedZone `field:"required" json:"delegatedZone" yaml:"delegatedZone"`
	// The delegation role in the parent account.
	DelegationRole awsiam.IRole `field:"required" json:"delegationRole" yaml:"delegationRole"`
	// The hosted zone id in the parent account.
	// Default: - no zone id.
	//
	ParentHostedZoneId *string `field:"optional" json:"parentHostedZoneId" yaml:"parentHostedZoneId"`
	// The hosted zone name in the parent account.
	// Default: - no zone name.
	//
	ParentHostedZoneName *string `field:"optional" json:"parentHostedZoneName" yaml:"parentHostedZoneName"`
	// The removal policy to apply to the record set.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The resource record cache time to live (TTL).
	// Default: Duration.days(2)
	//
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
}

