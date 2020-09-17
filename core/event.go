package core

import "github.com/gin-gonic/gin"

var Event *EventManager

type EventObject struct {
	Context *gin.Context
	Content interface{}
}


type EventManager struct {
	Event_UserCreate chan EventObject
	Event_UserEdit chan EventObject
	Event_UserDelete chan EventObject
	Event_RoleCreate chan EventObject
	Event_RoleEdit chan EventObject
	Event_RoleDelete chan EventObject
	Event_PermissionCreate chan EventObject
	Event_PermissionEdit chan EventObject
	Event_PermissionDelete chan EventObject
}


func init() {
	// Initialization Event
	// 初始化Event
	Event = &EventManager{
		Event_UserCreate:       make(chan EventObject, 10),
		Event_UserEdit:         make(chan EventObject, 10),
		Event_UserDelete:       make(chan EventObject, 10),
		Event_RoleCreate:       make(chan EventObject, 10),
		Event_RoleEdit:         make(chan EventObject, 10),
		Event_RoleDelete:       make(chan EventObject, 10),
		Event_PermissionCreate: make(chan EventObject, 10),
		Event_PermissionEdit:   make(chan EventObject, 10),
		Event_PermissionDelete: make(chan EventObject, 10),
	}
}


type Event_CallbackFunc func(EventObject)

func (this *EventManager) UserCreate (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_UserCreate
		callback(i)
	}
}

func (this *EventManager) UserEdit (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_UserEdit
		callback(i)
	}
}

func (this *EventManager) UserDelete (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_UserDelete
		callback(i)
	}
}

func (this *EventManager) RoleCreate (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_RoleCreate
		callback(i)
	}
}

func (this *EventManager) RoleEdit (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_RoleEdit
		callback(i)
	}
}

func (this *EventManager) RoleDelete (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_RoleDelete
		callback(i)
	}
}

func (this *EventManager) PermissionCreate (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_PermissionCreate
		callback(i)
	}
}

func (this *EventManager) PermissionEdit (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_PermissionEdit
		callback(i)
	}
}

func (this *EventManager) PermissionDelete (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_PermissionDelete
		callback(i)
	}
}



