# Shift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShiftId** | Pointer to **string** | A unique identifier for this vehicle&#39;s shift. | [optional] 
**StartAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**EndAddress** | Pointer to [**Address**](Address.md) | If this is omitted AND return_to_depot is true then the vehicle needs to return to its start_address. | [optional] 
**EarliestStart** | Pointer to **int64** | Earliest start of vehicle in seconds. It is recommended to use the unix timestamp. | [optional] [default to 0]
**LatestEnd** | Pointer to **int64** | Latest end of vehicle in seconds, i.e. the time the vehicle needs to be at its end location at latest. | [optional] [default to 9223372036854776000]
**ReturnToDepot** | Pointer to **bool** | If true, vehicle returns to its start location (or specified end location). If false, vehicle can end at any customer location that optimizes the objective function. | [optional] [default to true]
**Break** | Pointer to [**VehicleBreak**](VehicleBreak.md) |  | [optional] 

## Methods

### NewShift

`func NewShift() *Shift`

NewShift instantiates a new Shift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftWithDefaults

`func NewShiftWithDefaults() *Shift`

NewShiftWithDefaults instantiates a new Shift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShiftId

`func (o *Shift) GetShiftId() string`

GetShiftId returns the ShiftId field if non-nil, zero value otherwise.

### GetShiftIdOk

`func (o *Shift) GetShiftIdOk() (*string, bool)`

GetShiftIdOk returns a tuple with the ShiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftId

`func (o *Shift) SetShiftId(v string)`

SetShiftId sets ShiftId field to given value.

### HasShiftId

`func (o *Shift) HasShiftId() bool`

HasShiftId returns a boolean if a field has been set.

### GetStartAddress

`func (o *Shift) GetStartAddress() Address`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *Shift) GetStartAddressOk() (*Address, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *Shift) SetStartAddress(v Address)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *Shift) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *Shift) GetEndAddress() Address`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *Shift) GetEndAddressOk() (*Address, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *Shift) SetEndAddress(v Address)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *Shift) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetEarliestStart

`func (o *Shift) GetEarliestStart() int64`

GetEarliestStart returns the EarliestStart field if non-nil, zero value otherwise.

### GetEarliestStartOk

`func (o *Shift) GetEarliestStartOk() (*int64, bool)`

GetEarliestStartOk returns a tuple with the EarliestStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestStart

`func (o *Shift) SetEarliestStart(v int64)`

SetEarliestStart sets EarliestStart field to given value.

### HasEarliestStart

`func (o *Shift) HasEarliestStart() bool`

HasEarliestStart returns a boolean if a field has been set.

### GetLatestEnd

`func (o *Shift) GetLatestEnd() int64`

GetLatestEnd returns the LatestEnd field if non-nil, zero value otherwise.

### GetLatestEndOk

`func (o *Shift) GetLatestEndOk() (*int64, bool)`

GetLatestEndOk returns a tuple with the LatestEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEnd

`func (o *Shift) SetLatestEnd(v int64)`

SetLatestEnd sets LatestEnd field to given value.

### HasLatestEnd

`func (o *Shift) HasLatestEnd() bool`

HasLatestEnd returns a boolean if a field has been set.

### GetReturnToDepot

`func (o *Shift) GetReturnToDepot() bool`

GetReturnToDepot returns the ReturnToDepot field if non-nil, zero value otherwise.

### GetReturnToDepotOk

`func (o *Shift) GetReturnToDepotOk() (*bool, bool)`

GetReturnToDepotOk returns a tuple with the ReturnToDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToDepot

`func (o *Shift) SetReturnToDepot(v bool)`

SetReturnToDepot sets ReturnToDepot field to given value.

### HasReturnToDepot

`func (o *Shift) HasReturnToDepot() bool`

HasReturnToDepot returns a boolean if a field has been set.

### GetBreak

`func (o *Shift) GetBreak() VehicleBreak`

GetBreak returns the Break field if non-nil, zero value otherwise.

### GetBreakOk

`func (o *Shift) GetBreakOk() (*VehicleBreak, bool)`

GetBreakOk returns a tuple with the Break field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreak

`func (o *Shift) SetBreak(v VehicleBreak)`

SetBreak sets Break field to given value.

### HasBreak

`func (o *Shift) HasBreak() bool`

HasBreak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


