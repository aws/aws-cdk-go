package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A class used to generate resource arns for AppSync.
//
// Example:
//   var api IGraphqlApi
//   role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//
//   api.Grant(role, appsync.IamResource_Custom(jsii.String("types/Mutation/fields/updateExample")), jsii.String("appsync:GraphQL"))
//
type IamResource interface {
	// Return the Resource ARN.
	ResourceArns(api GraphqlApiBase) *[]*string
}

// The jsii proxy struct for IamResource
type jsiiProxy_IamResource struct {
	_ byte // padding
}

// Generate the resource names that accepts all types: `*`.
func IamResource_All() IamResource {
	_init_.Initialize()

	var returns IamResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.IamResource",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate the resource names given custom arns.
func IamResource_Custom(arns ...*string) IamResource {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range arns {
		args = append(args, a)
	}

	var returns IamResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.IamResource",
		"custom",
		args,
		&returns,
	)

	return returns
}

// Generate the resource names given a type and fields.
func IamResource_OfType(type_ *string, fields ...*string) IamResource {
	_init_.Initialize()

	if err := validateIamResource_OfTypeParameters(type_); err != nil {
		panic(err)
	}
	args := []interface{}{type_}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns IamResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.IamResource",
		"ofType",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamResource) ResourceArns(api GraphqlApiBase) *[]*string {
	if err := i.validateResourceArnsParameters(api); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"resourceArns",
		[]interface{}{api},
		&returns,
	)

	return returns
}

