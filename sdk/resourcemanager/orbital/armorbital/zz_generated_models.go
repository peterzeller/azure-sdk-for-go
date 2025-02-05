//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// AvailableContacts - Customer retrieves list of Available Contacts for a spacecraft resource. Later, one of the available contact can be selected to create
// a contact.
type AvailableContacts struct {
	// Properties of Contact resource.
	Properties *ContactInstanceProperties `json:"properties,omitempty"`

	// The reference to the spacecraft resource.
	Spacecraft *ResourceReference `json:"spacecraft,omitempty"`

	// READ-ONLY; Name of Azure Ground Station.
	GroundStationName *string `json:"groundStationName,omitempty" azure:"ro"`
}

// AvailableContactsListResult - Response for the ListAvailableContacts API service call.
type AvailableContactsListResult struct {
	// A list of available contacts
	Value []*AvailableContacts `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type AvailableContactsListResult.
func (a AvailableContactsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// AvailableGroundStation - GroundStations available to schedule Contacts
type AvailableGroundStation struct {
	// Azure region
	Location *string `json:"location,omitempty"`

	// The properties bag for this resource
	Properties *AvailableGroundStationProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of groundStation
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the ground station.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AvailableGroundStationListResult - Response for the AvailableGroundStations API service call.
type AvailableGroundStationListResult struct {
	// A list of ground station resources.
	Value []*AvailableGroundStation `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type AvailableGroundStationListResult.
func (a AvailableGroundStationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// AvailableGroundStationProperties - Properties object for Available groundstation.
type AvailableGroundStationProperties struct {
	// Altitude of the ground station
	AltitudeMeters *float32 `json:"altitudeMeters,omitempty"`

	// City of ground station.
	City *string `json:"city,omitempty"`

	// Latitude of the ground station in decimal degrees.
	LatitudeDegrees *float32 `json:"latitudeDegrees,omitempty"`

	// Longitude of the ground station in decimal degrees.
	LongitudeDegrees *float32 `json:"longitudeDegrees,omitempty"`

	// Ground station provider name.
	ProviderName *string `json:"providerName,omitempty"`
}

// AvailableGroundStationsGetOptions contains the optional parameters for the AvailableGroundStations.Get method.
type AvailableGroundStationsGetOptions struct {
	// placeholder for future optional parameters
}

// AvailableGroundStationsListByCapabilityOptions contains the optional parameters for the AvailableGroundStations.ListByCapability method.
type AvailableGroundStationsListByCapabilityOptions struct {
	// placeholder for future optional parameters
}

// CloudError - An error response from the service.
// Implements the error and azcore.HTTPResponse interfaces.
type CloudError struct {
	raw string
	// An error response from the service.
	InnerError *CloudErrorBody `json:"error,omitempty"`
}

// Error implements the error interface for type CloudError.
// The contents of the error text are not contractual and subject to change.
func (e CloudError) Error() string {
	return e.raw
}

// CloudErrorBody - An error response from the service.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// A list of additional details about the error.
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`

	// The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// Contact - Customer creates a contact resource for a spacecraft resource.
type Contact struct {
	ProxyResource
	// Properties of the Contact Resource.
	Properties *ContactsProperties `json:"properties,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Contact.
func (c Contact) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	c.ProxyResource.marshalInternal(objectMap)
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "properties", c.Properties)
	return json.Marshal(objectMap)
}

