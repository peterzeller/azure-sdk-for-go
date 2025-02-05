//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

const (
	module  = "armhardwaresecuritymodules"
	version = "v0.1.0"
)

// JSONWebKeyType - Provisioning state.
type JSONWebKeyType string

const (
	// JSONWebKeyTypeAllocating - A device is currently being allocated for the dedicated HSM resource.
	JSONWebKeyTypeAllocating JSONWebKeyType = "Allocating"
	// JSONWebKeyTypeCheckingQuota - Validating the subscription has sufficient quota to allocate a dedicated HSM device.
	JSONWebKeyTypeCheckingQuota JSONWebKeyType = "CheckingQuota"
	// JSONWebKeyTypeConnecting - The dedicated HSM is being connected to the virtual network.
	JSONWebKeyTypeConnecting JSONWebKeyType = "Connecting"
	// JSONWebKeyTypeDeleting - The dedicated HSM is currently being deleted.
	JSONWebKeyTypeDeleting JSONWebKeyType = "Deleting"
	// JSONWebKeyTypeFailed - Provisioning of the dedicated HSM has failed.
	JSONWebKeyTypeFailed JSONWebKeyType = "Failed"
	// JSONWebKeyTypeProvisioning - The dedicated HSM is currently being provisioned.
	JSONWebKeyTypeProvisioning JSONWebKeyType = "Provisioning"
	// JSONWebKeyTypeSucceeded - The dedicated HSM has been full provisioned.
	JSONWebKeyTypeSucceeded JSONWebKeyType = "Succeeded"
)

// PossibleJSONWebKeyTypeValues returns the possible values for the JSONWebKeyType const type.
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return []JSONWebKeyType{
		JSONWebKeyTypeAllocating,
		JSONWebKeyTypeCheckingQuota,
		JSONWebKeyTypeConnecting,
		JSONWebKeyTypeDeleting,
		JSONWebKeyTypeFailed,
		JSONWebKeyTypeProvisioning,
		JSONWebKeyTypeSucceeded,
	}
}

// ToPtr returns a *JSONWebKeyType pointing to the current value.
func (c JSONWebKeyType) ToPtr() *JSONWebKeyType {
	return &c
}

// SKUName - SKU of the dedicated HSM
type SKUName string

const (
	SKUNameSafeNetLunaNetworkHSMA790 SKUName = "SafeNet Luna Network HSM A790"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameSafeNetLunaNetworkHSMA790,
	}
}

// ToPtr returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}
