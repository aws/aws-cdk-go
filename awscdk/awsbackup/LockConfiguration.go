package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for AWS Backup Vault Lock.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lockConfiguration := &LockConfiguration{
//   	MinRetention: cdk.Duration_Minutes(jsii.Number(30)),
//
//   	// the properties below are optional
//   	ChangeableFor: cdk.Duration_*Minutes(jsii.Number(30)),
//   	MaxRetention: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html
//
type LockConfiguration struct {
	// The minimum retention period that the vault retains its recovery points.
	//
	// If this parameter is specified, any backup or copy job to the vault must
	// have a lifecycle policy with a retention period equal to or longer than
	// the minimum retention period. If the job's retention period is shorter than
	// that minimum retention period, then the vault fails that backup or copy job,
	// and you should either modify your lifecycle settings or use a different
	// vault. Recovery points already saved in the vault prior to Vault Lock are
	// not affected.
	MinRetention awscdk.Duration `field:"required" json:"minRetention" yaml:"minRetention"`
	// The duration before the lock date.
	//
	// AWS Backup enforces a 72-hour cooling-off period before Vault Lock takes
	// effect and becomes immutable.
	//
	// Before the lock date, you can delete Vault Lock from the vault or change
	// the Vault Lock configuration. On and after the lock date, the Vault Lock
	// becomes immutable and cannot be changed or deleted.
	ChangeableFor awscdk.Duration `field:"optional" json:"changeableFor" yaml:"changeableFor"`
	// The maximum retention period that the vault retains its recovery points.
	//
	// If this parameter is specified, any backup or copy job to the vault must
	// have a lifecycle policy with a retention period equal to or shorter than
	// the maximum retention period. If the job's retention period is longer than
	// that maximum retention period, then the vault fails the backup or copy job,
	// and you should either modify your lifecycle settings or use a different
	// vault. Recovery points already saved in the vault prior to Vault Lock are
	// not affected.
	MaxRetention awscdk.Duration `field:"optional" json:"maxRetention" yaml:"maxRetention"`
}

