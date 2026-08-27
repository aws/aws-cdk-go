package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The resource type to index.
//
// This is implemented as an enum-like class so that resource types the AWS Backup
// service adds in the future can be used before they are added to the CDK, e.g.
// `new IndexActionResourceType('EFS')`.
//
// Example:
//   var plan BackupPlan
//
//   plan.AddRule(backup.NewBackupPlanRule(&BackupPlanRuleProps{
//   	IndexActions: []BackupPlanIndexActionProps{
//   		&BackupPlanIndexActionProps{
//   			ResourceTypes: []IndexActionResourceType{
//   				backup.IndexActionResourceType_S3(),
//   			},
//   		},
//   	},
//   }))
//
// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/API_IndexAction.html
//
type IndexActionResourceType interface {
	// the resource type string value, e.g. `S3` or `EBS`.
	Value() *string
	// Returns the string representation of this resource type.
	ToString() *string
}

// The jsii proxy struct for IndexActionResourceType
type jsiiProxy_IndexActionResourceType struct {
	_ byte // padding
}

func (j *jsiiProxy_IndexActionResourceType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A custom resource type not yet supported as a static member of this class.
func NewIndexActionResourceType(value *string) IndexActionResourceType {
	_init_.Initialize()

	if err := validateNewIndexActionResourceTypeParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_IndexActionResourceType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_backup.IndexActionResourceType",
		[]interface{}{value},
		&j,
	)

	return &j
}

// A custom resource type not yet supported as a static member of this class.
func NewIndexActionResourceType_Override(i IndexActionResourceType, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_backup.IndexActionResourceType",
		[]interface{}{value},
		i,
	)
}

func IndexActionResourceType_EBS() IndexActionResourceType {
	_init_.Initialize()
	var returns IndexActionResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_backup.IndexActionResourceType",
		"EBS",
		&returns,
	)
	return returns
}

func IndexActionResourceType_S3() IndexActionResourceType {
	_init_.Initialize()
	var returns IndexActionResourceType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_backup.IndexActionResourceType",
		"S3",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IndexActionResourceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

