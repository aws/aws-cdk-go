package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch"
)

// Collection of grant methods for a IConsumableResourceRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var consumableResourceRef IConsumableResourceRef
//
//   consumableResourceGrants := awscdk.Aws_batch.ConsumableResourceGrants_FromConsumableResource(consumableResourceRef)
//
type ConsumableResourceGrants interface {
	Resource() interfacesawsbatch.IConsumableResourceRef
	// Grant the given identity custom permissions.
	Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant
	// Grant permission to describe this consumable resource and list the jobs that use it.
	//
	// This will grant the following permissions:
	//
	//  - batch:DescribeConsumableResource
	// - batch:ListJobsByConsumableResource.
	Read(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permission to change the total quantity of this consumable resource.
	//
	// This will grant the following permissions:
	//
	// - batch:UpdateConsumableResource.
	Update(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for ConsumableResourceGrants
type jsiiProxy_ConsumableResourceGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_ConsumableResourceGrants) Resource() interfacesawsbatch.IConsumableResourceRef {
	var returns interfacesawsbatch.IConsumableResourceRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for ConsumableResourceGrants.
func ConsumableResourceGrants_FromConsumableResource(resource interfacesawsbatch.IConsumableResourceRef) ConsumableResourceGrants {
	_init_.Initialize()

	if err := validateConsumableResourceGrants_FromConsumableResourceParameters(resource); err != nil {
		panic(err)
	}
	var returns ConsumableResourceGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.ConsumableResourceGrants",
		"fromConsumableResource",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumableResourceGrants) Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant {
	if err := c.validateActionsParameters(grantee, actions, options); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"actions",
		[]interface{}{grantee, actions, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumableResourceGrants) Read(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"read",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsumableResourceGrants) Update(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateUpdateParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"update",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

