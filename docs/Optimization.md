# Optimization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeInsertion** | Pointer to **bool** | If you use groups, you sometimes want to place orders without group assignment in the best position, i.e. sometimes in the middle of a group and not before or after the group. This is not allowed by default. However, if this field here is \&quot;true\&quot;, these orders (without a group assignment) can be inserted freely. | [optional] [default to false]

## Methods

### NewOptimization

`func NewOptimization() *Optimization`

NewOptimization instantiates a new Optimization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationWithDefaults

`func NewOptimizationWithDefaults() *Optimization`

NewOptimizationWithDefaults instantiates a new Optimization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeInsertion

`func (o *Optimization) GetFreeInsertion() bool`

GetFreeInsertion returns the FreeInsertion field if non-nil, zero value otherwise.

### GetFreeInsertionOk

`func (o *Optimization) GetFreeInsertionOk() (*bool, bool)`

GetFreeInsertionOk returns a tuple with the FreeInsertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeInsertion

`func (o *Optimization) SetFreeInsertion(v bool)`

SetFreeInsertion sets FreeInsertion field to given value.

### HasFreeInsertion

`func (o *Optimization) HasFreeInsertion() bool`

HasFreeInsertion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


