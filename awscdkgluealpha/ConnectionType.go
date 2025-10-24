package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of the glue connection.
//
// If you need to use a connection type that doesn't exist as a static member, you
// can instantiate a `ConnectionType` object, e.g: `new ConnectionType('NEW_TYPE')`.
//
// Example:
//   var securityGroup SecurityGroup
//   var subnet Subnet
//
//   glue.NewConnection(this, jsii.String("MyConnection"), &ConnectionProps{
//   	Type: glue.ConnectionType_NETWORK(),
//   	// The security groups granting AWS Glue inbound access to the data source within the VPC
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	// The VPC subnet which contains the data source
//   	Subnet: Subnet,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype
//
// Experimental.
type ConnectionType interface {
	// The name of this ConnectionType, as expected by Connection resource.
	// Experimental.
	Name() *string
	// The connection type name as expected by Connection resource.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectionType
type jsiiProxy_ConnectionType struct {
	_ byte // padding
}

func (j *jsiiProxy_ConnectionType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewConnectionType(name *string) ConnectionType {
	_init_.Initialize()

	if err := validateNewConnectionTypeParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectionType{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		[]interface{}{name},
		&j,
	)

	return &j
}

// Experimental.
func NewConnectionType_Override(c ConnectionType, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		[]interface{}{name},
		c,
	)
}

func ConnectionType_CUSTOM() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"CUSTOM",
		&returns,
	)
	return returns
}

func ConnectionType_FACEBOOKADS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"FACEBOOKADS",
		&returns,
	)
	return returns
}

func ConnectionType_GOOGLEADS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"GOOGLEADS",
		&returns,
	)
	return returns
}

func ConnectionType_GOOGLEANALYTICS4() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"GOOGLEANALYTICS4",
		&returns,
	)
	return returns
}

func ConnectionType_GOOGLESHEETS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"GOOGLESHEETS",
		&returns,
	)
	return returns
}

func ConnectionType_HUBSPOT() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"HUBSPOT",
		&returns,
	)
	return returns
}

func ConnectionType_INSTAGRAMADS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"INSTAGRAMADS",
		&returns,
	)
	return returns
}

func ConnectionType_INTERCOM() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"INTERCOM",
		&returns,
	)
	return returns
}

func ConnectionType_JDBC() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"JDBC",
		&returns,
	)
	return returns
}

func ConnectionType_JIRACLOUD() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"JIRACLOUD",
		&returns,
	)
	return returns
}

func ConnectionType_KAFKA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"KAFKA",
		&returns,
	)
	return returns
}

func ConnectionType_MARKETO() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"MARKETO",
		&returns,
	)
	return returns
}

func ConnectionType_MARKETPLACE() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"MARKETPLACE",
		&returns,
	)
	return returns
}

func ConnectionType_MONGODB() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"MONGODB",
		&returns,
	)
	return returns
}

func ConnectionType_NETSUITEERP() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"NETSUITEERP",
		&returns,
	)
	return returns
}

func ConnectionType_NETWORK() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"NETWORK",
		&returns,
	)
	return returns
}

func ConnectionType_SALESFORCE() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"SALESFORCE",
		&returns,
	)
	return returns
}

func ConnectionType_SALESFORCEMARKETINGCLOUD() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"SALESFORCEMARKETINGCLOUD",
		&returns,
	)
	return returns
}

func ConnectionType_SALESFORCEPARDOT() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"SALESFORCEPARDOT",
		&returns,
	)
	return returns
}

func ConnectionType_SAPODATA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"SAPODATA",
		&returns,
	)
	return returns
}

func ConnectionType_SERVICENOW() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"SERVICENOW",
		&returns,
	)
	return returns
}

func ConnectionType_SLACK() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"SLACK",
		&returns,
	)
	return returns
}

func ConnectionType_SNAPCHATADS() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"SNAPCHATADS",
		&returns,
	)
	return returns
}

func ConnectionType_STRIPE() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"STRIPE",
		&returns,
	)
	return returns
}

func ConnectionType_VIEW_VALIDATION_ATHENA() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"VIEW_VALIDATION_ATHENA",
		&returns,
	)
	return returns
}

func ConnectionType_VIEW_VALIDATION_REDSHIFT() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"VIEW_VALIDATION_REDSHIFT",
		&returns,
	)
	return returns
}

func ConnectionType_ZENDESK() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"ZENDESK",
		&returns,
	)
	return returns
}

func ConnectionType_ZOHOCRM() ConnectionType {
	_init_.Initialize()
	var returns ConnectionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.ConnectionType",
		"ZOHOCRM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConnectionType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

