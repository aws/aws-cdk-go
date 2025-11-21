package awsopensearchservice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopensearchservice"
)

// Collection of grant methods for a IDomainRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domainRef IDomainRef
//
//   domainGrants := awscdk.Aws_opensearchservice.DomainGrants_FromDomain(domainRef)
//
type DomainGrants interface {
	Resource() interfacesawsopensearchservice.IDomainRef
	// Grant read permissions for this domain and its contents to an IAM principal (Role/Group/User).
	Read(grantee awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for this domain and its contents to an IAM principal (Role/Group/User).
	ReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for this domain and its contents to an IAM principal (Role/Group/User).
	Write(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for DomainGrants
type jsiiProxy_DomainGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_DomainGrants) Resource() interfacesawsopensearchservice.IDomainRef {
	var returns interfacesawsopensearchservice.IDomainRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for DomainGrants.
func DomainGrants_FromDomain(resource interfacesawsopensearchservice.IDomainRef) DomainGrants {
	_init_.Initialize()

	if err := validateDomainGrants_FromDomainParameters(resource); err != nil {
		panic(err)
	}
	var returns DomainGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_opensearchservice.DomainGrants",
		"fromDomain",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainGrants) Read(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"read",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainGrants) ReadWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateReadWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"readWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainGrants) Write(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"write",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

