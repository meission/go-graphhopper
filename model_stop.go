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

// checks if the Stop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stop{}

// Stop struct for Stop
type Stop struct {
	// Specifies pickup or delivery address.
	Address *Address `json:"address,omitempty"`
	// Specifies the duration of the pickup or delivery in seconds, e.g. how long it takes unload items at the customer site.
	Duration *int64 `json:"duration,omitempty"`
	// Specifies the preparation time in seconds. It can be used to model parking lot search time since if you have 3 identical locations in a row, it only falls due once.
	PreparationTime *int64 `json:"preparation_time,omitempty"`
	// Specifies an array of time window objects (see time window object below). For example, if an item needs to be delivered between 7am and 10am then specify the array as follows: [ { \"earliest\": 25200, \"latest\" : 32400 } ] (starting the day from 0 in seconds).
	TimeWindows []TimeWindow `json:"time_windows,omitempty"`
	// Group this stop belongs to. See the group relation and [this post](https://discuss.graphhopper.com/t/4040) on how to utilize this.
	Group *string `json:"group,omitempty"`
}

// NewStop instantiates a new Stop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStop() *Stop {
	this := Stop{}
	var duration int64 = 0
	this.Duration = &duration
	var preparationTime int64 = 0
	this.PreparationTime = &preparationTime
	return &this
}

// NewStopWithDefaults instantiates a new Stop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopWithDefaults() *Stop {
	this := Stop{}
	var duration int64 = 0
	this.Duration = &duration
	var preparationTime int64 = 0
	this.PreparationTime = &preparationTime
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Stop) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stop) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Stop) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Stop) SetAddress(v Address) {
	o.Address = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Stop) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stop) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Stop) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *Stop) SetDuration(v int64) {
	o.Duration = &v
}

// GetPreparationTime returns the PreparationTime field value if set, zero value otherwise.
func (o *Stop) GetPreparationTime() int64 {
	if o == nil || IsNil(o.PreparationTime) {
		var ret int64
		return ret
	}
	return *o.PreparationTime
}

// GetPreparationTimeOk returns a tuple with the PreparationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stop) GetPreparationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PreparationTime) {
		return nil, false
	}
	return o.PreparationTime, true
}

// HasPreparationTime returns a boolean if a field has been set.
func (o *Stop) HasPreparationTime() bool {
	if o != nil && !IsNil(o.PreparationTime) {
		return true
	}

	return false
}

// SetPreparationTime gets a reference to the given int64 and assigns it to the PreparationTime field.
func (o *Stop) SetPreparationTime(v int64) {
	o.PreparationTime = &v
}

// GetTimeWindows returns the TimeWindows field value if set, zero value otherwise.
func (o *Stop) GetTimeWindows() []TimeWindow {
	if o == nil || IsNil(o.TimeWindows) {
		var ret []TimeWindow
		return ret
	}
	return o.TimeWindows
}

// GetTimeWindowsOk returns a tuple with the TimeWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stop) GetTimeWindowsOk() ([]TimeWindow, bool) {
	if o == nil || IsNil(o.TimeWindows) {
		return nil, false
	}
	return o.TimeWindows, true
}

// HasTimeWindows returns a boolean if a field has been set.
func (o *Stop) HasTimeWindows() bool {
	if o != nil && !IsNil(o.TimeWindows) {
		return true
	}

	return false
}

// SetTimeWindows gets a reference to the given []TimeWindow and assigns it to the TimeWindows field.
func (o *Stop) SetTimeWindows(v []TimeWindow) {
	o.TimeWindows = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *Stop) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stop) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Stop) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *Stop) SetGroup(v string) {
	o.Group = &v
}

func (o Stop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Stop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.PreparationTime) {
		toSerialize["preparation_time"] = o.PreparationTime
	}
	if !IsNil(o.TimeWindows) {
		toSerialize["time_windows"] = o.TimeWindows
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableStop struct {
	value *Stop
	isSet bool
}

func (v NullableStop) Get() *Stop {
	return v.value
}

func (v *NullableStop) Set(val *Stop) {
	v.value = val
	v.isSet = true
}

func (v NullableStop) IsSet() bool {
	return v.isSet
}

func (v *NullableStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStop(val *Stop) *NullableStop {
	return &NullableStop{value: val, isSet: true}
}

func (v NullableStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


