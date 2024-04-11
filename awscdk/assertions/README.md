# Assertions

If you're migrating from the old `@aws-cdk/assert` library, first use this migration guide to migrate from `@aws-cdk/assert` to `@aws-cdk/assertions` found in
[our GitHub repository](https://github.com/aws/aws-cdk/blob/v1-main/packages/@aws-cdk/assertions/MIGRATING.md). Then, you can migrate your application to AWS CDK v2 in order to use this library using [this guide](https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html).

Functions for writing test asserting against CDK applications, with focus on CloudFormation templates.

The `Template` class includes a set of methods for writing assertions against CloudFormation templates. Use one of the `Template.fromXxx()` static methods to create an instance of this class.

To create `Template` from CDK stack, start off with:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

stack := awscdk.NewStack()
// ...
template := awscdk.Template_FromStack(stack)
```

Alternatively, assertions can be run on an existing CloudFormation template -

```go
templateJson := "{ \"Resources\": ... }" /* The CloudFormation template as JSON serialized string. */
template := awscdk.Template_FromString(templateJson)
```

**Cyclical Resources Note**

If allowing cyclical references is desired, for example in the case of unprocessed Transform templates, supply TemplateParsingOptions and
set skipCyclicalDependenciesCheck to true. In all other cases, will fail on detecting cyclical dependencies.

## Full Template Match

The simplest assertion would be to assert that the template matches a given
template.

```go
template.TemplateMatches(map[string]map[string]map[string]interface{}{
	"Resources": map[string]map[string]interface{}{
		"BarLogicalId": map[string]interface{}{
			"Type": jsii.String("Foo::Bar"),
			"Properties": map[string]*string{
				"Baz": jsii.String("Qux"),
			},
		},
	},
})
```

By default, the `templateMatches()` API will use the an 'object-like' comparison,
which means that it will allow for the actual template to be a superset of the
given expectation. See [Special Matchers](#special-matchers) for details on how
to change this.

Snapshot testing is a common technique to store a snapshot of the output and
compare it during future changes. Since CloudFormation templates are human readable,
they are a good target for snapshot testing.

The `toJSON()` method on the `Template` can be used to produce a well formatted JSON
of the CloudFormation template that can be used as a snapshot.

See [Snapshot Testing in Jest](https://jestjs.io/docs/snapshot-testing) and [Snapshot
Testing in Java](https://json-snapshot.github.io/).

## Counting Resources

This module allows asserting the number of resources of a specific type found
in a template.

```go
template.ResourceCountIs(jsii.String("Foo::Bar"), jsii.Number(2))
```

You can also count the number of resources of a specific type whose `Properties`
section contains the specified properties:

```go
template.ResourcePropertiesCountIs(jsii.String("Foo::Bar"), map[string]interface{}{
	"Foo": jsii.String("Bar"),
	"Baz": jsii.Number(5),
	"Qux": []*string{
		jsii.String("Waldo"),
		jsii.String("Fred"),
	},
}, jsii.Number(1))
```

## Resource Matching & Retrieval

Beyond resource counting, the module also allows asserting that a resource with
specific properties are present.

The following code asserts that the `Properties` section of a resource of type
`Foo::Bar` contains the specified properties -

```go
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]interface{}{
	"Lorem": jsii.String("Ipsum"),
	"Baz": jsii.Number(5),
	"Qux": []*string{
		jsii.String("Waldo"),
		jsii.String("Fred"),
	},
})
```

You can also assert that the `Properties` section of all resources of type
`Foo::Bar` contains the specified properties -

```go
template.AllResourcesProperties(jsii.String("Foo::Bar"), map[string]interface{}{
	"Lorem": jsii.String("Ipsum"),
	"Baz": jsii.Number(5),
	"Qux": []*string{
		jsii.String("Waldo"),
		jsii.String("Fred"),
	},
})
```

Alternatively, if you would like to assert the entire resource definition, you
can use the `hasResource()` API.

```go
template.HasResource(jsii.String("Foo::Bar"), map[string]interface{}{
	"Properties": map[string]*string{
		"Lorem": jsii.String("Ipsum"),
	},
	"DependsOn": []*string{
		jsii.String("Waldo"),
		jsii.String("Fred"),
	},
})
```

You can also assert the definitions of all resources of a type using the
`allResources()` API.

```go
template.AllResources(jsii.String("Foo::Bar"), map[string]interface{}{
	"Properties": map[string]*string{
		"Lorem": jsii.String("Ipsum"),
	},
	"DependsOn": []*string{
		jsii.String("Waldo"),
		jsii.String("Fred"),
	},
})
```

Beyond assertions, the module provides APIs to retrieve matching resources.
The `findResources()` API is complementary to the `hasResource()` API, except,
instead of asserting its presence, it returns the set of matching resources.

By default, the `hasResource()` and `hasResourceProperties()` APIs perform deep
partial object matching. This behavior can be configured using matchers.
See subsequent section on [special matchers](#special-matchers).

## Output and Mapping sections

The module allows you to assert that the CloudFormation template contains an Output
that matches specific properties. The following code asserts that a template contains
an Output with a `logicalId` of `Foo` and the specified properties -

```go
expected := map[string]interface{}{
	"Value": jsii.String("Bar"),
	"Export": map[string]*string{
		"Name": jsii.String("ExportBaz"),
	},
}
template.HasOutput(jsii.String("Foo"), expected)
```

If you want to match against all Outputs in the template, use `*` as the `logicalId`.

```go
template.HasOutput(jsii.String("*"), map[string]interface{}{
	"Value": jsii.String("Bar"),
	"Export": map[string]*string{
		"Name": jsii.String("ExportBaz"),
	},
})
```

`findOutputs()` will return a set of outputs that match the `logicalId` and `props`,
and you can use the `'*'` special case as well.

```go
result := template.FindOutputs(jsii.String("*"), map[string]*string{
	"Value": jsii.String("Fred"),
})
expect(result.foo).toEqual(map[string]*string{
	"Value": jsii.String("Fred"),
	"Description": jsii.String("FooFred"),
})
expect(result.bar).toEqual(map[string]*string{
	"Value": jsii.String("Fred"),
	"Description": jsii.String("BarFred"),
})
```

The APIs `hasMapping()`, `findMappings()`, `hasCondition()`, and `hasCondtions()` provide similar functionalities.

## Special Matchers

The expectation provided to the `hasXxx()`, `findXxx()` and `templateMatches()`
APIs, besides carrying literal values, as seen in the above examples, also accept
special matchers.

They are available as part of the `Match` class.

### Object Matchers

The `Match.objectLike()` API can be used to assert that the target is a superset
object of the provided pattern.
This API will perform a deep partial match on the target.
Deep partial matching is where objects are matched partially recursively. At each
level, the list of keys in the target is a subset of the provided pattern.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": {
//           "Wobble": "Flob",
//           "Bob": "Cat"
//         }
//       }
//     }
//   }
// }

// The following will NOT throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Fred": awscdk.Match_objectLike(map[string]interface{}{
		"Wobble": jsii.String("Flob"),
	}),
})

// The following will throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Fred": awscdk.Match_objectLike(map[string]interface{}{
		"Brew": jsii.String("Coffee"),
	}),
})
```

