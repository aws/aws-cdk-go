package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Cross Account Zone Delegation record.
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
type CrossAccountZoneDelegationRecord interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CrossAccountZoneDelegationRecord
type jsiiProxy_CrossAccountZoneDelegationRecord struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CrossAccountZoneDelegationRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCrossAccountZoneDelegationRecord(scope constructs.Construct, id *string, props *CrossAccountZoneDelegationRecordProps) CrossAccountZoneDelegationRecord {
	_init_.Initialize()

	if err := validateNewCrossAccountZoneDelegationRecordParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CrossAccountZoneDelegationRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CrossAccountZoneDelegationRecord",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCrossAccountZoneDelegationRecord_Override(c CrossAccountZoneDelegationRecord, scope constructs.Construct, id *string, props *CrossAccountZoneDelegationRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CrossAccountZoneDelegationRecord",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CrossAccountZoneDelegationRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCrossAccountZoneDelegationRecord_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CrossAccountZoneDelegationRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountZoneDelegationRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

