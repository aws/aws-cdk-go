# AWS ServiceCatalogAppRegistry Construct Library

[AWS Service Catalog App Registry](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/appregistry.html)
enables organizations to create and manage repositores of applications and associated resources.

## Table Of Contents

* [Application](#application)
* [Attribute-Group](#attribute-group)
* [Associations](#associations)

  * [Associating application with an attribute group](#attribute-group-association)
  * [Associating application with a stack](#resource-association)

The `@aws-cdk/aws-servicecatalogappregistry` package contains resources that enable users to automate governance and management of their AWS resources at scale.

```go
import appreg "github.com/aws/aws-cdk-go/awscdk"
```

## Application

An AppRegistry application enables you to define your applications and associated resources.
The application name must be unique at the account level, but is mutable.

```go
application := appreg.NewApplication(this, jsii.String("MyFirstApplication"), &ApplicationProps{
	ApplicationName: jsii.String("MyFirstApplicationName"),
	Description: jsii.String("description for my application"),
})
```

An application that has been created outside of the stack can be imported into your CDK app.
Applications can be imported by their ARN via the `Application.fromApplicationArn()` API:

```go
importedApplication := appreg.Application_FromApplicationArn(this, jsii.String("MyImportedApplication"), jsii.String("arn:aws:servicecatalog:us-east-1:012345678910:/applications/0aqmvxvgmry0ecc4mjhwypun6i"))
```

## Attribute Group

An AppRegistry attribute group acts as a container for user-defined attributes for an application.
Metadata is attached in a machine-readble format to integrate with automated workflows and tools.

```go
attributeGroup := appreg.NewAttributeGroup(this, jsii.String("MyFirstAttributeGroup"), &AttributeGroupProps{
	AttributeGroupName: jsii.String("MyFirstAttributeGroupName"),
	Description: jsii.String("description for my attribute group"),
	 // the description is optional,
	Attributes: map[string]interface{}{
		"project": jsii.String("foo"),
		"team": []interface{}{
			jsii.String("member1"),
			jsii.String("member2"),
			jsii.String("member3"),
		},
		"public": jsii.Boolean(false),
		"stages": map[string]*string{
			"alpha": jsii.String("complete"),
			"beta": jsii.String("incomplete"),
			"release": jsii.String("not started"),
		},
	},
})
```

An attribute group that has been created outside of the stack can be imported into your CDK app.
Attribute groups can be imported by their ARN via the `AttributeGroup.fromAttributeGroupArn()` API:

```go
importedAttributeGroup := appreg.AttributeGroup_FromAttributeGroupArn(this, jsii.String("MyImportedAttrGroup"), jsii.String("arn:aws:servicecatalog:us-east-1:012345678910:/attribute-groups/0aqmvxvgmry0ecc4mjhwypun6i"))
```

## Associations

You can associate your appregistry application with attribute groups and resources.
Resources are CloudFormation stacks that you can associate with an application to group relevant
stacks together to enable metadata rich insights into your applications and resources.
A Cloudformation stack can only be associated with one appregistry application.
If a stack is associated with multiple applications in your app or is already associated with one,
CDK will fail at deploy time.

### Associating application with an attribute group

You can associate an attribute group with an application with the `associateAttributeGroup()` API:

```go
var application application
var attributeGroup attributeGroup

application.associateAttributeGroup(attributeGroup)
```

### Associating application with a Stack

You can associate a stack with an application with the `associateStack()` API:

```go
var application application
app := awscdk.NewApp()
myStack := awscdk.Newstack(app, jsii.String("MyStack"))
application.associateStack(myStack)
```
