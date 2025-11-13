package interfacesawsservicecatalog

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.AcceptedPortfolioShareReference",
		reflect.TypeOf((*AcceptedPortfolioShareReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.CloudFormationProductReference",
		reflect.TypeOf((*CloudFormationProductReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.CloudFormationProvisionedProductReference",
		reflect.TypeOf((*CloudFormationProvisionedProductReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IAcceptedPortfolioShareRef",
		reflect.TypeOf((*IAcceptedPortfolioShareRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceptedPortfolioShareRef", GoGetter: "AcceptedPortfolioShareRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IAcceptedPortfolioShareRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ICloudFormationProductRef",
		reflect.TypeOf((*ICloudFormationProductRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudFormationProductRef", GoGetter: "CloudFormationProductRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICloudFormationProductRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ICloudFormationProvisionedProductRef",
		reflect.TypeOf((*ICloudFormationProvisionedProductRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudFormationProvisionedProductRef", GoGetter: "CloudFormationProvisionedProductRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICloudFormationProvisionedProductRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ILaunchNotificationConstraintRef",
		reflect.TypeOf((*ILaunchNotificationConstraintRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "launchNotificationConstraintRef", GoGetter: "LaunchNotificationConstraintRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ILaunchNotificationConstraintRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ILaunchRoleConstraintRef",
		reflect.TypeOf((*ILaunchRoleConstraintRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "launchRoleConstraintRef", GoGetter: "LaunchRoleConstraintRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ILaunchRoleConstraintRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ILaunchTemplateConstraintRef",
		reflect.TypeOf((*ILaunchTemplateConstraintRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateConstraintRef", GoGetter: "LaunchTemplateConstraintRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ILaunchTemplateConstraintRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IPortfolioPrincipalAssociationRef",
		reflect.TypeOf((*IPortfolioPrincipalAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "portfolioPrincipalAssociationRef", GoGetter: "PortfolioPrincipalAssociationRef"},
		},
		func() interface{} {
			j := jsiiProxy_IPortfolioPrincipalAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IPortfolioProductAssociationRef",
		reflect.TypeOf((*IPortfolioProductAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "portfolioProductAssociationRef", GoGetter: "PortfolioProductAssociationRef"},
		},
		func() interface{} {
			j := jsiiProxy_IPortfolioProductAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IPortfolioRef",
		reflect.TypeOf((*IPortfolioRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "portfolioRef", GoGetter: "PortfolioRef"},
		},
		func() interface{} {
			j := jsiiProxy_IPortfolioRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IPortfolioShareRef",
		reflect.TypeOf((*IPortfolioShareRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "portfolioShareRef", GoGetter: "PortfolioShareRef"},
		},
		func() interface{} {
			j := jsiiProxy_IPortfolioShareRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IResourceUpdateConstraintRef",
		reflect.TypeOf((*IResourceUpdateConstraintRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceUpdateConstraintRef", GoGetter: "ResourceUpdateConstraintRef"},
		},
		func() interface{} {
			j := jsiiProxy_IResourceUpdateConstraintRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IServiceActionAssociationRef",
		reflect.TypeOf((*IServiceActionAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceActionAssociationRef", GoGetter: "ServiceActionAssociationRef"},
		},
		func() interface{} {
			j := jsiiProxy_IServiceActionAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IServiceActionRef",
		reflect.TypeOf((*IServiceActionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceActionRef", GoGetter: "ServiceActionRef"},
		},
		func() interface{} {
			j := jsiiProxy_IServiceActionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.IStackSetConstraintRef",
		reflect.TypeOf((*IStackSetConstraintRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stackSetConstraintRef", GoGetter: "StackSetConstraintRef"},
		},
		func() interface{} {
			j := jsiiProxy_IStackSetConstraintRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ITagOptionAssociationRef",
		reflect.TypeOf((*ITagOptionAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "tagOptionAssociationRef", GoGetter: "TagOptionAssociationRef"},
		},
		func() interface{} {
			j := jsiiProxy_ITagOptionAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ITagOptionRef",
		reflect.TypeOf((*ITagOptionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "tagOptionRef", GoGetter: "TagOptionRef"},
		},
		func() interface{} {
			j := jsiiProxy_ITagOptionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.LaunchNotificationConstraintReference",
		reflect.TypeOf((*LaunchNotificationConstraintReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.LaunchRoleConstraintReference",
		reflect.TypeOf((*LaunchRoleConstraintReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.LaunchTemplateConstraintReference",
		reflect.TypeOf((*LaunchTemplateConstraintReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.PortfolioPrincipalAssociationReference",
		reflect.TypeOf((*PortfolioPrincipalAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.PortfolioProductAssociationReference",
		reflect.TypeOf((*PortfolioProductAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.PortfolioReference",
		reflect.TypeOf((*PortfolioReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.PortfolioShareReference",
		reflect.TypeOf((*PortfolioShareReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ResourceUpdateConstraintReference",
		reflect.TypeOf((*ResourceUpdateConstraintReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ServiceActionAssociationReference",
		reflect.TypeOf((*ServiceActionAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.ServiceActionReference",
		reflect.TypeOf((*ServiceActionReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.StackSetConstraintReference",
		reflect.TypeOf((*StackSetConstraintReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.TagOptionAssociationReference",
		reflect.TypeOf((*TagOptionAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_servicecatalog.TagOptionReference",
		reflect.TypeOf((*TagOptionReference)(nil)).Elem(),
	)
}
