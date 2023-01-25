// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Inspector that maintains an attribute bag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   treeInspector := monocdk.NewTreeInspector()
//
// Experimental.
type TreeInspector interface {
	// Represents the bag of attributes as key-value pairs.
	// Experimental.
	Attributes() *map[string]interface{}
	// Adds attribute to bag.
	//
	// Keys should be added by convention to prevent conflicts
	// i.e. L1 constructs will contain attributes with keys prefixed with aws:cdk:cloudformation
	// Experimental.
	AddAttribute(key *string, value interface{})
}

// The jsii proxy struct for TreeInspector
type jsiiProxy_TreeInspector struct {
	_ byte // padding
}

func (j *jsiiProxy_TreeInspector) Attributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}


// Experimental.
func NewTreeInspector() TreeInspector {
	_init_.Initialize()

	j := jsiiProxy_TreeInspector{}

	_jsii_.Create(
		"monocdk.TreeInspector",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTreeInspector_Override(t TreeInspector) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.TreeInspector",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TreeInspector) AddAttribute(key *string, value interface{}) {
	if err := t.validateAddAttributeParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addAttribute",
		[]interface{}{key, value},
	)
}

