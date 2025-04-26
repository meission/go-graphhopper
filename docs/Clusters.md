# Clusters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | id of customer | [optional] 
**Center** | Pointer to [**ClusterCustomerAddress**](ClusterCustomerAddress.md) |  | [optional] 
**MinQuantity** | Pointer to **float32** | Specifies min. quantity of this cluster | [optional] 
**MaxQuantity** | Pointer to **float32** | Specifies max. quantity of this cluster | [optional] 

## Methods

### NewClusters

`func NewClusters() *Clusters`

NewClusters instantiates a new Clusters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClustersWithDefaults

`func NewClustersWithDefaults() *Clusters`

NewClustersWithDefaults instantiates a new Clusters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Clusters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Clusters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Clusters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Clusters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCenter

`func (o *Clusters) GetCenter() ClusterCustomerAddress`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *Clusters) GetCenterOk() (*ClusterCustomerAddress, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *Clusters) SetCenter(v ClusterCustomerAddress)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *Clusters) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetMinQuantity

`func (o *Clusters) GetMinQuantity() float32`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *Clusters) GetMinQuantityOk() (*float32, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *Clusters) SetMinQuantity(v float32)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *Clusters) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *Clusters) GetMaxQuantity() float32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *Clusters) GetMaxQuantityOk() (*float32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *Clusters) SetMaxQuantity(v float32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *Clusters) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


