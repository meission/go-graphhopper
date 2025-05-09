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
	"bytes"
	"fmt"
)

// checks if the JobRelation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobRelation{}

// JobRelation struct for JobRelation
type JobRelation struct {
	// Specifies the type of relation. It must be either of type `in_same_route`, `not_in_same_route`, `in_sequence`, `in_direct_sequence` or `neighbor`.  `in_same_route`: As the name suggest, it enforces the specified services or shipments to be in the same route. It can be specified as follows:  ```json {    \"type\": \"in_same_route\",    \"ids\": [\"serv_i_id\",\"serv_j_id\"] } ```  This enforces service i to be in the same route as service j no matter which vehicle will be employed. If a specific vehicle (driver) is required to conduct this, just add a `vehicle_id` like this:  ``` {    \"type\": \"in_same_route\",    \"ids\": [\"serv_i_id\",\"serv_j_id\"],    \"vehicle_id\": \"vehicle1\" } ```  This not only enforce service i and j to be in the same route, but also makes sure that both services are in the route of `vehicle1`.  *Tip*: This way initial loads and vehicle routes can be modelled. For example, if your vehicles are already on the road and new orders come in, then vehicles can still be rescheduled subject to the orders that have already been assigned to these vehicles.  `not_in_same_route`: It ensures that 2 or more orders are not transported by the same vehicle. It can be specified as follows:  ```json {    \"type\": \"not_in_same_route\",    \"ids\": [\"serv_i_id\",\"serv_j_id\"] } ```   `in_sequence`: This relation type enforces n jobs to be in sequence. It can be specified as  ```json {    \"type\": \"in_sequence\",    \"ids\": [\"serv_i_id\",\"serv_j_id\"] } ```  which means that service j need to be in the same route as service i AND it needs to occur somewhere after service i. As described above if a specific vehicle needs to conduct this, just add `vehicle_id`.   `in_direct_sequence`: This enforces n services or shipments to be in direct sequence. It can be specified as  ```json {    \"type\": \"in_direct_sequence\",    \"ids\": [\"serv_i_id\",\"serv_j_id\",\"serv_k_id\"] } ```  yielding service j to occur directly after service i, and service k to occur directly after service j i.e. in strong order. Again, a vehicle can be assigned a priority by adding a `vehicle_id` to the relation.  `neighbor`: This specifies a neighbor relationship, i.e., if services i and j are to be neighbors, i must be either immediately before or after j. I can be specified as follows:  ```json {     \"type\": \"neighbor\",     \"ids\": [\"serv_i_id\",\"serv_j_id\"] } ```  *Special IDs*: If you look at the previous example and you want service i to be the first in the route, use the special ID `start` as follows:  ```json {    \"type\": \"in_direct_sequence\",    \"ids\": [\"start\",\"serv_i_id\",\"serv_j_id\",\"serv_k_id\"] } ```  Latter enforces the direct sequence of i, j and k at the beginning of the route. If this sequence should be bound to the end of the route, use the special ID `end` like this:  ```json {    \"type\": \"in_direct_sequence\",    \"ids\": [\"serv_i_id\",\"service_j_id\",\"serv_k_id\",\"end\"] } ```  If you deal with services then you need to use the 'id' of your services in the field 'ids'. To also consider sequences of the pickups and deliveries of your shipments, you need to use a special ID, i.e. use the shipment id plus the keyword `_pickup` or `_delivery`. For example, to ensure that the pickup and delivery of the shipment with the id 'my_shipment' are direct neighbors, you need the following specification:  ``` {    \"type\": \"in_direct_sequence\",    \"ids\": [\"my_ship_pickup\",\"my_ship_delivery\"] } ```  
	Type string `json:"type"`
	// Specifies an array of shipment and/or service ids that are in relation. If you deal with services then you need to use the id of your services in ids. To also consider sequences of the pickups and deliveries of your shipments, you need to use a special ID, i.e. use your shipment id plus the keyword `_pickup` or `_delivery`. If you want to place a service or shipment activity at the beginning of your route, use the special ID `start`. In turn, use `end` to place it at the end of the route.
	Ids []string `json:"ids"`
	// Id of pre-assigned vehicle, i.e. the vehicle id that is determined to conduct the services and shipments in this relation.
	VehicleId *string `json:"vehicle_id,omitempty"`
}

type _JobRelation JobRelation

// NewJobRelation instantiates a new JobRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRelation(type_ string, ids []string) *JobRelation {
	this := JobRelation{}
	this.Type = type_
	this.Ids = ids
	return &this
}

// NewJobRelationWithDefaults instantiates a new JobRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobRelationWithDefaults() *JobRelation {
	this := JobRelation{}
	return &this
}

// GetType returns the Type field value
func (o *JobRelation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobRelation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobRelation) SetType(v string) {
	o.Type = v
}

// GetIds returns the Ids field value
func (o *JobRelation) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *JobRelation) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *JobRelation) SetIds(v []string) {
	o.Ids = v
}

// GetVehicleId returns the VehicleId field value if set, zero value otherwise.
func (o *JobRelation) GetVehicleId() string {
	if o == nil || IsNil(o.VehicleId) {
		var ret string
		return ret
	}
	return *o.VehicleId
}

// GetVehicleIdOk returns a tuple with the VehicleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRelation) GetVehicleIdOk() (*string, bool) {
	if o == nil || IsNil(o.VehicleId) {
		return nil, false
	}
	return o.VehicleId, true
}

// HasVehicleId returns a boolean if a field has been set.
func (o *JobRelation) HasVehicleId() bool {
	if o != nil && !IsNil(o.VehicleId) {
		return true
	}

	return false
}

// SetVehicleId gets a reference to the given string and assigns it to the VehicleId field.
func (o *JobRelation) SetVehicleId(v string) {
	o.VehicleId = &v
}

func (o JobRelation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobRelation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["ids"] = o.Ids
	if !IsNil(o.VehicleId) {
		toSerialize["vehicle_id"] = o.VehicleId
	}
	return toSerialize, nil
}

func (o *JobRelation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"ids",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varJobRelation := _JobRelation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJobRelation)

	if err != nil {
		return err
	}

	*o = JobRelation(varJobRelation)

	return err
}

type NullableJobRelation struct {
	value *JobRelation
	isSet bool
}

func (v NullableJobRelation) Get() *JobRelation {
	return v.value
}

func (v *NullableJobRelation) Set(val *JobRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRelation(val *JobRelation) *NullableJobRelation {
	return &NullableJobRelation{value: val, isSet: true}
}

func (v NullableJobRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


