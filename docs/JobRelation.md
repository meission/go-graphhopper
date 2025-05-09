# JobRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Specifies the type of relation. It must be either of type &#x60;in_same_route&#x60;, &#x60;not_in_same_route&#x60;, &#x60;in_sequence&#x60;, &#x60;in_direct_sequence&#x60; or &#x60;neighbor&#x60;.  &#x60;in_same_route&#x60;: As the name suggest, it enforces the specified services or shipments to be in the same route. It can be specified as follows:  &#x60;&#x60;&#x60;json {    \&quot;type\&quot;: \&quot;in_same_route\&quot;,    \&quot;ids\&quot;: [\&quot;serv_i_id\&quot;,\&quot;serv_j_id\&quot;] } &#x60;&#x60;&#x60;  This enforces service i to be in the same route as service j no matter which vehicle will be employed. If a specific vehicle (driver) is required to conduct this, just add a &#x60;vehicle_id&#x60; like this:  &#x60;&#x60;&#x60; {    \&quot;type\&quot;: \&quot;in_same_route\&quot;,    \&quot;ids\&quot;: [\&quot;serv_i_id\&quot;,\&quot;serv_j_id\&quot;],    \&quot;vehicle_id\&quot;: \&quot;vehicle1\&quot; } &#x60;&#x60;&#x60;  This not only enforce service i and j to be in the same route, but also makes sure that both services are in the route of &#x60;vehicle1&#x60;.  *Tip*: This way initial loads and vehicle routes can be modelled. For example, if your vehicles are already on the road and new orders come in, then vehicles can still be rescheduled subject to the orders that have already been assigned to these vehicles.  &#x60;not_in_same_route&#x60;: It ensures that 2 or more orders are not transported by the same vehicle. It can be specified as follows:  &#x60;&#x60;&#x60;json {    \&quot;type\&quot;: \&quot;not_in_same_route\&quot;,    \&quot;ids\&quot;: [\&quot;serv_i_id\&quot;,\&quot;serv_j_id\&quot;] } &#x60;&#x60;&#x60;   &#x60;in_sequence&#x60;: This relation type enforces n jobs to be in sequence. It can be specified as  &#x60;&#x60;&#x60;json {    \&quot;type\&quot;: \&quot;in_sequence\&quot;,    \&quot;ids\&quot;: [\&quot;serv_i_id\&quot;,\&quot;serv_j_id\&quot;] } &#x60;&#x60;&#x60;  which means that service j need to be in the same route as service i AND it needs to occur somewhere after service i. As described above if a specific vehicle needs to conduct this, just add &#x60;vehicle_id&#x60;.   &#x60;in_direct_sequence&#x60;: This enforces n services or shipments to be in direct sequence. It can be specified as  &#x60;&#x60;&#x60;json {    \&quot;type\&quot;: \&quot;in_direct_sequence\&quot;,    \&quot;ids\&quot;: [\&quot;serv_i_id\&quot;,\&quot;serv_j_id\&quot;,\&quot;serv_k_id\&quot;] } &#x60;&#x60;&#x60;  yielding service j to occur directly after service i, and service k to occur directly after service j i.e. in strong order. Again, a vehicle can be assigned a priority by adding a &#x60;vehicle_id&#x60; to the relation.  &#x60;neighbor&#x60;: This specifies a neighbor relationship, i.e., if services i and j are to be neighbors, i must be either immediately before or after j. I can be specified as follows:  &#x60;&#x60;&#x60;json {     \&quot;type\&quot;: \&quot;neighbor\&quot;,     \&quot;ids\&quot;: [\&quot;serv_i_id\&quot;,\&quot;serv_j_id\&quot;] } &#x60;&#x60;&#x60;  *Special IDs*: If you look at the previous example and you want service i to be the first in the route, use the special ID &#x60;start&#x60; as follows:  &#x60;&#x60;&#x60;json {    \&quot;type\&quot;: \&quot;in_direct_sequence\&quot;,    \&quot;ids\&quot;: [\&quot;start\&quot;,\&quot;serv_i_id\&quot;,\&quot;serv_j_id\&quot;,\&quot;serv_k_id\&quot;] } &#x60;&#x60;&#x60;  Latter enforces the direct sequence of i, j and k at the beginning of the route. If this sequence should be bound to the end of the route, use the special ID &#x60;end&#x60; like this:  &#x60;&#x60;&#x60;json {    \&quot;type\&quot;: \&quot;in_direct_sequence\&quot;,    \&quot;ids\&quot;: [\&quot;serv_i_id\&quot;,\&quot;service_j_id\&quot;,\&quot;serv_k_id\&quot;,\&quot;end\&quot;] } &#x60;&#x60;&#x60;  If you deal with services then you need to use the &#39;id&#39; of your services in the field &#39;ids&#39;. To also consider sequences of the pickups and deliveries of your shipments, you need to use a special ID, i.e. use the shipment id plus the keyword &#x60;_pickup&#x60; or &#x60;_delivery&#x60;. For example, to ensure that the pickup and delivery of the shipment with the id &#39;my_shipment&#39; are direct neighbors, you need the following specification:  &#x60;&#x60;&#x60; {    \&quot;type\&quot;: \&quot;in_direct_sequence\&quot;,    \&quot;ids\&quot;: [\&quot;my_ship_pickup\&quot;,\&quot;my_ship_delivery\&quot;] } &#x60;&#x60;&#x60;   | 
**Ids** | **[]string** | Specifies an array of shipment and/or service ids that are in relation. If you deal with services then you need to use the id of your services in ids. To also consider sequences of the pickups and deliveries of your shipments, you need to use a special ID, i.e. use your shipment id plus the keyword &#x60;_pickup&#x60; or &#x60;_delivery&#x60;. If you want to place a service or shipment activity at the beginning of your route, use the special ID &#x60;start&#x60;. In turn, use &#x60;end&#x60; to place it at the end of the route. | 
**VehicleId** | Pointer to **string** | Id of pre-assigned vehicle, i.e. the vehicle id that is determined to conduct the services and shipments in this relation. | [optional] 

## Methods

### NewJobRelation

`func NewJobRelation(type_ string, ids []string, ) *JobRelation`

NewJobRelation instantiates a new JobRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRelationWithDefaults

`func NewJobRelationWithDefaults() *JobRelation`

NewJobRelationWithDefaults instantiates a new JobRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JobRelation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobRelation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobRelation) SetType(v string)`

SetType sets Type field to given value.


### GetIds

`func (o *JobRelation) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *JobRelation) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *JobRelation) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetVehicleId

`func (o *JobRelation) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *JobRelation) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *JobRelation) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *JobRelation) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


