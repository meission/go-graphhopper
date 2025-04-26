# PreferredVehicle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleId** | Pointer to **string** | Id of the preferred vehicle. | [optional] 
**Priority** | Pointer to **int64** | Number between 1 and 10 which indicates the priority of the preferred vehicle. 1 indicates the highest priority, 10 the lowest. | [optional] [default to 2]

## Methods

### NewPreferredVehicle

`func NewPreferredVehicle() *PreferredVehicle`

NewPreferredVehicle instantiates a new PreferredVehicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferredVehicleWithDefaults

`func NewPreferredVehicleWithDefaults() *PreferredVehicle`

NewPreferredVehicleWithDefaults instantiates a new PreferredVehicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleId

`func (o *PreferredVehicle) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *PreferredVehicle) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *PreferredVehicle) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *PreferredVehicle) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### GetPriority

`func (o *PreferredVehicle) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PreferredVehicle) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PreferredVehicle) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PreferredVehicle) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