The `Match.objectEquals()` API can be used to assert a target as a deep exact
match.

### Presence and Absence

The `Match.absent()` matcher can be used to specify that a specific
value should not exist on the target. This can be used within `Match.objectLike()`
or outside of any matchers.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": {
//           "Wobble": "Flob",
//         }
//       }
//     }
//   }
// }

// The following will NOT throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Fred": awscdk.Match_objectLike(map[string]interface{}{
		"Bob": awscdk.Match_absent(),
	}),
})

// The following will throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Fred": awscdk.Match_objectLike(map[string]interface{}{
		"Wobble": awscdk.Match_absent(),
	}),
})
```

The `Match.anyValue()` matcher can be used to specify that a specific value should be found
at the location. This matcher will fail if when the target location has null-ish values
(i.e., `null` or `undefined`).

This matcher can be combined with any of the other matchers.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": {
//           "Wobble": ["Flob", "Flib"],
//         }
//       }
//     }
//   }
// }

// The following will NOT throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]map[string][]matcher{
	"Fred": map[string][]matcher{
		"Wobble": []matcher{
			awscdk.Match_anyValue(),
			awscdk.Match_anyValue(),
		},
	},
})

// The following will throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]map[string]matcher{
	"Fred": map[string]matcher{
		"Wimble": awscdk.Match_anyValue(),
	},
})
```

