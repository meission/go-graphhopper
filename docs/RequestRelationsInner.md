# RequestRelationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Specifies the type of relation. It must be either of type &#x60;in_sequence&#x60;, &#x60;in_direct_sequence&#x60; or &#x60;neighbor&#x60;.  | 
**Ids** | **[]string** | Specifies an array of shipment and/or service ids that are in relation. If you deal with services then you need to use the id of your services in ids. To also consider sequences of the pickups and deliveries of your shipments, you need to use a special ID, i.e. use your shipment id plus the keyword &#x60;_pickup&#x60; or &#x60;_delivery&#x60;. If you want to place a service or shipment activity at the beginning of your route, use the special ID &#x60;start&#x60;. In turn, use &#x60;end&#x60; to place it at the end of the route. | 
**VehicleId** | Pointer to **string** | Id of pre-assigned vehicle, i.e. the vehicle id that is determined to conduct the services and shipments in this relation. | [optional] 
**Groups** | **[]string** | An array of groups that should be related | 

## Methods

### NewRequestRelationsInner

`func NewRequestRelationsInner(type_ string, ids []string, groups []string, ) *RequestRelationsInner`

NewRequestRelationsInner instantiates a new RequestRelationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestRelationsInnerWithDefaults

`func NewRequestRelationsInnerWithDefaults() *RequestRelationsInner`

NewRequestRelationsInnerWithDefaults instantiates a new RequestRelationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestRelationsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestRelationsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestRelationsInner) SetType(v string)`

SetType sets Type field to given value.


### GetIds

`func (o *RequestRelationsInner) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *RequestRelationsInner) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *RequestRelationsInner) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetVehicleId

`func (o *RequestRelationsInner) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *RequestRelationsInner) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *RequestRelationsInner) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *RequestRelationsInner) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### GetGroups

`func (o *RequestRelationsInner) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RequestRelationsInner) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RequestRelationsInner) SetGroups(v []string)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


