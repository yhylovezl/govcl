
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

type TTreeNode struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTreeNode(AOwner *TTreeNodes) *TTreeNode {
    t := new(TTreeNode)
    t.instance = TreeNode_Create(CheckPtr(AOwner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsTreeNode(obj interface{}) *TTreeNode {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TTreeNode{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsTreeNode.
func TreeNodeFromInst(inst uintptr) *TTreeNode {
    return AsTreeNode(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsTreeNode.
func TreeNodeFromObj(obj IObject) *TTreeNode {
    return AsTreeNode(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTreeNode.
func TreeNodeFromUnsafePointer(ptr unsafe.Pointer) *TTreeNode {
    return AsTreeNode(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (t *TTreeNode) Free() {
    if t.instance != 0 {
        TreeNode_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTreeNode) Instance() uintptr {
    return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTreeNode) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTreeNode) IsValid() bool {
    return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TTreeNode) Is() TIs {
    return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TTreeNode) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTreeNodeClass() TClass {
    return TreeNode_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTreeNode) Assign(Source IObject) {
    TreeNode_Assign(t.instance, CheckPtr(Source))
}

func (t *TTreeNode) Collapse(Recurse bool) {
    TreeNode_Collapse(t.instance, Recurse)
}

func (t *TTreeNode) Delete() {
    TreeNode_Delete(t.instance)
}

func (t *TTreeNode) DisplayRect(TextOnly bool) TRect {
    return TreeNode_DisplayRect(t.instance, TextOnly)
}

func (t *TTreeNode) EditText() bool {
    return TreeNode_EditText(t.instance)
}

func (t *TTreeNode) Expand(Recurse bool) {
    TreeNode_Expand(t.instance, Recurse)
}

func (t *TTreeNode) IndexOf(Value *TTreeNode) int32 {
    return TreeNode_IndexOf(t.instance, CheckPtr(Value))
}

func (t *TTreeNode) MakeVisible() {
    TreeNode_MakeVisible(t.instance)
}

func (t *TTreeNode) MoveTo(Destination *TTreeNode, Mode TNodeAttachMode) {
    TreeNode_MoveTo(t.instance, CheckPtr(Destination), Mode)
}

func (t *TTreeNode) CustomSort(SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
    return TreeNode_CustomSort(t.instance, SortProc , Data , ARecurse)
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTreeNode) GetNamePath() string {
    return TreeNode_GetNamePath(t.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTreeNode) ClassType() TClass {
    return TreeNode_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTreeNode) ClassName() string {
    return TreeNode_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTreeNode) InstanceSize() int32 {
    return TreeNode_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTreeNode) InheritsFrom(AClass TClass) bool {
    return TreeNode_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTreeNode) Equals(Obj IObject) bool {
    return TreeNode_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTreeNode) GetHashCode() int32 {
    return TreeNode_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TTreeNode) ToString() string {
    return TreeNode_ToString(t.instance)
}

func (t *TTreeNode) AbsoluteIndex() int32 {
    return TreeNode_GetAbsoluteIndex(t.instance)
}

func (t *TTreeNode) Count() int32 {
    return TreeNode_GetCount(t.instance)
}

func (t *TTreeNode) Cut() bool {
    return TreeNode_GetCut(t.instance)
}

func (t *TTreeNode) SetCut(value bool) {
    TreeNode_SetCut(t.instance, value)
}

func (t *TTreeNode) Data() unsafe.Pointer {
    return TreeNode_GetData(t.instance)
}

func (t *TTreeNode) SetData(value unsafe.Pointer) {
    TreeNode_SetData(t.instance, value)
}

func (t *TTreeNode) Deleting() bool {
    return TreeNode_GetDeleting(t.instance)
}

// CN: 获取返回是否获取焦点。
// EN: Get Return to get focus.
func (t *TTreeNode) Focused() bool {
    return TreeNode_GetFocused(t.instance)
}

// CN: 设置返回是否获取焦点。
// EN: Set Return to get focus.
func (t *TTreeNode) SetFocused(value bool) {
    TreeNode_SetFocused(t.instance, value)
}

func (t *TTreeNode) Selected() bool {
    return TreeNode_GetSelected(t.instance)
}

func (t *TTreeNode) SetSelected(value bool) {
    TreeNode_SetSelected(t.instance, value)
}

func (t *TTreeNode) Expanded() bool {
    return TreeNode_GetExpanded(t.instance)
}

func (t *TTreeNode) SetExpanded(value bool) {
    TreeNode_SetExpanded(t.instance, value)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (t *TTreeNode) Handle() HWND {
    return TreeNode_GetHandle(t.instance)
}

func (t *TTreeNode) HasChildren() bool {
    return TreeNode_GetHasChildren(t.instance)
}

func (t *TTreeNode) SetHasChildren(value bool) {
    TreeNode_SetHasChildren(t.instance, value)
}

// CN: 获取图像在images中的索引。
// EN: .
func (t *TTreeNode) ImageIndex() int32 {
    return TreeNode_GetImageIndex(t.instance)
}

// CN: 设置图像在images中的索引。
// EN: .
func (t *TTreeNode) SetImageIndex(value int32) {
    TreeNode_SetImageIndex(t.instance, value)
}

func (t *TTreeNode) Index() int32 {
    return TreeNode_GetIndex(t.instance)
}

func (t *TTreeNode) IsVisible() bool {
    return TreeNode_GetIsVisible(t.instance)
}

func (t *TTreeNode) Level() int32 {
    return TreeNode_GetLevel(t.instance)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTreeNode) Owner() *TTreeNodes {
    return AsTreeNodes(TreeNode_GetOwner(t.instance))
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TTreeNode) Parent() *TTreeNode {
    return AsTreeNode(TreeNode_GetParent(t.instance))
}

func (t *TTreeNode) SelectedIndex() int32 {
    return TreeNode_GetSelectedIndex(t.instance)
}

func (t *TTreeNode) SetSelectedIndex(value int32) {
    TreeNode_SetSelectedIndex(t.instance, value)
}

func (t *TTreeNode) StateIndex() int32 {
    return TreeNode_GetStateIndex(t.instance)
}

func (t *TTreeNode) SetStateIndex(value int32) {
    TreeNode_SetStateIndex(t.instance, value)
}

// CN: 获取文本。
// EN: .
func (t *TTreeNode) Text() string {
    return TreeNode_GetText(t.instance)
}

// CN: 设置文本。
// EN: .
func (t *TTreeNode) SetText(value string) {
    TreeNode_SetText(t.instance, value)
}

func (t *TTreeNode) TreeView() *TWinControl {
    return AsWinControl(TreeNode_GetTreeView(t.instance))
}

func (t *TTreeNode) Item(Index int32) *TTreeNode {
    return AsTreeNode(TreeNode_GetItem(t.instance, Index))
}

func (t *TTreeNode) SetItem(Index int32, value *TTreeNode) {
    TreeNode_SetItem(t.instance, Index, CheckPtr(value))
}

