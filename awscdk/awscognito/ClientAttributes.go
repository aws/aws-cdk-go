package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A set of attributes, useful to set Read and Write attributes.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   clientWriteAttributes := (cognito.NewClientAttributes()).WithStandardAttributes(&StandardAttributesMask{
//   	Fullname: jsii.Boolean(true),
//   	Email: jsii.Boolean(true),
//   }).WithCustomAttributes(jsii.String("favouritePizza"), jsii.String("favouriteBeverage"))
//
//   clientReadAttributes := clientWriteAttributes.WithStandardAttributes(&StandardAttributesMask{
//   	EmailVerified: jsii.Boolean(true),
//   }).WithCustomAttributes(jsii.String("pointsEarned"))
//
//   pool.addClient(jsii.String("app-client"), &UserPoolClientOptions{
//   	// ...
//   	ReadAttributes: clientReadAttributes,
//   	WriteAttributes: clientWriteAttributes,
//   })
//
type ClientAttributes interface {
	// The list of attributes represented by this ClientAttributes.
	Attributes() *[]*string
	// Creates a custom ClientAttributes with the specified attributes.
	WithCustomAttributes(attributes ...*string) ClientAttributes
	// Creates a custom ClientAttributes with the specified attributes.
	WithStandardAttributes(attributes *StandardAttributesMask) ClientAttributes
}

// The jsii proxy struct for ClientAttributes
type jsiiProxy_ClientAttributes struct {
	_ byte // padding
}

// Creates a ClientAttributes with the specified attributes.
// Default: - a ClientAttributes object without any attributes.
//
func NewClientAttributes() ClientAttributes {
	_init_.Initialize()

	j := jsiiProxy_ClientAttributes{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.ClientAttributes",
		nil, // no parameters
		&j,
	)

	return &j
}

// Creates a ClientAttributes with the specified attributes.
// Default: - a ClientAttributes object without any attributes.
//
func NewClientAttributes_Override(c ClientAttributes) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.ClientAttributes",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_ClientAttributes) Attributes() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"attributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClientAttributes) WithCustomAttributes(attributes ...*string) ClientAttributes {
	args := []interface{}{}
	for _, a := range attributes {
		args = append(args, a)
	}

	var returns ClientAttributes

	_jsii_.Invoke(
		c,
		"withCustomAttributes",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClientAttributes) WithStandardAttributes(attributes *StandardAttributesMask) ClientAttributes {
	if err := c.validateWithStandardAttributesParameters(attributes); err != nil {
		panic(err)
	}
	var returns ClientAttributes

	_jsii_.Invoke(
		c,
		"withStandardAttributes",
		[]interface{}{attributes},
		&returns,
	)

	return returns
}

