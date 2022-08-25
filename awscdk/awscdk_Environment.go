// Version 2 of the AWS Cloud Development Kit library
package awscdk


// The deployment environment for a stack.
//
// Example:
//   // Passing a replication bucket created in a different stack.
//   app := awscdk.NewApp()
//   replicationStack := awscdk.Newstack(app, jsii.String("ReplicationStack"), &stackProps{
//   	env: &environment{
//   		region: jsii.String("us-west-1"),
//   	},
//   })
//   key := kms.NewKey(replicationStack, jsii.String("ReplicationKey"))
//   replicationBucket := s3.NewBucket(replicationStack, jsii.String("ReplicationBucket"), &bucketProps{
//   	// like was said above - replication buckets need a set physical name
//   	bucketName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
//   	encryptionKey: key,
//   })
//
//   // later...
//   // later...
//   codepipeline.NewPipeline(replicationStack, jsii.String("Pipeline"), &pipelineProps{
//   	crossRegionReplicationBuckets: map[string]iBucket{
//   		"us-west-1": replicationBucket,
//   	},
//   })
//
type Environment struct {
	// The AWS account ID for this environment.
	//
	// This can be either a concrete value such as `585191031104` or `Aws.ACCOUNT_ID` which
	// indicates that account ID will only be determined during deployment (it
	// will resolve to the CloudFormation intrinsic `{"Ref":"AWS::AccountId"}`).
	// Note that certain features, such as cross-stack references and
	// environmental context providers require concerete region information and
	// will cause this stack to emit synthesis errors.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The AWS region for this environment.
	//
	// This can be either a concrete value such as `eu-west-2` or `Aws.REGION`
	// which indicates that account ID will only be determined during deployment
	// (it will resolve to the CloudFormation intrinsic `{"Ref":"AWS::Region"}`).
	// Note that certain features, such as cross-stack references and
	// environmental context providers require concrete region information and
	// will cause this stack to emit synthesis errors.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