### Array Matchers

The `Match.arrayWith()` API can be used to assert that the target is equal to or a subset
of the provided pattern array.
This API will perform subset match on the target.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": ["Flob", "Cat"]
//       }
//     }
//   }
// }

// The following will NOT throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Fred": awscdk.Match_arrayWith([]interface{}{
		jsii.String("Flob"),
	}),
})

// The following will throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), awscdk.Match_ObjectLike(map[string]interface{}{
	"Fred": awscdk.Match_arrayWith([]interface{}{
		jsii.String("Wobble"),
	}),
}))
```

*Note:* The list of items in the pattern array should be in order as they appear in the
target array. Out of order will be recorded as a match failure.

Alternatively, the `Match.arrayEquals()` API can be used to assert that the target is
exactly equal to the pattern array.

### String Matchers

The `Match.stringLikeRegexp()` API can be used to assert that the target matches the
provided regular expression.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Template": "const includeHeaders = true;"
//       }
//     }
//   }
// }

// The following will NOT throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Template": awscdk.Match_stringLikeRegexp(jsii.String("includeHeaders = (true|false)")),
})

// The following will throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Template": awscdk.Match_stringLikeRegexp(jsii.String("includeHeaders = null")),
})
```

### Not Matcher

The not matcher inverts the search pattern and matches all patterns in the path that does
not match the pattern specified.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": ["Flob", "Cat"]
//       }
//     }
//   }
// }

// The following will NOT throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Fred": awscdk.Match_not([]interface{}{
		jsii.String("Flob"),
	}),
})

// The following will throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), awscdk.Match_ObjectLike(map[string]interface{}{
	"Fred": awscdk.Match_not([]interface{}{
		jsii.String("Flob"),
		jsii.String("Cat"),
	}),
}))
```

### Serialized JSON

Often, we find that some CloudFormation Resource types declare properties as a string,
but actually expect JSON serialized as a string.
For example, the [`BuildSpec` property of `AWS::CodeBuild::Project`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-buildspec),
the [`Definition` property of `AWS::StepFunctions::StateMachine`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definition),
to name a couple.

The `Match.serializedJson()` matcher allows deep matching within a stringified JSON.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Baz": "{ \"Fred\": [\"Waldo\", \"Willow\"] }"
//       }
//     }
//   }
// }

// The following will NOT throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Baz": awscdk.Match_serializedJson(map[string]matcher{
		"Fred": awscdk.Match_arrayWith([]interface{}{
			jsii.String("Waldo"),
		}),
	}),
})

// The following will throw an assertion error
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
	"Baz": awscdk.Match_serializedJson(map[string][]*string{
		"Fred": []*string{
			jsii.String("Waldo"),
			jsii.String("Johnny"),
		},
	}),
})
```

## Capturing Values

The matcher APIs documented above allow capturing values in the matching entry
(Resource, Output, Mapping, etc.). The following code captures a string from a
matching resource.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": ["Flob", "Cat"],
//         "Waldo": ["Qix", "Qux"],
//       }
//     }
//   }
// }

fredCapture := awscdk.NewCapture()
waldoCapture := awscdk.NewCapture()
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]interface{}{
	"Fred": fredCapture,
	"Waldo": []interface{}{
		jsii.String("Qix"),
		waldoCapture,
	}.([]interface{}),
})

fredCapture.AsArray() // returns ["Flob", "Cat"]
waldoCapture.AsString()
```

With captures, a nested pattern can also be specified, so that only targets
that match the nested pattern will be captured. This pattern can be literals or
further Matchers.

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar1": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": ["Flob", "Cat"],
//       }
//     }
//     "MyBar2": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": ["Qix", "Qux"],
//       }
//     }
//   }
// }

capture := awscdk.NewCapture(awscdk.Match_ArrayWith([]interface{}{
	jsii.String("Cat"),
}))
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]capture{
	"Fred": capture,
})

capture.AsArray()
```

When multiple resources match the given condition, each `Capture` defined in
the condition will capture all matching values. They can be paged through using
the `next()` API. The following example illustrates this -