// ContactInstanceProperties - Contact Instance Properties
type ContactInstanceProperties struct {
	// READ-ONLY; Azimuth of the antenna at the end of the contact in decimal degrees.
	EndAzimuthDegrees *float32 `json:"endAzimuthDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Spacecraft elevation above the horizon at contact end.
	EndElevationDegrees *float32 `json:"endElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Maximum elevation of the antenna during the contact in decimal degrees.
	MaximumElevationDegrees *float32 `json:"maximumElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Time to lost receiving a signal.
	RxEndTime *time.Time `json:"rxEndTime,omitempty" azure:"ro"`

	// READ-ONLY; Earliest time to receive a signal.
	RxStartTime *time.Time `json:"rxStartTime,omitempty" azure:"ro"`

	// READ-ONLY; Azimuth of the antenna at the start of the contact in decimal degrees.
	StartAzimuthDegrees *float32 `json:"startAzimuthDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Spacecraft elevation above the horizon at contact start.
	StartElevationDegrees *float32 `json:"startElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Time at which antenna transmit will be disabled.
	TxEndTime *time.Time `json:"txEndTime,omitempty" azure:"ro"`

	// READ-ONLY; Time at which antenna transmit will be enabled.
	TxStartTime *time.Time `json:"txStartTime,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ContactInstanceProperties.
func (c ContactInstanceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endAzimuthDegrees", c.EndAzimuthDegrees)
	populate(objectMap, "endElevationDegrees", c.EndElevationDegrees)
	populate(objectMap, "maximumElevationDegrees", c.MaximumElevationDegrees)
	populateTimeRFC3339(objectMap, "rxEndTime", c.RxEndTime)
	populateTimeRFC3339(objectMap, "rxStartTime", c.RxStartTime)
	populate(objectMap, "startAzimuthDegrees", c.StartAzimuthDegrees)
	populate(objectMap, "startElevationDegrees", c.StartElevationDegrees)
	populateTimeRFC3339(objectMap, "txEndTime", c.TxEndTime)
	populateTimeRFC3339(objectMap, "txStartTime", c.TxStartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactInstanceProperties.
func (c *ContactInstanceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endAzimuthDegrees":
			err = unpopulate(val, &c.EndAzimuthDegrees)
			delete(rawMsg, key)
		case "endElevationDegrees":
			err = unpopulate(val, &c.EndElevationDegrees)
			delete(rawMsg, key)
		case "maximumElevationDegrees":
			err = unpopulate(val, &c.MaximumElevationDegrees)
			delete(rawMsg, key)
		case "rxEndTime":
			err = unpopulateTimeRFC3339(val, &c.RxEndTime)
			delete(rawMsg, key)
		case "rxStartTime":
			err = unpopulateTimeRFC3339(val, &c.RxStartTime)
			delete(rawMsg, key)
		case "startAzimuthDegrees":
			err = unpopulate(val, &c.StartAzimuthDegrees)
			delete(rawMsg, key)
		case "startElevationDegrees":
			err = unpopulate(val, &c.StartElevationDegrees)
			delete(rawMsg, key)
		case "txEndTime":
			err = unpopulateTimeRFC3339(val, &c.TxEndTime)
			delete(rawMsg, key)
		case "txStartTime":
			err = unpopulateTimeRFC3339(val, &c.TxStartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ContactListResult - Response for the ListContacts API service call.
type ContactListResult struct {
	// A list of contact resources in a resource group.
	Value []*Contact `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ContactListResult.
func (c ContactListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// ContactParameters - Parameters that define the contact resource.
type ContactParameters struct {
	// REQUIRED; The reference to the contact profile resource.
	ContactProfile *ResourceReference `json:"contactProfile,omitempty"`

	// REQUIRED; End time of a contact.
	EndTime *time.Time `json:"endTime,omitempty"`

	// REQUIRED; Name of Azure Ground Station.
	GroundStationName *string `json:"groundStationName,omitempty"`

	// REQUIRED; Start time of a contact.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ContactParameters.
func (c ContactParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contactProfile", c.ContactProfile)
	populateTimeRFC3339(objectMap, "endTime", c.EndTime)
	populate(objectMap, "groundStationName", c.GroundStationName)
	populateTimeRFC3339(objectMap, "startTime", c.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactParameters.
func (c *ContactParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contactProfile":
			err = unpopulate(val, &c.ContactProfile)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &c.EndTime)
			delete(rawMsg, key)
		case "groundStationName":
			err = unpopulate(val, &c.GroundStationName)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &c.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ContactProfile - Customer creates a Contact Profile Resource, which will contain all of the configurations required for scheduling a contact.
type ContactProfile struct {
	TrackedResource
	// Properties of the spacecraft resource.
	Properties *ContactProfilesProperties `json:"properties,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfile.
func (c ContactProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	c.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "properties", c.Properties)
	return json.Marshal(objectMap)
}

// ContactProfileLink - Contact Profile link
type ContactProfileLink struct {
	// REQUIRED; Contact Profile Link Channel
	Channels []*ContactProfileLinkChannel `json:"channels,omitempty"`

	// REQUIRED; Direction (uplink or downlink)
	Direction *Direction `json:"direction,omitempty"`

	// REQUIRED; polarization. eg (RHCP, LHCP)
	Polarization *Polarization `json:"polarization,omitempty"`

	// Effective Isotropic Radiated Power (EIRP) in dBW.
	EirpdBW *float32 `json:"eirpdBW,omitempty"`

	// Gain To Noise Temperature in db/K.
	GainOverTemperature *float32 `json:"gainOverTemperature,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfileLink.
func (c ContactProfileLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "channels", c.Channels)
	populate(objectMap, "direction", c.Direction)
	populate(objectMap, "eirpdBW", c.EirpdBW)
	populate(objectMap, "gainOverTemperature", c.GainOverTemperature)
	populate(objectMap, "polarization", c.Polarization)
	return json.Marshal(objectMap)
}

// ContactProfileLinkChannel - Contact Profile Link Channel
type ContactProfileLinkChannel struct {
	// REQUIRED; Bandwidth in MHz
	BandwidthMHz *float32 `json:"bandwidthMHz,omitempty"`

	// REQUIRED; Center Frequency in MHz
	CenterFrequencyMHz *float32 `json:"centerFrequencyMHz,omitempty"`

	// REQUIRED; Customer End point to store/retrieve data during a contact.
	EndPoint *EndPoint `json:"endPoint,omitempty"`

	// Configuration for decoding
	DecodingConfiguration *string `json:"decodingConfiguration,omitempty"`

	// Configuration for demodulation
	DemodulationConfiguration *string `json:"demodulationConfiguration,omitempty"`

	// Configuration for encoding
	EncodingConfiguration *string `json:"encodingConfiguration,omitempty"`

	// Configuration for modulation
	ModulationConfiguration *string `json:"modulationConfiguration,omitempty"`
}

// ContactProfileListResult - Response for the ListContactProfiles API service call.
type ContactProfileListResult struct {
	// A list of contact profile resources in a resource group.
	Value []*ContactProfile `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfileListResult.
func (c ContactProfileListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// ContactProfilesBeginCreateOrUpdateOptions contains the optional parameters for the ContactProfiles.BeginCreateOrUpdate method.
type ContactProfilesBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesBeginDeleteOptions contains the optional parameters for the ContactProfiles.BeginDelete method.
type ContactProfilesBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesGetOptions contains the optional parameters for the ContactProfiles.Get method.
type ContactProfilesGetOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesListBySubscriptionOptions contains the optional parameters for the ContactProfiles.ListBySubscription method.
type ContactProfilesListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesListOptions contains the optional parameters for the ContactProfiles.List method.
type ContactProfilesListOptions struct {
	// placeholder for future optional parameters
}

// ContactProfilesProperties - List of Contact Profile Resource Properties.
type ContactProfilesProperties struct {
	// REQUIRED; Links of the Contact Profile
	Links []*ContactProfileLink `json:"links,omitempty"`

	// Auto track configuration.
	AutoTrackingConfiguration *AutoTrackingConfiguration `json:"autoTrackingConfiguration,omitempty"`

	// Minimum viable elevation for the contact in decimal degrees.
	MinimumElevationDegrees *float32 `json:"minimumElevationDegrees,omitempty"`

	// Minimum viable contact duration in ISO 8601 format.
	MinimumViableContactDuration *string `json:"minimumViableContactDuration,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfilesProperties.
func (c ContactProfilesProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoTrackingConfiguration", c.AutoTrackingConfiguration)
	populate(objectMap, "links", c.Links)
	populate(objectMap, "minimumElevationDegrees", c.MinimumElevationDegrees)
	populate(objectMap, "minimumViableContactDuration", c.MinimumViableContactDuration)
	return json.Marshal(objectMap)
}

// ContactProfilesUpdateTagsOptions contains the optional parameters for the ContactProfiles.UpdateTags method.
type ContactProfilesUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// ContactsBeginCreateOptions contains the optional parameters for the Contacts.BeginCreate method.
type ContactsBeginCreateOptions struct {
	// placeholder for future optional parameters
}

// ContactsBeginDeleteOptions contains the optional parameters for the Contacts.BeginDelete method.
type ContactsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// ContactsGetOptions contains the optional parameters for the Contacts.Get method.
type ContactsGetOptions struct {
	// placeholder for future optional parameters
}

// ContactsListOptions contains the optional parameters for the Contacts.List method.
type ContactsListOptions struct {
	// placeholder for future optional parameters
}

// ContactsProperties - Properties of the Contact Resource.
type ContactsProperties struct {
	// REQUIRED; The reference to the contact profile resource.
	ContactProfile *ResourceReference `json:"contactProfile,omitempty"`

	// REQUIRED; Azure Ground Station name.
	GroundStationName *string `json:"groundStationName,omitempty"`

	// REQUIRED; Reservation end time of a contact.
	ReservationEndTime *time.Time `json:"reservationEndTime,omitempty"`

	// REQUIRED; Reservation start time of a contact.
	ReservationStartTime *time.Time `json:"reservationStartTime,omitempty"`

	// READ-ONLY; Azimuth of the antenna at the end of the contact in decimal degrees.
	EndAzimuthDegrees *float32 `json:"endAzimuthDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Spacecraft elevation above the horizon at contact end.
	EndElevationDegrees *float32 `json:"endElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Any error message while scheduling a contact.
	ErrorMessage *string `json:"errorMessage,omitempty" azure:"ro"`

	// READ-ONLY; Maximum elevation of the antenna during the contact in decimal degrees.
	MaximumElevationDegrees *float32 `json:"maximumElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Receive end time of a contact.
	RxEndTime *time.Time `json:"rxEndTime,omitempty" azure:"ro"`

	// READ-ONLY; Receive start time of a contact.
	RxStartTime *time.Time `json:"rxStartTime,omitempty" azure:"ro"`

	// READ-ONLY; Azimuth of the antenna at the start of the contact in decimal degrees.
	StartAzimuthDegrees *float32 `json:"startAzimuthDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Spacecraft elevation above the horizon at contact start.
	StartElevationDegrees *float32 `json:"startElevationDegrees,omitempty" azure:"ro"`

	// READ-ONLY; Status of a contact.
	Status *Status `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; Transmit end time of a contact.
	TxEndTime *time.Time `json:"txEndTime,omitempty" azure:"ro"`

	// READ-ONLY; Transmit start time of a contact.
	TxStartTime *time.Time `json:"txStartTime,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ContactsProperties.
func (c ContactsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contactProfile", c.ContactProfile)
	populate(objectMap, "endAzimuthDegrees", c.EndAzimuthDegrees)
	populate(objectMap, "endElevationDegrees", c.EndElevationDegrees)
	populate(objectMap, "errorMessage", c.ErrorMessage)
	populate(objectMap, "groundStationName", c.GroundStationName)
	populate(objectMap, "maximumElevationDegrees", c.MaximumElevationDegrees)
	populateTimeRFC3339(objectMap, "reservationEndTime", c.ReservationEndTime)
	populateTimeRFC3339(objectMap, "reservationStartTime", c.ReservationStartTime)
	populateTimeRFC3339(objectMap, "rxEndTime", c.RxEndTime)
	populateTimeRFC3339(objectMap, "rxStartTime", c.RxStartTime)
	populate(objectMap, "startAzimuthDegrees", c.StartAzimuthDegrees)
	populate(objectMap, "startElevationDegrees", c.StartElevationDegrees)
	populate(objectMap, "status", c.Status)
	populateTimeRFC3339(objectMap, "txEndTime", c.TxEndTime)
	populateTimeRFC3339(objectMap, "txStartTime", c.TxStartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactsProperties.
func (c *ContactsProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contactProfile":
			err = unpopulate(val, &c.ContactProfile)
			delete(rawMsg, key)
		case "endAzimuthDegrees":
			err = unpopulate(val, &c.EndAzimuthDegrees)
			delete(rawMsg, key)
		case "endElevationDegrees":
			err = unpopulate(val, &c.EndElevationDegrees)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, &c.ErrorMessage)
			delete(rawMsg, key)
		case "groundStationName":
			err = unpopulate(val, &c.GroundStationName)
			delete(rawMsg, key)
		case "maximumElevationDegrees":
			err = unpopulate(val, &c.MaximumElevationDegrees)
			delete(rawMsg, key)
		case "reservationEndTime":
			err = unpopulateTimeRFC3339(val, &c.ReservationEndTime)
			delete(rawMsg, key)
		case "reservationStartTime":
			err = unpopulateTimeRFC3339(val, &c.ReservationStartTime)
			delete(rawMsg, key)
		case "rxEndTime":
			err = unpopulateTimeRFC3339(val, &c.RxEndTime)
			delete(rawMsg, key)
		case "rxStartTime":
			err = unpopulateTimeRFC3339(val, &c.RxStartTime)
			delete(rawMsg, key)
		case "startAzimuthDegrees":
			err = unpopulate(val, &c.StartAzimuthDegrees)
			delete(rawMsg, key)
		case "startElevationDegrees":
			err = unpopulate(val, &c.StartElevationDegrees)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &c.Status)
			delete(rawMsg, key)
		case "txEndTime":
			err = unpopulateTimeRFC3339(val, &c.TxEndTime)
			delete(rawMsg, key)
		case "txStartTime":
			err = unpopulateTimeRFC3339(val, &c.TxStartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// EndPoint - Customer End point to store/retrieve data during a contact.
type EndPoint struct {
	// REQUIRED; Name of an end point.
	EndPointName *string `json:"endPointName,omitempty"`

	// REQUIRED; IP Address.
	IPAddress *string `json:"ipAddress,omitempty"`

	// REQUIRED; TCP port to listen on to receive data.
	Port *string `json:"port,omitempty"`

	// REQUIRED; Protocol either UDP or TCP.
	Protocol *Protocol `json:"protocol,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write", "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual Machine", "Restart Virtual
	// Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// placeholder for future optional parameters
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a location
type ProxyResource struct {
	Resource
}

func (p ProxyResource) marshalInternal(objectMap map[string]interface{}) {
	p.Resource.marshalInternal(objectMap)
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "type", r.Type)
}

// ResourceIDListResult - Response for an API service call that lists the resource IDs of resources associated with another resource.
type ResourceIDListResult struct {
	// A list of Azure Resource IDs.
	Value []*ResourceIDListResultValueItem `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceIDListResult.
func (r ResourceIDListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

type ResourceIDListResultValueItem struct {
	// The Azure Resource ID
	ID *string `json:"id,omitempty"`
}

// ResourceReference - Resource Reference
type ResourceReference struct {
	// Resource ID.
	ID *string `json:"id,omitempty"`
}

// Spacecraft - Customer creates a spacecraft resource to schedule a contact.
type Spacecraft struct {
	TrackedResource
	// Spacecraft Properties
	Properties *SpacecraftsProperties `json:"properties,omitempty"`

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Spacecraft.
func (s Spacecraft) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "etag", s.Etag)
	populate(objectMap, "properties", s.Properties)
	return json.Marshal(objectMap)
}

// SpacecraftLink - Spacecraft Link
type SpacecraftLink struct {
	// REQUIRED; Bandwidth in MHz
	BandwidthMHz *float32 `json:"bandwidthMHz,omitempty"`

	// REQUIRED; Center Frequency in MHz
	CenterFrequencyMHz *float32 `json:"centerFrequencyMHz,omitempty"`

	// REQUIRED; Direction (uplink or downlink)
	Direction *Direction `json:"direction,omitempty"`

	// REQUIRED; polarization. eg (RHCP, LHCP)
	Polarization *Polarization `json:"polarization,omitempty"`
}

// SpacecraftListResult - Response for the ListSpacecrafts API service call.
type SpacecraftListResult struct {
	// A list of spacecraft resources in a resource group.
	Value []*Spacecraft `json:"value,omitempty"`

	// READ-ONLY; The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type SpacecraftListResult.
func (s SpacecraftListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// SpacecraftsBeginCreateOrUpdateOptions contains the optional parameters for the Spacecrafts.BeginCreateOrUpdate method.
type SpacecraftsBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsBeginDeleteOptions contains the optional parameters for the Spacecrafts.BeginDelete method.
type SpacecraftsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsBeginListAvailableContactsOptions contains the optional parameters for the Spacecrafts.BeginListAvailableContacts method.
type SpacecraftsBeginListAvailableContactsOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsGetOptions contains the optional parameters for the Spacecrafts.Get method.
type SpacecraftsGetOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsListBySubscriptionOptions contains the optional parameters for the Spacecrafts.ListBySubscription method.
type SpacecraftsListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsListOptions contains the optional parameters for the Spacecrafts.List method.
type SpacecraftsListOptions struct {
	// placeholder for future optional parameters
}

// SpacecraftsProperties - List of Spacecraft Resource Properties.
type SpacecraftsProperties struct {
	// REQUIRED; NORAD ID of the spacecraft.
	NoradID *string `json:"noradId,omitempty"`

	// Links of the Spacecraft
	Links []*SpacecraftLink `json:"links,omitempty"`

	// Title line of Two Line Element (TLE).
	TitleLine *string `json:"titleLine,omitempty"`

	// Line 1 of Two Line Element (TLE).
	TleLine1 *string `json:"tleLine1,omitempty"`

	// Line 2 of Two Line Element (TLE).
	TleLine2 *string `json:"tleLine2,omitempty"`

	// READ-ONLY; Authorization status of spacecraft.
	AuthorizationStatus *AuthorizationStatus `json:"authorizationStatus,omitempty" azure:"ro"`

	// READ-ONLY; Details of the authorization status.
	AuthorizationStatusExtended *string `json:"authorizationStatusExtended,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type SpacecraftsProperties.
func (s SpacecraftsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizationStatus", s.AuthorizationStatus)
	populate(objectMap, "authorizationStatusExtended", s.AuthorizationStatusExtended)
	populate(objectMap, "links", s.Links)
	populate(objectMap, "noradId", s.NoradID)
	populate(objectMap, "titleLine", s.TitleLine)
	populate(objectMap, "tleLine1", s.TleLine1)
	populate(objectMap, "tleLine2", s.TleLine2)
	return json.Marshal(objectMap)
}

// SpacecraftsUpdateTagsOptions contains the optional parameters for the Spacecrafts.UpdateTags method.
type SpacecraftsUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// TagsObject - Tags object for patch operations.
type TagsObject struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TagsObject.
func (t TagsObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags' and a 'location'
type TrackedResource struct {
	Resource
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	t.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (t TrackedResource) marshalInternal(objectMap map[string]interface{}) {
	t.Resource.marshalInternal(objectMap)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "tags", t.Tags)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
