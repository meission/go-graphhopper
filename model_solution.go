/*
GraphHopper Directions API

 Integrate A-to-B route planning, turn-by-turn navigation, route optimization, isochrone calculations, location clustering and other tools into your application.    ##### Authentication      1. [Sign up for GraphHopper](https://graphhopper.com/dashboard/#/signup)   2. [Create an API key](https://support.graphhopper.com/a/solutions/articles/44001976027)    Authenticate to the API by passing your key as a query parameter in every request.    ##### API Explorer    You can also try all API parts without registration in our [API explorer](https://explorer.graphhopper.com/).    ##### Client Libraries    To speed up development and make coding easier, we offer a [JavaScript client](https://github.com/graphhopper/directions-api-js-client) and a   [Java client](https://github.com/graphhopper/graphhopper/tree/master/client-hc).    ##### Optimize Response Speed    1. Reuse SSL/TLS sessions    You should utilize the SSL session to speed up responses after the initial response or use a library that does this. E.g. for Java the   [OkHttp library](https://square.github.io/okhttp/) automatically reuses SSL/TLS sessions and also the browser takes care of this automatically.   For python you can use the [`requests` library](https://requests.readthedocs.io/en/latest/user/advanced/): first you create a   session (`session = requests.Session()`) and then do requests only with this session instead of directly using \"requests\".    2. Bandwidth reduction    If you create your own client, make sure it supports http/2 and gzipped responses for best speed.   If you use the Matrix, the Route Optimization API or the  and want to solve large problems, we recommend you to reduce bandwidth   by [compressing your POST request](https://gist.github.com/karussell/82851e303ea7b3459b2dea01f18949f4) and specifying the header   as follows: `Content-Encoding: gzip`. This will also avoid the HTTP 413 error \"Request Entity Too Large\". 

API version: 1.0.0
Contact: support@graphhopper.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Solution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Solution{}

// Solution Only available if status field indicates `finished`.
type Solution struct {
	// Deprecated
	Costs *int32 `json:"costs,omitempty"`
	// Overall distance travelled in meter, i.e. the sum of each route's transport distance
	Distance *int32 `json:"distance,omitempty"`
	// Use `transport_time` instead.
	// Deprecated
	Time *int64 `json:"time,omitempty"`
	// Overall time travelled in seconds, i.e. the sum of each route's transport time.
	TransportTime *int64 `json:"transport_time,omitempty"`
	// Operation time of longest route in seconds.
	MaxOperationTime *int64 `json:"max_operation_time,omitempty"`
	// Overall waiting time in seconds.
	WaitingTime *int64 `json:"waiting_time,omitempty"`
	// Overall service time in seconds.
	ServiceDuration *int64 `json:"service_duration,omitempty"`
	// Overall preparation time in seconds.
	PreparationTime *int64 `json:"preparation_time,omitempty"`
	// Overall completion time in seconds, i.e. the sum of each routes/drivers operation time.
	CompletionTime *int64 `json:"completion_time,omitempty"`
	// Number of employed vehicles.
	NoVehicles *int32 `json:"no_vehicles,omitempty"`
	// Number of jobs that could not be assigned to final solution.
	NoUnassigned *int32 `json:"no_unassigned,omitempty"`
	// An array of routes
	Routes []Route `json:"routes,omitempty"`
	Unassigned *SolutionUnassigned `json:"unassigned,omitempty"`
}

// NewSolution instantiates a new Solution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolution() *Solution {
	this := Solution{}
	return &this
}

// NewSolutionWithDefaults instantiates a new Solution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolutionWithDefaults() *Solution {
	this := Solution{}
	return &this
}

// GetCosts returns the Costs field value if set, zero value otherwise.
// Deprecated
func (o *Solution) GetCosts() int32 {
	if o == nil || IsNil(o.Costs) {
		var ret int32
		return ret
	}
	return *o.Costs
}

// GetCostsOk returns a tuple with the Costs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Solution) GetCostsOk() (*int32, bool) {
	if o == nil || IsNil(o.Costs) {
		return nil, false
	}
	return o.Costs, true
}

// HasCosts returns a boolean if a field has been set.
func (o *Solution) HasCosts() bool {
	if o != nil && !IsNil(o.Costs) {
		return true
	}

	return false
}

// SetCosts gets a reference to the given int32 and assigns it to the Costs field.
// Deprecated
func (o *Solution) SetCosts(v int32) {
	o.Costs = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *Solution) GetDistance() int32 {
	if o == nil || IsNil(o.Distance) {
		var ret int32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetDistanceOk() (*int32, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *Solution) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int32 and assigns it to the Distance field.
func (o *Solution) SetDistance(v int32) {
	o.Distance = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
// Deprecated
func (o *Solution) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Solution) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Solution) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
// Deprecated
func (o *Solution) SetTime(v int64) {
	o.Time = &v
}

// GetTransportTime returns the TransportTime field value if set, zero value otherwise.
func (o *Solution) GetTransportTime() int64 {
	if o == nil || IsNil(o.TransportTime) {
		var ret int64
		return ret
	}
	return *o.TransportTime
}

// GetTransportTimeOk returns a tuple with the TransportTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetTransportTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TransportTime) {
		return nil, false
	}
	return o.TransportTime, true
}

// HasTransportTime returns a boolean if a field has been set.
func (o *Solution) HasTransportTime() bool {
	if o != nil && !IsNil(o.TransportTime) {
		return true
	}

	return false
}

// SetTransportTime gets a reference to the given int64 and assigns it to the TransportTime field.
func (o *Solution) SetTransportTime(v int64) {
	o.TransportTime = &v
}

// GetMaxOperationTime returns the MaxOperationTime field value if set, zero value otherwise.
func (o *Solution) GetMaxOperationTime() int64 {
	if o == nil || IsNil(o.MaxOperationTime) {
		var ret int64
		return ret
	}
	return *o.MaxOperationTime
}

// GetMaxOperationTimeOk returns a tuple with the MaxOperationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetMaxOperationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxOperationTime) {
		return nil, false
	}
	return o.MaxOperationTime, true
}

// HasMaxOperationTime returns a boolean if a field has been set.
func (o *Solution) HasMaxOperationTime() bool {
	if o != nil && !IsNil(o.MaxOperationTime) {
		return true
	}

	return false
}

// SetMaxOperationTime gets a reference to the given int64 and assigns it to the MaxOperationTime field.
func (o *Solution) SetMaxOperationTime(v int64) {
	o.MaxOperationTime = &v
}

// GetWaitingTime returns the WaitingTime field value if set, zero value otherwise.
func (o *Solution) GetWaitingTime() int64 {
	if o == nil || IsNil(o.WaitingTime) {
		var ret int64
		return ret
	}
	return *o.WaitingTime
}

// GetWaitingTimeOk returns a tuple with the WaitingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetWaitingTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.WaitingTime) {
		return nil, false
	}
	return o.WaitingTime, true
}

// HasWaitingTime returns a boolean if a field has been set.
func (o *Solution) HasWaitingTime() bool {
	if o != nil && !IsNil(o.WaitingTime) {
		return true
	}

	return false
}

// SetWaitingTime gets a reference to the given int64 and assigns it to the WaitingTime field.
func (o *Solution) SetWaitingTime(v int64) {
	o.WaitingTime = &v
}

// GetServiceDuration returns the ServiceDuration field value if set, zero value otherwise.
func (o *Solution) GetServiceDuration() int64 {
	if o == nil || IsNil(o.ServiceDuration) {
		var ret int64
		return ret
	}
	return *o.ServiceDuration
}

// GetServiceDurationOk returns a tuple with the ServiceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetServiceDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.ServiceDuration) {
		return nil, false
	}
	return o.ServiceDuration, true
}

// HasServiceDuration returns a boolean if a field has been set.
func (o *Solution) HasServiceDuration() bool {
	if o != nil && !IsNil(o.ServiceDuration) {
		return true
	}

	return false
}

// SetServiceDuration gets a reference to the given int64 and assigns it to the ServiceDuration field.
func (o *Solution) SetServiceDuration(v int64) {
	o.ServiceDuration = &v
}

// GetPreparationTime returns the PreparationTime field value if set, zero value otherwise.
func (o *Solution) GetPreparationTime() int64 {
	if o == nil || IsNil(o.PreparationTime) {
		var ret int64
		return ret
	}
	return *o.PreparationTime
}

// GetPreparationTimeOk returns a tuple with the PreparationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetPreparationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PreparationTime) {
		return nil, false
	}
	return o.PreparationTime, true
}

// HasPreparationTime returns a boolean if a field has been set.
func (o *Solution) HasPreparationTime() bool {
	if o != nil && !IsNil(o.PreparationTime) {
		return true
	}

	return false
}

// SetPreparationTime gets a reference to the given int64 and assigns it to the PreparationTime field.
func (o *Solution) SetPreparationTime(v int64) {
	o.PreparationTime = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *Solution) GetCompletionTime() int64 {
	if o == nil || IsNil(o.CompletionTime) {
		var ret int64
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetCompletionTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CompletionTime) {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *Solution) HasCompletionTime() bool {
	if o != nil && !IsNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given int64 and assigns it to the CompletionTime field.
func (o *Solution) SetCompletionTime(v int64) {
	o.CompletionTime = &v
}

// GetNoVehicles returns the NoVehicles field value if set, zero value otherwise.
func (o *Solution) GetNoVehicles() int32 {
	if o == nil || IsNil(o.NoVehicles) {
		var ret int32
		return ret
	}
	return *o.NoVehicles
}

// GetNoVehiclesOk returns a tuple with the NoVehicles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetNoVehiclesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoVehicles) {
		return nil, false
	}
	return o.NoVehicles, true
}

// HasNoVehicles returns a boolean if a field has been set.
func (o *Solution) HasNoVehicles() bool {
	if o != nil && !IsNil(o.NoVehicles) {
		return true
	}

	return false
}

// SetNoVehicles gets a reference to the given int32 and assigns it to the NoVehicles field.
func (o *Solution) SetNoVehicles(v int32) {
	o.NoVehicles = &v
}

// GetNoUnassigned returns the NoUnassigned field value if set, zero value otherwise.
func (o *Solution) GetNoUnassigned() int32 {
	if o == nil || IsNil(o.NoUnassigned) {
		var ret int32
		return ret
	}
	return *o.NoUnassigned
}

// GetNoUnassignedOk returns a tuple with the NoUnassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetNoUnassignedOk() (*int32, bool) {
	if o == nil || IsNil(o.NoUnassigned) {
		return nil, false
	}
	return o.NoUnassigned, true
}

// HasNoUnassigned returns a boolean if a field has been set.
func (o *Solution) HasNoUnassigned() bool {
	if o != nil && !IsNil(o.NoUnassigned) {
		return true
	}

	return false
}

// SetNoUnassigned gets a reference to the given int32 and assigns it to the NoUnassigned field.
func (o *Solution) SetNoUnassigned(v int32) {
	o.NoUnassigned = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *Solution) GetRoutes() []Route {
	if o == nil || IsNil(o.Routes) {
		var ret []Route
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetRoutesOk() ([]Route, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *Solution) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *Solution) SetRoutes(v []Route) {
	o.Routes = v
}

// GetUnassigned returns the Unassigned field value if set, zero value otherwise.
func (o *Solution) GetUnassigned() SolutionUnassigned {
	if o == nil || IsNil(o.Unassigned) {
		var ret SolutionUnassigned
		return ret
	}
	return *o.Unassigned
}

// GetUnassignedOk returns a tuple with the Unassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetUnassignedOk() (*SolutionUnassigned, bool) {
	if o == nil || IsNil(o.Unassigned) {
		return nil, false
	}
	return o.Unassigned, true
}

// HasUnassigned returns a boolean if a field has been set.
func (o *Solution) HasUnassigned() bool {
	if o != nil && !IsNil(o.Unassigned) {
		return true
	}

	return false
}

// SetUnassigned gets a reference to the given SolutionUnassigned and assigns it to the Unassigned field.
func (o *Solution) SetUnassigned(v SolutionUnassigned) {
	o.Unassigned = &v
}

func (o Solution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Solution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Costs) {
		toSerialize["costs"] = o.Costs
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TransportTime) {
		toSerialize["transport_time"] = o.TransportTime
	}
	if !IsNil(o.MaxOperationTime) {
		toSerialize["max_operation_time"] = o.MaxOperationTime
	}
	if !IsNil(o.WaitingTime) {
		toSerialize["waiting_time"] = o.WaitingTime
	}
	if !IsNil(o.ServiceDuration) {
		toSerialize["service_duration"] = o.ServiceDuration
	}
	if !IsNil(o.PreparationTime) {
		toSerialize["preparation_time"] = o.PreparationTime
	}
	if !IsNil(o.CompletionTime) {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if !IsNil(o.NoVehicles) {
		toSerialize["no_vehicles"] = o.NoVehicles
	}
	if !IsNil(o.NoUnassigned) {
		toSerialize["no_unassigned"] = o.NoUnassigned
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	if !IsNil(o.Unassigned) {
		toSerialize["unassigned"] = o.Unassigned
	}
	return toSerialize, nil
}

type NullableSolution struct {
	value *Solution
	isSet bool
}

func (v NullableSolution) Get() *Solution {
	return v.value
}

func (v *NullableSolution) Set(val *Solution) {
	v.value = val
	v.isSet = true
}

func (v NullableSolution) IsSet() bool {
	return v.isSet
}

func (v *NullableSolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolution(val *Solution) *NullableSolution {
	return &NullableSolution{value: val, isSet: true}
}

func (v NullableSolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


