# VehicleBreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Earliest** | **int64** | Specifies the earliest start time of the break in seconds. | 
**Latest** | **int64** | Specifies the latest start time of break in seconds. | 
**Duration** | **int64** | Specifies the duration of the break in seconds. | 
**MaxDrivingTime** | **int64** | Specifies the max driving time (in a row) without break in seconds. | 
**InitialDrivingTime** | Pointer to **int64** | Specifies the initial (current) driving time of a driver to allow dynamic adaptations in seconds. | [optional] 
**PossibleSplit** | Pointer to **[]int64** | Array specifying how a break duration (in seconds) can be split into several smaller breaks | [optional] 

## Methods

### NewVehicleBreak

`func NewVehicleBreak(earliest int64, latest int64, duration int64, maxDrivingTime int64, ) *VehicleBreak`

NewVehicleBreak instantiates a new VehicleBreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleBreakWithDefaults

`func NewVehicleBreakWithDefaults() *VehicleBreak`

NewVehicleBreakWithDefaults instantiates a new VehicleBreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEarliest

`func (o *VehicleBreak) GetEarliest() int64`

GetEarliest returns the Earliest field if non-nil, zero value otherwise.

### GetEarliestOk

`func (o *VehicleBreak) GetEarliestOk() (*int64, bool)`

GetEarliestOk returns a tuple with the Earliest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliest

`func (o *VehicleBreak) SetEarliest(v int64)`

SetEarliest sets Earliest field to given value.


### GetLatest

`func (o *VehicleBreak) GetLatest() int64`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *VehicleBreak) GetLatestOk() (*int64, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *VehicleBreak) SetLatest(v int64)`

SetLatest sets Latest field to given value.


### GetDuration

`func (o *VehicleBreak) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VehicleBreak) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VehicleBreak) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetMaxDrivingTime

`func (o *VehicleBreak) GetMaxDrivingTime() int64`

GetMaxDrivingTime returns the MaxDrivingTime field if non-nil, zero value otherwise.

### GetMaxDrivingTimeOk

`func (o *VehicleBreak) GetMaxDrivingTimeOk() (*int64, bool)`

GetMaxDrivingTimeOk returns a tuple with the MaxDrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDrivingTime

`func (o *VehicleBreak) SetMaxDrivingTime(v int64)`

SetMaxDrivingTime sets MaxDrivingTime field to given value.


### GetInitialDrivingTime

`func (o *VehicleBreak) GetInitialDrivingTime() int64`

GetInitialDrivingTime returns the InitialDrivingTime field if non-nil, zero value otherwise.

### GetInitialDrivingTimeOk

`func (o *VehicleBreak) GetInitialDrivingTimeOk() (*int64, bool)`

GetInitialDrivingTimeOk returns a tuple with the InitialDrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDrivingTime

`func (o *VehicleBreak) SetInitialDrivingTime(v int64)`

SetInitialDrivingTime sets InitialDrivingTime field to given value.

### HasInitialDrivingTime

`func (o *VehicleBreak) HasInitialDrivingTime() bool`

HasInitialDrivingTime returns a boolean if a field has been set.

### GetPossibleSplit

`func (o *VehicleBreak) GetPossibleSplit() []int64`

GetPossibleSplit returns the PossibleSplit field if non-nil, zero value otherwise.

### GetPossibleSplitOk

`func (o *VehicleBreak) GetPossibleSplitOk() (*[]int64, bool)`

GetPossibleSplitOk returns a tuple with the PossibleSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleSplit

`func (o *VehicleBreak) SetPossibleSplit(v []int64)`

SetPossibleSplit sets PossibleSplit field to given value.

### HasPossibleSplit

`func (o *VehicleBreak) HasPossibleSplit() bool`

HasPossibleSplit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