```go
// Given a template -
// {
//   "Resources": {
//     "MyBar": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": "Flob",
//       }
//     },
//     "MyBaz": {
//       "Type": "Foo::Bar",
//       "Properties": {
//         "Fred": "Quib",
//       }
//     }
//   }
// }

fredCapture := awscdk.NewCapture()
template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]capture{
	"Fred": fredCapture,
})

fredCapture.AsString() // returns "Flob"
fredCapture.Next() // returns true
fredCapture.AsString()
```

## Asserting Annotations

In addition to template matching, we provide an API for annotation matching.
[Annotations](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.Annotations.html)
can be added via the [Aspects](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.Aspects.html)
API. You can learn more about Aspects [here](https://docs.aws.amazon.com/cdk/v2/guide/aspects.html).

Say you have a `MyAspect` and a `MyStack` that uses `MyAspect`:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"

type myAspect struct {
}

func (this *myAspect) visit(node iConstruct) {
	if *node instanceof cdk.CfnResource && *node.CfnResourceType == "Foo::Bar" {
		this.error(*node, jsii.String("we do not want a Foo::Bar resource"))
	}
}

func (this *myAspect) error(node iConstruct, message *string) {
	cdk.Annotations_Of(*node).AddError(*message)
}

type myStack struct {
	stack
}

func newMyStack(scope construct, id *string) *myStack {
	this := &myStack{}
	cdk.NewStack_Override(this, scope, id)

	stack := cdk.NewStack()
	cdk.NewCfnResource(stack, jsii.String("Foo"), &CfnResourceProps{
		Type: jsii.String("Foo::Bar"),
		Properties: map[string]interface{}{
			"Fred": jsii.String("Thud"),
		},
	})
	cdk.Aspects_Of(stack).Add(NewMyAspect())
	return this
}
```

We can then assert that the stack contains the expected Error:

```go
// import { Annotations } from '@aws-cdk/assertions';

awscdk.Annotations_FromStack(stack).HasError(jsii.String("/Default/Foo"), jsii.String("we do not want a Foo::Bar resource"))
```

Here are the available APIs for `Annotations`:

* `hasError()`, `hasNoError()`, and `findError()`
* `hasWarning()`, `hasNoWarning()`, and `findWarning()`
* `hasInfo()`, `hasNoInfo()`, and `findInfo()`

The corresponding `findXxx()` API is complementary to the `hasXxx()` API, except instead
of asserting its presence, it returns the set of matching messages.

In addition, this suite of APIs is compatible with `Matchers` for more fine-grained control.
For example, the following assertion works as well:

```go
awscdk.Annotations_FromStack(stack).HasError(jsii.String("/Default/Foo"), awscdk.Match_StringLikeRegexp(jsii.String(".*Foo::Bar.*")))
```

## Asserting Stack tags

Tags applied to a `Stack` are not part of the rendered template: instead, they
are included as properties in the Cloud Assembly Manifest. To test that stacks
are tagged as expected, simple assertions can be written.

Given the following setup:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

app := awscdk.NewApp()
stack := awscdk.NewStack(app, jsii.String("MyStack"), &StackProps{
	Tags: map[string]*string{
		"tag-name": jsii.String("tag-value"),
	},
})
```

It is possible to test against these values:

```go
tags := awscdk.Tags_FromStack(stack)

// using a default 'objectLike' Matcher
tags.HasValues(map[string]*string{
	"tag-name": jsii.String("tag-value"),
})

// ... with Matchers embedded
tags.HasValues(map[string]matcher{
	"tag-name": awscdk.Match_stringLikeRegexp(jsii.String("value")),
})

// or another object Matcher at the top level
tags.HasValues(awscdk.Match_ObjectEquals(map[string]interface{}{
	"tag-name": awscdk.Match_anyValue(),
}))
```

When tags are not defined on the stack, it is represented as an empty object
rather than `undefined`. To make this more obvious, there is a `hasNone()`
method that can be used in place of `Match.exactly({})`. If `Match.absent()` is
passed, an error will result.

```go
// no tags present
awscdk.Tags_FromStack(stack).HasNone()

// don't use absent() at the top level, it won't work
expect(() => { Tags.fromStack(stack).hasValues(Match.absent()); }).toThrow(/will never match/i)
```
