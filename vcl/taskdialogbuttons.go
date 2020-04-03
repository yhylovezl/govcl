
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TTaskDialogButtons struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsTaskDialogButtons(obj interface{}) *TTaskDialogButtons {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TTaskDialogButtons{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsTaskDialogButtons.
func TaskDialogButtonsFromInst(inst uintptr) *TTaskDialogButtons {
    return AsTaskDialogButtons(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsTaskDialogButtons.
func TaskDialogButtonsFromObj(obj IObject) *TTaskDialogButtons {
    return AsTaskDialogButtons(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTaskDialogButtons.
func TaskDialogButtonsFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialogButtons {
    return AsTaskDialogButtons(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskDialogButtons) Instance() uintptr {
    return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskDialogButtons) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskDialogButtons) IsValid() bool {
    return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TTaskDialogButtons) Is() TIs {
    return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TTaskDialogButtons) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskDialogButtonsClass() TClass {
    return TaskDialogButtons_StaticClassType()
}

func (t *TTaskDialogButtons) Add() *TTaskDialogBaseButtonItem {
    return AsTaskDialogBaseButtonItem(TaskDialogButtons_Add(t.instance))
}

func (t *TTaskDialogButtons) FindButton(AModalResult TModalResult) *TTaskDialogBaseButtonItem {
    return AsTaskDialogBaseButtonItem(TaskDialogButtons_FindButton(t.instance, AModalResult))
}

// CN: 组件所有者。
// EN: component owner.
func (t *TTaskDialogButtons) Owner() *TObject {
    return AsObject(TaskDialogButtons_Owner(t.instance))
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskDialogButtons) Assign(Source IObject) {
    TaskDialogButtons_Assign(t.instance, CheckPtr(Source))
}

func (t *TTaskDialogButtons) BeginUpdate() {
    TaskDialogButtons_BeginUpdate(t.instance)
}

// CN: 清除。
// EN: .
func (t *TTaskDialogButtons) Clear() {
    TaskDialogButtons_Clear(t.instance)
}

func (t *TTaskDialogButtons) Delete(Index int32) {
    TaskDialogButtons_Delete(t.instance, Index)
}

func (t *TTaskDialogButtons) EndUpdate() {
    TaskDialogButtons_EndUpdate(t.instance)
}

func (t *TTaskDialogButtons) FindItemID(ID int32) *TCollectionItem {
    return AsCollectionItem(TaskDialogButtons_FindItemID(t.instance, ID))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskDialogButtons) GetNamePath() string {
    return TaskDialogButtons_GetNamePath(t.instance)
}

func (t *TTaskDialogButtons) Insert(Index int32) *TCollectionItem {
    return AsCollectionItem(TaskDialogButtons_Insert(t.instance, Index))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskDialogButtons) ClassType() TClass {
    return TaskDialogButtons_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskDialogButtons) ClassName() string {
    return TaskDialogButtons_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskDialogButtons) InstanceSize() int32 {
    return TaskDialogButtons_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskDialogButtons) InheritsFrom(AClass TClass) bool {
    return TaskDialogButtons_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskDialogButtons) Equals(Obj IObject) bool {
    return TaskDialogButtons_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskDialogButtons) GetHashCode() int32 {
    return TaskDialogButtons_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TTaskDialogButtons) ToString() string {
    return TaskDialogButtons_ToString(t.instance)
}

func (t *TTaskDialogButtons) DefaultButton() *TTaskDialogBaseButtonItem {
    return AsTaskDialogBaseButtonItem(TaskDialogButtons_GetDefaultButton(t.instance))
}

func (t *TTaskDialogButtons) SetDefaultButton(value *TTaskDialogBaseButtonItem) {
    TaskDialogButtons_SetDefaultButton(t.instance, CheckPtr(value))
}

func (t *TTaskDialogButtons) Count() int32 {
    return TaskDialogButtons_GetCount(t.instance)
}

func (t *TTaskDialogButtons) Items(Index int32) *TTaskDialogBaseButtonItem {
    return AsTaskDialogBaseButtonItem(TaskDialogButtons_GetItems(t.instance, Index))
}

func (t *TTaskDialogButtons) SetItems(Index int32, value *TTaskDialogBaseButtonItem) {
    TaskDialogButtons_SetItems(t.instance, Index, CheckPtr(value))
}

