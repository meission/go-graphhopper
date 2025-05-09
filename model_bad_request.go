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

// checks if the BadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequest{}

// BadRequest struct for BadRequest
type BadRequest struct {
	// Short error message
	Message *string `json:"message,omitempty"`
	// Optional error information.
	Hints []ErrorMessage `json:"hints,omitempty"`
	// status
	Status *string `json:"status,omitempty"`
}

// NewBadRequest instantiates a new BadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequest() *BadRequest {
	this := BadRequest{}
	var status string = "finished"
	this.Status = &status
	return &this
}

// NewBadRequestWithDefaults instantiates a new BadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestWithDefaults() *BadRequest {
	this := BadRequest{}
	var status string = "finished"
	this.Status = &status
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BadRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BadRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BadRequest) SetMessage(v string) {
	o.Message = &v
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *BadRequest) GetHints() []ErrorMessage {
	if o == nil || IsNil(o.Hints) {
		var ret []ErrorMessage
		return ret
	}
	return o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequest) GetHintsOk() ([]ErrorMessage, bool) {
	if o == nil || IsNil(o.Hints) {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *BadRequest) HasHints() bool {
	if o != nil && !IsNil(o.Hints) {
		return true
	}

	return false
}

// SetHints gets a reference to the given []ErrorMessage and assigns it to the Hints field.
func (o *BadRequest) SetHints(v []ErrorMessage) {
	o.Hints = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BadRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BadRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BadRequest) SetStatus(v string) {
	o.Status = &v
}

func (o BadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Hints) {
		toSerialize["hints"] = o.Hints
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableBadRequest struct {
	value *BadRequest
	isSet bool
}

func (v NullableBadRequest) Get() *BadRequest {
	return v.value
}

func (v *NullableBadRequest) Set(val *BadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequest(val *BadRequest) *NullableBadRequest {
	return &NullableBadRequest{value: val, isSet: true}
}

func (v NullableBadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


