package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Represents server-side encryption for a DynamoDB table.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-west-2"),
//   	},
//   })
//
//   tableKey := kms.NewKey(stack, jsii.String("Key"))
//   replicaKeyArns := map[string]*string{
//   	"us-east-1": jsii.String("arn:aws:kms:us-east-1:123456789012:key/g24efbna-az9b-42ro-m3bp-cq249l94fca6"),
//   	"us-east-2": jsii.String("arn:aws:kms:us-east-2:123456789012:key/h90bkasj-bs1j-92wp-s2ka-bh857d60bkj8"),
//   }
//
//   globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Encryption: dynamodb.TableEncryptionV2_CustomerManagedKey(tableKey, replicaKeyArns),
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   		},
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   		},
//   	},
//   })
//
type TableEncryptionV2 interface {
	ReplicaKeyArns() *map[string]*string
	TableKey() awskms.IKey
	Type() TableEncryption
}

// The jsii proxy struct for TableEncryptionV2
type jsiiProxy_TableEncryptionV2 struct {
	_ byte // padding
}

func (j *jsiiProxy_TableEncryptionV2) ReplicaKeyArns() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"replicaKeyArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableEncryptionV2) TableKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"tableKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableEncryptionV2) Type() TableEncryption {
	var returns TableEncryption
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Configure server-side encryption using a DynamoDB owned key.
func TableEncryptionV2_AwsManagedKey() TableEncryptionV2 {
	_init_.Initialize()

	var returns TableEncryptionV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.TableEncryptionV2",
		"awsManagedKey",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Configure server-side encryption using customer managed keys.
func TableEncryptionV2_CustomerManagedKey(tableKey awskms.IKey, replicaKeyArns *map[string]*string) TableEncryptionV2 {
	_init_.Initialize()

	if err := validateTableEncryptionV2_CustomerManagedKeyParameters(tableKey); err != nil {
		panic(err)
	}
	var returns TableEncryptionV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.TableEncryptionV2",
		"customerManagedKey",
		[]interface{}{tableKey, replicaKeyArns},
		&returns,
	)

	return returns
}

// Configure server-side encryption using a DynamoDB owned key.
func TableEncryptionV2_DynamoOwnedKey() TableEncryptionV2 {
	_init_.Initialize()

	var returns TableEncryptionV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.TableEncryptionV2",
		"dynamoOwnedKey",
		nil, // no parameters
		&returns,
	)

	return returns
}

