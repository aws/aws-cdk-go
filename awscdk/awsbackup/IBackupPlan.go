package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackup"
	"github.com/aws/constructs-go/constructs/v10"
)

// A backup plan.
type IBackupPlan interface {
	interfacesawsbackup.IBackupPlanRef
	awscdk.IResource
	// The identifier of the backup plan.
	BackupPlanId() *string
}

// The jsii proxy for IBackupPlan
type jsiiProxy_IBackupPlan struct {
	internal.Type__interfacesawsbackupIBackupPlanRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBackupPlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IBackupPlan) BackupPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupPlan) BackupPlanRef() *interfacesawsbackup.BackupPlanReference {
	var returns *interfacesawsbackup.BackupPlanReference
	_jsii_.Get(
		j,
		"backupPlanRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupPlan) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupPlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBackupPlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

